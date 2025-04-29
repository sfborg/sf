package fdwca

import (
	"context"
	"errors"
	"log/slog"

	"github.com/sfborg/sflib/pkg/coldp"
	"github.com/sfborg/sflib/pkg/dwca"
	"golang.org/x/sync/errgroup"
)

// importVernacular imports data from DwCA vernacular extension.
func (fd *fdwca) importVernacular(idx int, ext *dwca.Extension) error {
	var err error
	slog.Info("Importing data from DwCA vernacular extension")
	chOut := make(chan []coldp.Vernacular)
	g, ctx := errgroup.WithContext(context.Background())

	// loader
	g.Go(func() error {
		defer close(chOut)
		err := fd.dwca.LoadVernacular(ctx, idx, ext, chOut)
		return err
	})

	g.Go(func() error {
		return fd.writeVernData(ctx, chOut)
	})

	err = g.Wait()
	if err != nil && !errors.Is(err, context.Canceled) {
		return err
	}

	return nil
}

func (fd *fdwca) writeVernData(
	ctx context.Context,
	chOut <-chan []coldp.Vernacular,
) error {
	var err error
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case v, ok := <-chOut:
			if !ok {
				return nil
			}
			err = fd.sfga.InsertVernaculars(v)
			if err != nil {
				for range chOut {
				}
				return err
			}
		}
	}
}
