package fsfga

import (
	"log/slog"

	"github.com/sfborg/sf/pkg/config"
	"github.com/sfborg/sf/pkg/from"
	"github.com/sfborg/sf/pkg/sf"
	"github.com/sfborg/sflib"
	"github.com/sfborg/sflib/pkg/arch"
	"github.com/sfborg/sflib/pkg/sfga"
)

type fsfga struct {
	cfg         config.Config
	sfga        sfga.Archive
	sfgaCurrent sfga.Archive
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
	fs.sfgaCurrent, err = fs.InitSfga()
	if err != nil {
		return err
	}

	slog.Info("Downloading outdated SFGA archive")
	err = fs.sfga.Fetch(src, fs.cfg.ImportDir)
	if err != nil {
		return &arch.ErrExtract{Path: src, Err: err}
	}

	_, err = fs.sfga.Connect()
	if err != nil {
		return err
	}

	slog.Info("Transferring data from outdated to current SFGA")
	err = fs.sfga.Update(fs.sfgaCurrent)
	if err != nil {
		return err
	}

	err = fs.sfgaCurrent.Export(dst, fs.cfg.WithZipOutput)
	if err != nil {
		return err
	}

	return nil
}
