package dsl

import (
	"fmt"
)

type Node interface {
	String() string
}

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
