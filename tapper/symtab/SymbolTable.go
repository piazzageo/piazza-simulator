package symtab

import "fmt"

type Symbol string

func (s Symbol) String() string {
	return string(s)
}

type SymbolTable struct {
	Symbols map[Symbol]DslType
}

func NewSymbolTable() *SymbolTable {
	syms := map[Symbol]DslType{}

	st := &SymbolTable{
		Symbols: syms,
	}

	return st
}

func (st *SymbolTable) Init() {
	st.add("int8", &NumberDslType{Flavor: S8})
	st.add("int16", &NumberDslType{Flavor: S16})
	st.add("int32", &NumberDslType{Flavor: S32})
	st.add("int64", &NumberDslType{Flavor: S64})
	st.add("uint8", &NumberDslType{Flavor: U8})
	st.add("uint16", &NumberDslType{Flavor: U16})
	st.add("uint32", &NumberDslType{Flavor: U32})
	st.add("uint64", &NumberDslType{Flavor: U64})
	st.add("float32", &NumberDslType{Flavor: F32})
	st.add("float64", &NumberDslType{Flavor: F64})
	st.add("bool", &BoolDslType{})
	st.add("string", &StringDslType{})
	st.add("any", &AnyDslType{})
}

func (st *SymbolTable) String() string {
	s := fmt.Sprintf("%d:", len(st.Symbols))
	for k, v := range st.Symbols {
		s += fmt.Sprintf(" [%s:%v]", k, v)
	}
	return s
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

func (st *SymbolTable) has(s Symbol) bool {
	_, ok := st.Symbols[s]
	return ok
}

func (st *SymbolTable) remove(s Symbol) {
	delete(st.Symbols, s)
}
