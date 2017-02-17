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
	name string
	node TypeNode
}

var typeTestData = []typeTestItem{
	typeTestItem{
		name: "MyInt",
		node: NewTypeNodeInt(),
	},
	typeTestItem{
		name: "MyFloat",
		node: NewTypeNodeFloat(),
	},
	typeTestItem{
		name: "MyBool",
		node: NewTypeNodeBool(),
	},
	typeTestItem{
		name: "MyString",
		node: NewTypeNodeString(),
	},
	typeTestItem{
		name: "MyMapInt",
		node: NewTypeNodeMap(NewTypeNodeString(), NewTypeNodeInt()),
	},
	typeTestItem{
		name: "MyMapPoint",
		node: NewTypeNodeMap(NewTypeNodeString(), NewTypeNodeName("Point")),
	},
	typeTestItem{
		name: "MySliceInt",
		node: NewTypeNodeSlice(NewTypeNodeInt()),
	},
	typeTestItem{
		name: "MySlicePoint",
		node: NewTypeNodeSlice(NewTypeNodeName("Point")),
	},
	typeTestItem{
		name: "MyArray10Float",
		node: NewTypeNodeArray(NewTypeNodeFloat(), 10),
	},
	typeTestItem{
		name: "MyArray4Point",
		node: NewTypeNodeArray(NewTypeNodeName("Point"), 4),
	},
	typeTestItem{
		name: "Point",
		node: &TypeNodeStruct{
			Fields: map[string]TypeNode{
				"x": NewTypeNodeFloat(),
				"y": NewTypeNodeFloat(),
			},
		},
	},
	typeTestItem{
		name: "MyStruct",
		node: &TypeNodeStruct{
			Fields: map[string]TypeNode{
				"alpha": NewTypeNodeString(),
				"beta":  NewTypeNodeName("Point"),
				"gamma": NewTypeNodeName("MyStruct"),
			},
		},
	},
}

func TestTypeParser(t *testing.T) {
	assert := assert.New(t)

	tt := TypeTokenizer{}
	typeTable, err := tt.ParseJson(typeTestString)
	assert.NoError(err)
	assert.NotNil(typeTable)

	//log.Printf("===== %#v ====", typeTable)

	for _, item := range typeTestData {
		if typeTable.isBuiltin(item.name) {
			continue
		}

		//	log.Printf("===== %s ====", item.name)

		tte, ok := typeTable.Types[item.name]
		assert.True(ok)
		assert.Equal(item.name, tte.Name)
	}
}
