package fdwca_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/sfborg/sf/config"
	"github.com/sfborg/sf/internal/util"
	"github.com/sfborg/sf/pkg/from/fdwca"
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
	testDir, err = os.MkdirTemp("", "sf-from-dwca")
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
	dir := filepath.Join(testDir, "sf-from-dwca")

	err := os.Mkdir(dir, 0755)
	assert.Nil(err)
	defer os.RemoveAll(dir)

	tests := []struct {
		msg, file                            string
		taxons, synonyms, names, vern, distr int
	}{
		{"col", "col-mini.zip", 73, 6, 79, 19, 18},
		// two vernacular files
		{"2vern", "two-vern.tar.gz", 73, 6, 79, 19, 18},
	}

	for _, v := range tests {
		path := filepath.Join("../../../testdata/dwca", v.file)
		out := filepath.Join(dir, "output")
		opts := []config.Option{
			config.OptCacheDir(dir),
			config.OptBatchSize(5),
		}
		cfg := config.New(opts...)
		util.PrepareFileStructure(cfg)

		libOpts := cfg.OptsSflib()

		fd := fdwca.New(cfg)
		err := fd.Import(path, out)
		assert.Nil(err)
		sf := sflib.NewSfga(libOpts...)
		sf.SetDb(out + ".sqlite")
		db, err := sf.Connect()
		assert.Nil(err)

		var count int
		err = db.QueryRow("select count(*) from taxon").Scan(&count)
		assert.Nil(err)
		assert.Equal(v.taxons, count)

		err = db.QueryRow("select count(*) from synonym").Scan(&count)
		assert.Nil(err)
		assert.Equal(v.synonyms, count)

		err = db.QueryRow("select count(*) from name").Scan(&count)
		assert.Nil(err)
		assert.Equal(v.names, count)

		err = db.QueryRow("select count(*) from vernacular").Scan(&count)
		assert.Nil(err)
		assert.Equal(v.vern, count)

		err = db.QueryRow("select count(*) from distribution").Scan(&count)
		assert.Nil(err)
		assert.Equal(v.distr, count)
	}
}
