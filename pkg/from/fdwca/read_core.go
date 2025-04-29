package fdwca

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"os"
	"strings"

	"github.com/dustin/go-humanize"
	"github.com/sfborg/sflib/pkg/coldp"
	"golang.org/x/sync/errgroup"
)

var refs []coldp.Reference

type NuRef struct {
	NameUsages []coldp.NameUsage
	References []coldp.Reference
}

func (fd *fdwca) importCore() error {
	var err error
	slog.Info("Importing data from DwCA core")
	ch := make(chan coldp.Data)
	g, ctx := errgroup.WithContext(context.Background())

	// loader
	g.Go(func() error {
		defer close(ch)
		err := fd.dwca.LoadCore(ctx, ch)
		return err
	})

	g.Go(func() error {
		return fd.writeCoreData(ctx, ch)
	})

	err = g.Wait()
	if err != nil && !errors.Is(err, context.Canceled) {
		return err
	}

	return nil
}

func (fd *fdwca) writeCoreData(
	ctx context.Context,
	ch <-chan coldp.Data,
) error {
	var err error
	var rows int
	ids := make(map[string]struct{})

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case data, ok := <-ch:
			if !ok {
				fmt.Fprintf(os.Stderr, "\r%s", strings.Repeat(" ", 50))
				fmt.Fprintf(os.Stderr, "\rSaved %s core records\n", humanize.Comma(int64(rows)))
				return nil
			}
			if len(data.NameUsages) > 0 {
				dedup := make([]coldp.NameUsage, 0, len(data.NameUsages))
				rows += len(data.NameUsages)
				fmt.Fprintf(os.Stderr, "\r%s", strings.Repeat(" ", 50))
				fmt.Fprintf(os.Stderr, "\rSaved %s records", humanize.Comma(int64(rows)))
				for _, v := range data.NameUsages {
					if _, ok := ids[v.ID]; ok {
						slog.Error("Duplicated TaxonID, skipping", "id", v.ID)
					} else {
						ids[v.ID] = struct{}{}
						dedup = append(dedup, v)
					}
				}
				err = fd.sfga.InsertNameUsages(data.NameUsages)
				if err != nil {
					return err
				}
			}
			if len(data.References) > 0 {
				err = fd.sfga.InsertReferences(data.References)
				if err != nil {
					return err
				}
			}
		}
	}
}
