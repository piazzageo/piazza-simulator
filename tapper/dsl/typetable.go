package dsl

import "fmt"

type TypeTableEntry struct {
	Name string
	Node Node
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

var builtinTypes map[string]Node

func (st *TypeTable) Init() error {
	var err error

	builtinTypes = map[string]Node{
		"int8":    NewNodeNumberType(FlavorS8),
		"int16":   NewNodeNumberType(FlavorS16),
		"int32":   NewNodeNumberType(FlavorS32),
		"int64":   NewNodeNumberType(FlavorS64),
		"uint8":   NewNodeNumberType(FlavorU8),
		"uint16":  NewNodeNumberType(FlavorU16),
		"uint32":  NewNodeNumberType(FlavorU32),
		"uint64":  NewNodeNumberType(FlavorU64),
		"float32": NewNodeNumberType(FlavorF32),
		"float64": NewNodeNumberType(FlavorF64),
		"bool":    NewNodeBoolType(),
		"string":  NewNodeStringType(),
		"any":     NewNodeAnyType(),
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

func (st *TypeTable) set(name string, node Node) error {
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
