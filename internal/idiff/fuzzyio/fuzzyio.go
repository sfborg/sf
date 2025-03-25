package fuzzyio

import (
	"sort"

	"github.com/dvirsky/levenshtein"
	"github.com/sfborg/sf/pkg/diff"
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

func (f *fuzzyio) FindExact(stem string) []string {
	return f.find(stem, 0)
}
func (f *fuzzyio) FindFuzzy(stem string) []string {
	return f.find(stem, 1)
}

func (f *fuzzyio) find(stem string, maxDist int) []string {
	stems := f.trie.FuzzyMatches(stem, maxDist)
	resMap := make(map[string]struct{})
	for i := range stems {
		cs := f.canonicals[stems[i]]
		for i := range cs {
			resMap[cs[i]] = struct{}{}
		}
	}
	res := make([]string, len(resMap))
	var i int
	for k := range resMap {
		res[i] = k
		i++
	}
	return res
}
