package tdwca

import (
	"context"
	"errors"
	"os"
	"path/filepath"

	"github.com/sfborg/sflib/pkg/coldp"
	"github.com/sfborg/sflib/pkg/dwca"
	"golang.org/x/sync/errgroup"
)

func (t *tdwca) convertMeta() error {
	meta, err := t.sfga.LoadMeta()
	if err != nil {
		return err
	}

	err = t.dwca.WriteEML(meta)
	if err != nil {
		return err
	}

	return nil
}

func (t *tdwca) convertNameUsage() error {
	var err error
	ch := make(chan coldp.NameUsage)
	gr, ctx := errgroup.WithContext(context.Background())

	gr.Go(func() error {
		return t.dwca.WriteCore(ctx, ch)
	})

	err = t.sfga.LoadNameUsages(ctx, ch)
	if err != nil {
		return err
	}
	close(ch)

	if err = gr.Wait(); err != nil && !errors.Is(err, context.Canceled) {
		return err
	}
	return nil
}

func (t *tdwca) convertVernacular() error {
	var err error
	ch := make(chan coldp.Vernacular)
	gr, ctx := errgroup.WithContext(context.Background())

	gr.Go(func() error {
		return t.dwca.WriteVernaculars(ctx, ch)
	})

	err = t.sfga.LoadVernaculars(ctx, ch)
	if err != nil {
		return err
	}
	close(ch)

	if err = gr.Wait(); err != nil && !errors.Is(err, context.Canceled) {
		return err
	}
	return nil
}

func (t *tdwca) convertDistribution() error {
	var err error
	ch := make(chan coldp.Distribution)
	gr, ctx := errgroup.WithContext(context.Background())

	gr.Go(func() error {
		return t.dwca.WriteDistributions(ctx, ch)
	})

	err = t.sfga.LoadDistributions(ctx, ch)
	if err != nil {
		return err
	}
	close(ch)

	if err = gr.Wait(); err != nil && !errors.Is(err, context.Canceled) {
		return err
	}
	return nil
}

func (t *tdwca) genMetaDwCA() error {
	vernPath := filepath.Join(t.cfg.OutputDir, "VernacularName.csv")
	hasVern := fileExists(vernPath)

	distrPath := filepath.Join(t.cfg.OutputDir, "Distribution.csv")
	hasDistr := fileExists(distrPath)

	meta := dwca.BuildMeta(hasVern, hasDistr)
	return t.dwca.WriteMeta(meta)
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}
