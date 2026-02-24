package txsv

import (
	"log/slog"

	"github.com/sfborg/sf/config"
	"github.com/sfborg/sf/pkg/sf"
	"github.com/sfborg/sflib"
	"github.com/sfborg/sflib/pkg/sfga"
	"github.com/sfborg/sflib/pkg/xsv"
)

type txsv struct {
	cfg  config.Config
	sfga sfga.Archive
	xsv  xsv.Archive
}

func New(cfg config.Config) sf.ToConvertor {
	res := txsv{
		cfg:  cfg,
		sfga: sflib.NewSfga(),
		xsv:  sflib.NewXsv(),
	}
	return &res
}

func (t *txsv) Export(src, dst string) error {
	var err error
	if err = t.xsv.Create(t.cfg.OutputDir); err != nil {
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

	slog.Info("Exporting NameUsage file")
	err = t.convertNameUsages()
	if err != nil {
		return err
	}

	err = t.xsv.Export(dst, t.cfg.WithZipOutput)
	if err != nil {
		return err
	}

	slog.Info("Conversion to CSV is completed")
	return nil
}
