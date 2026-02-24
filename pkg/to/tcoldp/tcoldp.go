package tcoldp

import (
	"log/slog"

	"github.com/sfborg/sf/config"
	"github.com/sfborg/sf/pkg/sf"
	"github.com/sfborg/sflib"
	"github.com/sfborg/sflib/pkg/coldp"
	"github.com/sfborg/sflib/pkg/sfga"
)

type tcoldp struct {
	cfg   config.Config
	sfga  sfga.Archive
	coldp coldp.Archive
}

func New(cfg config.Config) sf.ToConvertor {
	res := tcoldp{
		cfg:   cfg,
		sfga:  sflib.NewSfga(),
		coldp: sflib.NewColdp(),
	}
	return &res
}

func (t *tcoldp) Export(src, dst string) error {
	var err error
	if err = t.coldp.Create(t.cfg.OutputDir); err != nil {
		return err
	}

	err = t.sfga.Fetch(src, t.cfg.ImportDir)
	if err != nil {
		return err
	}
	_, err = t.sfga.Connect()
	if err != nil {
		return err
	}

	slog.Info("Exporting Metadata file")
	err = t.convertMeta()
	if err != nil {
		return err
	}

	slog.Info("Exporting Author file")
	err = t.convertAuthors()
	if err != nil {
		return err
	}

	slog.Info("Exporting Distribution file")
	err = t.convertDistributions()
	if err != nil {
		return err
	}

	slog.Info("Exporting Media file")
	err = t.convertMedia()
	if err != nil {
		return err
	}

	if t.cfg.ColdpNameUsage {
		slog.Info("Exporting NameUsage file")
		err = t.convertNameUsages()
		if err != nil {
			return err
		}
	} else {
		slog.Info("Exporting Name file")
		err = t.convertNames()
		if err != nil {
			return err
		}

		slog.Info("Exporting Synonym file")
		err = t.convertSynonyms()
		if err != nil {
			return err
		}

		slog.Info("Exporting Taxon file")
		err = t.convertTaxa()
		if err != nil {
			return err
		}
	}

	slog.Info("Exporting Name Relationship file")
	err = t.convertNameRelationships()
	if err != nil {
		return err
	}

	slog.Info("Exporting Reference file")
	err = t.convertReferences()
	if err != nil {
		return err
	}

	slog.Info("Exporting Species Estimate file")
	err = t.convertSpeciesEstimates()
	if err != nil {
		return err
	}

	slog.Info("Exporting Species Interaction file")
	err = t.convertSpeciesInteractions()
	if err != nil {
		return err
	}

	slog.Info("Exporting Taxon Concept Relation file")
	err = t.convertTaxonConceptRelations()
	if err != nil {
		return err
	}

	slog.Info("Exporting Taxon Property file")
	err = t.convertTaxonProperties()
	if err != nil {
		return err
	}

	slog.Info("Exporting Treatment file")
	err = t.convertTreatments()
	if err != nil {
		return err
	}

	slog.Info("Exporting Type Material file")
	err = t.convertTypeMaterials()
	if err != nil {
		return err
	}

	slog.Info("Exporting Vernacular file")
	err = t.convertVernaculars()
	if err != nil {
		return err
	}

	err = t.coldp.Export(dst, true)
	if err != nil {
		return err
	}

	return nil
}
