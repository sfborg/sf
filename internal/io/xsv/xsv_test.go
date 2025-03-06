package xsv

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/gnames/gnsys"
	"github.com/sfborg/sf/internal/io/sysio"
	"github.com/sfborg/sf/pkg/config"
	"github.com/stretchr/testify/assert"
)

func TestImport(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		name      string
		src       string
		out       string
		errNotNil bool
	}{
		{
			name:      "csv",
			src:       "ioc-bird.csv",
			errNotNil: false,
		},
		{
			name:      "tsv",
			src:       "ioc-bird.tsv",
			errNotNil: false,
		},
		{
			name:      "psv",
			src:       "ioc-bird.psv",
			errNotNil: false,
		},
		{
			name:      "non-existing-file",
			src:       "non-existing.csv",
			errNotNil: true,
		},
	}

	dir, err := os.MkdirTemp("", "xsv-")
	assert.Nil(err)
	defer os.RemoveAll(dir)

	out := filepath.Join(dir, "test")
	for _, v := range tests {
		src := filepath.Join("..", "..", "..", "testdata", "csv", v.src)
		cfg := config.New()
		x := &xsv{
			cfg: cfg,
		}
		err := sysio.PrepareFileStructure(cfg)
		assert.Nil(err)

		err = x.Import(src, out)
		assert.Equal(v.errNotNil, err != nil)

		exists, err := gnsys.FileExists(out + ".sqlite")
		assert.Nil(err)
		assert.True(exists)
	}
}
