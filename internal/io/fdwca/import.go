package fdwca

import (
	"github.com/gnames/gnsys"
	"github.com/sfborg/sflib/pkg/arch"
	"github.com/sfborg/sflib/pkg/sflib"
)

func (d *fdwca) Import(src, out string) error {
	var err error

	src, err = d.Download(src)
	if err != nil {
		return err
	}

	exists, _ := gnsys.FileExists(src)

	if !exists {
		return &arch.ErrFileNotFound{Path: src}
	}

	d.sfga, err = d.InitSfga()
	if err != nil {
		return err
	}

	dwca := sflib.NewDwca()

	err = dwca.Import(src, d.cfg.DataDir)
	if err != nil {
		return &arch.ErrExtract{Path: src, Err: err}
	}

	err = d.importNamesUsage()
	if err != nil {
		return err
	}

	err = d.sfga.Export(out, d.cfg.WithZipOutput)
	if err != nil {
		return err
	}

	return nil
}
