package symtab

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
		ElemType: &IntDslType{Form: U8},
		Len:      17,
	}

	assert.Equal("ARRAY(17, INT(U8))", t1.String())

	t2 := &StructDslType{
		KVPairs: map[string]DslType{
			"i": &IntDslType{Form: S32},
			"s": &StringDslType{},
			"f": &FloatDslType{},
			"b": &BoolDslType{},
			"a": &AnyDslType{},
		},
	}

	t2s := t2.String()
	assert.Contains(t2s, "STRUCT(")
	assert.Contains(t2s, "i:INT(S32)")
	assert.Contains(t2s, "s:STRING")
	assert.Contains(t2s, "f:FLOAT")
	assert.Contains(t2s, "b:BOOL")
	assert.Contains(t2s, "a:ANY")
	assert.Contains(t2s, ")")

	t3 := &MapDslType{
		KeyType:   &StringDslType{},
		ValueType: &BoolDslType{},
	}

	assert.Equal("MAP[STRING]BOOL", t3.String())

	t4 := &SliceDslType{
		ElemType: &IntDslType{Form: U32},
	}

	assert.Equal("SLICE(INT(U32))", t4.String())
}
