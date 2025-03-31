package config

import (
	"os"
	"path/filepath"

	"github.com/gnames/gnfmt"
	"github.com/sfborg/sflib/pkg/coldp"
)

// Config contains configuration data of the app.
type Config struct {
	// CacheDir is a path to working directory. Files in this directory
	// are cleaned up before each use of the app.
	CacheDir string

	// DownloadDir is the path to the directory where downloaded files are
	// stored
	DownloadDir string

	// DataDir is a path to a directory where source files are moved
	// or extracted to.
	DataDir string

	// SfgaDir is a path where SFGA database is created. When
	// the database is ready it is exported to output file.
	SfgaDir string

	// DiffSrcDir is a path of a directory where the source SFGA file resides.
	// This file is to be compared with the target SFGA file.
	DiffSrcDir string

	// DiffRefDir is a path to a directory where the target SFGA file resides.
	// This source SFGA file will be compared with the target file.
	DiffRefDir string

	// DiffWorkDir contains data necessary for comparing data of source and
	// target SFGA files. It can be a suffix trie data, bloom filter backup etc.
	DiffWorkDir string

	// DiffSourceTaxon defines a taxon in the source file that limits comparison
	// to the children of the taxon.
	DiffSourceTaxon string

	// DiffTargetTaxon defines a taxon in the target file that limits comparison
	// to the children of the taxon.
	DiffTargetTaxon string

	// NomCode tells which Nomenclatural Code to insert to all records of
	// coldp.Name records, as well as setting up GNparser code mode.
	// If imported data alread has the Code information, the data has a
	// precedence.
	NomCode coldp.NomCode

	// BadRow sets how to process rows with wrong number of fields in CSV
	// files. By default it is set to process such rows. Other options are
	// to return an error, or skip them.
	BadRow gnfmt.BadRow

	// BatchSize determines the size of slices to import into SFGA.
	BatchSize int

	// Number of concurrent jobs.
	JobsNum int

	// WithQuotes can be used to parse faster tab- or pipe-delimited
	// files where fields never escaped by quotes.
	WithQuotes bool

	// WithZipOutput indicates that zipped archives have to be created.
	WithZipOutput bool

	// WithDetails indicates that GNparser detailed data will be used to
	// populate name fields (eg. data like  Uninomial, Genus, SpecificEpithet,
	// CombinationAuthorship etc).
	WithDetails bool
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

func OptNomCode(code coldp.NomCode) Option {
	return func(c *Config) {
		c.NomCode = code
	}
}

func OptWithoutQuotes(b bool) Option {
	return func(c *Config) {
		c.WithQuotes = b
	}
}

func OptBadRow(br gnfmt.BadRow) Option {
	return func(c *Config) {
		c.BadRow = br
	}
}

func OptWithDetails(b bool) Option {
	return func(c *Config) {
		c.WithDetails = b
	}
}

func OptWithZipOutput(b bool) Option {
	return func(c *Config) {
		c.WithZipOutput = b
	}
}

func New(opts ...Option) Config {
	tmpDir := os.TempDir()
	cacheDir, err := os.UserCacheDir()
	if err != nil {
		cacheDir = tmpDir
	}

	workDir := filepath.Join(cacheDir, "sfborg", "sf-keep")
	cacheDir = filepath.Join(cacheDir, "sfborg", "sf")

	res := Config{
		CacheDir:    cacheDir,
		DiffWorkDir: workDir,
		BadRow:      gnfmt.ProcessBadRow,
		BatchSize:   50_000,
		JobsNum:     5,
	}

	for _, opt := range opts {
		opt(&res)
	}

	res.DownloadDir = filepath.Join(res.CacheDir, "import", "download")
	res.DataDir = filepath.Join(res.CacheDir, "import", "src")
	res.SfgaDir = filepath.Join(res.CacheDir, "import", "sfga")
	res.DiffSrcDir = filepath.Join(res.CacheDir, "diff", "src")
	res.DiffRefDir = filepath.Join(res.CacheDir, "diff", "trg")

	return res
}
