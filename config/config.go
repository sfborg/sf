package config

import (
	"os"
	"path/filepath"

	"github.com/gnames/gnfmt"
	"github.com/gnames/gnlib/ent/nomcode"
	sflibCfg "github.com/sfborg/sflib/config"
)

// Config contains configuration data of the app.
type Config struct {
	// CacheDir is a path to working directory. Files in this directory
	// are cleaned up before each use of the app.
	CacheDir string

	// DownloadDir is the path to the directory where downloaded files are
	// stored
	DownloadDir string

	// ImportDir is a path to a directory where source files are moved
	// or extracted to.
	ImportDir string

	// OutputDir is a cache directory where new archive files are created. When
	// all is ready the files will be exported to an output file.
	OutputDir string

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

	// ColdpNameUsage tells ToColdp covertor to combine name, taxon, synonym
	// data into name usage.
	ColdpNameUsage bool

	// NomCode tells which Nomenclatural Code to insert to all records of
	// coldp.Name records, as well as setting up GNparser code mode.
	// If imported data alread has the Code information, the data has a
	// precedence.
	NomCode nomcode.Code

	// BadRow sets how to process rows with wrong number of fields in CSV
	// files. By default it is set to process such rows. Other options are
	// to return an error, or skip them.
	BadRow gnfmt.BadRow

	// BatchSize determines the size of slices to import into SFGA.
	BatchSize int

	// Number of concurrent jobs.
	JobsNum int

	// WithParents can be used to attempt creation of parent/child tree out of
	// flat classification.
	WithParents bool

	// WithQuotes can be used to parse faster tab- or pipe-delimited
	// files where fields never escaped by quotes.
	WithQuotes bool

	// WithZipOutput indicates that zipped archives have to be created.
	WithZipOutput bool

	// WithParser indicates that GNparser detailed data will be used to
	// populate name fields (eg. data like  Uninomial, Genus, SpecificEpithet,
	// CombinationAuthorship etc).
	WithParser bool
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

func OptColdpNameUsage(b bool) Option {
	return func(c *Config) {
		c.ColdpNameUsage = b
	}
}

func OptNomCode(code nomcode.Code) Option {
	return func(c *Config) {
		c.NomCode = code
	}
}

func OptBatchSize(i int) Option {
	return func(c *Config) {
		c.BatchSize = i
	}
}

func OptJobsNum(i int) Option {
	return func(c *Config) {
		c.JobsNum = i
	}
}

func OptWithQuotes(b bool) Option {
	return func(c *Config) {
		c.WithQuotes = b
	}
}

func OptBadRow(br gnfmt.BadRow) Option {
	return func(c *Config) {
		c.BadRow = br
	}
}

func OptNoParser(b bool) Option {
	return func(c *Config) {
		c.WithParser = !b
	}
}

func OptWithParents(b bool) Option {
	return func(c *Config) {
		c.WithParents = b
	}
}

func OptWithZipOutput(b bool) Option {
	return func(c *Config) {
		c.WithZipOutput = b
	}
}

func (c Config) OptsSflib() []sflibCfg.Option {
	opts := []sflibCfg.Option{
		sflibCfg.OptBadRow(c.BadRow),
		sflibCfg.OptWithQuotes(c.WithQuotes),
		sflibCfg.OptBatchSize(c.BatchSize),
		sflibCfg.OptJobsNum(c.JobsNum),
		sflibCfg.OptNomCode(c.NomCode),
		sflibCfg.OptWithParents(c.WithParents),
	}
	return opts
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
		WithParser:  true,
	}

	for _, opt := range opts {
		opt(&res)
	}

	res.DownloadDir = filepath.Join(res.CacheDir, "import", "download")
	res.ImportDir = filepath.Join(res.CacheDir, "import", "src")
	res.OutputDir = filepath.Join(res.CacheDir, "import", "sfga")
	res.DiffSrcDir = filepath.Join(res.CacheDir, "diff", "src")
	res.DiffRefDir = filepath.Join(res.CacheDir, "diff", "trg")

	return res
}
