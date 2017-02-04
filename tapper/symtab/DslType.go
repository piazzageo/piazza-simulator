package symtab

import "fmt"

//---------------------------------------------------------------------

type Flavor int

const (
	U8 Flavor = iota
	U16
	U32
	U64
	S8
	S16
	S32
	S64
	F32
	F64
)

func (f Flavor) String() string {
	switch f {
	case U8:
		return "U8"
	case U16:
		return "U16"
	case U32:
		return "U32"
	case U64:
		return "U64"
	case S8:
		return "S8"
	case S16:
		return "S16"
	case S32:
		return "S32"
	case S64:
		return "S64"
	case F32:
		return "F32"
	case F64:
		return "F64"
	default:
		panic(8)
	}
}

//---------------------------------------------------------------------

type DslType interface {
	String() string
}

type SymbolDslType struct {
	Symbol Symbol
}

func (t *SymbolDslType) String() string {
	return fmt.Sprintf("SYM(%v)", t.Symbol)
}

type MapDslType struct {
	KeyType   DslType
	ValueType DslType
}

func (t *MapDslType) String() string {
	return fmt.Sprintf("MAP[%v]%v", t.KeyType, t.ValueType)
}

type StructDslType struct {
	Fields map[Symbol]DslType
}

func (t *StructDslType) String() string {
	s := ""
	for k, v := range t.Fields {
		if s != "" {
			s += ", "
		}
		s += fmt.Sprintf("%s:%v", k, v)
	}
	return fmt.Sprintf("STRUCT(%s)", s)
}

type ArrayDslType struct {
	ElemType DslType
	Len      int
}

func (t *ArrayDslType) String() string {
	return fmt.Sprintf("ARRAY(%d, %v)", t.Len, t.ElemType)
}

type SliceDslType struct {
	ElemType DslType
}

func (t *SliceDslType) String() string {
	return fmt.Sprintf("SLICE(%v)", t.ElemType)
}

type NumberDslType struct {
	Flavor Flavor
}

func (t *NumberDslType) String() string {
	return fmt.Sprintf("NUMBER(%v)", t.Flavor)
}

type BoolDslType struct {
}

func (t *BoolDslType) String() string {
	return "BOOL"
}

type StringDslType struct {
}

func (t *StringDslType) String() string {
	return "STRING"
}

type AnyDslType struct {
}

func (t *AnyDslType) String() string {
	return "ANY"
}
