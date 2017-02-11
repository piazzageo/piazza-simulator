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
	typeTable, err := p.ParseJson(s)
	assert.NoError(err)
	assert.NotNil(typeTable)

	data := map[string]Node{
		// the built-in types
		"uint8":   NewNodeNumberType(FlavorU8),
		"uint16":  NewNodeNumberType(FlavorU16),
		"uint32":  NewNodeNumberType(FlavorU32),
		"uint64":  NewNodeNumberType(FlavorU64),
		"int8":    NewNodeNumberType(FlavorS8),
		"int16":   NewNodeNumberType(FlavorS16),
		"int32":   NewNodeNumberType(FlavorS32),
		"int64":   NewNodeNumberType(FlavorS64),
		"float32": NewNodeNumberType(FlavorF32),
		"float64": NewNodeNumberType(FlavorF64),
		"bool":    NewNodeBoolType(),
		"string":  NewNodeStringType(),
		"any":     NewNodeAnyType(),

		// the types from our decl block above
		"MyIntU8":          NewNodeNumberType(FlavorU8),
		"MyIntU16":         NewNodeNumberType(FlavorU16),
		"MyIntU32":         NewNodeNumberType(FlavorU32),
		"MyIntU64":         NewNodeNumberType(FlavorU64),
		"MyInt8":           NewNodeNumberType(FlavorS8),
		"MyInt16":          NewNodeNumberType(FlavorS16),
		"MyInt32":          NewNodeNumberType(FlavorS32),
		"MyInt64":          NewNodeNumberType(FlavorS64),
		"MyFloat32":        NewNodeNumberType(FlavorF32),
		"MyFloat64":        NewNodeNumberType(FlavorF64),
		"MyBool":           NewNodeBoolType(),
		"MyString":         NewNodeStringType(),
		"MyAny":            NewNodeAnyType(),
		"MyMapInt32":       NewNodeMapType(NewNodeStringType(), NewNodeNumberType(FlavorS32)),
		"MyMapPoint":       NewNodeMapType(NewNodeStringType(), NewNodeUserType("Point")),
		"MySliceIntU8":     NewNodeSliceType(NewNodeNumberType(FlavorU8)),
		"MySlicePoint":     NewNodeSliceType(NewNodeUserType("Point")),
		"MyArray10Float32": NewNodeArrayType(NewNodeNumberType(FlavorF32), 10),
		"MyArray4Point":    NewNodeArrayType(NewNodeUserType("Point"), 4),
		"Point":            NewNodeStructType(map[string]bool{"x32": true, "y64": true}),
		"Point.x32":        NewNodeNumberType(FlavorF32),
		"Point.y64":        NewNodeNumberType(FlavorF64),
		"MyStruct":         NewNodeStructType(map[string]bool{"alpha": true, "beta": true, "gamma": true, "delta": true}),
		"MyStruct.alpha":   NewNodeStringType(),
		"MyStruct.beta":    NewNodeUserType("Point"),
		"MyStruct.gamma":   NewNodeUserType("MyStruct"),
		"MyStruct.delta":   NewNodeAnyType(),
	}

	assert.Len(typeTable.Types, len(data))

	dels := []string{}
	for name, tnode := range data {
		//log.Printf("========= %s ================", name)
		v, ok := typeTable.Types[name]
		assert.True(ok)
		assert.Equal(name, v.Name)
		assert.EqualValues(tnode, v.Node)
		dels = append(dels, name)
	}

	for _, v := range dels {
		delete(typeTable.Types, v)
	}
	assert.Len(typeTable.Types, 0)
}
