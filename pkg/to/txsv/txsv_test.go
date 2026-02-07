package txsv_test

import (
	"encoding/csv"
	"io"
	"os"
	"path/filepath"
	"testing"

	"github.com/sfborg/sf/config"
	"github.com/sfborg/sf/internal/util"
	"github.com/sfborg/sf/pkg/to/txsv"
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
	testDir, err = os.MkdirTemp("", "sf-to-txsv")
	if err != nil {
		panic(err)
	}
}

func teardownGlobal() {
	os.RemoveAll(testDir)
}

func TestExport(t *testing.T) {
	assert := assert.New(t)

	dir := filepath.Join(testDir, "txsv-export")
	err := os.Mkdir(dir, 0755)
	assert.Nil(err)

	src := filepath.Join("..", "..", "..", "testdata", "sfga", "ptero-v0.4.1.sqlite")
	out := filepath.Join(dir, "name_usage.csv")

	cfg := config.New(config.OptCacheDir(dir))
	err = util.PrepareFileStructure(cfg)
	assert.Nil(err)

	tx := txsv.New(cfg)
	err = tx.Export(src, out)
	assert.Nil(err)

	// Verify the output file exists.
	info, err := os.Stat(out)
	assert.Nil(err)
	assert.Greater(info.Size(), int64(0))

	// Parse CSV and count rows.
	f, err := os.Open(out)
	assert.Nil(err)
	defer f.Close()

	r := csv.NewReader(f)
	var rows int
	for {
		_, err := r.Read()
		if err == io.EOF {
			break
		}
		assert.Nil(err)
		rows++
	}

	// 1700 taxa + 1 header row.
	assert.Equal(1701, rows)
}

func TestExportZip(t *testing.T) {
	assert := assert.New(t)

	dir := filepath.Join(testDir, "txsv-export-zip")
	err := os.Mkdir(dir, 0755)
	assert.Nil(err)

	src := filepath.Join("..", "..", "..", "testdata", "sfga", "ptero-v0.4.1.sqlite")
	out := filepath.Join(dir, "name_usage.csv")

	cfg := config.New(
		config.OptCacheDir(dir),
		config.OptWithZipOutput(true),
	)
	err = util.PrepareFileStructure(cfg)
	assert.Nil(err)

	tx := txsv.New(cfg)
	err = tx.Export(src, out)
	assert.Nil(err)

	// Verify the zip output exists and is non-empty.
	zipPath := out + ".zip"
	info, err := os.Stat(zipPath)
	assert.Nil(err)
	assert.Greater(info.Size(), int64(0))
}

func TestExportBadSource(t *testing.T) {
	assert := assert.New(t)

	dir := filepath.Join(testDir, "txsv-bad-src")
	err := os.Mkdir(dir, 0755)
	assert.Nil(err)

	cfg := config.New(config.OptCacheDir(dir))
	err = util.PrepareFileStructure(cfg)
	assert.Nil(err)

	tx := txsv.New(cfg)
	err = tx.Export("/nonexistent/path.sqlite", filepath.Join(dir, "out.csv"))
	assert.NotNil(err)
}
