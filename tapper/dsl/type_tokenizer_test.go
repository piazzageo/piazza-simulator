package dsl

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// this is a DeclBlock, containing both struct decls and string decls
var typeTestString = `
{ 
    "MyInt": "int",
    "MyFloat": "float",
    "MyBool": "bool",
    "MyString": "string",
    "MyAny": "any",

    "MyMapInt": "[map]int",
    "MyMapPoint": "[map]Point",
    "MySliceInt": "[]int",
    "MySlicePoint": "[]Point",
    "MyArray10Float": "[10]float",
    "MyArray4Point": "[4]Point",

    "Point": {
        "x": "float",
        "y": "float"
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
		name:  "MyInt",
		token: []Token{Token{Line: 1, Column: 4, Id: 21, Text: "int"}},
		node:  NewNodeIntType(),
	},
	typeTestItem{
		name:  "MyFloat",
		token: []Token{Token{Line: 1, Column: 6, Id: 21, Text: "float"}},
		node:  NewNodeFloatType(),
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
		name: "MyMapInt",
		token: []Token{
			Token{Line: 1, Column: 2, Id: 25, Text: "[map]"},
			Token{Line: 1, Column: 9, Id: 21, Text: "int"},
		},
		node: NewNodeMapType(NewNodeStringType(), NewNodeIntType()),
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
		name: "MySliceInt",
		token: []Token{
			Token{Line: 1, Column: 2, Id: 23, Text: "[]"},
			Token{Line: 1, Column: 6, Id: 21, Text: "int"},
		},
		node: NewNodeSliceType(NewNodeIntType()),
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
		name: "MyArray10Float",
		token: []Token{
			Token{Line: 1, Column: 2, Id: 24, Text: "[10]", Value: 10},
			Token{Line: 1, Column: 10, Id: 21, Text: "float"},
		},
		node: NewNodeArrayType(NewNodeFloatType(), 10),
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
		name:  "Point.x",
		token: []Token{Token{Line: 1, Column: 6, Id: 21, Text: "float"}},
		node:  NewNodeFloatType(),
	},
	typeTestItem{
		name:  "Point.y",
		token: []Token{Token{Line: 1, Column: 6, Id: 21, Text: "float"}},
		node:  NewNodeFloatType(),
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

	//log.Printf("===== %#v ====", typeTable)

	for _, item := range typeTestData {
		if typeTable.isBuiltin(item.name) {
			continue
		}

		//log.Printf("===== %s ====", item.name)

		tte, ok := typeTable.Types[item.name]
		assert.True(ok)
		assert.Equal(item.name, tte.Name)
		assert.EqualValues(item.token, tte.Token)
	}
}
