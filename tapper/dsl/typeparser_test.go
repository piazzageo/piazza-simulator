package dsl

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"encoding/json"
)

func Test20(t *testing.T) {
	assert := assert.New(t)

	ok, siz := matchArrayPrefix("[]")
	assert.False(ok)
	ok, siz = matchArrayPrefix(" []")
	assert.False(ok)
	ok, siz = matchArrayPrefix("[}")
	assert.False(ok)
	ok, siz = matchArrayPrefix("[3a]")
	assert.False(ok)
	ok, siz = matchArrayPrefix("[3]")
	assert.True(ok)
	assert.Equal(3, siz)
	ok, siz = matchArrayPrefix("[32]")
	assert.True(ok)
	assert.Equal(32, siz)
}

func Test21(t *testing.T) {
	//	assert := assert.New(t)

	// this is a DeclBlock, containing 3 interfaces (2 StructDecls and 1 StringDecl)
	s := `
{ 
    "Point": {
        "x32": "float32",
        "y64": "float64"
    },
    "MyIntU16": "uint16",
    "MyMapStringIntS32": "[map]int32",
    "MyMapStringPoint": "[map]Point",
    "MySlicePoint": "[]Point",
    "MyArrya4Point": "[4]Point",
    "Message": {
        "astring": "string",
        "aPoint": "Point",
        "aMyMapStringPoint": "MyMapStringPoint"
    }
}`

	declBlock := &DeclBlock{}

	err := json.Unmarshal([]byte(s), declBlock)
	if err != nil {
		panic(err)
	}

	p := &TypeParser{}
	p.Parse(declBlock)
}
