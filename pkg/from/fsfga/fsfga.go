package fsfga

import (
	"log/slog"

	"github.com/sfborg/sf/config"
	"github.com/sfborg/sf/pkg/from"
	"github.com/sfborg/sf/pkg/sf"
	"github.com/sfborg/sflib"
	"github.com/sfborg/sflib/pkg/arch"
	"github.com/sfborg/sflib/pkg/sfga"
)

type fsfga struct {
	cfg            config.Config
	sfga           sfga.Archive
	sfgaCurrentVer sfga.Archive
	*from.Shared
}

func New(cfg config.Config) sf.FromConvertor {
	res := fsfga{
		cfg:    cfg,
		sfga:   sflib.NewSfga(),
		Shared: from.New(cfg),
	}
	return &res
}

func (fs *fsfga) Import(src, dst string) error {
	var err error

	slog.Info("Creating empty SFGA of the current version")
	fs.sfgaCurrentVer, err = fs.InitSfga()
	if err != nil {
		return err
	}

	slog.Info("Getting SFGA archive")
	err = fs.sfga.Fetch(src, fs.cfg.ImportDir)
	if err != nil {
		return &arch.ErrExtract{Path: src, Err: err}
	}

	_, err = fs.sfga.Connect()
	if err != nil {
		return err
	}

	slog.Info("Transferring data from to current SFGA version")
	err = fs.sfga.Update(fs.sfgaCurrentVer, fs.cfg.WithParents)
	if err != nil {
		return err
	}
	withZip := fs.cfg.WithZipOutput && !fs.cfg.WithParents

	err = fs.sfgaCurrentVer.Export(dst, withZip)
	if err != nil {
		return err
	}
	if !fs.cfg.WithParents {
		return nil
	}

	slog.Info("Creating parent/child hierarchy")

	slog.Info("Creating empty SFGA for parent/child relationships")
	fs.sfgaCurrentVer, err = fs.InitSfga()
	if err != nil {
		return err
	}

	return nil
}
