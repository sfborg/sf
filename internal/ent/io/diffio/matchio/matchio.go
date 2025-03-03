package matchio

import (
	"database/sql"

	"github.com/sfborg/sf/internal/ent/diff"
	"github.com/sfborg/sf/internal/ent/io/diffio/dbio"
	"github.com/sfborg/sf/internal/ent/io/diffio/exactio"
	"github.com/sfborg/sf/internal/ent/io/diffio/fuzzyio"
)

type matchio struct {
	db diff.DBase
	e  diff.Exact
	f  diff.Fuzzy
}

func New() diff.Matcher {
	res := matchio{
		db: dbio.New(),
		e:  exactio.New(),
		f:  fuzzyio.New(),
	}
	return &res
}

func (m *matchio) Init(db *sql.DB, recs []diff.Record) error {
	var err error
	m.db.Init(db)
	m.e.Init(recs)
	err = m.f.Init(recs)
	if err != nil {
		return err
	}
	return nil
}

func (m *matchio) Match(diff.Record) ([]diff.Record, error)         { return nil, nil }
func (m *matchio) MatchExact(string) ([]diff.Record, error)         { return nil, nil }
func (m *matchio) MatchFuzzy(string, string) ([]diff.Record, error) { return nil, nil }
