package dsl

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//---------------------------------------------------------------------------

func TestNodes(t *testing.T) {
	assert := assert.New(t)

	t1 := NewNodeArrayType(NewNodeIntType(), 17)

	assert.Equal("ARRAY(17, INT)", t1.String())

	t2 := &NodeStruct{
		Fields: map[string]Node{
			"aa": NewNodeBoolType(),
			"bb": NewNodeBoolType(),
			"cc": NewNodeBoolType(),
		},
	}

	t2s := t2.String()
	assert.Contains(t2s, "STRUCT(")
	assert.Contains(t2s, "aa")
	assert.Contains(t2s, "bb")
	assert.Contains(t2s, "cc")
	assert.Contains(t2s, ")")

	t3 := NewNodeMapType(NewNodeStringType(), NewNodeBoolType())

	assert.Equal("MAP[STRING]BOOL", t3.String())

	t4 := NewNodeSliceType(NewNodeIntType())

	assert.Equal("SLICE(INT)", t4.String())
}

func TestNodeEquality(t *testing.T) {
	assert := assert.New(t)

	//eq := func(a, b Node) bool {
	//	return a == b
	//}

	a := NewNodeSliceType(NewNodeIntType())
	b := NewNodeSliceType(NewNodeIntType())
	c := NewNodeStringType()
	d := NewNodeFloatType()
	assert.Equal(a, b)
	assert.NotEqual(a, c)
	assert.NotEqual(a, d)
	assert.NotEqual(c, d)
	assert.EqualValues(a, b)
}
