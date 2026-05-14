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

	cacheDir, err := os.UserCacheDir()
	assert.Nil(err)
	base := filepath.Join(cacheDir, "sfborg", "sf-test")
	err = os.MkdirAll(base, 0755)
	assert.Nil(err)
	testDir, err := os.MkdirTemp(base, "idiff-test")
	assert.Nil(err)
	defer os.RemoveAll(testDir)
	t.Setenv("TMPDIR", testDir)

	src := filepath.Join("../../testdata", "diff",
		"test-a.sqlite")
	ref := filepath.Join("../../testdata", "diff",
		"test-b.sqlite")
	cfg := config.New(config.OptCacheDir(testDir))

	err = util.PrepareFileStructure(cfg)
	out := filepath.Join(testDir, "test")
	assert.Nil(err)
	diff := idiff.New(cfg)
	err = diff.Compare(src, ref, out)
	assert.Nil(err)

	exists, err := gnsys.FileExists(out + ".sqlite")
	assert.Nil(err)
	assert.True(exists)
}
