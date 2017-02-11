package dsl

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// this is a DeclBlock, containing both struct decls and string decls
var typeTestString = `
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

type typeTestItem struct {
	name  string
	token []Token
	node  Node
}

var typeTestData = []typeTestItem{
	typeTestItem{
		name:  "MyIntU8",
		token: []Token{Token{Line: 1, Column: 6, Id: 21, Text: "uint8"}},
		node:  NewNodeNumberType(FlavorU8),
	},
	typeTestItem{
		name:  "MyIntU16",
		token: []Token{Token{Line: 1, Column: 7, Id: 21, Text: "uint16"}},
		node:  NewNodeNumberType(FlavorU16),
	},
	typeTestItem{
		name:  "MyIntU32",
		token: []Token{Token{Line: 1, Column: 7, Id: 21, Text: "uint32"}},
		node:  NewNodeNumberType(FlavorU32),
	},
	typeTestItem{
		name:  "MyIntU64",
		token: []Token{Token{Line: 1, Column: 7, Id: 21, Text: "uint64"}},
		node:  NewNodeNumberType(FlavorU64),
	},
	typeTestItem{
		name:  "MyInt8",
		token: []Token{Token{Line: 1, Column: 5, Id: 21, Text: "int8"}},
		node:  NewNodeNumberType(FlavorS8),
	},
	typeTestItem{
		name:  "MyInt16",
		token: []Token{Token{Line: 1, Column: 6, Id: 21, Text: "int16"}},
		node:  NewNodeNumberType(FlavorS16),
	},
	typeTestItem{
		name:  "MyInt32",
		token: []Token{Token{Line: 1, Column: 6, Id: 21, Text: "int32"}},
		node:  NewNodeNumberType(FlavorS32),
	},
	typeTestItem{
		name:  "MyInt64",
		token: []Token{Token{Line: 1, Column: 6, Id: 21, Text: "int64"}},
		node:  NewNodeNumberType(FlavorS64),
	},
	typeTestItem{
		name:  "MyFloat32",
		token: []Token{Token{Line: 1, Column: 8, Id: 21, Text: "float32"}},
		node:  NewNodeNumberType(FlavorF32),
	},
	typeTestItem{
		name:  "MyFloat64",
		token: []Token{Token{Line: 1, Column: 8, Id: 21, Text: "float64"}},
		node:  NewNodeNumberType(FlavorF64),
	},
	typeTestItem{
		name:  "MyBool",
		token: []Token{Token{Line: 1, Column: 5, Id: 21, Text: "bool"}},
		node:  NewNodeBoolType(),
	},
	typeTestItem{
		name:  "MyString",
		token: []Token{Token{Line: 1, Column: 7, Id: 21, Text: "string"}},
		node:  NewNodeStringType(),
	},
	typeTestItem{
		name:  "MyAny",
		token: []Token{Token{Line: 1, Column: 4, Id: 21, Text: "any"}},
		node:  NewNodeAnyType(),
	},
	typeTestItem{
		name: "MyMapInt32",
		token: []Token{
			Token{Line: 1, Column: 2, Id: 25, Text: "[map]"},
			Token{Line: 1, Column: 11, Id: 21, Text: "int32"},
		},
		node: NewNodeMapType(NewNodeStringType(), NewNodeNumberType(FlavorS32)),
	},
	typeTestItem{
		name: "MyMapPoint",
		token: []Token{
			Token{Line: 1, Column: 2, Id: 25, Text: "[map]"},
			Token{Line: 1, Column: 11, Id: 21, Text: "Point"},
		},
		node: NewNodeMapType(NewNodeStringType(), NewNodeUserType("Point")),
	},
	typeTestItem{
		name: "MySliceIntU8",
		token: []Token{
			Token{Line: 1, Column: 2, Id: 23, Text: "[]"},
			Token{Line: 1, Column: 8, Id: 21, Text: "uint8"},
		},
		node: NewNodeSliceType(NewNodeNumberType(FlavorU8)),
	},
	typeTestItem{
		name: "MySlicePoint",
		token: []Token{
			Token{Line: 1, Column: 2, Id: 23, Text: "[]"},
			Token{Line: 1, Column: 8, Id: 21, Text: "Point"},
		},
		node: NewNodeSliceType(NewNodeUserType("Point")),
	},
	typeTestItem{
		name: "MyArray10Float32",
		token: []Token{
			Token{Line: 1, Column: 2, Id: 24, Text: "[10]", Value: 10},
			Token{Line: 1, Column: 12, Id: 21, Text: "float32"},
		},
		node: NewNodeArrayType(NewNodeNumberType(FlavorF32), 10),
	},
	typeTestItem{
		name: "MyArray4Point",
		token: []Token{
			Token{Line: 1, Column: 2, Id: 24, Text: "[4]", Value: 4},
			Token{Line: 1, Column: 9, Id: 21, Text: "Point"},
		},
		node: NewNodeArrayType(NewNodeUserType("Point"), 4),
	},
	typeTestItem{
		name:  "Point.x32",
		token: []Token{Token{Line: 1, Column: 8, Id: 21, Text: "float32"}},
		node:  NewNodeNumberType(FlavorF32),
	},
	typeTestItem{
		name:  "Point.y64",
		token: []Token{Token{Line: 1, Column: 8, Id: 21, Text: "float64"}},
		node:  NewNodeNumberType(FlavorF64),
	},
	typeTestItem{
		name:  "MyStruct",
		token: nil,
		node:  NewNodeStructType(map[string]bool{"alpha": true, "beta": true, "gamma": true, "delta": true}),
	},
	typeTestItem{
		name:  "MyStruct.alpha",
		token: []Token{Token{Line: 1, Column: 7, Id: 21, Text: "string"}},
		node:  NewNodeStringType(),
	},
	typeTestItem{
		name:  "MyStruct.beta",
		token: []Token{Token{Line: 1, Column: 6, Id: 21, Text: "Point"}},
		node:  NewNodeUserType("Point"),
	},
	typeTestItem{
		name:  "MyStruct.gamma",
		token: []Token{Token{Line: 1, Column: 9, Id: 21, Text: "MyStruct"}},
		node:  NewNodeUserType("MyStruct"),
	},
	typeTestItem{
		name:  "MyStruct.delta",
		token: []Token{Token{Line: 1, Column: 4, Id: 21, Text: "any"}},
		node:  NewNodeAnyType(),
	},
}

func TestTypeTokenizer(t *testing.T) {
	assert := assert.New(t)

	tt, err := NewTypeTokenizer()
	assert.NoError(err)
	typeTable, err := tt.ParseJson(typeTestString)
	assert.NoError(err)
	assert.NotNil(typeTable)

	//	assert.Len(typeTable.Types, len(data))

	for _, item := range typeTestData {
		if typeTable.isBuiltin(item.name) {
			continue
		}

		tte, ok := typeTable.Types[item.name]
		assert.True(ok)
		assert.Equal(item.name, tte.Name)
		assert.EqualValues(item.token, tte.Token)
	}
}
