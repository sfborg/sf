package diffio

import (
	"github.com/sfborg/sf/internal/ent/diff"
	"github.com/sfborg/sf/pkg/config"
	"github.com/sfborg/sflib/ent/sfga"
)

type diffio struct {
	cfg      config.Config
	src, trg sfga.Archive
}

func New(cfg config.Config, src, trg sfga.Archive) diff.Diff {
	res := diffio{cfg: cfg, src: src, trg: trg}
	return &res
}
