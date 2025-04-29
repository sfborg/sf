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
func (fd *fdwca) importDistr(idx int, ext *dwca.Extension) error {
	var err error
	slog.Info("Importing data from DwCA distribution extension")
	chOut := make(chan coldp.Data)
	g, ctx := errgroup.WithContext(context.Background())

	// loader
	g.Go(func() error {
		defer close(chOut)
		err := fd.dwca.LoadDistribution(ctx, idx, ext, chOut)
		return err
	})

	g.Go(func() error {
		return fd.writeDistrData(ctx, chOut)
	})

	err = g.Wait()
	if err != nil && !errors.Is(err, context.Canceled) {
		return err
	}

	return nil
}

func (fd *fdwca) writeDistrData(
	ctx context.Context,
	chOut <-chan coldp.Data,
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
			err = fd.sfga.InsertDistributions(v.Distributions)
			if err != nil {
				for range chOut {
				}
				return err
			}
			if len(v.References) > 0 {
				err = fd.sfga.InsertReferences(v.References)
				if err != nil {
					for range chOut {
					}
					return err
				}
			}
		}
	}
}
