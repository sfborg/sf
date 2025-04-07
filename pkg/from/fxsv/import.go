package fxsv

func (fx *fxsv) Import(src, out string) error {
	var err error

	err = fx.xsv.Fetch(src, fx.cfg.ImportDir)
	if err != nil {
		return err
	}

	fx.sfga, err = fx.InitSfga()
	if err != nil {
		return err
	}

	err = fx.importNamesUsage()
	if err != nil {
		return err
	}

	err = fx.sfga.Export(out, fx.cfg.WithZipOutput)
	if err != nil {
		return err
	}

	return nil
}
