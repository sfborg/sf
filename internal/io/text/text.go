package text

import (
	sf "github.com/sfborg/sf/pkg"
	"github.com/sfborg/sf/pkg/config"
)

type text struct {
	cfg config.Config
}

func New(cfg config.Config) sf.Importer {
	res := text{
		cfg: cfg,
	}
	return &res
}
