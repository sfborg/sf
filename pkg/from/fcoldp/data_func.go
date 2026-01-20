package fcoldp

import (
	"context"
	"path/filepath"

	"github.com/sfborg/sflib/pkg/coldp"
	"github.com/sfborg/sflib/pkg/parser"
	"github.com/sfborg/sflib/pkg/sfga"
	"golang.org/x/sync/errgroup"
)

func (fc *fcoldp) importNameData(path string) error {
	chIn := make(chan coldp.Name)
	chOut := make(chan []coldp.Name)
	var err error

	g, _ := errgroup.WithContext(context.Background())

	g.Go(func() error {
		var err error
		for names := range chOut {
			err = fc.sfga.InsertNames(names)
			if err != nil {
				return err
			}
		}
		return nil
	})

	names := make([]coldp.Name, 0, fc.cfg.BatchSize)
	g.Go(func() error {
		defer func() {
			close(chOut)
		}()
		for n := range chIn {
			if fc.cfg.WithDetails {
				code := parser.ParserCode(fc.cfg.NomCode, n.Code)
				p := <-fc.parserPool[code]
				n.Amend(p)
				fc.parserPool[code] <- p
			}
			names = append(names, n)
			if len(names) >= fc.cfg.BatchSize {
				chOut <- names
				names = make([]coldp.Name, 0, fc.cfg.BatchSize)
			}
		}
		if len(names) > 0 {
			chOut <- names
		}
		return nil
	})

	err = coldp.Read(fc.coldp.Config(), path, chIn)
	if err != nil {
		return err
	}
	close(chIn)

	if err = g.Wait(); err != nil {
		return err
	}

	return nil
}

func (fc *fcoldp) importNameUsageData(path string) error {
	chIn := make(chan coldp.NameUsage)
	chOut := make(chan []coldp.NameUsage)
	var err error

	g, _ := errgroup.WithContext(context.Background())

	g.Go(func() error {
		var err error
		for nus := range chOut {
			err = fc.sfga.InsertNameUsages(nus)
			if err != nil {
				return err
			}
		}
		return nil
	})

	nus := make([]coldp.NameUsage, 0, fc.cfg.BatchSize)
	g.Go(func() error {
		defer func() {
			close(chOut)
		}()
		for n := range chIn {
			if fc.cfg.WithDetails {
				code := parser.ParserCode(fc.cfg.NomCode, n.Code)
				p := <-fc.parserPool[code]
				n.Amend(p)
				fc.parserPool[code] <- p
			}
			nus = append(nus, n)
			if len(nus) >= fc.cfg.BatchSize {
				chOut <- nus
				nus = make([]coldp.NameUsage, 0, fc.cfg.BatchSize)
			}
		}
		if len(nus) > 0 {
			chOut <- nus
		}
		return nil
	})

	err = coldp.Read(fc.coldp.Config(), path, chIn)
	if err != nil {
		return err
	}
	close(chIn)

	if err = g.Wait(); err != nil {
		return err
	}

	return nil
}
func importData[T coldp.DataLoader](
	fc *fcoldp,
	path string,
	c coldp.Archive,
	insertFunc func(sfga.Archive, []T) error) error {
	chIn := make(chan T)
	var err error

	ctx, cancel := context.WithCancel(context.Background())
	g, ctx := errgroup.WithContext(ctx)
	defer cancel()

	g.Go(func() error {
		return insert(fc.sfga, fc.cfg.BatchSize, chIn, insertFunc)
	})

	ext := filepath.Ext(path)

	switch ext {
	case ".bib", ".jsonl", ".json":
		// skip for now
	default:
		err = coldp.Read(c.Config(), path, chIn)
		if err != nil {
			return err
		}
	}

	close(chIn)
	if err = g.Wait(); err != nil {
		return err
	}

	return nil
}

func insert[T coldp.DataLoader](
	s sfga.Archive,
	batchSize int,
	ch <-chan T,
	insertFunc func(sfga.Archive, []T) error,
) error {
	var err error
	names := make([]T, 0, batchSize)
	var count int

	for n := range ch {
		count++
		names = append(names, n)
		if count >= batchSize {
			err = insertFunc(s, names)
			count = 0
			names = names[:0]
			if err != nil {
				return err
			}
		}
	}

	err = insertFunc(s, names[:count])
	if err != nil {
		return err
	}
	return nil
}
