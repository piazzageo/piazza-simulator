package dsl

import (
	"fmt"
)

//---------------------------------------------------------------------

type Node interface {
	String() string
	Type() TypeNode
}

//---------------------------------------------------------------------

type NodeMultiply struct {
	Typ   TypeNode
	Left  Node
	Right Node
}

func NewNodeMultiply(left Node, right Node) *NodeMultiply {
	n := &NodeMultiply{
		Left:  left,
		Right: right,
	}
	return n
}

func (n *NodeMultiply) Type() TypeNode {
	return n.Typ
}

func (n *NodeMultiply) String() string {
	return fmt.Sprintf("(* %v %v)", n.Left, n.Right)
}

//---------------------------------------------------------------------

type NodeAdd struct {
	Typ   TypeNode
	Left  Node
	Right Node
}

func NewNodeAdd(left Node, right Node) *NodeAdd {
	n := &NodeAdd{
		Left:  left,
		Right: right,
	}
	return n
}

func (n *NodeAdd) Type() TypeNode {
	return n.Typ
}

func (n *NodeAdd) String() string {
	return fmt.Sprintf("(+ %v %v)", n.Left, n.Right)
}

//---------------------------------------------------------------------------

type NodeSymbol struct {
	Typ  TypeNode
	Name string
}

func NewNodeSymbol(name string) *NodeSymbol {
	n := &NodeSymbol{
		Name: name,
	}
	return n
}

func (n *NodeSymbol) Type() TypeNode {
	return n.Typ
}

func (t *NodeSymbol) String() string {
	return fmt.Sprintf("SYMBOL(%s)", t.Name)
}

//---------------------------------------------------------------------------

type NodeIntConstant struct {
	Value int
}

func NewNodeIntConstant(value int) *NodeIntConstant {
	n := &NodeIntConstant{
		Value: value,
	}
	return n
}

func (n *NodeIntConstant) Type() TypeNode {
	return NewTypeNodeInt()
}

func (t *NodeIntConstant) String() string {
	return fmt.Sprintf("INTCONSTANT(%d)", t.Value)
}

//---------------------------------------------------------------------------

type NodeFloatConstant struct {
	Value float64
}

func NewNodeFloatConstant(value float64) *NodeFloatConstant {
	n := &NodeFloatConstant{
		Value: value,
	}
	return n
}

func (n *NodeFloatConstant) Type() TypeNode {
	return NewTypeNodeFloat()
}

func (t *NodeFloatConstant) String() string {
	return fmt.Sprintf("FLOATCONSTANT(%f)", t.Value)
}

//---------------------------------------------------------------------------

type NodeBoolConstant struct {
	Value bool
}

func NewNodeBoolConstant(value bool) *NodeBoolConstant {
	n := &NodeBoolConstant{
		Value: value,
	}
	return n
}

func (n *NodeBoolConstant) Type() TypeNode {
	return NewTypeNodeBool()
}

func (t *NodeBoolConstant) String() string {
	return fmt.Sprintf("BOOLCONSTANT(%t)", t.Value)
}

//---------------------------------------------------------------------------

type NodeStringConstant struct {
	Value string
}

func NewNodeStringConstant(value string) *NodeStringConstant {
	n := &NodeStringConstant{
		Value: value,
	}
	return n
}

func (n *NodeStringConstant) Type() TypeNode {
	return NewTypeNodeString()
}

func (t *NodeStringConstant) String() string {
	return fmt.Sprintf("STRINGCONSTANT(%s)", t.Value)
}

//---------------------------------------------------------------------------

type NodeStructField struct {
	Typ  TypeNode
	Name string
}

func NewNodeStructField(name string) *NodeStructField {
	n := &NodeStructField{
		Name: name,
	}
	return n
}

func (n *NodeStructField) Type() TypeNode {
	return n.Typ
}

func (t *NodeStructField) String() string {
	return fmt.Sprintf("FIELD(%s)", t.Name)
}

//---------------------------------------------------------------------------

type NodeArrayIndex struct {
	Typ   TypeNode
	Index int
}

func NewNodeArrayIndex(index int) *NodeArrayIndex {
	n := &NodeArrayIndex{
		Index: index,
	}
	return n
}

func (n *NodeArrayIndex) Type() TypeNode {
	return n.Typ
}

func (t *NodeArrayIndex) String() string {
	return fmt.Sprintf("INDEX(%d)", t.Index)
}

//---------------------------------------------------------------------------

type NodeMapKey struct {
	Typ TypeNode
	Key string
}

func NewNodeMapKey(key string) *NodeMapKey {
	n := &NodeMapKey{
		Key: key,
	}
	return n
}

func (n *NodeMapKey) Type() TypeNode {
	return n.Typ
}

func (t *NodeMapKey) String() string {
	return fmt.Sprintf("MAPKEY(%s)", t.Key)
}

//---------------------------------------------------------------------------
