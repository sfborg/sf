package diffio

import (
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
	dirState := gnsys.GetDirState(d.workDir)
	if dirState != gnsys.DirNotEmpty || true {
		d.setWorkRef()
	}

	defer d.src.Close()
	defer d.ref.Close()

	return nil
}

func (d *diffio) setWorkRef() error {
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

	d.setMatcher()
	return nil
}

func (d *diffio) setMatcher() error {
	m := matchio.New()
	q := `
SELECT
	col__id, gn__scientific_name_string, gn__canonical_stemmed
FROM name
`
	rows, err := d.ref.Db().Query(q)
	if err != nil {
		return err
	}
	defer rows.Close()
	var recs []diff.Record
	for rows.Next() {
		var r diff.Record
		err = rows.Scan(&r.ID, &r.Name, &r.CanonicalStemmed)
		if err != nil {
			return err
		}
		recs = append(recs, r)
	}
	err = m.Init(d.ref.Db(), recs)
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
