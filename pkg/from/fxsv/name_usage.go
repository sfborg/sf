package fxsv

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"strings"

	"github.com/dustin/go-humanize"
	"github.com/gnames/gnlib"
	"github.com/sfborg/sflib/pkg/coldp"
	"golang.org/x/sync/errgroup"
)

func (fx *fxsv) importNamesUsage() error {
	var err error
	file := filepath.Base(fx.xsv.FilePath())
	slog.Info("Importing names from file", "file", file)
	ch := make(chan coldp.NameUsage)

	g, ctx := errgroup.WithContext(context.Background())

	// loader
	g.Go(func() error {
		defer close(ch)
		err := fx.xsv.Load(ctx, ch, fx.cfg.JobsNum, fx.cfg.NomCode)
		return err
	})

	g.Go(func() error {
		return fx.write(ctx, ch)
	})

	err = g.Wait()
	if err != nil && !errors.Is(err, context.Canceled) {
		return err
	}

	return nil
}

func (fx *fxsv) write(ctx context.Context, chOut <-chan coldp.NameUsage) error {
	var err error
	ch := gnlib.ChunkChannel(ctx, chOut, fx.cfg.BatchSize)

	var rows int
	ids := make(map[string]struct{})

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case chunk, ok := <-ch:
			if !ok {
				rows += len(chunk)
				fmt.Fprintf(os.Stderr, "\r%s", strings.Repeat(" ", 50))
				fmt.Fprintf(os.Stderr, "\rProcessed %s rows\n", humanize.Comma(int64(rows)))
				return nil
			}
			dedup := make([]coldp.NameUsage, 0, len(chunk))
			rows += len(chunk)
			fmt.Fprintf(os.Stderr, "\r%s", strings.Repeat(" ", 50))
			fmt.Fprintf(os.Stderr, "\rProcessed %s rows", humanize.Comma(int64(rows)))
			for _, v := range chunk {
				if _, ok := ids[v.ID]; ok {
					slog.Error("Duplicate ID, skipping", "id", v.ID)
				} else {
					ids[v.ID] = struct{}{}
					dedup = append(dedup, v)
				}
			}
			err = fx.sfga.InsertNameUsages(chunk)
			if err != nil {
				return err
			}
		}
	}
}
