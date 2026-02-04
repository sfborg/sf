package txsv

import (
	"context"
	"errors"
	"path/filepath"

	"github.com/sfborg/sflib/pkg/coldp"
	"golang.org/x/sync/errgroup"
)

func (t *txsv) convertNameUsages() error {
	var err error
	ch := make(chan coldp.NameUsage)
	gr, ctx := errgroup.WithContext(context.Background())
	namePath := filepath.Join(t.cfg.OutputDir, "name_usage.csv")

	gr.Go(func() error {
		err = t.xsv.Write(ctx, ch, namePath)
		return err
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
