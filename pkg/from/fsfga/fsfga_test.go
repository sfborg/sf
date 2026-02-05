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
	code := m.Run() // Run all tests
	teardownGlobal()
	os.Exit(code)
}

func setupGlobal() {
	var err error
	testDir, err = os.MkdirTemp("", "fsfga-test")
	if err != nil {
		panic(err)
	}
}

func teardownGlobal() {
	var err error
	err = os.RemoveAll(testDir)
	if err != nil {
		panic(err)
	}
}

func TestUpdate(t *testing.T) {
	assert := assert.New(t)
	src := "../../../testdata/sfga/ptero-v0.4.1.sqlite"
	dst := filepath.Join(testDir, "test")
	cfg := config.New()
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
	var res string
	err = a.Db().QueryRow("select count(*) from taxon").Scan(&res)
	assert.Nil(err)
	assert.Equal("1700", res)
}
