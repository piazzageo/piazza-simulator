package dsl

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//--------------------------

func TestTypeTable(t *testing.T) {
	assert := assert.New(t)
	var err error

	tt := NewTypeTable()
	assert.Equal(0, tt.size())

	// add a symbol
	err = tt.add("myint")
	assert.NoError(err)
	assert.Equal(1, tt.size())

	// fail to add a symbol
	err = tt.add("myint")
	assert.Error(err)

	// set token
	err = tt.setToken("myint", []*Token{&Token{Text: "myint", Id: TokenSymbol}})
	assert.NoError(err)
	assert.Equal([]*Token{&Token{Text: "myint", Id: TokenSymbol}}, tt.getToken("myint"))

	// fail to set token
	err = tt.setToken("myint", []*Token{&Token{Text: "myint", Id: TokenSymbol}})
	assert.Error(err)
	err = tt.setToken("foofoofoo", []*Token{&Token{Text: "myint", Id: TokenSymbol}})
	assert.Error(err)

	// set node
	err = tt.setNode("myint", NewNodeNumberType(FlavorS32))
	assert.NoError(err)
	assert.Equal(NewNodeNumberType(FlavorS32), tt.getNode("myint"))

	// fail to set node
	err = tt.setNode("myint", NewNodeNumberType(FlavorS32))
	assert.Error(err)
	err = tt.setNode("barbarbar", NewNodeNumberType(FlavorS32))
	assert.Error(err)

	// test has()
	assert.True(tt.has("myint"))
	assert.False(tt.has("foofoofoo"))
}
