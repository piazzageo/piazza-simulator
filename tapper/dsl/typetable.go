package dsl

import "fmt"

type TypeTableEntry struct {
	Name string
	Node TNode
}

func (e *TypeTableEntry) String() string {
	s := fmt.Sprintf("%s:", e.Name)
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
	st.set("int8", &TNodeNumber{Flavor: FlavorS8})
	st.set("int16", &TNodeNumber{Flavor: FlavorS16})
	st.set("int32", &TNodeNumber{Flavor: FlavorS32})
	st.set("int64", &TNodeNumber{Flavor: FlavorS64})
	st.set("uint8", &TNodeNumber{Flavor: FlavorU8})
	st.set("uint16", &TNodeNumber{Flavor: FlavorU16})
	st.set("uint32", &TNodeNumber{Flavor: FlavorU32})
	st.set("uint64", &TNodeNumber{Flavor: FlavorU64})
	st.set("float32", &TNodeNumber{Flavor: FlavorF32})
	st.set("float64", &TNodeNumber{Flavor: FlavorF64})
	st.set("bool", &TNodeBool{})
	st.set("string", &TNodeString{})
	st.set("any", &TNodeAny{})
}

func (st *TypeTable) String() string {
	s := fmt.Sprintf("size: %d\n", len(st.Types))
	for _, tte := range st.Types {
		s += fmt.Sprintf("  %v\n", tte)
	}
	return s
}

func (st *TypeTable) set(name string, node TNode) error {
	v, ok := st.Types[name]
	if !ok {
		st.Types[name] = &TypeTableEntry{
			Name: name,
			Node: node,
		}
		return nil
	}

	if v.Node != nil {
		return fmt.Errorf("type table entry already has node set: %s", name)
	}

	st.Types[name].Node = node

	return nil
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
