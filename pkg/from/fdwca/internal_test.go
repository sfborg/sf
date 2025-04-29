package fdwca

import (
	"testing"

	"github.com/sfborg/sflib/pkg/dwca"
	"github.com/stretchr/testify/assert"
)

func TestFieldMap(t *testing.T) {
	assert := assert.New(t)
	fields := []dwca.Field{
		{Index: "0", Idx: 0, Term: "CapField"},
		{Index: "1", Idx: 1, Term: "lowcase"},
		{Index: "2", Idx: 2, Term: "UPPERCASE"},
		{Index: "3", Idx: 3, Term: "some:Prefix"},
		{Index: "4", Idx: 4, Term: "A:NewPrefix"},
	}

	res := fieldsMap(fields)

	idx, ok := res["unknown"]
	assert.False(ok)
	assert.Equal(idx, 0)

	idx, ok = res["capfield"]
	assert.True(ok)
	assert.Equal(idx, 0)
	assert.Equal(1, res["lowcase"])
	assert.Equal(2, res["uppercase"])
	assert.Equal(3, res["prefix"])
	assert.Equal(4, res["newprefix"])
}
