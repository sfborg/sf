package tcoldp_test

import (
	"archive/zip"
	"bufio"
	"io"
	"os"
	"path/filepath"
	"testing"

	"github.com/sfborg/sf/config"
	"github.com/sfborg/sf/internal/util"
	"github.com/sfborg/sf/pkg/to/tcoldp"
	"github.com/stretchr/testify/assert"
)

var testDir string

func TestMain(m *testing.M) {
	setupGlobal()
	code := m.Run()
	teardownGlobal()
	os.Exit(code)
}

func setupGlobal() {
	var err error
	testDir, err = os.MkdirTemp("", "sf-to-tcoldp")
	if err != nil {
		panic(err)
	}
}

func teardownGlobal() {
	os.RemoveAll(testDir)
}

func TestExport(t *testing.T) {
	assert := assert.New(t)

	dir := filepath.Join(testDir, "tcoldp-export")
	err := os.Mkdir(dir, 0755)
	assert.Nil(err)

	src := filepath.Join("..", "..", "..", "testdata", "sfga", "ptero-v0.5.1.sqlite")
	out := filepath.Join(dir, "output.zip")

	cfg := config.New(config.OptCacheDir(dir))
	err = util.PrepareFileStructure(cfg)
	assert.Nil(err)

	tc := tcoldp.New(cfg)
	err = tc.Export(src, out)
	assert.Nil(err)

	// Verify the output zip exists.
	_, err = os.Stat(out)
	assert.Nil(err)

	// Open the zip and inspect contents.
	zr, err := zip.OpenReader(out)
	assert.Nil(err)
	defer zr.Close()

	files := make(map[string]struct{})
	for _, f := range zr.File {
		files[f.Name] = struct{}{}
	}

	// Default mode (ColdpNameUsage=false) exports separate name, taxon, synonym files.
	assert.Contains(files, "meta.yaml")
	assert.Contains(files, "name_usage.txt")
	assert.Contains(files, "taxon.txt")
	assert.Contains(files, "synonym.txt")
	assert.Contains(files, "distribution.txt")
	assert.Contains(files, "reference.txt")
}

func TestExportNameUsageMode(t *testing.T) {
	assert := assert.New(t)

	dir := filepath.Join(testDir, "tcoldp-export-nu")
	err := os.Mkdir(dir, 0755)
	assert.Nil(err)

	src := filepath.Join("..", "..", "..", "testdata", "sfga", "ptero-v0.5.1.sqlite")
	out := filepath.Join(dir, "output.zip")

	cfg := config.New(
		config.OptCacheDir(dir),
		config.OptColdpNameUsage(true),
	)
	err = util.PrepareFileStructure(cfg)
	assert.Nil(err)

	tc := tcoldp.New(cfg)
	err = tc.Export(src, out)
	assert.Nil(err)

	// Verify the output zip exists.
	_, err = os.Stat(out)
	assert.Nil(err)

	// Open the zip and inspect contents.
	zr, err := zip.OpenReader(out)
	assert.Nil(err)
	defer zr.Close()

	files := make(map[string]struct{})
	for _, f := range zr.File {
		files[f.Name] = struct{}{}
	}

	// NameUsage mode should have name_usage.txt but not separate taxon/synonym.
	assert.Contains(files, "meta.yaml")
	assert.Contains(files, "name_usage.txt")
	assert.Contains(files, "distribution.txt")
}

// TestExportVirusHierarchy verifies that a flat-hierarchy SFGA file with no
// metadata produces a valid CoLDP zip — without a "null" meta.yaml.
func TestExportVirusHierarchy(t *testing.T) {
	assert := assert.New(t)

	dir := filepath.Join(testDir, "tcoldp-export-virus")
	err := os.Mkdir(dir, 0755)
	assert.Nil(err)

	src := filepath.Join("..", "..", "..", "testdata", "sfga", "virus-flat-hier.sqlite")
	out := filepath.Join(dir, "output.zip")

	cfg := config.New(config.OptCacheDir(dir))
	err = util.PrepareFileStructure(cfg)
	assert.Nil(err)

	tc := tcoldp.New(cfg)
	err = tc.Export(src, out)
	assert.Nil(err)

	_, err = os.Stat(out)
	assert.Nil(err)

	zr, err := zip.OpenReader(out)
	assert.Nil(err)
	defer zr.Close()

	files := make(map[string]*zip.File)
	for _, f := range zr.File {
		files[f.Name] = f
	}

	// No metadata in source — meta.yaml must be present with filename as title.
	assert.Contains(files, "meta.yaml")
	metaContent := readZipText(t, files["meta.yaml"])
	assert.Contains(metaContent, "virus-flat-hier")
	assert.NotContains(metaContent, "null")

	// 84 taxa, 0 synonyms.
	assert.Contains(files, "taxon.txt")
	taxonRows := countTSVDataRows(t, files["taxon.txt"])
	assert.Equal(84, taxonRows)

	assert.Contains(files, "name_usage.txt")
	nameRows := countTSVDataRows(t, files["name_usage.txt"])
	assert.Equal(84, nameRows)
}

func readZipText(t *testing.T, f *zip.File) string {
	t.Helper()
	rc, err := f.Open()
	assert.Nil(t, err)
	defer rc.Close()
	bs, err := io.ReadAll(rc)
	assert.Nil(t, err)
	return string(bs)
}

// countTSVDataRows counts non-header lines in a zip entry (tab-separated).
func countTSVDataRows(t *testing.T, f *zip.File) int {
	t.Helper()
	rc, err := f.Open()
	assert.Nil(t, err)
	defer rc.Close()

	scanner := bufio.NewScanner(rc)
	count := -1 // skip header line
	for scanner.Scan() {
		if scanner.Text() != "" {
			count++
		}
	}
	assert.Nil(t, scanner.Err())
	if count < 0 {
		return 0
	}
	return count
}


func TestExportBadSource(t *testing.T) {
	assert := assert.New(t)

	dir := filepath.Join(testDir, "tcoldp-bad-src")
	err := os.Mkdir(dir, 0755)
	assert.Nil(err)

	cfg := config.New(config.OptCacheDir(dir))
	err = util.PrepareFileStructure(cfg)
	assert.Nil(err)

	tc := tcoldp.New(cfg)
	err = tc.Export("/nonexistent/path.sqlite", filepath.Join(dir, "out.zip"))
	assert.NotNil(err)
}
