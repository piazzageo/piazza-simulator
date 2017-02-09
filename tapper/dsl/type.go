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

type TNode interface {
	String() string
}

//type TNodeSymbol struct {
//	Symbol string
//}
//
//func (t *TNodeSymbol) String() string {
//	return fmt.Sprintf("SYM(%v)", t.Symbol)
//}

type TNodeMap struct {
	KeyType   TNode
	ValueType TNode
}

func (t *TNodeMap) String() string {
	return fmt.Sprintf("MAP[%v]%v", t.KeyType, t.ValueType)
}

type TNodeUserType struct {
	Name string
}

func (t *TNodeUserType) String() string {
	return fmt.Sprintf("USERTYPE(%s)", t.Name)
}

type TNodeStruct struct {
	Fields map[string]TNode
}

func (t *TNodeStruct) String() string {
	s := ""
	for k, v := range t.Fields {
		if s != "" {
			s += ", "
		}
		s += fmt.Sprintf("%s:%v", k, v)
	}
	return fmt.Sprintf("STRUCT(%s)", s)
}

type TNodeArray struct {
	ElemType TNode
	Len      int
}

func (t *TNodeArray) String() string {
	return fmt.Sprintf("ARRAY(%d, %v)", t.Len, t.ElemType)
}

type TNodeSlice struct {
	ElemType TNode
}

func (t *TNodeSlice) String() string {
	return fmt.Sprintf("SLICE(%v)", t.ElemType)
}

type TNodeNumber struct {
	Flavor Flavor
}

func (t *TNodeNumber) String() string {
	return fmt.Sprintf("NUMBER(%v)", t.Flavor)
}

type TNodeBool struct {
}

func (t *TNodeBool) String() string {
	return "BOOL"
}

type TNodeString struct {
}

func (t *TNodeString) String() string {
	return "STRING"
}

type TNodeAny struct {
}

func (t *TNodeAny) String() string {
	return "ANY"
}
