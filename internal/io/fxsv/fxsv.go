package fxsv

import (
	"github.com/gnames/gnparser"
	"github.com/sfborg/sf/internal/io/fromx"
	sf "github.com/sfborg/sf/pkg"
	"github.com/sfborg/sf/pkg/config"
	"github.com/sfborg/sflib/pkg/sfga"
)

type fxsv struct {
	cfg     config.Config
	sfga    sfga.Archive
	csvPath string
	sf.FromX
	parserPool map[string]chan gnparser.GNparser
}

func New(cfg config.Config) sf.FromX {
	res := fxsv{
		cfg:   cfg,
		FromX: fromx.New(cfg),

		parserPool: sf.ParserPool(cfg.JobsNum, cfg.WithDetails),
	}
	return &res
}
