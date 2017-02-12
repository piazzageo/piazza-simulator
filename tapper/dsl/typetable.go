package dsl

import "fmt"

type TypeTableEntry struct {
	Name  string
	Token []Token
	Node  Node
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

var builtinTypes map[string]Node

func (st *TypeTable) setBuiltins() error {
	var err error

	builtinTypes = map[string]Node{
		"int":    NewNodeIntType(),
		"float":  NewNodeFloatType(),
		"bool":   NewNodeBoolType(),
		"string": NewNodeStringType(),
		"any":    NewNodeAnyType(),
	}

	for k, v := range builtinTypes {
		err = st.add(k)
		if err != nil {
			return err
		}
		err = st.setNode(k, v)
		if err != nil {
			return err
		}
	}

	return nil
}

//---------------------------------------------------------------------------

func (e *TypeTableEntry) String() string {
	s := fmt.Sprintf("%s:", e.Name)
	s += fmt.Sprintf("[%v]", e.Node)
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

func (st *TypeTable) add(name string) error {
	if st.has(name) {
		return fmt.Errorf("type table entry already exists: %s", name)
	}
	st.Types[name] = &TypeTableEntry{
		Name: name,
	}
	return nil
}

func (st *TypeTable) setNode(name string, node Node) error {
	if !st.has(name) {
		return fmt.Errorf("type table entry does not exist: %s", name)
	}
	if st.Types[name].Node != nil {
		return fmt.Errorf("type table entry already has node: %s", name)
	}
	st.Types[name].Node = node
	return nil
}

func (st *TypeTable) setToken(name string, token []Token) error {
	if !st.has(name) {
		return fmt.Errorf("type table entry does not exist: %s", name)
	}
	if st.Types[name].Token != nil {
		return fmt.Errorf("type table entry already has token: %s", name)
	}
	st.Types[name].Token = token
	return nil
}

func (st *TypeTable) getNode(s string) Node {
	v, ok := st.Types[s]
	if !ok {
		return nil
	}
	return v.Node
}

func (st *TypeTable) getToken(s string) []Token {
	v, ok := st.Types[s]
	if !ok {
		return nil
	}
	return v.Token
}

func (st *TypeTable) has(s string) bool {
	_, ok := st.Types[s]
	return ok
}
