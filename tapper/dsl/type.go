package dsl

import "fmt"

//---------------------------------------------------------------------

type Flavor int

const (
	FlavorU8 Flavor = iota
	FlavorU16
	FlavorU32
	FlavorU64
	FlavorS8
	FlavorS16
	FlavorS32
	FlavorS64
	FlavorF32
	FlavorF64
)

var flavorStrings map[Flavor]string

func init() {
	flavorStrings = map[Flavor]string{
		FlavorU8:  "U8",
		FlavorU16: "U16",
		FlavorU32: "U32",
		FlavorU64: "U64",
		FlavorS8:  "S8",
		FlavorS16: "S16",
		FlavorS32: "S32",
		FlavorS64: "S64",
		FlavorF32: "F32",
		FlavorF64: "F64",
	}
}

func (f Flavor) String() string {
	return flavorStrings[f]
}

//---------------------------------------------------------------------

type Type interface {
	String() string
}

type SymbolType struct {
	Symbol Symbol
}

func (t *SymbolType) String() string {
	return fmt.Sprintf("SYM(%v)", t.Symbol)
}

type MapType struct {
	KeyType   Type
	ValueType Type
}

func (t *MapType) String() string {
	return fmt.Sprintf("MAP[%v]%v", t.KeyType, t.ValueType)
}

type StructType struct {
	Fields map[Symbol]Type
}

func (t *StructType) String() string {
	s := ""
	for k, v := range t.Fields {
		if s != "" {
			s += ", "
		}
		s += fmt.Sprintf("%s:%v", k, v)
	}
	return fmt.Sprintf("STRUCT(%s)", s)
}

type ArrayType struct {
	ElemType Type
	Len      int
}

func (t *ArrayType) String() string {
	return fmt.Sprintf("ARRAY(%d, %v)", t.Len, t.ElemType)
}

type SliceType struct {
	ElemType Type
}

func (t *SliceType) String() string {
	return fmt.Sprintf("SLICE(%v)", t.ElemType)
}

type NumberType struct {
	Flavor Flavor
}

func (t *NumberType) String() string {
	return fmt.Sprintf("NUMBER(%v)", t.Flavor)
}

type BoolType struct {
}

func (t *BoolType) String() string {
	return "BOOL"
}

type StringType struct {
}

func (t *StringType) String() string {
	return "STRING"
}

type AnyType struct {
}

func (t *AnyType) String() string {
	return "ANY"
}
