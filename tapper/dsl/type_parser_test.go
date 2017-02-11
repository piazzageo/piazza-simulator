package dsl

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test50(t *testing.T) {
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
    "MyFloat32": "float32",
    "MyFloat64": "float64",

    "MyBool": "bool",
    "MyString": "string",
    "MyAny": "any",

    "MyMapInt32": "[map]int32",
    "MyMapPoint": "[map]Point",
    "MySliceIntU8": "[]uint8",
    "MySlicePoint": "[]Point",
    "MyArray10Float32": "[10]float32",
    "MyArray4Point": "[4]Point",

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

	p, err := NewTypeTokenizer()
	assert.NoError(err)
	typeTable, err := p.ParseJson(s) REMOVE THIS LINE AND USE THE TESTCASES BELOW DIRECTLY
	assert.NoError(err)
	assert.NotNil(typeTable)

	type item struct {
		name  string
		token []Token
		node  Node
	}

	data := []item{
		item{
			name:  "MyIntU8",
			token: []Token{Token{Line: 1, Column: 6, Id: 21, Text: "uint8"}},
			node:  NewNodeNumberType(FlavorU8),
		},
		item{
			name:  "MyIntU16",
			token: []Token{Token{Line: 1, Column: 7, Id: 21, Text: "uint16"}},
			node:  NewNodeNumberType(FlavorU16),
		},
		item{
			name:  "MyIntU32",
			token: []Token{Token{Line: 1, Column: 7, Id: 21, Text: "uint32"}},
			node:  NewNodeNumberType(FlavorU32),
		},
		item{
			name:  "MyIntU64",
			token: []Token{Token{Line: 1, Column: 7, Id: 21, Text: "uint64"}},
			node:  NewNodeNumberType(FlavorU64),
		},
		item{
			name:  "MyInt8",
			token: []Token{Token{Line: 1, Column: 5, Id: 21, Text: "int8"}},
			node:  NewNodeNumberType(FlavorS8),
		},
		item{
			name:  "MyInt16",
			token: []Token{Token{Line: 1, Column: 6, Id: 21, Text: "int16"}},
			node:  NewNodeNumberType(FlavorS16),
		},
		item{
			name:  "MyInt32",
			token: []Token{Token{Line: 1, Column: 6, Id: 21, Text: "int32"}},
			node:  NewNodeNumberType(FlavorS32),
		},
		item{
			name:  "MyInt64",
			token: []Token{Token{Line: 1, Column: 6, Id: 21, Text: "int64"}},
			node:  NewNodeNumberType(FlavorS64),
		},
		item{
			name:  "MyFloat32",
			token: []Token{Token{Line: 1, Column: 8, Id: 21, Text: "float32"}},
			node:  NewNodeNumberType(FlavorF32),
		},
		item{
			name:  "MyFloat64",
			token: []Token{Token{Line: 1, Column: 8, Id: 21, Text: "float64"}},
			node:  NewNodeNumberType(FlavorF64),
		},
		item{
			name:  "MyBool",
			token: []Token{Token{Line: 1, Column: 5, Id: 21, Text: "bool"}},
			node:  NewNodeBoolType(),
		},
		item{
			name:  "MyString",
			token: []Token{Token{Line: 1, Column: 7, Id: 21, Text: "string"}},
			node:  NewNodeStringType(),
		},
		item{
			name:  "MyAny",
			token: []Token{Token{Line: 1, Column: 4, Id: 21, Text: "any"}},
			node:  NewNodeAnyType(),
		},
		item{
			name: "MyMapInt32",
			token: []Token{
				Token{Line: 1, Column: 2, Id: 25, Text: "[map]"},
				Token{Line: 1, Column: 11, Id: 21, Text: "int32"},
			},
			node: NewNodeMapType(NewNodeStringType(), NewNodeNumberType(FlavorS32)),
		},
		item{
			name: "MyMapPoint",
			token: []Token{
				Token{Line: 1, Column: 2, Id: 25, Text: "[map]"},
				Token{Line: 1, Column: 11, Id: 21, Text: "Point"},
			},
			node: NewNodeMapType(NewNodeStringType(), NewNodeUserType("Point")),
		},
		item{
			name: "MySliceIntU8",
			token: []Token{
				Token{Line: 1, Column: 2, Id: 23, Text: "[]"},
				Token{Line: 1, Column: 8, Id: 21, Text: "uint8"},
			},
			node: NewNodeSliceType(NewNodeNumberType(FlavorU8)),
		},
		item{
			name: "MySlicePoint",
			token: []Token{
				Token{Line: 1, Column: 2, Id: 23, Text: "[]"},
				Token{Line: 1, Column: 8, Id: 21, Text: "Point"},
			},
			node: NewNodeSliceType(NewNodeUserType("Point")),
		},
		item{
			name: "MyArray10Float32",
			token: []Token{
				Token{Line: 1, Column: 2, Id: 24, Text: "[10]", Value: 10},
				Token{Line: 1, Column: 12, Id: 21, Text: "float32"},
			},
			node: NewNodeArrayType(NewNodeNumberType(FlavorF32), 10),
		},
		item{
			name: "MyArray4Point",
			token: []Token{
				Token{Line: 1, Column: 2, Id: 24, Text: "[4]", Value: 4},
				Token{Line: 1, Column: 9, Id: 21, Text: "Point"},
			},
			node: NewNodeArrayType(NewNodeUserType("Point"), 4),
		},
		item{
			name:  "Point.x32",
			token: []Token{Token{Line: 1, Column: 8, Id: 21, Text: "float32"}},
			node:  NewNodeNumberType(FlavorF32),
		},
		item{
			name:  "Point.y64",
			token: []Token{Token{Line: 1, Column: 8, Id: 21, Text: "float64"}},
			node:  NewNodeNumberType(FlavorF64),
		},
		item{
			name: "MyStruct",
			node: NewNodeStructType(map[string]bool{"alpha": true, "beta": true, "gamma": true, "delta": true}),
		},
		item{
			name:  "MyStruct.alpha",
			token: []Token{Token{Line: 1, Column: 7, Id: 21, Text: "string"}},
			node:  NewNodeStringType(),
		},
		item{
			name:  "MyStruct.beta",
			token: []Token{Token{Line: 1, Column: 6, Id: 21, Text: "Point"}},
			node:  NewNodeUserType("Point"),
		},
		item{
			name:  "MyStruct.gamma",
			token: []Token{Token{Line: 1, Column: 9, Id: 21, Text: "MyStruct"}},
			node:  NewNodeUserType("MyStruct"),
		},
		item{
			name:  "MyStruct.delta",
			token: []Token{Token{Line: 1, Column: 4, Id: 21, Text: "any"}},
			node:  NewNodeAnyType(),
		},
	}

	//	assert.Len(typeTable.Types, len(data))

	for _, tc := range data {
		//log.Printf("========= %s ================", name)
		v, ok := typeTable.Types[tc.name]
		assert.True(ok)
		assert.Equal(tc.name, v.Name)
		assert.EqualValues(tc.node, v.Node)
	}
}
