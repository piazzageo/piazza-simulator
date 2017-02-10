package dsl

import (
	"fmt"
)

//---------------------------------------------------------------------

type Node interface {
	String() string
}

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

type NodeMultiply struct {
	left  Node
	right Node
}

func (n *NodeMultiply) String() string {
	return fmt.Sprintf("(* %v %v)", n.left, n.right)
}

type NodeAdd struct {
	left  Node
	right Node
}

func (n *NodeAdd) String() string {
	return fmt.Sprintf("(+ %v %v)", n.left, n.right)
}

type NodeSymbol struct {
	name string
}

func (n *NodeSymbol) String() string {
	return fmt.Sprintf("%s", n.name)
}

//---------------------------------------------------------------------------

type NodeMapType struct {
	KeyType   Node
	ValueType Node
}

func (t *NodeMapType) String() string {
	return fmt.Sprintf("MAP[%v]%v", t.KeyType, t.ValueType)
}

type NodeUserType struct {
	Name string
}

func (t *NodeUserType) String() string {
	return fmt.Sprintf("USERTYPE(%s)", t.Name)
}

type NodeStructType struct {
	Fields map[string]bool
}

func (t *NodeStructType) String() string {
	s := ""
	for k, _ := range t.Fields {
		if s != "" {
			s += ", "
		}
		s += fmt.Sprintf("%v", k)
	}
	return fmt.Sprintf("STRUCT(%s)", s)
}

type NodeArrayType struct {
	ElemType Node
	Len      int
}

func (t *NodeArrayType) String() string {
	return fmt.Sprintf("ARRAY(%d, %v)", t.Len, t.ElemType)
}

type NodeSliceType struct {
	ElemType Node
}

func (t *NodeSliceType) String() string {
	return fmt.Sprintf("SLICE(%v)", t.ElemType)
}

type NodeNumberType struct {
	Flavor Flavor
}

func (t *NodeNumberType) String() string {
	return fmt.Sprintf("NUMBER(%v)", t.Flavor)
}

type NodeBoolType struct {
}

func (t *NodeBoolType) String() string {
	return "BOOL"
}

type NodeStringType struct {
}

func (t *NodeStringType) String() string {
	return "STRING"
}

type NodeAnyType struct {
}

func (t *NodeAnyType) String() string {
	return "ANY"
}
