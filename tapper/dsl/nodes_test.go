package dsl

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//---------------------------------------------------------------------------

func Test30(t *testing.T) {
	assert := assert.New(t)

	t1 := NodeArrayType{
		ElemType: &NodeNumberType{Flavor: FlavorU8},
		Len:      17,
	}

	assert.Equal("ARRAY(17, NUMBER(U8))", t1.String())

	t2 := &NodeStructType{
		Fields: map[string]bool{"aa": true, "bb": true, "cc": true},
	}

	t2s := t2.String()
	assert.Contains(t2s, "STRUCT(")
	assert.Contains(t2s, "aa")
	assert.Contains(t2s, "bb")
	assert.Contains(t2s, "cc")
	assert.Contains(t2s, ")")

	t3 := &NodeMapType{
		KeyType:   &NodeStringType{},
		ValueType: &NodeBoolType{},
	}

	assert.Equal("MAP[STRING]BOOL", t3.String())

	t4 := &NodeSliceType{
		ElemType: &NodeNumberType{Flavor: FlavorU32},
	}

	assert.Equal("SLICE(NUMBER(U32))", t4.String())
}
