package matchio

import (
	"database/sql"

	"github.com/gnames/gnlib/ent/verifier"
	"github.com/sfborg/sf/internal/ent/diff"
	"github.com/sfborg/sf/internal/io/diffio/dbio"
	"github.com/sfborg/sf/internal/io/diffio/exactio"
	"github.com/sfborg/sf/internal/io/diffio/fuzzyio"
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

func (m *matchio) MatchExact(canonical string) ([]diff.Record, error) {
	var err error
	var res []diff.Record
	if m.e.Find(canonical) {
		res, err = m.db.Select(canonical)
	}
	for i := range res {
		res[i].MatchType = verifier.Exact
	}
	return res, err
}

func (m *matchio) MatchFuzzy(can, stem string) ([]diff.Record, error) {
	var res []diff.Record
	var canonicals []string
	if canonicals = m.f.FindExact(stem); len(canonicals) > 0 {
		return m.fetchCanonicals(can, canonicals, true)
	}
	if canonicals = m.f.FindFuzzy(stem); len(canonicals) > 0 {
		return m.fetchCanonicals(can, canonicals, false)
	}
	return res, nil
}

func (m *matchio) fetchCanonicals(
	can string,
	cans []string,
	noCheck bool,
) ([]diff.Record, error) {
	var err error
	var recs, res []diff.Record
	for i := range cans {
		recs, err = m.db.Select(cans[i])
		if err != nil {
			return res, err
		}

		for ii := range recs {
			ed := fuzzyio.EditDistance(can, recs[ii].CanonicalSimple, noCheck)
			if ed < 0 {
				continue
			}

			recs[ii].EditDistance = ed
			recs[ii].MatchType = verifier.Fuzzy
			res = append(res, recs[ii])
		}
	}
	return res, err
}
