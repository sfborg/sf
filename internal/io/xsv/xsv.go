package xsv

import (
	sf "github.com/sfborg/sf/pkg"
	"github.com/sfborg/sf/pkg/config"
	"github.com/sfborg/sflib/ent/sfga"
)

type xsv struct {
	cfg     config.Config
	sfga    sfga.Archive
	csvPath string
}

func New(cfg config.Config) sf.Importer {
	res := xsv{
		cfg: cfg,
	}
	return &res
}
