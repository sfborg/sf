package tdwca_test

import (
	"archive/zip"
	"encoding/csv"
	"io"
	"os"
	"path/filepath"
	"testing"

	"github.com/sfborg/sf/config"
	"github.com/sfborg/sf/internal/util"
	"github.com/sfborg/sf/pkg/to/tdwca"
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
	testDir, err = os.MkdirTemp("", "sf-to-dwca")
	if err != nil {
		panic(err)
	}
}

func teardownGlobal() {
	os.RemoveAll(testDir)
}

func TestExport(t *testing.T) {
	assert := assert.New(t)

	dir := filepath.Join(testDir, "dwca-export")
	err := os.Mkdir(dir, 0755)
	assert.Nil(err)

	src := filepath.Join("..", "..", "..", "testdata", "sfga", "ptero-v0.4.1.sqlite")
	out := filepath.Join(dir, "output.zip")

	cfg := config.New(config.OptCacheDir(dir))
	err = util.PrepareFileStructure(cfg)
	assert.Nil(err)

	td := tdwca.New(cfg)
	err = td.Export(src, out)
	assert.Nil(err)

	// Verify the output zip exists.
	_, err = os.Stat(out)
	assert.Nil(err)

	// Open the zip and inspect contents.
	zr, err := zip.OpenReader(out)
	assert.Nil(err)
	defer zr.Close()

	files := make(map[string]*zip.File)
	for _, f := range zr.File {
		files[f.Name] = f
	}

	// Must contain eml.xml, meta.xml, Taxon.csv.
	assert.Contains(files, "eml.xml")
	assert.Contains(files, "meta.xml")
	assert.Contains(files, "Taxon.csv")

	// Distribution.csv should be present (6118 records in source).
	assert.Contains(files, "Distribution.csv")

	// VernacularName.csv should NOT be present (0 vernaculars in source).
	assert.NotContains(files, "VernacularName.csv")

	// Verify Taxon.csv row count: 1700 taxa + 1282 synonyms = 2982 data + 1 header.
	coreRows := countCSVRows(t, files["Taxon.csv"])
	assert.Equal(2983, coreRows)

	// Verify Distribution.csv row count: 6118 data + 1 header.
	distrRows := countCSVRows(t, files["Distribution.csv"])
	assert.Equal(6119, distrRows)

	// Verify meta.xml mentions Distribution but not VernacularName.
	metaContent := readZipFile(t, files["meta.xml"])
	assert.Contains(metaContent, "Distribution.csv")
	assert.NotContains(metaContent, "VernacularName.csv")
}

func countCSVRows(t *testing.T, f *zip.File) int {
	t.Helper()
	rc, err := f.Open()
	assert.Nil(t, err)
	defer rc.Close()

	r := csv.NewReader(rc)
	var count int
	for {
		_, err := r.Read()
		if err == io.EOF {
			break
		}
		assert.Nil(t, err)
		count++
	}
	return count
}

func readZipFile(t *testing.T, f *zip.File) string {
	t.Helper()
	rc, err := f.Open()
	assert.Nil(t, err)
	defer rc.Close()

	bs, err := io.ReadAll(rc)
	assert.Nil(t, err)
	return string(bs)
}
