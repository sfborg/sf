package diffio

import (
	"fmt"
	"io"
	"log/slog"
	"os"
	"path/filepath"

	"github.com/gnames/gnsys"
	"github.com/gnames/gnuuid"
	"github.com/sfborg/sf/internal/ent/diff"
	"github.com/sfborg/sf/internal/ent/io/diffio/matchio"
	"github.com/sfborg/sflib/ent/sfga"
	"github.com/sfborg/sflib/io/sfgaio"
)

var recQ = `
SELECT
	col__id, gn__scientific_name_string, gn__canonical_simple,
	gn__canonical_stemmed
FROM name
`

func (d *diffio) Compare(src, ref string) error {
	var err error
	slog.Info("Getting reference file", "path", src)
	d.src, err = d.initSfga(src, d.cfg.DiffSrcDir)
	if err != nil {
		return err
	}
	slog.Info("Getting comparison file", "path", ref)
	d.ref, err = d.initSfga(ref, d.cfg.DiffTrgDir)
	if err != nil {
		return err
	}

	refPath := d.ref.DbPath()
	slog.Info("Creating a hash signature for reference")
	d.refUUID, err = gnuuid.FromFile(refPath)
	if err != nil {
		return err
	}
	slog.Info(
		"Generated UUID v5 for the reference",
		"file",
		filepath.Base(refPath),
		"uuid",
		d.refUUID,
	)

	d.workDir = filepath.Join(d.cfg.DiffWorkDir, d.refUUID.String())
	// dirState := gnsys.GetDirState(d.workDir)
	// if dirState != gnsys.DirNotEmpty {
	d.setRefSpace()
	// }

	defer d.src.Close()
	defer d.ref.Close()

	srcRec, err := d.sourceRecords()
	for _, v := range srcRec {
		res, err := d.matcher.Match(v)
		if err != nil {
			return err
		}
		l := len(res)
		if l == 0 {
			fmt.Printf("%s: %s\n", v.ID, v.Name)
		} else {
			for i := range res {
				fmt.Printf("%s: %s => %s: %s (%s)\n", v.ID, v.Name, res[i].ID, res[i].Name, res[i].MatchType)
			}
		}
	}

	return nil
}

func (d *diffio) sourceRecords() ([]diff.Record, error) {
	rows, err := d.src.Db().Query(recQ)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var res []diff.Record
	for rows.Next() {
		var r diff.Record
		err = rows.Scan(
			&r.ID, &r.Name, &r.CanonicalSimple, &r.CanonicalStemmed,
		)
		if err != nil {
			return nil, err
		}
		res = append(res, r)
	}
	return res, nil
}

func (d *diffio) setRefSpace() error {
	dir := d.workDir
	var err error
	_ = os.RemoveAll(dir)
	slog.Info("Optimizing reference for comparison")
	err = gnsys.MakeDir(dir)
	if err != nil {
		return err
	}
	newRef := filepath.Join(dir, "schema.sqlite")
	copyFile(d.ref.DbPath(), newRef)
	d.ref.SetDb(newRef)
	_, err = d.ref.Connect()
	if err != nil {
		return err
	}

	err = d.setMatcher()
	if err != nil {
		return err
	}
	return nil
}

func (d *diffio) setMatcher() error {
	d.matcher = matchio.New()
	rows, err := d.ref.Db().Query(recQ)
	if err != nil {
		return err
	}
	defer rows.Close()
	var recs []diff.Record
	for rows.Next() {
		var r diff.Record
		err = rows.Scan(&r.ID, &r.Name, &r.CanonicalSimple, &r.CanonicalStemmed)
		if err != nil {
			return err
		}
		recs = append(recs, r)
	}
	err = d.matcher.Init(d.ref.Db(), recs)
	if err != nil {
		return err
	}
	return nil
}

func copyFile(path1, path2 string) error {
	sourceFile, err := os.Open(path1)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destFile, err := os.Create(path2)
	if err != nil {
		return err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, sourceFile)
	if err != nil {
		return err
	}

	err = destFile.Sync()
	if err != nil {
		return err
	}
	return nil
}

func (d *diffio) initSfga(path, dir string) (sfga.Archive, error) {
	sfga := sfgaio.New()
	err := sfga.Import(path, dir)
	if err != nil {
		return nil, err
	}
	_, err = sfga.Connect()
	if err != nil {
		return nil, err
	}
	return sfga, nil
}
