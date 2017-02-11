package dsl

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTypeTokenizer(t *testing.T) {
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

	tt, err := NewTypeTokenizer()
	assert.NoError(err)
	typeTable, err := tt.ParseJson(s)
	assert.NoError(err)
	assert.NotNil(typeTable)

	data := map[string][]Token{
		// the built-in types don't have token trees, so skip those

		// these are the types from our decl block above

		"MyIntU8":   []Token{Token{Line: 1, Column: 6, Id: 21, Text: "uint8"}},
		"MyIntU16":  []Token{Token{Line: 1, Column: 7, Id: 21, Text: "uint16"}},
		"MyIntU32":  []Token{Token{Line: 1, Column: 7, Id: 21, Text: "uint32"}},
		"MyIntU64":  []Token{Token{Line: 1, Column: 7, Id: 21, Text: "uint64"}},
		"MyInt8":    []Token{Token{Line: 1, Column: 5, Id: 21, Text: "int8"}},
		"MyInt16":   []Token{Token{Line: 1, Column: 6, Id: 21, Text: "int16"}},
		"MyInt32":   []Token{Token{Line: 1, Column: 6, Id: 21, Text: "int32"}},
		"MyInt64":   []Token{Token{Line: 1, Column: 6, Id: 21, Text: "int64"}},
		"MyFloat32": []Token{Token{Line: 1, Column: 8, Id: 21, Text: "float32"}},
		"MyFloat64": []Token{Token{Line: 1, Column: 8, Id: 21, Text: "float64"}},
		"MyBool":    []Token{Token{Line: 1, Column: 5, Id: 21, Text: "bool"}},
		"MyString":  []Token{Token{Line: 1, Column: 7, Id: 21, Text: "string"}},
		"MyAny":     []Token{Token{Line: 1, Column: 4, Id: 21, Text: "any"}},
		"MyMapInt32": []Token{
			Token{Line: 1, Column: 2, Id: 25, Text: "[map]"},
			Token{Line: 1, Column: 11, Id: 21, Text: "int32"},
		},
		"MyMapPoint": []Token{
			Token{Line: 1, Column: 2, Id: 25, Text: "[map]"},
			Token{Line: 1, Column: 11, Id: 21, Text: "Point"},
		},
		"MySliceIntU8": []Token{
			Token{Line: 1, Column: 2, Id: 23, Text: "[]"},
			Token{Line: 1, Column: 8, Id: 21, Text: "uint8"},
		},
		"MySlicePoint": []Token{
			Token{Line: 1, Column: 2, Id: 23, Text: "[]"},
			Token{Line: 1, Column: 8, Id: 21, Text: "Point"},
		},
		"MyArray10Float32": []Token{
			Token{Line: 1, Column: 2, Id: 24, Text: "[10]", Value: 10},
			Token{Line: 1, Column: 12, Id: 21, Text: "float32"},
		},
		"MyArray4Point": []Token{
			Token{Line: 1, Column: 2, Id: 24, Text: "[4]", Value: 4},
			Token{Line: 1, Column: 9, Id: 21, Text: "Point"},
		},
		"Point":          nil,
		"Point.x32":      []Token{Token{Line: 1, Column: 8, Id: 21, Text: "float32"}},
		"Point.y64":      []Token{Token{Line: 1, Column: 8, Id: 21, Text: "float64"}},
		"MyStruct":       nil,
		"MyStruct.alpha": []Token{Token{Line: 1, Column: 7, Id: 21, Text: "string"}},
		"MyStruct.beta":  []Token{Token{Line: 1, Column: 6, Id: 21, Text: "Point"}},
		"MyStruct.gamma": []Token{Token{Line: 1, Column: 9, Id: 21, Text: "MyStruct"}},
		"MyStruct.delta": []Token{Token{Line: 1, Column: 4, Id: 21, Text: "any"}},
	}

	//	assert.Len(typeTable.Types, len(data))

	for name, token := range data {
		if typeTable.isBuiltin(name) {
			continue
		}

		tte, ok := typeTable.Types[name]
		assert.True(ok)
		assert.Equal(name, tte.Name)
		assert.EqualValues(token, tte.Token)
	}
}
