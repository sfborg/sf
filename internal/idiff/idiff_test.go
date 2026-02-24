package idiff_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/gnames/gnsys"
	"github.com/sfborg/sf/config"
	"github.com/sfborg/sf/internal/idiff"
	"github.com/sfborg/sf/internal/util"
	"github.com/stretchr/testify/assert"
)

func TestCompare(t *testing.T) {
	assert := assert.New(t)

	src := filepath.Join("../../testdata", "diff",
		"test-a.sqlite")
	ref := filepath.Join("../../testdata", "diff",
		"test-b.sqlite")
	cfg := config.New()

	dir, err := os.MkdirTemp("", "sf-test")
	assert.Nil(err)
	defer os.RemoveAll(dir)

	out := filepath.Join(dir, "test")
	err = util.PrepareFileStructure(cfg)
	assert.Nil(err)
	diff := idiff.New(cfg)
	err = diff.Compare(src, ref, out)
	assert.Nil(err)

	exists, err := gnsys.FileExists(out + ".sqlite")
	assert.Nil(err)
	assert.True(exists)
}
