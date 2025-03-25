package fcoldp_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/gnames/gnsys"
	"github.com/sfborg/sf/internal/util"
	"github.com/sfborg/sf/pkg/config"
	"github.com/sfborg/sf/pkg/from/fcoldp"
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
	testDir, err = os.MkdirTemp("", "coldp-test")
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

func TestImport(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		name   string
		src    string
		out    string
		errNil bool
	}{
		{
			name:   "ptero",
			src:    "ptero-yaml.coldp.zip",
			errNil: true,
		},
		{
			name:   "non-existing-file",
			src:    "non-existing.csv",
			errNil: false,
		},
	}

	dir := filepath.Join(testDir, "fcoldp")
	out := filepath.Join(dir, "test")
	err := os.Mkdir(dir, 0777)
	assert.Nil(err)

	for _, v := range tests {
		err = gnsys.CleanDir(dir)
		assert.Nil(err)

		src := filepath.Join("..", "..", "..", "testdata", "coldp", v.src)
		cfg := config.New(config.OptCacheDir(dir))
		fc := fcoldp.New(cfg)
		err := util.PrepareFileStructure(cfg)
		assert.Nil(err)

		err = fc.Import(src, out)
		assert.Equal(v.errNil, err == nil)

		if err != nil {
			continue
		}
		exists, err := gnsys.FileExists(out + ".sqlite")
		assert.Nil(err)
		assert.True(exists)

		sfga := sflib.NewSfga()
		sfga.SetDb(out + ".sqlite")
		db, err := sfga.Connect()
		assert.Nil(err)
		var count int
		err = db.QueryRow("SELECT count(*) FROM name").Scan(&count)
		assert.Nil(err)
		assert.Greater(count, 5, v.name)
	}
}
