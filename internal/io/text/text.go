package text

import (
	"github.com/sfborg/sf/internal/io/fromx"
	sf "github.com/sfborg/sf/pkg"
	"github.com/sfborg/sf/pkg/config"
)

type text struct {
	cfg config.Config
	sf.FromX
}

func New(cfg config.Config) sf.FromX {
	res := text{
		cfg:   cfg,
		FromX: fromx.New(cfg),
	}

	return &res
}
