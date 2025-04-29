package fdwca

func (fd *fdwca) Import(src, out string) error {
	var err error

	fd.sfga, err = fd.InitSfga()
	if err != nil {
		return err
	}

	err = fd.dwca.Fetch(src, fd.cfg.ImportDir)
	if err != nil {
		return err
	}

	err = fd.importCore()
	if err != nil {
		return err
	}

	err = fd.sfga.Export(out, fd.cfg.WithZipOutput)
	if err != nil {
		return err
	}

	return nil
}
