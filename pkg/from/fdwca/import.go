package fdwca

import "log/slog"

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

	err = fd.importEML()
	if err != nil {
		return err
	}

	err = fd.importCore()
	if err != nil {
		return err
	}

	err = fd.importExtensions()
	if err != nil {
		return err
	}

	slog.Info("Preparing SFGA file")
	err = fd.sfga.Export(out, fd.cfg.WithZipOutput)
	if err != nil {
		return err
	}

	return nil
}
