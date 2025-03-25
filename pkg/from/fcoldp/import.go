package fcoldp

import (
	"github.com/sfborg/sflib/pkg/arch"
)

func (fc *fcoldp) Import(src, out string) error {
	var err error

	fc.sfga, err = fc.InitSfga()
	if err != nil {
		return err
	}

	err = fc.coldp.Fetch(src, fc.cfg.DataDir)
	if err != nil {
		return &arch.ErrExtract{Path: src, Err: err}
	}

	err = fc.importMeta()
	if err != nil {
		return err
	}

	err = fc.importData()
	if err != nil {
		return err
	}

	err = fc.sfga.Export(out, fc.cfg.WithZipOutput)
	if err != nil {
		return err
	}

	return nil
}
