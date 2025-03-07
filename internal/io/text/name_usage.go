package text

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"log/slog"
	"os"
	"strings"
	"sync"

	"github.com/dustin/go-humanize"
	"github.com/gnames/coldp/ent/coldp"
	"github.com/gnames/gnlib"
	"github.com/gnames/gnparser"
	"golang.org/x/sync/errgroup"
)

func (t *text) importNamesUsage() error {
	var err error
	slog.Info("Importing names from file", "file", t.textPath)
	chIn := make(chan string)
	chOut := make(chan coldp.NameUsage)
	var wgProcess sync.WaitGroup

	g, ctx := errgroup.WithContext(context.Background())

	// loader
	g.Go(func() error {
		err := t.read(ctx, chIn, t.textPath)
		close(chIn)
		return err
	})

	for range t.cfg.JobsNum {
		wgProcess.Add(1)
		g.Go(func() error {
			defer wgProcess.Done()
			return t.process(ctx, chIn, chOut)
		})
	}

	g.Go(func() error {
		return t.write(ctx, chOut)
	})

	wgProcess.Wait()
	close(chOut)

	err = g.Wait()
	if err != nil && !errors.Is(err, context.Canceled) {
		return err
	}

	return nil
}

func (t *text) read(ctx context.Context, chIn chan<- string, path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case chIn <- scanner.Text():
		}
	}

	err = scanner.Err()
	if err != nil {
		return err
	}
	return nil
}

func (t *text) process(
	ctx context.Context,
	chIn <-chan string,
	chOut chan<- coldp.NameUsage,
) error {
	code := t.getNomCode()
	p := <-t.parserPool[code]
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case line, ok := <-chIn:
			if !ok {
				t.parserPool[code] <- p
				return nil
			}
			chOut <- t.processLine(p, line)
		}
	}
}

func (t *text) write(ctx context.Context, chOut <-chan coldp.NameUsage) error {
	var err error
	ch := gnlib.ChunkChannel(chOut, t.cfg.BatchSize)

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
			err = t.sfga.InsertNameUsages(chunk)
			if err != nil {
				return err
			}
		}
	}
}

func (t *text) getNomCode() string {
	res := "any"
	codeVal := t.cfg.NomCode
	switch codeVal {
	case coldp.Botanical, coldp.Cultivars:
		res = "botanical"
	}
	return res
}

func (t *text) processLine(p gnparser.GNparser, line string) coldp.NameUsage {
	prsd := p.ParseName(line).Flatten()
	res := coldp.NameUsage{
		ID:                        prsd.VerbatimID,
		GlobalID:                  prsd.VerbatimID,
		ScientificName:            line,
		Authorship:                prsd.Authorship,
		ScientificNameString:      line,
		ParseQuality:              coldp.ToInt(prsd.ParseQuality),
		CanonicalSimple:           prsd.CanonicalSimple,
		CanonicalFull:             prsd.CanonicalFull,
		CanonicalStemmed:          prsd.CanonicalStemmed,
		Cardinality:               coldp.ToInt(prsd.Cardinality),
		Virus:                     coldp.ToBool(prsd.Virus),
		Hybrid:                    prsd.Hybrid,
		Surrogate:                 prsd.Surrogate,
		Authors:                   prsd.Authors,
		GnID:                      prsd.VerbatimID,
		Rank:                      coldp.NewRank(prsd.Rank),
		Uninomial:                 prsd.Uninomial,
		GenericName:               prsd.Genus,
		InfragenericEpithet:       prsd.Subgenus,
		SpecificEpithet:           prsd.Species,
		InfraspecificEpithet:      prsd.Infraspecies,
		CultivarEpithet:           prsd.CultivarEpithet,
		CombinationAuthorship:     prsd.CombinationAuthorship,
		CombinationExAuthorship:   prsd.CombinationExAuthorship,
		CombinationAuthorshipYear: prsd.CombinationAuthorshipYear,
		BasionymAuthorship:        prsd.BasionymAuthorship,
		BasionymExAuthorship:      prsd.BasionymExAuthorship,
		BasionymAuthorshipYear:    prsd.BasionymAuthorshipYear,
	}

	return res
}
