package fcoldp

func (fc *fcoldp) importMeta() error {
	meta, err := fc.coldp.Meta()
	if err != nil {
		return err
	}

	err = fc.sfga.InsertMeta(meta)
	if err != nil {
		return err
	}
	return nil
}
