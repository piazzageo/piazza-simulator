package dsl

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test21(t *testing.T) {
	assert := assert.New(t)

	// this is a DeclBlock, containing both struct decls and string decls
	s := `
{ 
    "MyIntU8": "uint8",
    "MyIntU16": "uint16",
    "MyIntU32": "uint32",
    "MyIntU64": "uint64",
    "MyInt8": "int8",
    "MyInt16": "int16",
    "MyInt32": "int32",
    "MyInt64": "int64",
    "MyFloatt32": "float32",
    "MyFloat64": "float64",

    "MyBool": "bool",
    "MyString": "string",
    "MyAny": "any",

    "MyMapInt32": "[map]int32",
    "MyMapPoint": "[map]Point",
    "MySliceIntU8": "[]uint8",
    "MySlicePoint": "[]Point",
    "MyArray10Float32": "[4]float32",
    "MyArrya4Point": "[4]Point",

    "Point": {
        "x32": "float32",
        "y64": "float64"
    },

    "MyStruct": {
        "alpha": "string",
        "beta": "Point",
        "gamma": "MyStruct",
        "delta": "any"
    }
}`

	p := NewTypeParser()
	err := p.ParseJson(s)
	assert.NoError(err)

	log.Printf("TT: %v", p.typeTable)
}
