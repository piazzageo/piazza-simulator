package dsl

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type evalTestItem struct {
	expr      ExprNode
	env       *Environment
	typeTable *TypeTable
	result    *ExprValue
}

var evalTestData = []evalTestItem{
	evalTestItem{ // "a * (b + c )"
		expr: NewExprNodeMultiply(
			NewExprNodeAdd(
				NewExprNodeSymbolRef("c", IntType),
				NewExprNodeSymbolRef("b", IntType)),
			NewExprNodeSymbolRef("a", IntType)),
		typeTable: &TypeTable{
			Types: map[string]*TypeTableEntry{
				"a": &TypeTableEntry{Name: "a", Typ: NewTypeNodeInt()},
				"b": &TypeTableEntry{Name: "b", Typ: NewTypeNodeInt()},
				"c": &TypeTableEntry{Name: "c", Typ: NewTypeNodeInt()},
			},
		},
		env: &Environment{
			values: map[string]*ExprValue{
				"a": &ExprValue{Type: IntType, Value: 2},
				"b": &ExprValue{Type: IntType, Value: 3},
				"c": &ExprValue{Type: IntType, Value: 4},
			},
		},
		result: &ExprValue{Type: IntType, Value: 14},
	},
	evalTestItem{ // "a.foo + 1"
		expr: NewExprNodeMultiply(
			NewExprNodeAdd(
				NewExprNodeSymbolRef("c", IntType),
				NewExprNodeSymbolRef("b", IntType)),
			NewExprNodeSymbolRef("a", IntType)),
		typeTable: &TypeTable{
			Types: map[string]*TypeTableEntry{
				"a": &TypeTableEntry{Name: "a", Typ: NewTypeNodeInt()},
				"b": &TypeTableEntry{Name: "b", Typ: NewTypeNodeInt()},
				"c": &TypeTableEntry{Name: "c", Typ: NewTypeNodeInt()},
			},
		},
		env: &Environment{
			values: map[string]*ExprValue{
				"a": &ExprValue{Type: IntType, Value: 2},
				"b": &ExprValue{Type: IntType, Value: 3},
				"c": &ExprValue{Type: IntType, Value: 4},
			},
		},
		result: &ExprValue{Type: IntType, Value: 14},
	},
}

func TestEval(t *testing.T) {
	assert := assert.New(t)

	for _, item := range evalTestData {
		eval := &Eval{}
		value, err := eval.Evaluate(item.expr, item.env)
		assert.NoError(err)
		assert.Equal(item.result, value)
	}
}
