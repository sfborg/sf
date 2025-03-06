package xsv

import (
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/gnames/gnsys"
	"github.com/sfborg/sflib/ent/sfga"
	"github.com/sfborg/sflib/io/sfgaio"
)

func (x *xsv) Import(src, out string) error {
	var err error

	if strings.HasPrefix(src, "http") {
		dir, err := os.MkdirTemp("", "sf-")
		if err != nil {
			return err
		}
		defer os.RemoveAll(dir)
		src, err = gnsys.Download(src, dir, true)
		if err != nil {
			return err
		}
	}

	err = x.copy(src, x.cfg.ImporterSfgaDir)
	if err != nil {
		return err
	}

	x.sfga, err = x.initSfga()
	if err != nil {
		return err
	}

	err = x.importNamesUsage()
	if err != nil {
		return err
	}

	err = x.sfga.Export(out, true)

	return nil
}

func (x *xsv) copy(src, dstDir string) error {
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	dstPath := filepath.Join(dstDir, filepath.Base(src))

	dstFile, err := os.Create(dstPath)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		return err
	}

	x.csvPath = dstPath

	return nil
}

func (x *xsv) initSfga() (sfga.Archive, error) {
	sfga := sfgaio.New()
	err := sfga.Create(x.cfg.ImporterSfgaDir)
	if err != nil {
		return nil, err
	}
	_, err = sfga.Connect()
	if err != nil {
		return nil, err
	}
	return sfga, nil
}
