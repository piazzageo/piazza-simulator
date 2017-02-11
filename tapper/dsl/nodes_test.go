package dsl

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//---------------------------------------------------------------------------

func TestNodes(t *testing.T) {
	assert := assert.New(t)

	t1 := NewNodeArrayType(NewNodeNumberType(FlavorU8), 17)

	assert.Equal("ARRAY(17, NUMBER(U8))", t1.String())

	t2 := NewNodeStructType(map[string]bool{"aa": true, "bb": true, "cc": true})

	t2s := t2.String()
	assert.Contains(t2s, "STRUCT(")
	assert.Contains(t2s, "aa")
	assert.Contains(t2s, "bb")
	assert.Contains(t2s, "cc")
	assert.Contains(t2s, ")")

	t3 := NewNodeMapType(NewNodeStringType(), NewNodeBoolType())

	assert.Equal("MAP[STRING]BOOL", t3.String())

	t4 := NewNodeSliceType(NewNodeNumberType(FlavorU32))

	assert.Equal("SLICE(NUMBER(U32))", t4.String())
}

func TestNodeEquality(t *testing.T) {
	assert := assert.New(t)

	//eq := func(a, b Node) bool {
	//	return a == b
	//}

	a := NewNodeSliceType(NewNodeNumberType(FlavorU32))
	b := NewNodeSliceType(NewNodeNumberType(FlavorU32))
	c := NewNodeStringType()
	d := NewNodeSliceType(NewNodeNumberType(FlavorS32))
	assert.Equal(a, b)
	assert.NotEqual(a, c)
	assert.NotEqual(a, d)
	assert.EqualValues(a, b)
}
