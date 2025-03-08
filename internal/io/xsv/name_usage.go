package xsv

import (
	"context"
	"log/slog"
	"strings"
	"sync"

	"github.com/gnames/coldp/ent/coldp"
	"github.com/gnames/gnfmt/gncsv"
	csvConfig "github.com/gnames/gnfmt/gncsv/config"
	"github.com/gnames/gnlib"
	"github.com/gnames/gnparser"
)

func (x *xsv) importNamesUsage() error {
	chIn := make(chan []string)
	chOut := make(chan coldp.NameUsage)
	var wg, wg2 sync.WaitGroup
	wg.Add(1)
	wg2.Add(1)

	opts := []csvConfig.Option{
		csvConfig.OptPath(x.csvPath),
		csvConfig.OptBadRowMode(x.cfg.BadRow),
		csvConfig.OptWithQuotes(!x.cfg.WithoutQuotes),
	}

	// create new config with required options.
	cfg, err := csvConfig.New(opts...)
	if err != nil {
		return err
	}
	csv := gncsv.New(cfg)

	headers := coldp.NormalizeHeaders(cfg.Headers)

	go x.reader(headers, chIn, chOut, &wg)
	go x.writer(chOut, &wg2)

	_, err = csv.Read(context.Background(), chIn)
	if err != nil {
		return err
	}
	close(chIn)

	wg.Wait()
	close(chOut)

	wg2.Wait()
	return nil
}

func getVal(headers map[string]int, row []string, field string) string {
	if idx, ok := headers[field]; ok {
		res := row[idx]
		return strings.TrimSpace(res)
	}
	return ""
}

func (x *xsv) reader(
	headers map[string]int, // headers is normalized headers
	chIn <-chan []string,
	chOut chan<- coldp.NameUsage,
	wg *sync.WaitGroup,
) {
	p := gnparser.New(gnparser.NewConfig(gnparser.OptWithDetails(true)))

	defer wg.Done()
	var ids = make(map[string]struct{})
	for row := range chIn {
		id := getVal(headers, row, "id")
		if _, ok := ids[id]; ok {
			slog.Error("Duplicate ID", "ID", id)
			continue
		} else {
			ids[id] = struct{}{}
		}

		nu := coldp.NameUsage{
			ID:                id,
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
			NameStatus:                coldp.NewNomStatus(getVal(headers, row, "namestatus")),
			AccordingToID:             getVal(headers, row, "accordingtoid"),
			AccordingToPage:           getVal(headers, row, "accordingtopage"),
			AccordingToPageLink:       getVal(headers, row, "accordingtopagelink"),
			ReferenceID:               getVal(headers, row, "referenceid"),
			Scrutinizer:               getVal(headers, row, "scrutinizer"),
			ScrutinizerID:             getVal(headers, row, "scrutinizerid"),
			ScrutinizerDate:           getVal(headers, row, "scrutinizerdate"),
			Species:                   getVal(headers, row, "species"),
			Section:                   getVal(headers, row, "section"),
			Subgenus:                  getVal(headers, row, "subgenus"),
			Genus:                     getVal(headers, row, "genus"),
			Subtribe:                  getVal(headers, row, "subtribe"),
			Tribe:                     getVal(headers, row, "tribe"),
			Subfamily:                 getVal(headers, row, "subfamily"),
			Family:                    getVal(headers, row, "family"),
			Superfamily:               getVal(headers, row, "superfamily"),
			Suborder:                  getVal(headers, row, "suborder"),
			Order:                     getVal(headers, row, "order"),
			Subclass:                  getVal(headers, row, "subclass"),
			Class:                     getVal(headers, row, "class"),
			Subphylum:                 getVal(headers, row, "subphylum"),
			Phylum:                    getVal(headers, row, "phylum"),
			Kingdom:                   getVal(headers, row, "kingdom"),
			Link:                      getVal(headers, row, "link"),
			NameRemarks:               getVal(headers, row, "nameremarks"),
			Remarks:                   getVal(headers, row, "remarks"),
			Modified:                  getVal(headers, row, "modified"),
			ModifiedBy:                getVal(headers, row, "modifiedby"),
		}

		nu.ScientificNameString = nu.ScientificName

		if nu.Authorship != "" &&
			!strings.HasSuffix(nu.ScientificName, nu.Authorship) {
			nu.ScientificNameString += " " + nu.Authorship
		}
		prsd := p.ParseName(nu.ScientificNameString).Flatten()
		if prsd.Parsed {
			nu.CanonicalSimple = prsd.CanonicalSimple
			nu.CanonicalFull = prsd.CanonicalFull
			nu.CanonicalStemmed = prsd.CanonicalStemmed
			nu.Rank = coldp.NewRank(pick(nu.Rank.String(), prsd.Rank))
			nu.Uninomial = pick(nu.Uninomial, prsd.Uninomial)
			nu.GenericName = pick(nu.GenericName, prsd.Genus)
			nu.InfragenericEpithet = pick(nu.InfragenericEpithet, prsd.Subgenus)
			nu.SpecificEpithet = pick(nu.SpecificEpithet, prsd.Species)
			nu.InfraspecificEpithet = pick(nu.InfraspecificEpithet, prsd.Infraspecies)
			nu.CultivarEpithet = pick(nu.CultivarEpithet, prsd.CultivarEpithet)
			nu.CombinationAuthorship = pick(nu.CombinationAuthorship, prsd.CombinationAuthorship)
			nu.CombinationExAuthorship = pick(
				nu.CombinationExAuthorship,
				prsd.CombinationExAuthorship,
			)
			nu.CombinationAuthorshipYear = pick(
				nu.CombinationAuthorshipYear,
				prsd.CombinationAuthorshipYear,
			)
			nu.BasionymAuthorship = pick(nu.BasionymAuthorship, prsd.BasionymAuthorship)
			nu.BasionymExAuthorship = pick(nu.BasionymExAuthorship, prsd.BasionymExAuthorship)
			nu.BasionymAuthorshipYear = pick(nu.BasionymAuthorshipYear, prsd.BasionymAuthorshipYear)
		}

		chOut <- nu
	}
}

func (x *xsv) writer(chOut <-chan coldp.NameUsage, wg *sync.WaitGroup) {
	defer wg.Done()
	ch := gnlib.ChunkChannel(chOut, x.cfg.BatchSize)
	for chunk := range ch {
		x.sfga.InsertNameUsages(chunk)
	}
}

func pick(a, b string) string {
	if a != "" {
		return a
	}
	return b
}
