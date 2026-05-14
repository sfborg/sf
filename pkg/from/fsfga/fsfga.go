package fsfga

import (
	"context"
	"log/slog"

	"github.com/sfborg/sf/config"
	"github.com/sfborg/sf/pkg/from"
	"github.com/sfborg/sf/pkg/sf"
	"github.com/sfborg/sflib"
	"github.com/sfborg/sflib/pkg/arch"
	"github.com/sfborg/sflib/pkg/sfga"
)

type fsfga struct {
	cfg  config.Config
	sfga sfga.Archive
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
	slog.Info("Getting SFGA archive")
	err := fs.sfga.Fetch(src, fs.cfg.ImportDir)
	if err != nil {
		return &arch.ErrExtract{Path: src, Err: err}
	}

	_, err = fs.sfga.Connect()
	if err != nil {
		return err
	}

	slog.Info("Migrating SFGA to current version")
	migrated, err := fs.sfga.Migrate(fs.cfg.OutputDir)
	if err != nil {
		return err
	}

	if fs.cfg.WithParents {
		slog.Info("Building parent/child hierarchy")
		if err = migrated.AddParents(context.Background()); err != nil {
			return err
		}
	}

	withZip := fs.cfg.WithZipOutput && !fs.cfg.WithParents
	return migrated.Export(dst, withZip)
}
