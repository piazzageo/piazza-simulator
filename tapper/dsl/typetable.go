package dsl

import "fmt"

type Symbol string

func (s Symbol) String() string {
	return string(s)
}

type SymbolTable struct {
	Symbols map[Symbol]Type
}

func NewSymbolTable() *SymbolTable {
	syms := map[Symbol]Type{}

	st := &SymbolTable{
		Symbols: syms,
	}

	return st
}

func (st *SymbolTable) Init() {
	st.add("int8", &NumberType{Flavor: FlavorS8})
	st.add("int16", &NumberType{Flavor: FlavorS16})
	st.add("int32", &NumberType{Flavor: FlavorS32})
	st.add("int64", &NumberType{Flavor: FlavorS64})
	st.add("uint8", &NumberType{Flavor: FlavorU8})
	st.add("uint16", &NumberType{Flavor: FlavorU16})
	st.add("uint32", &NumberType{Flavor: FlavorU32})
	st.add("uint64", &NumberType{Flavor: FlavorU64})
	st.add("float32", &NumberType{Flavor: FlavorF32})
	st.add("float64", &NumberType{Flavor: FlavorF64})
	st.add("bool", &BoolType{})
	st.add("string", &StringType{})
	st.add("any", &AnyType{})
}

func (st *SymbolTable) String() string {
	s := fmt.Sprintf("%d:", len(st.Symbols))
	for k, v := range st.Symbols {
		s += fmt.Sprintf(" [%s:%v]", k, v)
	}
	return s
}

func (st *SymbolTable) add(s Symbol, t Type) {
	st.Symbols[s] = t
}

func (st *SymbolTable) get(s Symbol) Type {
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
