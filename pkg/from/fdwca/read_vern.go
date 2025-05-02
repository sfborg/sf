package fdwca

import (
	"context"
	"errors"
	"log/slog"

	"github.com/sfborg/sflib/pkg/coldp"
	"github.com/sfborg/sflib/pkg/dwca"
	"golang.org/x/sync/errgroup"
)

func (fd *fdwca) importVernacular(idx int, ext *dwca.Extension) error {
	var err error
	slog.Info("Importing data from DwCA vernacular extension")
	ch := make(chan []coldp.Vernacular)
	g, ctx := errgroup.WithContext(context.Background())

	// loader
	g.Go(func() error {
		defer close(ch)
		err := fd.dwca.LoadVernacular(ctx, idx, ext, ch)
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

// 	chIn := make(chan []string)
// 	chOut := make(chan []coldp.Vernacular)
//
// 	g, ctx := errgroup.WithContext(context.Background())
//
// 	g.Go(func() error {
// 		defer close(chOut)
// 		return fd.vernWorker(ctx, ext, chIn, chOut)
// 	})
//
// 	g.Go(func() error {
// 		return fd.writeVernData(ctx, chOut)
// 	})
//
// 	_, err := fd.d.ExtensionStream(ctx, idx, chIn)
// 	if err != nil {
// 		return err
// 	}
//
// 	if err := g.Wait(); err != nil {
// 		return err
// 	}
// 	return nil
// }
//
// func (fd *fdwca) vernWorker(
// 	ctx context.Context,
// 	ext *meta.Extension,
// 	chIn <-chan []string,
// 	chOut chan<- []coldp.Vernacular,
// ) error {
// 	fieldsMap := fieldsMap(ext.Fields)
// 	coreID := ext.CoreID.Idx
//
// 	batch := make([]coldp.Vernacular, 0, fd.cfg.BatchSize)
// 	for v := range chIn {
// 		vrn := fd.processVernRow(v, coreID, fieldsMap)
// 		if vrn == nil {
// 			continue
// 		}
// 		if len(batch) == fd.cfg.BatchSize {
// 			chOut <- batch
// 			batch = make([]coldp.Vernacular, 0, fd.cfg.BatchSize)
// 		}
// 		batch = append(batch, *vrn)
//
// 		select {
// 		case <-ctx.Done():
// 			return ctx.Err()
// 		default:
// 		}
// 	}
// 	chOut <- batch
// 	return nil
// }
//
// func (fd *fdwca) processVernRow(
// 	row []string,
// 	coreID int,
// 	fieldsMap map[string]int,
// ) *coldp.Vernacular {
// 	var res coldp.Vernacular
// 	res.TaxonID = row[coreID]
//
// 	res.Name = fieldVal(row, fieldsMap, "vernacularname")
// 	if res.Name == "" {
// 		return nil
// 	}
//
// 	lang := fieldVal(row, fieldsMap, "language")
// 	if len(lang) > 3 {
// 		res.Language = gnlib.LangCode(lang)
// 	} else {
// 		res.Language = lang
// 	}
// 	res.Area = fieldVal(row, fieldsMap, "locality")
// 	res.Country = fieldVal(row, fieldsMap, "countrycode")
// 	return &res
// }
//
// func (fd *fdwca) writeVernData(
// 	ctx context.Context,
// 	chOut <-chan []coldp.Vernacular,
// ) error {
// 	var err error
// 	for d := range chOut {
// 		// write to db
// 		err = fd.s.InsertVernaculars(d)
// 		if err != nil {
// 			for range chOut {
// 			}
// 			return err
// 		}
//
// 		select {
// 		case <-ctx.Done():
// 			return ctx.Err()
// 		default:
// 		}
// 	}
// 	return nil
// }
