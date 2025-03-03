package fuzzyio

import (
	"sort"

	"github.com/dvirsky/levenshtein"
	"github.com/sfborg/sf/internal/ent/diff"
)

type fuzzyio struct {
	trie       *levenshtein.MinTree
	canonicals map[string][]string
}

func New() diff.Fuzzy {
	res := fuzzyio{}
	res.canonicals = make(map[string][]string)
	return &res
}

func (f *fuzzyio) Init(recs []diff.Record) error {
	var err error
	stems := make([]string, len(recs))
	for i := range recs {
		stem := recs[i].CanonicalStemmed
		stems[i] = stem
		f.canonicals[stem] = append(f.canonicals[stem], recs[i].CanonicalSimple)
	}
	sort.Strings(stems)
	f.trie, err = levenshtein.NewMinTree(stems)
	if err != nil {
		return err
	}
	return nil
}

func (f *fuzzyio) FindExact(string) []string { return nil }
func (f *fuzzyio) FindFuzzy(string) []string { return nil }
