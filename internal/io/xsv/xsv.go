package xsv

import (
	"github.com/sfborg/sf/internal/io/fromx"
	sf "github.com/sfborg/sf/pkg"
	"github.com/sfborg/sf/pkg/config"
	"github.com/sfborg/sflib/ent/sfga"
)

type xsv struct {
	cfg     config.Config
	sfga    sfga.Archive
	csvPath string
	sf.FromX
}

func New(cfg config.Config) sf.FromX {
	res := xsv{
		cfg:   cfg,
		FromX: fromx.New(cfg),
	}
	return &res
}
