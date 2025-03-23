package fxsv

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"os"
	"strings"
	"sync"

	"github.com/dustin/go-humanize"
	"github.com/gnames/gnfmt/gncsv"
	csvConfig "github.com/gnames/gnfmt/gncsv/config"
	"github.com/gnames/gnlib"
	"github.com/gnames/gnparser"
	"github.com/sfborg/sflib/pkg/coldp"
	"golang.org/x/sync/errgroup"
)

func (x *fxsv) importNamesUsage() error {
	slog.Info("Importing data from file", "file", x.csvPath)
	chIn := make(chan []string)
	chOut := make(chan coldp.NameUsage)
	var wgProcess sync.WaitGroup

	csv, err := x.getCSV()
	if err != nil {
		return err
	}

	headers := coldp.NormalizeHeaders(csv.Headers())

	g, ctx := errgroup.WithContext(context.Background())

	// loader
	g.Go(func() error {
		_, err := csv.Read(ctx, chIn)
		close(chIn)
		return err
	})

	for range x.cfg.JobsNum {
		wgProcess.Add(1)
		g.Go(func() error {
			defer wgProcess.Done()
			return x.process(ctx, headers, chIn, chOut)
		})
	}

	g.Go(func() error {
		return x.write(ctx, chOut)
	})

	wgProcess.Wait()
	close(chOut)

	err = g.Wait()
	if err != nil && !errors.Is(err, context.Canceled) {
		return err
	}

	return nil
}

func getVal(headers map[string]int, row []string, field string) string {
	if idx, ok := headers[field]; ok {
		res := row[idx]
		return strings.TrimSpace(res)
	}
	return ""
}

func (x *fxsv) process(
	ctx context.Context,
	headers map[string]int, // headers is normalized headers
	chIn <-chan []string,
	chOut chan<- coldp.NameUsage,
) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case row, ok := <-chIn:
			if !ok {
				return nil
			}
			x.processRow(row, headers, chOut)
		}
	}
}

func (x *fxsv) processRow(
	row []string,
	headers map[string]int,
	chOut chan<- coldp.NameUsage,
) {
	rowCode := coldp.NewNomCode(getVal(headers, row, "code"))
	code := x.getNomCode(rowCode)

	p := <-x.parserPool[code]

	nu := getNameUsage(p, headers, row)
	x.parserPool[code] <- p

	chOut <- nu
}

func (x *fxsv) write(ctx context.Context, chOut <-chan coldp.NameUsage) error {
	var err error
	ch := gnlib.ChunkChannel(chOut, x.cfg.BatchSize)

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
			err = x.sfga.InsertNameUsages(chunk)
			if err != nil {
				return err
			}
		}
	}
}

func pick(a, b string) string {
	if a != "" {
		return a
	}
	return b
}

func (x *fxsv) getCSV() (gncsv.GnCSV, error) {
	opts := []csvConfig.Option{
		csvConfig.OptPath(x.csvPath),
		csvConfig.OptBadRowMode(x.cfg.BadRow),
	}
	cfg, err := csvConfig.New(opts...)
	if err != nil {
		return nil, err
	}

	return gncsv.New(cfg), nil
}

func (x *fxsv) getNomCode(rowCode coldp.NomCode) string {
	res := "any"
	codeVal := x.cfg.NomCode
	switch codeVal {
	case coldp.Botanical, coldp.Cultivars:
		res = "botanical"
	}
	switch rowCode {
	case coldp.Botanical, coldp.Cultivars:
		res = "botanical"
	}
	return res
}

func getNameUsage(
	p gnparser.GNparser,
	headers map[string]int,
	row []string,
) coldp.NameUsage {
	res := coldp.NameUsage{
		ID:                getVal(headers, row, "id"),
		AlternativeID:     getVal(headers, row, "alternativeid"),
		NameAlternativeID: getVal(headers, row, "namealternativeid"),
		LocalID:           getVal(headers, row, "localid"),
		GlobalID:          getVal(headers, row, "globalid"),
		SourceID:          getVal(headers, row, "sourceid"),
		ParentID:          getVal(headers, row, "parentid"),
		BasionymID:        getVal(headers, row, "basyonymid"),
		TaxonomicStatus: coldp.NewTaxonomicStatus(
			getVal(headers, row, "taxonomicstatus"),
		),
		ScientificName:            getVal(headers, row, "scientificname"),
		Authorship:                getVal(headers, row, "authorship"),
		ScientificNameString:      getVal(headers, row, "scientificnamestring"),
		Rank:                      coldp.NewRank(getVal(headers, row, "rank")),
		Notho:                     coldp.NewNamePart(getVal(headers, row, "notho")),
		Uninomial:                 getVal(headers, row, "uninomial"),
		GenericName:               getVal(headers, row, "genericname"),
		InfragenericEpithet:       getVal(headers, row, "infragenericepithet"),
		SpecificEpithet:           getVal(headers, row, "specificepithet"),
		InfraspecificEpithet:      getVal(headers, row, "infraspecificepithet"),
		CultivarEpithet:           getVal(headers, row, "cultivarepithet"),
		CombinationAuthorship:     getVal(headers, row, "combinationauthorship"),
		CombinationAuthorshipID:   getVal(headers, row, "combinationauthorshipid"),
		CombinationExAuthorship:   getVal(headers, row, "combinationexauthorship"),
		CombinationExAuthorshipID: getVal(headers, row, "combinationexauthorshipid"),
		CombinationAuthorshipYear: getVal(headers, row, "combinationauthorshipyear"),
		BasionymAuthorship:        getVal(headers, row, "basionymauthorship"),
		BasionymAuthorshipID:      getVal(headers, row, "basionymauthorshipid"),
		BasionymExAuthorship:      getVal(headers, row, "basionymexauthorship"),
		BasionymExAuthorshipID:    getVal(headers, row, "basionymexauthorshipid"),
		BasionymAuthorshipYear:    getVal(headers, row, "basionymauthorshipyear"),
		NamePhrase:                getVal(headers, row, "namephrase"),
		NameReferenceID:           getVal(headers, row, "namereferenceid"),
		PublishedInYear:           getVal(headers, row, "publishedinyear"),
		PublishedInPage:           getVal(headers, row, "publishedinpage"),
		PublishedInPageLink:       getVal(headers, row, "publishedinpagelink"),
		Gender:                    coldp.NewGender(getVal(headers, row, "gender")),
		Etymology:                 getVal(headers, row, "etymology"),
		Code:                      coldp.NewNomCode(getVal(headers, row, "code")),
		NameStatus: coldp.NewNomStatus(
			getVal(headers, row, "namestatus"),
		),
		AccordingToID:       getVal(headers, row, "accordingtoid"),
		AccordingToPage:     getVal(headers, row, "accordingtopage"),
		AccordingToPageLink: getVal(headers, row, "accordingtopagelink"),
		ReferenceID:         getVal(headers, row, "referenceid"),
		Scrutinizer:         getVal(headers, row, "scrutinizer"),
		ScrutinizerID:       getVal(headers, row, "scrutinizerid"),
		ScrutinizerDate:     getVal(headers, row, "scrutinizerdate"),
		Species:             getVal(headers, row, "species"),
		Section:             getVal(headers, row, "section"),
		Subgenus:            getVal(headers, row, "subgenus"),
		Genus:               getVal(headers, row, "genus"),
		Subtribe:            getVal(headers, row, "subtribe"),
		Tribe:               getVal(headers, row, "tribe"),
		Subfamily:           getVal(headers, row, "subfamily"),
		Family:              getVal(headers, row, "family"),
		Superfamily:         getVal(headers, row, "superfamily"),
		Suborder:            getVal(headers, row, "suborder"),
		Order:               getVal(headers, row, "order"),
		Subclass:            getVal(headers, row, "subclass"),
		Class:               getVal(headers, row, "class"),
		Subphylum:           getVal(headers, row, "subphylum"),
		Phylum:              getVal(headers, row, "phylum"),
		Kingdom:             getVal(headers, row, "kingdom"),
		Link:                getVal(headers, row, "link"),
		NameRemarks:         getVal(headers, row, "nameremarks"),
		Remarks:             getVal(headers, row, "remarks"),
		Modified:            getVal(headers, row, "modified"),
		ModifiedBy:          getVal(headers, row, "modifiedby"),
	}

	// add parsed data
	res.ScientificNameString = res.ScientificName

	if res.Authorship != "" &&
		!strings.HasSuffix(res.ScientificName, res.Authorship) {
		res.ScientificNameString += " " + res.Authorship
	}
	prsd := p.ParseName(res.ScientificNameString).Flatten()
	if prsd.Parsed {
		res.ParseQuality = coldp.ToInt(prsd.ParseQuality)
		res.CanonicalSimple = prsd.CanonicalSimple
		res.CanonicalFull = prsd.CanonicalFull
		res.CanonicalStemmed = prsd.CanonicalStemmed
		res.Cardinality = coldp.ToInt(prsd.Cardinality)
		res.Virus = coldp.ToBool(prsd.Virus)
		res.Hybrid = prsd.Hybrid
		res.Surrogate = prsd.Surrogate
		res.Authors = prsd.Authors
		res.GnID = prsd.VerbatimID

		res.Authorship = pick(res.Authorship, prsd.Authorship)
		res.Rank = coldp.NewRank(pick(res.Rank.String(), prsd.Rank))
		res.Uninomial = pick(res.Uninomial, prsd.Uninomial)
		res.GenericName = pick(res.GenericName, prsd.Genus)
		res.InfragenericEpithet = pick(res.InfragenericEpithet, prsd.Subgenus)
		res.SpecificEpithet = pick(res.SpecificEpithet, prsd.Species)
		res.InfraspecificEpithet = pick(
			res.InfraspecificEpithet,
			prsd.Infraspecies,
		)
		res.CultivarEpithet = pick(res.CultivarEpithet, prsd.CultivarEpithet)

		res.CombinationAuthorship = pick(
			res.CombinationAuthorship,
			prsd.CombinationAuthorship,
		)
		res.CombinationExAuthorship = pick(
			res.CombinationExAuthorship,
			prsd.CombinationExAuthorship,
		)
		res.CombinationAuthorshipYear = pick(
			res.CombinationAuthorshipYear,
			prsd.CombinationAuthorshipYear,
		)

		res.BasionymAuthorship = pick(
			res.BasionymAuthorship,
			prsd.BasionymAuthorship,
		)
		res.BasionymExAuthorship = pick(
			res.BasionymExAuthorship,
			prsd.BasionymExAuthorship,
		)
		res.BasionymAuthorshipYear = pick(
			res.BasionymAuthorshipYear,
			prsd.BasionymAuthorshipYear,
		)
	}
	return res
}
