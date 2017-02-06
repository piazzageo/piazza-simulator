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

	st0 := NewSymbolTable()
	assert.Equal("0:", st0.String())

	st0.add("myint", &NumberType{Flavor: FlavorS32})
	assert.Equal("1: [myint:NUMBER(S32)]", st0.String())

	assert.True(st0.has("myint"))

	typ := st0.get("myint")
	assert.Equal("NUMBER(S32)", typ.String())

	assert.Nil(st0.get("frobnitz"))
	assert.False(st0.has("frobnitz"))

	// delete should be safe if no such symbol
	st0.remove("frobnitz")

	st0.remove("myint")
	assert.Equal("0:", st0.String())
}
