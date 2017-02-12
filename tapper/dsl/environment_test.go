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
	err = tt.add("i")
	assert.NoError(err)
	err = tt.setNode("i", NewNodeIntType())
	assert.NoError(err)
	err = tt.add("f")
	assert.NoError(err)
	err = tt.setNode("f", NewNodeFloatType())
	assert.NoError(err)

	env := NewEnvironment(tt)

	env.set("i", 12)
	assert.Equal(12, env.get("i"))
}
