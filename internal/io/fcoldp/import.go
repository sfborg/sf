package fcoldp

import (
	"github.com/gnames/gnsys"
	"github.com/sfborg/sflib/pkg/arch"
	"github.com/sfborg/sflib/pkg/sflib"
)

func (fc *fcoldp) Import(src, out string) error {
	var err error

	src, err = fc.Download(src)
	if err != nil {
		return err
	}

	exists, _ := gnsys.FileExists(src)

	if !exists {
		return &arch.ErrFileNotFound{Path: src}
	}

	fc.sfga, err = fc.InitSfga()
	if err != nil {
		return err
	}

	coldp := sflib.NewColdp()

	err = coldp.Import(src, fc.cfg.DataDir)
	if err != nil {
		return &arch.ErrExtract{Path: src, Err: err}
	}

	err = fc.importMeta(coldp)
	if err != nil {
		return err
	}

	err = fc.importData(coldp)
	if err != nil {
		return err
	}

	err = fc.sfga.Export(out, fc.cfg.WithZipOutput)
	if err != nil {
		return err
	}

	return nil
}
