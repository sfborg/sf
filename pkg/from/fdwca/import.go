package fdwca

func (fd *fdwca) Import(src, out string) error {
	var err error

	err = fd.dwca.Fetch(src, fd.cfg.ImportDir)
	if err != nil {
		return err
	}

	fd.sfga, err = fd.InitSfga()
	if err != nil {
		return err
	}

	err = fd.importNamesUsage()
	if err != nil {
		return err
	}

	err = fd.sfga.Export(out, fd.cfg.WithZipOutput)
	if err != nil {
		return err
	}

	return nil
}
