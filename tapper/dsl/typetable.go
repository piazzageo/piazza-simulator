package dsl

import "fmt"

type TypeTableEntry struct {
	Name string
	//Token []Token
	Typ TypeNode
}

type TypeTable struct {
	Types map[string]*TypeTableEntry
}

func NewTypeTable() (*TypeTable, error) {
	entries := map[string]*TypeTableEntry{}

	tt := &TypeTable{
		Types: entries,
	}

	err := tt.setBuiltins()
	if err != nil {
		return nil, err
	}
	return tt, nil
}

//---------------------------------------------------------------------------

type BaseType int

const (
	IntType BaseType = iota
	FloatType
	BoolType
	StringType
)

var builtinTypes map[string]TypeNode

func (st *TypeTable) setBuiltins() error {
	var err error

	builtinTypes = map[string]TypeNode{
		"int":    NewTypeNodeInt(),
		"float":  NewTypeNodeFloat(),
		"bool":   NewTypeNodeBool(),
		"string": NewTypeNodeString(),
	}

	for k, v := range builtinTypes {
		err = st.addNode(k, v)
		if err != nil {
			return err
		}
	}

	return nil
}

//---------------------------------------------------------------------------

func (e *TypeTableEntry) String() string {
	s := fmt.Sprintf("%s:", e.Name)
	s += fmt.Sprintf("[%v]", e.Typ)
	return s
}

func (st *TypeTable) String() string {
	s := fmt.Sprintf("size: %d\n", len(st.Types))
	for _, tte := range st.Types {
		s += fmt.Sprintf("  %v\n", tte)
	}
	return s
}

func (st *TypeTable) size() int {
	return len(st.Types)
}

func (st *TypeTable) isBuiltin(name string) bool {
	_, ok := builtinTypes[name]
	return ok
}

func (st *TypeTable) addNode(name string, node TypeNode) error {
	if st.has(name) {
		return fmt.Errorf("type table entry already exists: %s", name)
	}
	st.Types[name] = &TypeTableEntry{
		Name: name,
		Typ:  node,
	}
	return nil
}

func (st *TypeTable) getNode(s string) TypeNode {
	v, ok := st.Types[s]
	if !ok {
		return nil
	}
	return v.Typ
}

func (st *TypeTable) has(s string) bool {
	_, ok := st.Types[s]
	return ok
}
