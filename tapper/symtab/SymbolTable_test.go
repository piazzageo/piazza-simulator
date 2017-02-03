package symtab

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

	st0 := NewSymbolTable(nil)
	assert.Equal("Parent:false, Children:0", st0.String())

	st0.add("myint", &IntDslType{Form: S32})
	assert.Equal("Parent:false, Children:0 [myint:INT(S32)]", st0.String())

	typ := st0.get("myint")
	assert.Equal("INT(S32)", typ.String())

	assert.Nil(st0.get("frobnitz"))

	// delete should be safe if no such symbol
	st0.remove("frobnitz")

	st0.remove("myint")
	assert.Equal("Parent:false, Children:0", st0.String())

	st1 := st0.addScope()
	assert.Equal("Parent:false, Children:1", st0.String())
	assert.Equal("Parent:true, Children:0", st1.String())
}
