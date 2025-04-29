package fcoldp

import (
	"github.com/sfborg/sf/config"
	"github.com/sfborg/sf/pkg/from"
	"github.com/sfborg/sf/pkg/sf"
	"github.com/sfborg/sflib"
	"github.com/sfborg/sflib/pkg/coldp"
	"github.com/sfborg/sflib/pkg/sfga"
)

type fcoldp struct {
	cfg   config.Config
	sfga  sfga.Archive
	coldp coldp.Archive
	*from.Shared
}

func New(cfg config.Config) sf.FromConvertor {
	res := fcoldp{
		cfg:    cfg,
		coldp:  sflib.NewColdp(),
		Shared: from.New(cfg),
	}
	return &res
}
