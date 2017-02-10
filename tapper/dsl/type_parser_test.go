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
		"uint8":   &NodeNumberType{Flavor: FlavorU8},
		"uint16":  &NodeNumberType{Flavor: FlavorU16},
		"uint32":  &NodeNumberType{Flavor: FlavorU32},
		"uint64":  &NodeNumberType{Flavor: FlavorU64},
		"int8":    &NodeNumberType{Flavor: FlavorS8},
		"int16":   &NodeNumberType{Flavor: FlavorS16},
		"int32":   &NodeNumberType{Flavor: FlavorS32},
		"int64":   &NodeNumberType{Flavor: FlavorS64},
		"float32": &NodeNumberType{Flavor: FlavorF32},
		"float64": &NodeNumberType{Flavor: FlavorF64},
		"bool":    &NodeBoolType{},
		"string":  &NodeStringType{},
		"any":     &NodeAnyType{},

		// the types from our decl block above
		"MyIntU8":          &NodeNumberType{Flavor: FlavorU8},
		"MyIntU16":         &NodeNumberType{Flavor: FlavorU16},
		"MyIntU32":         &NodeNumberType{Flavor: FlavorU32},
		"MyIntU64":         &NodeNumberType{Flavor: FlavorU64},
		"MyInt8":           &NodeNumberType{Flavor: FlavorS8},
		"MyInt16":          &NodeNumberType{Flavor: FlavorS16},
		"MyInt32":          &NodeNumberType{Flavor: FlavorS32},
		"MyInt64":          &NodeNumberType{Flavor: FlavorS64},
		"MyFloat32":        &NodeNumberType{Flavor: FlavorF32},
		"MyFloat64":        &NodeNumberType{Flavor: FlavorF64},
		"MyBool":           &NodeBoolType{},
		"MyString":         &NodeStringType{},
		"MyAny":            &NodeAnyType{},
		"MyMapInt32":       &NodeMapType{KeyType: &NodeStringType{}, ValueType: &NodeNumberType{Flavor: FlavorS32}},
		"MyMapPoint":       &NodeMapType{KeyType: &NodeStringType{}, ValueType: &NodeUserType{Name: "Point"}},
		"MySliceIntU8":     &NodeSliceType{ElemType: &NodeNumberType{Flavor: FlavorU8}},
		"MySlicePoint":     &NodeSliceType{ElemType: &NodeUserType{Name: "Point"}},
		"MyArray10Float32": &NodeArrayType{Len: 10, ElemType: &NodeNumberType{Flavor: FlavorF32}},
		"MyArray4Point":    &NodeArrayType{Len: 4, ElemType: &NodeUserType{Name: "Point"}},
		"Point":            &NodeStructType{Fields: map[string]bool{"x32": true, "y64": true}},
		"Point.x32":        &NodeNumberType{Flavor: FlavorF32},
		"Point.y64":        &NodeNumberType{Flavor: FlavorF64},
		"MyStruct":         &NodeStructType{Fields: map[string]bool{"alpha": true, "beta": true, "gamma": true, "delta": true}},
		"MyStruct.alpha":   &NodeStringType{},
		"MyStruct.beta":    &NodeUserType{Name: "Point"},
		"MyStruct.gamma":   &NodeUserType{Name: "MyStruct"},
		"MyStruct.delta":   &NodeAnyType{},
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
