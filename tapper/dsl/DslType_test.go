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

	t1 := &ArrayDslType{
		ElemType: &NumberDslType{Flavor: U8},
		Len:      17,
	}

	assert.Equal("ARRAY(17, NUMBER(U8))", t1.String())

	t2 := &StructDslType{
		Fields: map[Symbol]DslType{
			"i": &NumberDslType{Flavor: S32},
			"s": &StringDslType{},
			"f": &NumberDslType{Flavor: F32},
			"b": &BoolDslType{},
			"a": &AnyDslType{},
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

	t3 := &MapDslType{
		KeyType:   &StringDslType{},
		ValueType: &BoolDslType{},
	}

	assert.Equal("MAP[STRING]BOOL", t3.String())

	t4 := &SliceDslType{
		ElemType: &NumberDslType{Flavor: U32},
	}

	assert.Equal("SLICE(NUMBER(U32))", t4.String())
}
