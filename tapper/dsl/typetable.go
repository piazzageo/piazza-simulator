package dsl

import "fmt"

type TypeTableEntry struct {
	Name   string
	Tokens []Token
	Node   TNode
}

func (e *TypeTableEntry) String() string {
	s := fmt.Sprintf("%s:", e.Name)
	for _, v := range e.Tokens {
		s += fmt.Sprintf(" <%v>", v)
	}
	s += fmt.Sprintf("[%v]", e.Node)
	return s
}

type TypeTable struct {
	Types map[string]*TypeTableEntry
}

func NewTypeTable() *TypeTable {
	entries := map[string]*TypeTableEntry{}

	st := &TypeTable{
		Types: entries,
	}

	return st
}

func (st *TypeTable) Init() {
	st.insertFull("int8", nil, &TNodeNumber{Flavor: FlavorS8})
	st.insertFull("int16", nil, &TNodeNumber{Flavor: FlavorS16})
	st.insertFull("int32", nil, &TNodeNumber{Flavor: FlavorS32})
	st.insertFull("int64", nil, &TNodeNumber{Flavor: FlavorS64})
	st.insertFull("uint8", nil, &TNodeNumber{Flavor: FlavorU8})
	st.insertFull("uint16", nil, &TNodeNumber{Flavor: FlavorU16})
	st.insertFull("uint32", nil, &TNodeNumber{Flavor: FlavorU32})
	st.insertFull("uint64", nil, &TNodeNumber{Flavor: FlavorU64})
	st.insertFull("float32", nil, &TNodeNumber{Flavor: FlavorF32})
	st.insertFull("float64", nil, &TNodeNumber{Flavor: FlavorF64})
	st.insertFull("bool", nil, &TNodeBool{})
	st.insertFull("string", nil, &TNodeString{})
	st.insertFull("any", nil, &TNodeAny{})
}

func (st *TypeTable) String() string {
	s := fmt.Sprintf("%d:", len(st.Types))
	for _, v := range st.Types {
		s += fmt.Sprintf(" %v", v)
	}
	return s
}

func (st *TypeTable) insert(s string) {
	st.insertFull(s, nil, nil)
}

func (st *TypeTable) insertFull(s string, tok []Token, node TNode) {
	st.Types[s] = &TypeTableEntry{
		Name:   s,
		Node:   node,
		Tokens: tok,
	}
}

func (st *TypeTable) get(s string) *TypeTableEntry {
	v, ok := st.Types[s]
	if !ok {
		return nil
	}
	return v
}

func (st *TypeTable) has(s string) bool {
	_, ok := st.Types[s]
	return ok
}
