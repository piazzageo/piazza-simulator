package dsl

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//--------------------------

func TestTypeTable(t *testing.T) {
	assert := assert.New(t)
	var err error

	tt, err := NewTypeTable()
	assert.NoError(err)
	assert.Equal(4, tt.size())

	// add a symbol
	err = tt.addNode("myint", NewTypeNodeInt())
	assert.NoError(err)
	assert.Equal(5, tt.size())

	// add a symbol again
	err = tt.addNode("myint", NewTypeNodeInt())
	assert.Error(err)
	assert.Equal(5, tt.size())

	// test has()
	assert.True(tt.has("myint"))
	assert.False(tt.has("foofoofoo"))
}
