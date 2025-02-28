package diffio

import (
	"log/slog"

	"github.com/sfborg/sflib/ent/sfga"
	"github.com/sfborg/sflib/io/sfgaio"
)

func (d *diffio) Compare(src, trg string) error {
	var err error
	slog.Info("Getting reference file", "path", src)
	d.src, err = d.initSfga(src, d.cfg.DiffSrcDir)
	if err != nil {
		return err
	}
	slog.Info("Getting comparison file", "path", trg)
	d.trg, err = d.initSfga(trg, d.cfg.DiffTrgDir)
	if err != nil {
		return err
	}

	slog.Info("Optimizing reference for comparison")

	defer d.src.Close()
	defer d.trg.Close()

	return nil
}

func (d *diffio) initSfga(path, dir string) (sfga.Archive, error) {
	sfga := sfgaio.New()
	err := sfga.Import(path, dir)
	if err != nil {
		return nil, err
	}
	_, err = sfga.Connect()
	if err != nil {
		return nil, err
	}
	return sfga, nil
}
