package matchio

import (
	"strings"

	"github.com/gnames/gnlib/ent/verifier"
	"github.com/sfborg/sf/pkg/diff"
)

func (m *matchio) Match(rec diff.Record) ([]diff.Record, error) {
	res, err := m.MatchExact(rec.CanonicalSimple)
	if len(res) > 0 || err != nil {
		return res, err
	}

	res, err = m.MatchFuzzy(rec.CanonicalSimple, rec.CanonicalStemmed)
	if len(res) > 0 || err != nil {
		return res, err
	}

	if rec.Cardinality > 1 {
		res, err = m.partialMatch(rec.CanonicalSimple, rec.CanonicalStemmed)
	}

	return res, err
}

type canPair struct {
	can, stem   string
	cardinality int
}

func (m *matchio) partialMatch(can, stem string) ([]diff.Record, error) {
	var res []diff.Record
	var err error
	cans := partialCombos(can, stem)
	for i := range cans {
		res, err = m.MatchExact(cans[i].can)
		if len(res) > 0 || err != nil {
			for i := range res {
				res[i].MatchType = verifier.PartialExact
			}
			return res, err
		}

		if cans[i].cardinality == 1 {
			break
		}

		res, err = m.MatchFuzzy(cans[i].can, cans[i].stem)
		if len(res) > 0 || err != nil {
			for i := range res {
				res[i].MatchType = verifier.PartialFuzzy
			}
			return res, err
		}
	}

	return res, err
}

func partialCombos(can, stem string) []canPair {
	canWs := strings.Split(can, " ")
	stemWs := strings.Split(stem, " ")
	switch len(canWs) {
	case 2:
		return []canPair{
			{
				can:         canWs[0],
				stem:        stemWs[0],
				cardinality: 1,
			},
		}
	case 3:
		return []canPair{
			{
				can:         canWs[0] + " " + canWs[2],
				stem:        stemWs[0] + " " + stemWs[2],
				cardinality: 2,
			},
			{
				can:         canWs[0] + " " + canWs[1],
				stem:        stemWs[0] + " " + stemWs[1],
				cardinality: 2,
			},
			{
				can:         canWs[0],
				stem:        canWs[0],
				cardinality: 1,
			},
		}
	case 4:
		return []canPair{
			{
				can:         canWs[0] + " " + canWs[3],
				stem:        stemWs[0] + " " + stemWs[3],
				cardinality: 2,
			},
			{
				can:         canWs[0] + " " + canWs[2],
				stem:        stemWs[0] + " " + stemWs[2],
				cardinality: 2,
			},
			{
				can:         canWs[0] + " " + canWs[1] + " " + canWs[2],
				stem:        stemWs[0] + " " + stemWs[1] + " " + stemWs[2],
				cardinality: 2,
			},
			{
				can:         canWs[0] + " " + canWs[1],
				stem:        stemWs[0] + " " + stemWs[1],
				cardinality: 2,
			},
			{
				can:         canWs[0],
				stem:        canWs[0],
				cardinality: 1,
			},
		}
	default:
		return nil
	}
}
