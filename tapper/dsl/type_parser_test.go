package dsl

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test61(t *testing.T) {
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

	data := map[string]TNode{
		// the built-in types
		"uint8":   &TNodeNumber{Flavor: FlavorU8},
		"uint16":  &TNodeNumber{Flavor: FlavorU16},
		"uint32":  &TNodeNumber{Flavor: FlavorU32},
		"uint64":  &TNodeNumber{Flavor: FlavorU64},
		"int8":    &TNodeNumber{Flavor: FlavorS8},
		"int16":   &TNodeNumber{Flavor: FlavorS16},
		"int32":   &TNodeNumber{Flavor: FlavorS32},
		"int64":   &TNodeNumber{Flavor: FlavorS64},
		"float32": &TNodeNumber{Flavor: FlavorF32},
		"float64": &TNodeNumber{Flavor: FlavorF64},
		"bool":    &TNodeBool{},
		"string":  &TNodeString{},
		"any":     &TNodeAny{},

		// the types from our decl block above
		"MyIntU8":          &TNodeNumber{Flavor: FlavorU8},
		"MyIntU16":         &TNodeNumber{Flavor: FlavorU16},
		"MyIntU32":         &TNodeNumber{Flavor: FlavorU32},
		"MyIntU64":         &TNodeNumber{Flavor: FlavorU64},
		"MyInt8":           &TNodeNumber{Flavor: FlavorS8},
		"MyInt16":          &TNodeNumber{Flavor: FlavorS16},
		"MyInt32":          &TNodeNumber{Flavor: FlavorS32},
		"MyInt64":          &TNodeNumber{Flavor: FlavorS64},
		"MyFloat32":        &TNodeNumber{Flavor: FlavorF32},
		"MyFloat64":        &TNodeNumber{Flavor: FlavorF64},
		"MyBool":           &TNodeBool{},
		"MyString":         &TNodeString{},
		"MyAny":            &TNodeAny{},
		"MyMapInt32":       &TNodeMap{KeyType: &TNodeString{}, ValueType: &TNodeNumber{Flavor: FlavorS32}},
		"MyMapPoint":       &TNodeMap{KeyType: &TNodeString{}, ValueType: &TNodeUserType{Name: "Point"}},
		"MySliceIntU8":     &TNodeSlice{ElemType: &TNodeNumber{Flavor: FlavorU8}},
		"MySlicePoint":     &TNodeSlice{ElemType: &TNodeUserType{Name: "Point"}},
		"MyArray10Float32": &TNodeArray{Len: 10, ElemType: &TNodeNumber{Flavor: FlavorF32}},
		"MyArray4Point":    &TNodeArray{Len: 4, ElemType: &TNodeUserType{Name: "Point"}},
		"Point":            &TNodeStruct{Fields: map[string]bool{"x32": true, "y64": true}},
		"Point.x32":        &TNodeNumber{Flavor: FlavorF32},
		"Point.y64":        &TNodeNumber{Flavor: FlavorF64},
		"MyStruct":         &TNodeStruct{Fields: map[string]bool{"alpha": true, "beta": true, "gamma": true, "delta": true}},
		"MyStruct.alpha":   &TNodeString{},
		"MyStruct.beta":    &TNodeUserType{Name: "Point"},
		"MyStruct.gamma":   &TNodeUserType{Name: "MyStruct"},
		"MyStruct.delta":   &TNodeAny{},
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
