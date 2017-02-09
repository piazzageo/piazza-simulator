package dsl

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//--------------------------

func Test10(t *testing.T) {
	assert := assert.New(t)
	assert.True(!false)
}

func Test11(t *testing.T) {
	assert := assert.New(t)
	var err error

	st0 := NewTypeTable()
	assert.Equal("size: 0\n", st0.String())

	err = st0.add("myint")
	assert.NoError(err)
	err = st0.set("myint", &TNodeNumber{Flavor: FlavorS32})
	assert.NoError(err)
	assert.Equal("size: 1\n  myint:[NUMBER(S32)]\n", st0.String())

	assert.True(st0.has("myint"))

	typ := st0.get("myint")
	assert.Equal("myint:[NUMBER(S32)]", typ.String())

	assert.Nil(st0.get("frobnitz"))
	assert.False(st0.has("frobnitz"))
}
