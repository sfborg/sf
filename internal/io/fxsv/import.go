package fxsv

import (
	"github.com/gnames/gnsys"
	"github.com/sfborg/sflib/pkg/arch"
)

func (fx *fxsv) Import(src, out string) error {
	var err error

	src, err = fx.Download(src)
	if err != nil {
		return err
	}

	exists, _ := gnsys.FileExists(src)

	if !exists {
		return &arch.ErrFileNotFound{Path: src}
	}

	fx.csvPath = src

	fx.sfga, err = fx.InitSfga()
	if err != nil {
		return err
	}

	err = fx.importNamesUsage()
	if err != nil {
		return err
	}

	err = fx.sfga.Export(out, fx.cfg.WithZipOutput)
	if err != nil {
		return err
	}

	return nil
}
