package fxsv

import (
	"log/slog"

	"github.com/sfborg/sf/pkg/from/fsfga"
)

func (fx *fxsv) Import(src, out string) error {
	var err error

	err = fx.xsv.Fetch(src, fx.cfg.ImportDir)
	if err != nil {
		return err
	}

	fx.sfga, err = fx.InitSfga()
	if err != nil {
		return err
	}

	err = fx.importNamesUsage()
	if err != nil {
		return err
	}

	if !fx.cfg.WithParents {
		return fx.sfga.Export(out, fx.cfg.WithZipOutput)
	}

	err = fx.sfga.Export(out, false)
	if err != nil {
		return err
	}

	slog.Info("Converting flat hierarchy to parent/child tree")
	fs := fsfga.New(fx.cfg)
	return fs.Import(out+".sqlite", out)
}
