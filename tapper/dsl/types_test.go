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
		Fields: map[string]bool{"aa": true, "bb": true, "cc": true},
	}

	t2s := t2.String()
	assert.Contains(t2s, "STRUCT(")
	assert.Contains(t2s, "aa")
	assert.Contains(t2s, "bb")
	assert.Contains(t2s, "cc")
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
