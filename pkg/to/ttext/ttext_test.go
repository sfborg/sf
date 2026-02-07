package ttext_test

import (
	"bufio"
	"os"
	"path/filepath"
	"testing"

	"github.com/sfborg/sf/config"
	"github.com/sfborg/sf/internal/util"
	"github.com/sfborg/sf/pkg/to/ttext"
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
	testDir, err = os.MkdirTemp("", "sf-to-ttext")
	if err != nil {
		panic(err)
	}
}

func teardownGlobal() {
	os.RemoveAll(testDir)
}

func TestExport(t *testing.T) {
	assert := assert.New(t)

	dir := filepath.Join(testDir, "ttext-export")
	err := os.Mkdir(dir, 0755)
	assert.Nil(err)

	src := filepath.Join("..", "..", "..", "testdata", "sfga", "ptero-v0.4.1.sqlite")
	out := filepath.Join(dir, "names.txt")

	cfg := config.New(config.OptCacheDir(dir))
	err = util.PrepareFileStructure(cfg)
	assert.Nil(err)

	tt := ttext.New(cfg)
	err = tt.Export(src, out)
	assert.Nil(err)

	// Verify the output file exists.
	info, err := os.Stat(out)
	assert.Nil(err)
	assert.Greater(info.Size(), int64(0))

	// Count lines and verify content.
	lines := readLines(t, out)

	// The ptero-v0.4.1.sqlite has 2982 name usages (1700 taxa + 1282 synonyms).
	assert.Equal(2982, len(lines))

	// Verify lines contain scientific names (non-empty strings).
	for _, line := range lines {
		assert.NotEmpty(line)
	}
}

func TestExportZip(t *testing.T) {
	assert := assert.New(t)

	dir := filepath.Join(testDir, "ttext-export-zip")
	err := os.Mkdir(dir, 0755)
	assert.Nil(err)

	src := filepath.Join("..", "..", "..", "testdata", "sfga", "ptero-v0.4.1.sqlite")
	out := filepath.Join(dir, "names.txt")

	cfg := config.New(
		config.OptCacheDir(dir),
		config.OptWithZipOutput(true),
	)
	err = util.PrepareFileStructure(cfg)
	assert.Nil(err)

	tt := ttext.New(cfg)
	err = tt.Export(src, out)
	assert.Nil(err)

	// Verify the output zip exists and is non-empty.
	zipPath := out + ".zip"
	info, err := os.Stat(zipPath)
	assert.Nil(err)
	assert.Greater(info.Size(), int64(0))
}

func TestExportBadSource(t *testing.T) {
	assert := assert.New(t)

	dir := filepath.Join(testDir, "ttext-bad-src")
	err := os.Mkdir(dir, 0755)
	assert.Nil(err)

	cfg := config.New(config.OptCacheDir(dir))
	err = util.PrepareFileStructure(cfg)
	assert.Nil(err)

	tt := ttext.New(cfg)
	err = tt.Export("/nonexistent/path.sqlite", filepath.Join(dir, "out.txt"))
	assert.NotNil(err)
}

func readLines(t *testing.T, path string) []string {
	t.Helper()
	f, err := os.Open(path)
	assert.Nil(t, err)
	defer f.Close()

	var lines []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	assert.Nil(t, scanner.Err())
	return lines
}
