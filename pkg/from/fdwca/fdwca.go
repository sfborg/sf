package fdwca

import (
	"github.com/sfborg/sf/config"
	"github.com/sfborg/sf/pkg/from"
	"github.com/sfborg/sf/pkg/sf"
	"github.com/sfborg/sflib"
	libCfg "github.com/sfborg/sflib/config"
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
	opts := []libCfg.Option{
		libCfg.OptBadRow(cfg.BadRow),
		libCfg.OptCode(cfg.NomCode),
		libCfg.OptJobsNum(cfg.JobsNum),
	}

	res := fdwca{
		cfg:    cfg,
		dwca:   sflib.NewDwca(opts...),
		Shared: from.New(cfg),
	}
	return &res
}
