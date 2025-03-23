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
	"github.com/sfborg/sf/internal/io/diffio/matchio"
	"github.com/sfborg/sflib/pkg/sfga"
	"github.com/sfborg/sflib/pkg/sflib"
)

var recQ = `
SELECT
	col__id, gn__scientific_name_string, gn__canonical_simple,
	gn__canonical_stemmed
FROM name
`

func (d *diffio) Compare(src, ref, out string) error {
	var err error
	slog.Info("Getting reference file", "path", src)
	d.src, err = d.initSfga(src, d.cfg.DiffSrcDir)
	if err != nil {
		return err
	}
	slog.Info("Getting comparison file", "path", ref)
	d.ref, err = d.initSfga(ref, d.cfg.DiffRefDir)
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

	file := filepath.Base(out) + ".sqlite"
	slog.Info("Comparing, saving results to name_match table", "output", file)
	srcRec, err := d.sourceRecords()
	qRes := `
INSERT INTO name_match
  (
    col__name_id, gn__scientific_name_string,
		ref_col__name_id, ref_gn__scientific_name_string,
    gn__match_id
  )
  VALUES (?, ?, ?, ?, ?)
`
	qEmpty := `
INSERT INTO name_match
  (
    col__name_id, gn__scientific_name_string
  )
  VALUES (?, ?)
`

	for i, v := range srcRec {
		if (i+1)%100_000 == 0 {
			fmt.Printf("Processed %d names\n", i+1)
		}
		var res []diff.Record
		if v.CanonicalSimple == "" {
			continue
		}
		res, err = d.matcher.Match(v)
		if err != nil {
			return err
		}
		l := len(res)
		if l == 0 {
			_, err = d.src.Db().Exec(qEmpty, v.ID, v.Name)
			if err != nil {
				return err
			}
		} else {
			for i := range res {
				match := matchTypeID(res[i].MatchType.String())
				_, err := d.src.Db().Exec(qRes,
					v.ID, v.Name,
					res[i].ID, res[i].Name, match,
				)
				if err != nil {
					return err
				}
			}
		}
	}
	err = d.src.Export(out, false)
	if err != nil {
		return err
	}
	return nil
}

func matchTypeID(s string) string {
	switch s {
	case "Exact", "ExactSpeciesGroup", "Virus":
		return "EXACT"
	case "PartialExact":
		return "EXACT_PARTIAL"
	case "Fuzzy", "FuzzyRelaxed", "FuzzySpeciesGroup", "FuzzySpeciesGroupRelaxed":
		return "FUZZY"
	case "PartialFuzzy", "PartialFuzzyRelaxed":
		return "FUZZY_PARTIAL"
	default:
		return ""
	}
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
	sfga := sflib.NewSfga()
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
