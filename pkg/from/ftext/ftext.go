package ftext

import (
	"github.com/sfborg/sf/pkg/config"
	"github.com/sfborg/sf/pkg/from"
	"github.com/sfborg/sf/pkg/sf"
	"github.com/sfborg/sflib"
	"github.com/sfborg/sflib/pkg/sfga"
	"github.com/sfborg/sflib/pkg/text"
)

type ftext struct {
	cfg  config.Config
	sfga sfga.Archive
	text text.Archive
	*from.Shared
}

func New(cfg config.Config) sf.FromConvertor {
	res := ftext{
		cfg:    cfg,
		Shared: from.New(cfg),
		text:   sflib.NewText(),
	}
	return &res
}
