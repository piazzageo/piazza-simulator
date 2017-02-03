package symtab

import "fmt"

type DslType interface {
	String() string
}

type MapDslType struct {
	KeyType   DslType
	ValueType DslType
}

func (t *MapDslType) String() string {
	return fmt.Sprintf("MAP[%v]%v", t.KeyType, t.ValueType)
}

type StructDslType struct {
	KVPairs map[string]DslType
}

func (t *StructDslType) String() string {
	s := ""
	for k, v := range t.KVPairs {
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

type IntForm int

const (
	U8 IntForm = iota
	U32
	U64
	S8
	S32
	S64
)

type IntDslType struct {
	Form IntForm
}

func (t *IntDslType) String() string {
	return fmt.Sprintf("INT(%v)", t.Form)
}

func (f IntForm) String() string {
	switch f {
	case U8:
		return "U8"
	case U32:
		return "U32"
	case U64:
		return "U64"
	case S8:
		return "S8"
	case S32:
		return "S32"
	case S64:
		return "S64"
	default:
		panic(8)
	}
}

type FloatDslType struct {
}

func (t *FloatDslType) String() string {
	return "FLOAT"
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
