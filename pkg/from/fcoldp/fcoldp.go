package fcoldp

import (
	"sync"

	"github.com/gnames/gnlib/ent/nomcode"
	"github.com/sfborg/sf/config"
	"github.com/sfborg/sf/pkg/from"
	"github.com/sfborg/sf/pkg/sf"
	"github.com/sfborg/sflib"
	sflCfg "github.com/sfborg/sflib/config"
	"github.com/sfborg/sflib/pkg/coldp"
	"github.com/sfborg/sflib/pkg/parser"
	"github.com/sfborg/sflib/pkg/sfga"
)

type fcoldp struct {
	cfg   config.Config
	sfga  sfga.Archive
	coldp coldp.Archive
	*from.Shared
	parserPool map[nomcode.Code]*sync.Pool
}

func New(cfg config.Config) sf.FromConvertor {
	opts := []sflCfg.Option{
		sflCfg.OptWithQuotes(cfg.WithQuotes),
	}
	res := fcoldp{
		cfg:        cfg,
		coldp:      sflib.NewColdp(opts...),
		Shared:     from.New(cfg),
		parserPool: parser.Pool(cfg.JobsNum),
	}
	return &res
}
