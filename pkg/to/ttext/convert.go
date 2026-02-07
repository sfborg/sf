package ttext

import (
	"context"
	"errors"

	"github.com/sfborg/sflib/pkg/coldp"
	"golang.org/x/sync/errgroup"
)

func (t *ttext) convertNameUsages() error {
	var err error
	ch := make(chan coldp.NameUsage)
	gr, ctx := errgroup.WithContext(context.Background())

	gr.Go(func() error {
		err = t.txt.Write(ctx, ch)
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
