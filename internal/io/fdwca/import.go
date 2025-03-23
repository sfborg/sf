package fdwca

import (
	"github.com/gnames/gnsys"
	"github.com/sfborg/sflib/pkg/arch"
	"github.com/sfborg/sflib/pkg/sflib"
)

func (fd *fdwca) Import(src, out string) error {
	var err error

	src, err = fd.Download(src)
	if err != nil {
		return err
	}

	exists, _ := gnsys.FileExists(src)

	if !exists {
		return &arch.ErrFileNotFound{Path: src}
	}

	fd.sfga, err = fd.InitSfga()
	if err != nil {
		return err
	}

	dwca := sflib.NewDwca()

	err = dwca.Import(src, fd.cfg.DataDir)
	if err != nil {
		return &arch.ErrExtract{Path: src, Err: err}
	}

	err = fd.importNamesUsage()
	if err != nil {
		return err
	}

	err = fd.sfga.Export(out, fd.cfg.WithZipOutput)
	if err != nil {
		return err
	}

	return nil
}
