package fcoldp

import (
	"fmt"

	"github.com/gnames/gnsys"
)

func (cd *fcoldp) Import(src, out string) error {
	var err error

	src, err = cd.Download(src)
	if err != nil {
		return err
	}

	exists, _ := gnsys.FileExists(src)

	if !exists {
		return fmt.Errorf("file does not exist '%s'", src)
	}

	err = cd.Extract(src)
	if err != nil {
		return err
	}

	cd.sfga, err = cd.InitSfga()
	if err != nil {
		return err
	}

	err = cd.importColdp()
	if err != nil {
		return err
	}

	err = cd.sfga.Export(out, cd.cfg.WithZipOutput)
	if err != nil {
		return err
	}

	return nil
}
