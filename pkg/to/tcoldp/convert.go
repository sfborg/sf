package tcoldp

import (
	"context"
	"errors"
	"path/filepath"

	"github.com/sfborg/sflib/pkg/coldp"
	"golang.org/x/sync/errgroup"
)

func (t *tcoldp) convertMeta() error {
	meta, err := t.sfga.LoadMeta()
	if err != nil {
		return err
	}

	filePath := filepath.Join(t.cfg.OutputDir, "meta.yaml")
	err = t.coldp.WriteMeta(meta, filePath)
	if err != nil {
		return err
	}

	return nil
}

func (t *tcoldp) convertAuthors() error {
	var err error
	ch := make(chan coldp.Author)
	gr, ctx := errgroup.WithContext(context.Background())
	namePath := filepath.Join(t.cfg.OutputDir, "author.txt")

	gr.Go(func() error {
		err = coldp.Write(ctx, ch, namePath)
		return err
	})

	err = t.sfga.LoadAuthors(ctx, ch)
	if err != nil {
		return err
	}
	close(ch)

	if err = gr.Wait(); err != nil && !errors.Is(err, context.Canceled) {
		return err
	}
	return nil
}

func (t *tcoldp) convertDistributions() error {
	var err error
	ch := make(chan coldp.Distribution)
	gr, ctx := errgroup.WithContext(context.Background())
	namePath := filepath.Join(t.cfg.OutputDir, "distribution.txt")

	gr.Go(func() error {
		err = coldp.Write(ctx, ch, namePath)
		return err
	})

	err = t.sfga.LoadDistributions(ctx, ch)
	if err != nil {
		return err
	}
	close(ch)

	if err = gr.Wait(); err != nil && !errors.Is(err, context.Canceled) {
		return err
	}
	return nil
}

func (t *tcoldp) convertMedia() error {
	var err error
	ch := make(chan coldp.Media)
	gr, ctx := errgroup.WithContext(context.Background())
	namePath := filepath.Join(t.cfg.OutputDir, "media.txt")

	gr.Go(func() error {
		err = coldp.Write(ctx, ch, namePath)
		return err
	})

	err = t.sfga.LoadMedia(ctx, ch)
	if err != nil {
		return err
	}
	close(ch)

	if err = gr.Wait(); err != nil && !errors.Is(err, context.Canceled) {
		return err
	}
	return nil
}

func (t *tcoldp) convertNameRelationships() error {
	var err error
	ch := make(chan coldp.NameRelation)
	gr, ctx := errgroup.WithContext(context.Background())
	namePath := filepath.Join(t.cfg.OutputDir, "name_relationship.txt")

	gr.Go(func() error {
		err = coldp.Write(ctx, ch, namePath)
		return err
	})

	err = t.sfga.LoadNameRelationships(ctx, ch)
	if err != nil {
		return err
	}
	close(ch)

	if err = gr.Wait(); err != nil && !errors.Is(err, context.Canceled) {
		return err
	}
	return nil
}

func (t *tcoldp) convertNames() error {
	var err error
	ch := make(chan coldp.Name)
	gr, ctx := errgroup.WithContext(context.Background())
	namePath := filepath.Join(t.cfg.OutputDir, "name_usage.txt")

	gr.Go(func() error {
		err = coldp.Write(ctx, ch, namePath)
		return err
	})

	err = t.sfga.LoadNames(ctx, ch)
	if err != nil {
		return err
	}
	close(ch)

	if err = gr.Wait(); err != nil && !errors.Is(err, context.Canceled) {
		return err
	}
	return nil
}

func (t *tcoldp) convertNameUsages() error {
	var err error
	ch := make(chan coldp.NameUsage)
	gr, ctx := errgroup.WithContext(context.Background())
	namePath := filepath.Join(t.cfg.OutputDir, "name_usage.txt")

	gr.Go(func() error {
		err = coldp.Write(ctx, ch, namePath)
		return err
	})

	err = t.sfga.LoadNameUsages(ctx, ch)
	if err != nil {
		return err
	}
	close(ch)

	if err = gr.Wait(); err != nil && !errors.Is(err, context.Canceled) {
		return err
	}
	return nil
}

func (t *tcoldp) convertReferences() error {
	var err error
	ch := make(chan coldp.Reference)
	gr, ctx := errgroup.WithContext(context.Background())
	namePath := filepath.Join(t.cfg.OutputDir, "reference.txt")

	gr.Go(func() error {
		err = coldp.Write(ctx, ch, namePath)
		return err
	})

	err = t.sfga.LoadReferences(ctx, ch)
	if err != nil {
		return err
	}
	close(ch)

	if err = gr.Wait(); err != nil && !errors.Is(err, context.Canceled) {
		return err
	}
	return nil
}

func (t *tcoldp) convertSpeciesEstimates() error {
	var err error
	ch := make(chan coldp.SpeciesEstimate)
	gr, ctx := errgroup.WithContext(context.Background())
	namePath := filepath.Join(t.cfg.OutputDir, "species_estimate.txt")

	gr.Go(func() error {
		err = coldp.Write(ctx, ch, namePath)
		return err
	})

	err = t.sfga.LoadSpeciesEstimates(ctx, ch)
	if err != nil {
		return err
	}
	close(ch)

	if err = gr.Wait(); err != nil && !errors.Is(err, context.Canceled) {
		return err
	}
	return nil
}

func (t *tcoldp) convertSpeciesInteractions() error {
	var err error
	ch := make(chan coldp.SpeciesInteraction)
	gr, ctx := errgroup.WithContext(context.Background())
	namePath := filepath.Join(t.cfg.OutputDir, "species_interaction.txt")

	gr.Go(func() error {
		err = coldp.Write(ctx, ch, namePath)
		return err
	})

	err = t.sfga.LoadSpeciesInteractions(ctx, ch)
	if err != nil {
		return err
	}
	close(ch)

	if err = gr.Wait(); err != nil && !errors.Is(err, context.Canceled) {
		return err
	}
	return nil
}

func (t *tcoldp) convertSynonyms() error {
	var err error
	ch := make(chan coldp.Synonym)
	gr, ctx := errgroup.WithContext(context.Background())
	namePath := filepath.Join(t.cfg.OutputDir, "synonym.txt")

	gr.Go(func() error {
		err = coldp.Write(ctx, ch, namePath)
		return err
	})

	err = t.sfga.LoadSynonyms(ctx, ch)
	if err != nil {
		return err
	}
	close(ch)

	if err = gr.Wait(); err != nil && !errors.Is(err, context.Canceled) {
		return err
	}
	return nil
}

func (t *tcoldp) convertTaxa() error {
	var err error
	ch := make(chan coldp.Taxon)
	gr, ctx := errgroup.WithContext(context.Background())
	namePath := filepath.Join(t.cfg.OutputDir, "taxon.txt")

	gr.Go(func() error {
		err = coldp.Write(ctx, ch, namePath)
		return err
	})

	err = t.sfga.LoadTaxa(ctx, ch)
	if err != nil {
		return err
	}
	close(ch)

	if err = gr.Wait(); err != nil && !errors.Is(err, context.Canceled) {
		return err
	}
	return nil
}

func (t *tcoldp) convertTaxonConceptRelations() error {
	var err error
	ch := make(chan coldp.TaxonConceptRelation)
	gr, ctx := errgroup.WithContext(context.Background())
	namePath := filepath.Join(t.cfg.OutputDir, "taxon_concept_relation.txt")

	gr.Go(func() error {
		err = coldp.Write(ctx, ch, namePath)
		return err
	})

	err = t.sfga.LoadTaxonConceptRelations(ctx, ch)
	if err != nil {
		return err
	}
	close(ch)

	if err = gr.Wait(); err != nil && !errors.Is(err, context.Canceled) {
		return err
	}
	return nil
}

func (t *tcoldp) convertTaxonProperties() error {
	var err error
	ch := make(chan coldp.TaxonProperty)
	gr, ctx := errgroup.WithContext(context.Background())
	namePath := filepath.Join(t.cfg.OutputDir, "taxon_property.txt")

	gr.Go(func() error {
		err = coldp.Write(ctx, ch, namePath)
		return err
	})

	err = t.sfga.LoadTaxonProperties(ctx, ch)
	if err != nil {
		return err
	}
	close(ch)

	if err = gr.Wait(); err != nil && !errors.Is(err, context.Canceled) {
		return err
	}
	return nil
}

func (t *tcoldp) convertTreatments() error {
	var err error
	ch := make(chan coldp.Treatment)
	gr, ctx := errgroup.WithContext(context.Background())
	namePath := filepath.Join(t.cfg.OutputDir, "treatment.txt")

	gr.Go(func() error {
		err = coldp.Write(ctx, ch, namePath)
		return err
	})

	err = t.sfga.LoadTreatments(ctx, ch)
	if err != nil {
		return err
	}
	close(ch)

	if err = gr.Wait(); err != nil && !errors.Is(err, context.Canceled) {
		return err
	}
	return nil
}

func (t *tcoldp) convertTypeMaterials() error {
	var err error
	ch := make(chan coldp.TypeMaterial)
	gr, ctx := errgroup.WithContext(context.Background())
	namePath := filepath.Join(t.cfg.OutputDir, "type_material.txt")

	gr.Go(func() error {
		err = coldp.Write(ctx, ch, namePath)
		return err
	})

	err = t.sfga.LoadTypeMaterials(ctx, ch)
	if err != nil {
		return err
	}
	close(ch)

	if err = gr.Wait(); err != nil && !errors.Is(err, context.Canceled) {
		return err
	}
	return nil
}

func (t *tcoldp) convertVernaculars() error {
	var err error
	ch := make(chan coldp.Vernacular)
	gr, ctx := errgroup.WithContext(context.Background())
	namePath := filepath.Join(t.cfg.OutputDir, "vernacular.txt")

	gr.Go(func() error {
		err = coldp.Write(ctx, ch, namePath)
		return err
	})

	err = t.sfga.LoadVernaculars(ctx, ch)
	if err != nil {
		return err
	}
	close(ch)

	if err = gr.Wait(); err != nil && !errors.Is(err, context.Canceled) {
		return err
	}
	return nil
}
