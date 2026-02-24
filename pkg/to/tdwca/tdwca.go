package tdwca

import (
	"log/slog"

	"github.com/sfborg/sf/config"
	"github.com/sfborg/sf/pkg/sf"
	"github.com/sfborg/sflib"
	"github.com/sfborg/sflib/pkg/dwca"
	"github.com/sfborg/sflib/pkg/sfga"
)

type tdwca struct {
	cfg  config.Config
	sfga sfga.Archive
	dwca dwca.Archive
}

func New(cfg config.Config) sf.ToConvertor {
	res := tdwca{
		cfg:  cfg,
		sfga: sflib.NewSfga(),
		dwca: sflib.NewDwca(),
	}
	return &res
}

func (t *tdwca) Export(src, dst string) error {
	var err error
	if err = t.dwca.Create(t.cfg.OutputDir); err != nil {
		return err
	}

	err = t.sfga.Fetch(src, t.cfg.ImportDir)
	if err != nil {
		return err
	}
	_, err = t.sfga.Connect()
	if err != nil {
		return err
	}

	slog.Info("Exporting Metadata file")
	err = t.convertMeta()
	if err != nil {
		return err
	}

	slog.Info("Exporting DwCA Core")
	err = t.convertNameUsage()
	if err != nil {
		return err
	}

	slog.Info("Exporting Vernaculars")
	err = t.convertVernacular()
	if err != nil {
		return err
	}

	slog.Info("Exporting Distributions")
	err = t.convertDistribution()
	if err != nil {
		return err
	}

	slog.Info("Generate meta.xml")
	err = t.genMetaDwCA()
	if err != nil {
		return err
	}

	err = t.dwca.Export(dst, true)
	if err != nil {
		return err
	}

	return nil
}
