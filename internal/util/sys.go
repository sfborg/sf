package util

import (
	"github.com/gnames/gnsys"
	"github.com/sfborg/sf/config"
)

func PrepareFileStructure(cfg config.Config) error {
	var err error
	root := cfg.CacheDir
	err = gnsys.MakeDir(root)
	if err != nil {
		return err
	}
	// create cfg.DiffWorkDir if does not exist
	err = gnsys.MakeDir(cfg.DiffWorkDir)
	if err != nil {
		return err
	}

	err = gnsys.CleanDir(root)
	if err != nil {
		return err
	}
	dirs := []string{
		cfg.DownloadDir,
		cfg.DiffSrcDir,
		cfg.DiffRefDir,
		cfg.ImportDir,
		cfg.OutputDir,
	}
	for _, v := range dirs {
		err = gnsys.MakeDir(v)
		if err != nil {
			return err
		}
	}

	return nil
}
