package dsl

import "testing"

type evalTestItem struct {
	node   ExprNode
	values map[string]*ExprValue
	typs   map[string]TypeNode
	result interface{}
}

var evalTestData = []evalTestItem{
	evalTestItem{ // "a * (b + c )"
		node: NewExprNodeMultiply(
			NewExprNodeAdd(NewExprNodeSymbolRef("c", IntType), NewExprNodeSymbolRef("b", IntType)),
			NewExprNodeSymbolRef("a", IntType)),
		values: map[string]*ExprValue{
			"a": &ExprValue{Type: IntType, Value: 2},
			"b": &ExprValue{Type: IntType, Value: 3},
			"c": &ExprValue{Type: IntType, Value: 4},
		},
		typs: map[string]TypeNode{
			"a": NewTypeNodeInt(),
			"b": NewTypeNodeInt(),
			"c": NewTypeNodeInt(),
		},
		result: 14,
	},
}

func TestEval(t *testing.T) {
	/*
		assert := assert.New(t)

		for _, item := range evalTestData {

			typeTable, err := NewTypeTable()
			assert.NoError(err)

			env := NewEnvironment(typeTable)

			for k, v := range item.typs {
				err = typeTable.addNode(k, v)
				assert.NoError(err)
			}

			for k, v := range item.values {
				env.set(k, v)
			}

			eval := &Eval{}
			_, err = eval.Evaluate(item.node, env)
			assert.NoError(err)
			//assert.NotNil(result)
			//assert.EqualValues(item.result, result)
		}
	*/
}
