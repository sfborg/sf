package text

import (
	"github.com/sfborg/sf/internal/io/fromx"
	sf "github.com/sfborg/sf/pkg"
	"github.com/sfborg/sf/pkg/config"
	"github.com/sfborg/sflib/ent/sfga"
)

type text struct {
	cfg  config.Config
	sfga sfga.Archive
	sf.FromX
	txtPath string
}

func New(cfg config.Config) sf.FromX {
	res := text{
		cfg:   cfg,
		FromX: fromx.New(cfg),
	}
	return &res
}
