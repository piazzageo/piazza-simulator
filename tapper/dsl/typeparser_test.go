package dsl

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test21(t *testing.T) {
	assert := assert.New(t)

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

	p := &TypeParser{}
	err := p.ParseJson(s)
	assert.NoError(err)
}
