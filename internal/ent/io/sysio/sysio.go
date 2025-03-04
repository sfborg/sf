package sysio

import (
	"github.com/gnames/gnsys"
	"github.com/sfborg/sf/pkg/config"
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
		cfg.DiffSrcDir,
		cfg.DiffTrgDir,
	}
	for _, v := range dirs {
		err = gnsys.MakeDir(v)
		if err != nil {
			return err
		}
	}

	return nil
}
