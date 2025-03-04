package dbio

import (
	"database/sql"

	"github.com/sfborg/sf/internal/ent/diff"
)

type dbio struct {
	db *sql.DB
}

func New() diff.DBase {
	res := dbio{}
	return &res
}

func (d *dbio) Init(db *sql.DB) {
	d.db = db
}
func (d *dbio) Select(canonical string) ([]diff.Record, error) {
	q := `
SELECT col__id, col__scientific_name, gn__canonical_simple,
	gn__canonical_full, gn__canonical_stemmed, col__authorship
	FROM name
	WHERE gn__canonical_simple = ?
`
	rows, err := d.db.Query(q, canonical)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var res []diff.Record
	for rows.Next() {
		var r diff.Record
		err = rows.Scan(&r.ID, &r.Name, &r.CanonicalSimple,
			&r.CanonicalFull, &r.CanonicalStemmed, &r.Authors,
		)
		if err != nil {
			return nil, err
		}
		res = append(res, r)
	}
	return res, nil
}
