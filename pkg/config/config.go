package config

import (
	"os"
	"path/filepath"
)

// Config contains configuration data of the app.
type Config struct {
	// CacheDir is a path to working directory. Files in this directory
	// are cleaned up before each use of the app.
	CacheDir string

	// DiffSrcDir is a path of a directory where the source SFGA file resides.
	// This file is to be compared with the target SFGA file.
	DiffSrcDir string

	// DiffTrgDir is a path to a directory where the target SFGA file resides.
	// This source SFGA file will be compared with the target file.
	DiffTrgDir string

	// DiffWorkDir contains data necessary for comparing data of source and
	// target SFGA files. It can be a suffix trie data, bloom filter backup etc.
	DiffWorkDir string

	// DiffSourceTaxon defines a taxon in the source file that limits comparison
	// to the children of the taxon.
	DiffSourceTaxon string

	// DiffTargetTaxon defines a taxon in the target file that limits comparison
	// to the children of the taxon.
	DiffTargetTaxon string
}

// Option type is used for all options sent to the config file.
type Option func(*Config)

func OptCacheDir(s string) Option {
	return func(c *Config) {
		c.CacheDir = s
	}
}

func OptDiffSourceTaxon(s string) Option {
	return func(c *Config) {
		c.DiffSourceTaxon = s
	}
}

func OptDiffTargetTaxon(s string) Option {
	return func(c *Config) {
		c.DiffTargetTaxon = s
	}
}

func New(opts ...Option) Config {
	tmpDir := os.TempDir()
	cacheDir, err := os.UserCacheDir()
	if err != nil {
		cacheDir = tmpDir
	}

	cacheDir = filepath.Join(cacheDir, "sfborg", "sf")

	res := Config{
		CacheDir: cacheDir,
	}

	for _, opt := range opts {
		opt(&res)
	}

	res.DiffSrcDir = filepath.Join(res.CacheDir, "diff", "src")
	res.DiffTrgDir = filepath.Join(res.CacheDir, "diff", "trg")
	res.DiffWorkDir = filepath.Join(res.CacheDir, "diff", "work")

	return res
}
