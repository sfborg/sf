package fdwca

import (
	"github.com/sfborg/sf/pkg/config"
	"github.com/sfborg/sf/pkg/from"
	"github.com/sfborg/sf/pkg/sf"
	"github.com/sfborg/sflib"
	"github.com/sfborg/sflib/pkg/dwca"
	"github.com/sfborg/sflib/pkg/sfga"
)

type fdwca struct {
	cfg  config.Config
	sfga sfga.Archive
	dwca dwca.Archive
	*from.Shared
}

func New(cfg config.Config) sf.FromConvertor {
	res := fdwca{
		cfg:    cfg,
		dwca:   sflib.NewDwca(),
		Shared: from.New(cfg),
	}
	return &res
}
