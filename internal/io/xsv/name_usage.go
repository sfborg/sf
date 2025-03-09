package xsv

import (
	"context"
	"fmt"
	"log/slog"
	"strings"
	"sync"

	"github.com/gnames/coldp/ent/coldp"
	"github.com/gnames/gnfmt/gncsv"
	csvConfig "github.com/gnames/gnfmt/gncsv/config"
	"github.com/gnames/gnlib"
	"github.com/gnames/gnparser"
	"golang.org/x/sync/errgroup"
)

func (x *xsv) importNamesUsage() error {
	chIn := make(chan []string)
	chOut := make(chan coldp.NameUsage)
	var wgProcess sync.WaitGroup

	csv, err := x.getCSV()
	if err != nil {
		return err
	}

	headers := coldp.NormalizeHeaders(csv.Headers())

	g, _ := errgroup.WithContext(context.Background())

	// loader
	g.Go(func() error {
		_, err := csv.Read(context.Background(), chIn)
		close(chIn)
		return err
	})

	for range x.cfg.JobsNum {
		wgProcess.Add(1)
		g.Go(func() error {
			defer wgProcess.Done()
			x.process(headers, chIn, chOut)
			return nil
		})
	}

	g.Go(func() error {
		x.write(chOut)
		return nil
	})

	wgProcess.Wait()
	close(chOut)

	err = g.Wait()
	if err != nil {
		return err
	}
	return err
}

func getVal(headers map[string]int, row []string, field string) string {
	if idx, ok := headers[field]; ok {
		res := row[idx]
		return strings.TrimSpace(res)
	}
	return ""
}

func (x *xsv) process(
	headers map[string]int, // headers is normalized headers
	chIn <-chan []string,
	chOut chan<- coldp.NameUsage,
) {

	var ids = make(map[string]struct{})
	for row := range chIn {
		id := getVal(headers, row, "id")
		if _, ok := ids[id]; ok {
			slog.Error("Duplicate ID", "ID", id)
			continue
		} else {
			ids[id] = struct{}{}
		}

		rowCode := coldp.NewNomCode(getVal(headers, row, "code"))
		code := x.getNomCode(rowCode)

		p := <-x.parserPool[code]

		nu := processRow(p, headers, row)

		x.parserPool[code] <- p
		chOut <- nu
	}
}

func (x *xsv) write(chOut <-chan coldp.NameUsage) {
	ch := gnlib.ChunkChannel(chOut, x.cfg.BatchSize)

	var count int
	for chunk := range ch {
		count++
		fmt.Printf("Processed %d rows\n", x.cfg.BatchSize*count)
		x.sfga.InsertNameUsages(chunk)
	}
}

func pick(a, b string) string {
	if a != "" {
		return a
	}
	return b
}

func (x *xsv) getCSV() (gncsv.GnCSV, error) {
	opts := []csvConfig.Option{
		csvConfig.OptPath(x.csvPath),
		csvConfig.OptBadRowMode(x.cfg.BadRow),
		csvConfig.OptWithQuotes(!x.cfg.WithoutQuotes),
	}
	cfg, err := csvConfig.New(opts...)
	if err != nil {
		return nil, err
	}

	return gncsv.New(cfg), nil
}

func (x *xsv) getNomCode(rowCode coldp.NomCode) string {
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

func processRow(
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
		res.CanonicalSimple = prsd.CanonicalSimple
		res.CanonicalFull = prsd.CanonicalFull
		res.CanonicalStemmed = prsd.CanonicalStemmed
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
		res.BasionymAuthorship = pick(res.BasionymAuthorship, prsd.BasionymAuthorship)
		res.BasionymExAuthorship = pick(res.BasionymExAuthorship, prsd.BasionymExAuthorship)
		res.BasionymAuthorshipYear = pick(res.BasionymAuthorshipYear, prsd.BasionymAuthorshipYear)
	}
	return res
}
