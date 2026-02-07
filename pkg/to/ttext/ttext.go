package ttext

import (
	"log/slog"

	"github.com/sfborg/sf/config"
	"github.com/sfborg/sf/pkg/sf"
	"github.com/sfborg/sflib"
	"github.com/sfborg/sflib/pkg/sfga"
	"github.com/sfborg/sflib/pkg/text"
)

type ttext struct {
	cfg  config.Config
	sfga sfga.Archive
	txt  text.Archive
}

func New(cfg config.Config) sf.ToConvertor {
	res := ttext{
		cfg:  cfg,
		sfga: sflib.NewSfga(),
		txt:  sflib.NewText(),
	}
	return &res
}

func (t *ttext) Export(src, dst string) error {
	var err error
	err = t.sfga.Fetch(src, t.cfg.ImportDir)
	if err != nil {
		return err
	}
	_, err = t.sfga.Connect()
	if err != nil {
		return err
	}

	err = t.txt.Create(t.cfg.OutputDir)
	if err != nil {
		return err
	}

	slog.Info("Exporting NameUsage file")
	err = t.convertNameUsages()
	if err != nil {
		return err
	}

	err = t.txt.Export(dst, t.cfg.WithZipOutput)
	if err != nil {
		return err
	}

	slog.Info("Conversion to text is completed")
	return nil
}
