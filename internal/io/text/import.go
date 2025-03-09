package text

import (
	"fmt"

	"github.com/gnames/gnsys"
)

func (t *text) Import(src, out string) error {
	var err error

	src, err = t.Download(src)
	if err != nil {
		return err
	}

	exists, _ := gnsys.FileExists(src)

	if !exists {
		return fmt.Errorf("file does not exist '%s'", src)
	}

	t.textPath = src

	t.sfga, err = t.InitSfga()
	if err != nil {
		return err
	}

	err = t.importNamesUsage()
	if err != nil {
		return err
	}

	err = t.sfga.Export(out, t.cfg.WithZipOutput)
	if err != nil {
		return err
	}

	return nil
}
