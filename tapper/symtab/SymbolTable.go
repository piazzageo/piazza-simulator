package symtab

import "fmt"

type Symbol string

func (s Symbol) String() string {
	return string(s)
}

type SymbolTable struct {
	Children []*SymbolTable
	Parent   *SymbolTable

	Symbols map[Symbol]DslType
}

func NewSymbolTable(parent *SymbolTable) *SymbolTable {
	syms := map[Symbol]DslType{}

	st := &SymbolTable{
		Children: []*SymbolTable{},
		Parent:   parent,
		Symbols:  syms,
	}
	return st
}

func (st *SymbolTable) String() string {
	s := fmt.Sprintf("Parent:%t, Children:%d", st.Parent != nil, len(st.Children))
	for k, v := range st.Symbols {
		s += fmt.Sprintf(" [%s:%v]", k, v)
	}
	return s
}

func (st *SymbolTable) addScope() *SymbolTable {
	stNew := NewSymbolTable(st)
	st.Children = append(st.Children, stNew)
	return stNew
}

func (st *SymbolTable) add(s Symbol, t DslType) {
	st.Symbols[s] = t
}

func (st *SymbolTable) get(s Symbol) DslType {
	v, ok := st.Symbols[s]
	if !ok {
		return nil
	}
	return v
}

func (st *SymbolTable) remove(s Symbol) {
	delete(st.Symbols, s)
}
