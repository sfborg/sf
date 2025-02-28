package config

import (
	"os"
	"path/filepath"
)

type Config struct {
	CacheDir    string
	DiffSrcDir  string
	DiffTrgDir  string
	DiffWorkDir string
}

type Option func(*Config)

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
