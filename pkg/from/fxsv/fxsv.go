package fxsv

import (
	"github.com/sfborg/sf/pkg/config"
	"github.com/sfborg/sf/pkg/from"
	"github.com/sfborg/sf/pkg/sf"
	"github.com/sfborg/sflib"
	"github.com/sfborg/sflib/pkg/sfga"
	"github.com/sfborg/sflib/pkg/xsv"
)

type fxsv struct {
	cfg  config.Config
	sfga sfga.Archive
	xsv  xsv.Archive
	*from.Shared
}

func New(cfg config.Config) sf.FromConvertor {
	res := fxsv{
		cfg:    cfg,
		Shared: from.NewShared(cfg),
		xsv:    sflib.NewXsv(),
	}
	return &res
}
