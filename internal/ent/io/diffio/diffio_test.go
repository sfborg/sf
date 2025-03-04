package diffio_test

import (
	"testing"

	"github.com/sfborg/sf/internal/ent/io/diffio"
	"github.com/sfborg/sf/internal/ent/io/sysio"
	"github.com/sfborg/sf/pkg/config"
	"github.com/stretchr/testify/assert"
)

func TestCompare(t *testing.T) {
	assert := assert.New(t)

	src := "../../../../testdata/diff/test-a.sqlite"
	ref := "../../../../testdata/diff/test-b.sqlite"
	cfg := config.New()
	sysio.PrepareFileStructure(cfg)
	diff := diffio.New(cfg)
	err := diff.Compare(src, ref)
	assert.Nil(err)
}
