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

var builtinTypes map[string]TNode

func (st *TypeTable) Init() error {
	var err error

	builtinTypes = map[string]TNode{
		"int8":    &TNodeNumber{Flavor: FlavorS8},
		"int16":   &TNodeNumber{Flavor: FlavorS16},
		"int32":   &TNodeNumber{Flavor: FlavorS32},
		"int64":   &TNodeNumber{Flavor: FlavorS64},
		"uint8":   &TNodeNumber{Flavor: FlavorU8},
		"uint16":  &TNodeNumber{Flavor: FlavorU16},
		"uint32":  &TNodeNumber{Flavor: FlavorU32},
		"uint64":  &TNodeNumber{Flavor: FlavorU64},
		"float32": &TNodeNumber{Flavor: FlavorF32},
		"float64": &TNodeNumber{Flavor: FlavorF64},
		"bool":    &TNodeBool{},
		"string":  &TNodeString{},
		"any":     &TNodeAny{},
	}

	for k, v := range builtinTypes {
		err = st.add(k)
		if err != nil {
			return err
		}
		err = st.set(k, v)
		if err != nil {
			return err
		}
	}

	return nil
}

func (st *TypeTable) String() string {
	s := fmt.Sprintf("size: %d\n", len(st.Types))
	for _, tte := range st.Types {
		s += fmt.Sprintf("  %v\n", tte)
	}
	return s
}

func (st *TypeTable) isBuiltin(name string) bool {
	_, ok := builtinTypes[name]
	return ok
}

func (st *TypeTable) add(name string) error {
	if st.has(name) {
		return fmt.Errorf("type table entry already exists: %s", name)
	}
	st.Types[name] = &TypeTableEntry{
		Name: name,
	}
	return nil
}

func (st *TypeTable) set(name string, node TNode) error {
	if !st.has(name) {
		return fmt.Errorf("type table entry does not exist: %s", name)
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
