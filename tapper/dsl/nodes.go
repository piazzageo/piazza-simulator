package dsl

import (
	"fmt"
)

//---------------------------------------------------------------------

type NodeCore struct {
	left  Node
	right Node
	value interface{}
}

func (n *NodeCore) Left() Node {
	return n.left
}

func (n *NodeCore) Right() Node {
	return n.right
}

func (n *NodeCore) Value() interface{} {
	return n.value
}

type Node interface {
	String() string
	Left() Node
	Right() Node
	Value() interface{}
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
	NodeCore
}

func NewNodeMultiply(left Node, right Node) *NodeMultiply {
	n := &NodeMultiply{}
	n.left = left
	n.right = right
	return n
}

func (n *NodeMultiply) String() string {
	return fmt.Sprintf("(* %v %v)", n.left, n.right)
}

type NodeAdd struct {
	NodeCore
}

func NewNodeAdd(left Node, right Node) *NodeAdd {
	n := &NodeAdd{}
	n.left = left
	n.right = right
	return n
}

func (n *NodeAdd) String() string {
	return fmt.Sprintf("(+ %v %v)", n.left, n.right)
}

type NodeSymbol struct {
	NodeCore
}

func NewNodeSymbol(name string) *NodeSymbol {
	n := &NodeSymbol{}
	n.value = name
	return n
}

func (n *NodeSymbol) String() string {
	return fmt.Sprintf("%s", n.value.(string))
}

//---------------------------------------------------------------------------

type NodeMapType struct {
	// left is key type
	// right is value type
	NodeCore
}

func NewNodeMapType(elemType Node, valueType Node) *NodeMapType {
	n := &NodeMapType{}
	n.left = elemType
	n.right = valueType
	return n
}

func (t *NodeMapType) String() string {
	return fmt.Sprintf("MAP[%v]%v", t.left, t.right)
}

type NodeUserType struct {
	NodeCore
}

func NewNodeUserType(name string) *NodeUserType {
	n := &NodeUserType{}
	n.value = name
	return n
}

func (t *NodeUserType) String() string {
	return fmt.Sprintf("USERTYPE(%s)", t.value.(string))
}

type NodeStructType struct {
	NodeCore
}

func NewNodeStructType(fields map[string]bool) *NodeStructType {
	n := &NodeStructType{}
	n.value = fields
	return n
}

func (t *NodeStructType) String() string {
	fields := t.value.(map[string]bool)
	s := ""
	for k, _ := range fields {
		if s != "" {
			s += ", "
		}
		s += fmt.Sprintf("%v", k)
	}
	return fmt.Sprintf("STRUCT(%s)", s)
}

type NodeArrayType struct {
	// left is elem type
	// value is length (siz)
	NodeCore
}

func NewNodeArrayType(elemType Node, siz int) *NodeArrayType {
	n := &NodeArrayType{}
	n.left = elemType
	n.value = siz
	return n
}
func (t *NodeArrayType) String() string {
	return fmt.Sprintf("ARRAY(%d, %v)", t.value.(int), t.left)
}

type NodeSliceType struct {
	NodeCore
}

func NewNodeSliceType(elemType Node) *NodeSliceType {
	n := &NodeSliceType{}
	n.left = elemType
	return n
}

func (t *NodeSliceType) String() string {
	return fmt.Sprintf("SLICE(%v)", t.left)
}

type NodeNumberType struct {
	// value is Flavor
	NodeCore
}

func NewNodeNumberType(flavor Flavor) *NodeNumberType {
	n := &NodeNumberType{}
	n.value = flavor
	return n
}

func (t *NodeNumberType) String() string {
	return fmt.Sprintf("NUMBER(%v)", t.value.(Flavor))
}

type NodeBoolType struct {
	NodeCore
}

func NewNodeBoolType() *NodeBoolType {
	return &NodeBoolType{}
}

func (t *NodeBoolType) String() string {
	return "BOOL"
}

type NodeStringType struct {
	NodeCore
}

func NewNodeStringType() *NodeStringType {
	n := &NodeStringType{}
	return n
}

func (t *NodeStringType) String() string {
	return "STRING"
}

type NodeAnyType struct {
	NodeCore
}

func NewNodeAnyType() *NodeAnyType {
	return &NodeAnyType{}
}

func (t *NodeAnyType) String() string {
	return "ANY"
}
