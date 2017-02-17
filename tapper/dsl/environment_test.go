package dsl

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//---------------------------------------------------------------------------

func TestEnvironment(t *testing.T) {
	assert := assert.New(t)

	// setup
	tt, err := NewTypeTable()
	assert.NoError(err)
	err = tt.addNode("i", NewTypeNodeInt())
	assert.NoError(err)
	err = tt.addNode("f", NewTypeNodeFloat())
	assert.NoError(err)

	env := NewEnvironment(tt)

	env.set("i", &ExprValue{Type: IntType, Value: 12})
	assert.Equal(&ExprValue{Type: IntType, Value: 12}, env.get("i"))
}
