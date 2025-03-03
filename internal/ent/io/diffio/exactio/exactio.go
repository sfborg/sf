package exactio

import (
	"log/slog"
	"sync"

	"github.com/devopsfaith/bloomfilter"
	baseBloomfilter "github.com/devopsfaith/bloomfilter/bloomfilter"
	"github.com/sfborg/sf/internal/ent/diff"
)

type exactio struct {
	canonical     *baseBloomfilter.Bloomfilter
	canonicalSize uint
	mux           sync.Mutex
}

func New() diff.Exact {
	res := exactio{}
	return &res
}

func (e *exactio) Init(recs []diff.Record) {
	slog.Info("Setting up data for exact matching")
	e.canonicalSize = uint(len(recs))
	cfg := bloomfilter.Config{
		N:        e.canonicalSize,
		P:        0.00001,
		HashName: bloomfilter.HASHER_OPTIMAL,
	}
	bf := baseBloomfilter.New(cfg)

	for i := range recs {
		bf.Add([]byte(recs[i].CanonicalSimple))
	}
	e.canonical = bf
}

func (e *exactio) Find(string) bool {
	return false
}
