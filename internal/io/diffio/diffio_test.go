package diffio_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/gnames/gnsys"
	"github.com/sfborg/sf/internal/io/diffio"
	"github.com/sfborg/sf/internal/io/sysio"
	"github.com/sfborg/sf/pkg/config"
	"github.com/stretchr/testify/assert"
)

func TestCompare(t *testing.T) {
	assert := assert.New(t)

	src := filepath.Join("..", "..", "..", "testdata", "diff",
		"test-a.sqlite")
	ref := filepath.Join("..", "..", "..", "testdata", "diff",
		"test-b.sqlite")
	cfg := config.New()

	dir, err := os.MkdirTemp("", "sf")
	assert.Nil(err)
	defer os.RemoveAll(dir)

	out := filepath.Join(dir, "test")
	sysio.PrepareFileStructure(cfg)
	diff := diffio.New(cfg)
	err = diff.Compare(src, ref, out)
	assert.Nil(err)

	exists, err := gnsys.FileExists(out + ".sqlite")
	assert.Nil(err)
	assert.True(exists)
}
