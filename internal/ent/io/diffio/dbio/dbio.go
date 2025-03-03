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
func (d *dbio) Select(string) ([]diff.Record, error) {
	return nil, nil
}
