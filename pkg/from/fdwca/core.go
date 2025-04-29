package fdwca

import (
	"context"
	"path/filepath"
	"strings"
	"sync"

	"github.com/gnames/gnuuid"
	"github.com/sfborg/sflib/pkg/coldp"
	"github.com/sfborg/sflib/pkg/dwca"
	"golang.org/x/sync/errgroup"
)

var refs []coldp.Reference

func (fd *fdwca) importCore() error {
	chIn := make(chan []string)
	chOut := make(chan []coldp.NameUsage)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	g, ctx := errgroup.WithContext(ctx)
	var wg sync.WaitGroup
	wg.Add(1)

	g.Go(func() error {
		defer wg.Done()
		return fd.coreWorker(ctx, chIn, chOut)
	})

	g.Go(func() error {
		return fd.writeCoreData(ctx, chOut)
	})

	// close chOut when all workers are done
	go func() {
		wg.Wait()
		close(chOut)
	}()

	_, err := fd.dwca.CoreStream(ctx, chIn)
	if err != nil {
		return err
	}

	if err := g.Wait(); err != nil {
		return err
	}
	return nil
}

func (fd *fdwca) coreWorker(
	ctx context.Context,
	chIn chan []string,
	chOut chan []coldp.NameUsage,
) error {
	fieldsMap := fieldsMap(fd.dwca.Meta().Core.Fields)
	coreID := fd.dwca.Meta().Core.ID.Idx

	batch := make([]coldp.NameUsage, 0, fd.cfg.BatchSize)
	for v := range chIn {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			row := fd.processCoreRow(v, coreID, fieldsMap)
			if len(batch) == fd.cfg.BatchSize {
				chOut <- batch
				batch = make([]coldp.NameUsage, 0, fd.cfg.BatchSize)
			}
			batch = append(batch, row)
		}
	}
	chOut <- batch
	return nil
}

func (f *fdwca) writeCoreData(
	ctx context.Context,
	chOut chan []coldp.NameUsage,
) error {
	var err error
	for nu := range chOut {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			// write to db
			err = f.sfga.InsertNameUsages(nu)
			if err != nil {
				for range chOut {
				}
				return err
			}
			if len(refs) > 0 {
				err = f.sfga.InsertReferences(refs)
				refs = refs[0:0]
			}
		}
	}
	return nil
}

func fieldsMap(fields []dwca.Field) map[string]int {
	fieldsMap := make(map[string]int)
	for _, v := range fields {
		term := filepath.Base(v.Term)
		term = strings.ToLower(term)
		_, noPrefix, found := strings.Cut(term, ":")
		if found {
			term = noPrefix
		}
		fieldsMap[term] = v.Idx
	}
	return fieldsMap
}

func fieldVal(
	row []string,
	fielsMap map[string]int,
	name string) string {
	if idx, ok := fielsMap[name]; ok {
		if idx >= len(row) {
			return ""
		}
		return row[idx]
	}
	return ""
}

func (fd *fdwca) processCoreRow(
	row []string,
	idIdx int,
	fieldsMap map[string]int,
) coldp.NameUsage {
	res := coldp.NameUsage{ID: row[idIdx]}
	res.SourceID = fieldVal(row, fieldsMap, "datasetid")
	res.ScientificName = fieldVal(row, fieldsMap, "scientificname")
	res.Authorship = fieldVal(row, fieldsMap, "scientificnameauthorship")

	res.ScientificNameString = fieldVal(row, fieldsMap, "scientificnamestring")
	res.LocalID = fieldVal(row, fieldsMap, "localid")
	res.GlobalID = fieldVal(row, fieldsMap, "globalid")

	res.ParentID = fieldVal(row, fieldsMap, "parentnameusageid")
	code := fieldVal(row, fieldsMap, "nomenclaturalcode")
	res.Code = coldp.NewNomCode(code)
	rank := fieldVal(row, fieldsMap, "taxonrank")
	res.Rank = coldp.NewRank(rank)
	res = addTaxonomicStatus(res, row, fieldsMap)

	res.Notho = coldp.NewNamePart(fieldVal(row, fieldsMap, "notho"))
	res.Kingdom = fieldVal(row, fieldsMap, "kingdom")
	res.Phylum = fieldVal(row, fieldsMap, "phylum")
	res.Class = fieldVal(row, fieldsMap, "class")
	res.Order = fieldVal(row, fieldsMap, "order")
	res.Superfamily = fieldVal(row, fieldsMap, "superfamily")
	res.Family = fieldVal(row, fieldsMap, "family")
	res.Subfamily = fieldVal(row, fieldsMap, "subfamily")
	res.Tribe = fieldVal(row, fieldsMap, "tribe")
	res.Genus = fieldVal(row, fieldsMap, "genericName")
	res.InfragenericEpithet = fieldVal(row, fieldsMap, "infragenericepither")
	res.InfraspecificEpithet = fieldVal(row, fieldsMap, "infraspecificepither")
	res.CultivarEpithet = fieldVal(row, fieldsMap, "cultivarepithet")
	res.AccordingToID = fieldVal(row, fieldsMap, "nameaccordingto")
	res.NameStatus = coldp.NewNomStatus(fieldVal(row, fieldsMap, "nomenclaturalstatus"))
	res.Remarks = fieldVal(row, fieldsMap, "taxonremarks")

	citation := fieldVal(row, fieldsMap, "namepublishedin")
	if citation != "" {
		res.ReferenceID = gnuuid.New(citation).String()
		refs = append(refs, coldp.Reference{ID: res.ReferenceID, Citation: citation})
	}

	return res
}

func addTaxonomicStatus(
	nu coldp.NameUsage,
	row []string,
	fieldMap map[string]int,
) coldp.NameUsage {
	nu.TaxonomicStatus = coldp.UnknownTaxSt

	// if neither acceptedNameUsageID or taxonomicStatus are given
	// we cannot assert taxonomic status
	_, hasStatus := fieldMap["taxonomicstatus"]
	_, hasAccepted := fieldMap["acceptednameusageid"]
	if !hasStatus && !hasAccepted {
		return nu
	}

	// use provided addTaxonomicStatus if possible
	ts := fieldVal(row, fieldMap, "taxonomicstatus")
	nu.TaxonomicStatus = coldp.NewTaxonomicStatus(ts)
	if nu.TaxonomicStatus != coldp.UnknownTaxSt {
		return nu
	}

	acceptedNameUsageID := fieldVal(row, fieldMap, "acceptednameusageid")
	if acceptedNameUsageID != "" && nu.ID != acceptedNameUsageID {
		// for NameUsage's synonyms accepted ID goes to parent, and
		// synonymy is expressed by taxonomic status
		nu.ParentID = acceptedNameUsageID
		nu.TaxonomicStatus = coldp.SynonymTS
		return nu
	}

	nu.TaxonomicStatus = coldp.AcceptedTS
	return nu
}
