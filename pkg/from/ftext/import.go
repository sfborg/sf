package ftext

import (
	"github.com/sfborg/sf/internal/util"
)

func (ft *ftext) Import(src, out string) error {
	var err error
	err = util.PrepareFileStructure(ft.cfg)

	err = ft.text.Fetch(src, ft.cfg.ImportDir)
	if err != nil {
		return err
	}

	ft.sfga, err = ft.InitSfga()
	if err != nil {
		return err
	}

	err = ft.importNamesUsage()
	if err != nil {
		return err
	}

	err = ft.sfga.Export(out, ft.cfg.WithZipOutput)
	if err != nil {
		return err
	}

	return nil
}
