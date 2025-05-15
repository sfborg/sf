package fcoldp

import (
	"github.com/sfborg/sflib/pkg/arch"
)

func (fc *fcoldp) Import(src, out string) error {
	var err error

	// create empty SFGA database and connect to it
	fc.sfga, err = fc.InitSfga()
	if err != nil {
		return err
	}

	// get CoLDP file, unzip into working directory
	err = fc.coldp.Fetch(src, fc.cfg.ImportDir)
	if err != nil {
		return &arch.ErrExtract{Path: src, Err: err}
	}

	// import metadata to SFGA
	err = fc.importMeta()
	if err != nil {
		return err
	}

	// import all other data to SFGA
	err = fc.importData()
	if err != nil {
		return err
	}

	// create sql and sqlite files with optional zip compression
	err = fc.sfga.Export(out, fc.cfg.WithZipOutput)
	if err != nil {
		return err
	}

	return nil
}
