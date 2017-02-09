package dsl

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//--------------------------

func Test00(t *testing.T) {
	assert := assert.New(t)
	assert.True(!false)
}

func Test01(t *testing.T) {
	assert := assert.New(t)

	t1 := &TNodeArray{
		ElemType: &TNodeNumber{Flavor: FlavorU8},
		Len:      17,
	}

	assert.Equal("ARRAY(17, NUMBER(U8))", t1.String())

	t2 := &TNodeStruct{
		Fields: map[string]TNode{
			"i": &TNodeNumber{Flavor: FlavorS32},
			"s": &TNodeString{},
			"f": &TNodeNumber{Flavor: FlavorF32},
			"b": &TNodeBool{},
			"a": &TNodeAny{},
		},
	}

	t2s := t2.String()
	assert.Contains(t2s, "STRUCT(")
	assert.Contains(t2s, "i:NUMBER(S32)")
	assert.Contains(t2s, "s:STRING")
	assert.Contains(t2s, "f:NUMBER(F32)")
	assert.Contains(t2s, "b:BOOL")
	assert.Contains(t2s, "a:ANY")
	assert.Contains(t2s, ")")

	t3 := &TNodeMap{
		KeyType:   &TNodeString{},
		ValueType: &TNodeBool{},
	}

	assert.Equal("MAP[STRING]BOOL", t3.String())

	t4 := &TNodeSlice{
		ElemType: &TNodeNumber{Flavor: FlavorU32},
	}

	assert.Equal("SLICE(NUMBER(U32))", t4.String())
}
