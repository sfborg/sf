package ftext_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/gnames/gnsys"
	"github.com/sfborg/sf/pkg/config"
	"github.com/sfborg/sf/pkg/from/ftext"
	"github.com/sfborg/sflib"
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
			name:      "txt",
			src:       "names.txt",
			errNotNil: false,
		},
		{
			name:      "non-existing-file",
			src:       "non-existing.txt",
			errNotNil: true,
		},
	}

	dir, err := os.MkdirTemp("", "txt-")
	assert.Nil(err)
	defer os.RemoveAll(dir)

	out := filepath.Join(dir, "test")
	for _, v := range tests {
		src := filepath.Join("..", "..", "..", "testdata", "text", v.src)
		cfg := config.New()
		a := ftext.New(cfg)

		err = a.Import(src, out)
		assert.Equal(v.errNotNil, err != nil)

		if err == nil {
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
}
