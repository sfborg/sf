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

func fieldVal(headers map[string]int, row []string, field string) string {
	if idx, ok := headers[field]; ok {
		res := row[idx]
		return res
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
		id := fieldVal(headers, row, "id")
		if _, ok := ids[id]; ok {
			slog.Error("Duplicate ID", "ID", id)
			continue
		} else {
			ids[id] = struct{}{}
		}
		nu := coldp.NameUsage{
			ID:                id,
			AlternativeID:     fieldVal(headers, row, "alternativeid"),
			NameAlternativeID: fieldVal(headers, row, "namealternativeid"),
			LocalID:           fieldVal(headers, row, "localid"),
			GlobalID:          fieldVal(headers, row, "globalid"),
			SourceID:          fieldVal(headers, row, "sourceid"),
			ParentID:          fieldVal(headers, row, "parentid"),
			BasionymID:        fieldVal(headers, row, "basyonymid"),
			TaxonomicStatus: coldp.NewTaxonomicStatus(
				fieldVal(headers, row, "taxonomicstatus"),
			),
			ScientificName:            fieldVal(headers, row, "scientificname"),
			Authorship:                fieldVal(headers, row, "authorship"),
			ScientificNameString:      fieldVal(headers, row, "scientificnamestring"),
			Rank:                      coldp.NewRank(fieldVal(headers, row, "rank")),
			Notho:                     coldp.NewNamePart(fieldVal(headers, row, "notho")),
			Uninomial:                 fieldVal(headers, row, "uninomial"),
			GenericName:               fieldVal(headers, row, "genericname"),
			InfragenericEpithet:       fieldVal(headers, row, "infragenericepithet"),
			SpecificEpithet:           fieldVal(headers, row, "specificepithet"),
			InfraspecificEpithet:      fieldVal(headers, row, "infraspecificepithet"),
			CultivarEpithet:           fieldVal(headers, row, "cultivarepithet"),
			CombinationAuthorship:     fieldVal(headers, row, "combinationauthorship"),
			CombinationAuthorshipID:   fieldVal(headers, row, "combinationauthorshipid"),
			CombinationExAuthorship:   fieldVal(headers, row, "combinationexauthorship"),
			CombinationExAuthorshipID: fieldVal(headers, row, "combinationexauthorshipid"),
			CombinationAuthorshipYear: fieldVal(headers, row, "combinationauthorshipyear"),
			BasionymAuthorship:        fieldVal(headers, row, "basionymauthorship"),
			BasionymAuthorshipID:      fieldVal(headers, row, "basionymauthorshipid"),
			BasionymExAuthorship:      fieldVal(headers, row, "basionymexauthorship"),
			BasionymExAuthorshipID:    fieldVal(headers, row, "basionymexauthorshipid"),
			BasionymAuthorshipYear:    fieldVal(headers, row, "basionymauthorshipyear"),
			NamePhrase:                fieldVal(headers, row, "namephrase"),
			NameReferenceID:           fieldVal(headers, row, "namereferenceid"),
			PublishedInYear:           fieldVal(headers, row, "publishedinyear"),
			PublishedInPage:           fieldVal(headers, row, "publishedinpage"),
			PublishedInPageLink:       fieldVal(headers, row, "publishedinpagelink"),
			Gender:                    coldp.NewGender(fieldVal(headers, row, "gender")),
			Etymology:                 fieldVal(headers, row, "etymology"),
			Code:                      coldp.NewNomCode(fieldVal(headers, row, "code")),
			NameStatus:                coldp.NewNomStatus(fieldVal(headers, row, "namestatus")),
			AccordingToID:             fieldVal(headers, row, "accordingtoid"),
			AccordingToPage:           fieldVal(headers, row, "accordingtopage"),
			AccordingToPageLink:       fieldVal(headers, row, "accordingtopagelink"),
			ReferenceID:               fieldVal(headers, row, "referenceid"),
			Scrutinizer:               fieldVal(headers, row, "scrutinizer"),
			ScrutinizerID:             fieldVal(headers, row, "scrutinizerid"),
			ScrutinizerDate:           fieldVal(headers, row, "scrutinizerdate"),
			Species:                   fieldVal(headers, row, "species"),
			Section:                   fieldVal(headers, row, "section"),
			Subgenus:                  fieldVal(headers, row, "subgenus"),
			Genus:                     fieldVal(headers, row, "genus"),
			Subtribe:                  fieldVal(headers, row, "subtribe"),
			Tribe:                     fieldVal(headers, row, "tribe"),
			Subfamily:                 fieldVal(headers, row, "subfamily"),
			Family:                    fieldVal(headers, row, "family"),
			Superfamily:               fieldVal(headers, row, "superfamily"),
			Suborder:                  fieldVal(headers, row, "suborder"),
			Order:                     fieldVal(headers, row, "order"),
			Subclass:                  fieldVal(headers, row, "subclass"),
			Class:                     fieldVal(headers, row, "class"),
			Subphylum:                 fieldVal(headers, row, "subphylum"),
			Phylum:                    fieldVal(headers, row, "phylum"),
			Kingdom:                   fieldVal(headers, row, "kingdom"),
			Link:                      fieldVal(headers, row, "link"),
			NameRemarks:               fieldVal(headers, row, "nameremarks"),
			Remarks:                   fieldVal(headers, row, "remarks"),
			Modified:                  fieldVal(headers, row, "modified"),
			ModifiedBy:                fieldVal(headers, row, "modifiedby"),
		}
		nu.ScientificNameString = nu.ScientificName

		if nu.Authorship != "" &&
			!strings.HasSuffix(nu.ScientificName, nu.Authorship) {
			nu.ScientificNameString += " " + nu.Authorship
		}
		pres := p.ParseName(nu.ScientificNameString)
		if pres.Parsed {
			nu.CanonicalSimple = pres.Canonical.Simple
			nu.CanonicalFull = pres.Canonical.Full
			nu.CanonicalStemmed = pres.Canonical.Stemmed
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
