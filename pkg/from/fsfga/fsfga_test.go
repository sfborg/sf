package fsfga_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/sfborg/sf/config"
	"github.com/sfborg/sf/internal/util"
	"github.com/sfborg/sf/pkg/from/fsfga"
	"github.com/sfborg/sflib"
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
	cacheDir, err := os.UserCacheDir()
	if err != nil {
		panic(err)
	}
	base := filepath.Join(cacheDir, "sfborg", "sf-test")
	if err = os.MkdirAll(base, 0755); err != nil {
		panic(err)
	}
	testDir, err = os.MkdirTemp(base, "fsfga-test")
	if err != nil {
		panic(err)
	}
}

func teardownGlobal() {
	err := os.RemoveAll(testDir)
	if err != nil {
		panic(err)
	}
}

// setTmpDir overrides TMPDIR so sflib's internal os.MkdirTemp calls stay on
// the same filesystem as the cache dir, preventing cross-device rename errors.
func setTmpDir(t *testing.T) {
	t.Helper()
	t.Setenv("TMPDIR", testDir)
}

func TestMigrate(t *testing.T) {
	setTmpDir(t)
	assert := assert.New(t)
	src := "../../../testdata/sfga/ptero-v0.4.1.sqlite"
	dst := filepath.Join(testDir, "test")
	cfg := config.New(config.OptCacheDir(testDir))
	err := util.PrepareFileStructure(cfg)
	assert.Nil(err)
	fs := fsfga.New(cfg)
	err = fs.Import(src, dst)
	assert.Nil(err)
	a := sflib.NewSfga()
	a.SetDb(dst + ".sqlite")
	_, err = a.Connect()
	assert.Nil(err)
	assert.True(a.Ping())
	var count string
	err = a.Db().QueryRow("SELECT count(*) FROM taxon").Scan(&count)
	assert.Nil(err)
	assert.Equal("1700", count)
}

func TestMigrateWithParents(t *testing.T) {
	setTmpDir(t)
	assert := assert.New(t)
	src := "../../../testdata/sfga/virus-flat-hier.sqlite"
	dst := filepath.Join(testDir, "virus-parents")
	cfg := config.New(config.OptCacheDir(testDir), config.OptWithParents(true))
	err := util.PrepareFileStructure(cfg)
	assert.Nil(err)
	fs := fsfga.New(cfg)
	err = fs.Import(src, dst)
	assert.Nil(err)
	a := sflib.NewSfga()
	a.SetDb(dst + ".sqlite")
	_, err = a.Connect()
	assert.Nil(err)
	assert.True(a.Ping())
	var origCount, newCount int
	err = a.Db().QueryRow("SELECT COUNT(*) FROM taxon").Scan(&newCount)
	assert.Nil(err)
	assert.Greater(newCount, origCount,
		"AddParents should insert classification taxa, growing the taxon table")
	var hasParents int
	err = a.Db().QueryRow(`
		SELECT COUNT(*) FROM taxon
		WHERE col__parent_id IS NOT NULL AND col__parent_id != ''
	`).Scan(&hasParents)
	assert.Nil(err)
	assert.Greater(hasParents, 0, "taxa should have parent_id set after hierarchy build")
}
