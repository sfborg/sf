package fcoldp

import (
	"github.com/sfborg/sf/internal/io/fromx"
	sf "github.com/sfborg/sf/pkg"
	"github.com/sfborg/sf/pkg/config"
	"github.com/sfborg/sflib/pkg/dwca"
	"github.com/sfborg/sflib/pkg/sfga"
)

type fcoldp struct {
	cfg  config.Config
	sfga sfga.Archive
	dwca dwca.Archive
	sf.FromX
}

func New(cfg config.Config) sf.FromX {
	res := fcoldp{
		cfg:   cfg,
		FromX: fromx.New(cfg),
	}
	return &res
}
