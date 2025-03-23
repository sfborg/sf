package fxsv

import (
	"fmt"

	"github.com/gnames/gnsys"
)

func (x *fxsv) Import(src, out string) error {
	var err error

	src, err = x.Download(src)
	if err != nil {
		return err
	}

	exists, _ := gnsys.FileExists(src)

	if !exists {
		return fmt.Errorf("file does not exist '%s'", src)
	}

	x.csvPath = src

	x.sfga, err = x.InitSfga()
	if err != nil {
		return err
	}

	err = x.importNamesUsage()
	if err != nil {
		return err
	}

	err = x.sfga.Export(out, x.cfg.WithZipOutput)
	if err != nil {
		return err
	}

	return nil
}
