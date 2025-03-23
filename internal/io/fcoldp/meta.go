package fcoldp

import "github.com/sfborg/sflib/pkg/coldp"

func (fc *fcoldp) importMeta(coldp coldp.Archive) error {
	meta, err := coldp.Meta()
	if err != nil {
		return err
	}

	err = fc.sfga.InsertMeta(meta)
	if err != nil {
		return err
	}
	return nil
}
