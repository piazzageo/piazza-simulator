package dsl

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type evalTestItem struct {
	node   Node
	vars   map[string]varItem
	result interface{}
}

type varItem struct {
	value    interface{}
	datatype TypeNode
}

var evalTestData = []evalTestItem{
	evalTestItem{ // "a * (b + c )"
		node: NewNodeMultiply(
			NewNodeAdd(NewNodeSymbol("c"), NewNodeSymbol("b")),
			NewNodeSymbol("a")),
		vars: map[string]varItem{
			"a": varItem{value: 2, datatype: NewTypeNodeInt()},
			"b": varItem{value: 3, datatype: NewTypeNodeInt()},
			"c": varItem{value: 4, datatype: NewTypeNodeInt()},
		},
		result: 14,
	},
}

func TestEval(t *testing.T) {
	assert := assert.New(t)

	for _, item := range evalTestData {

		typeTable, err := NewTypeTable()
		assert.NoError(err)

		env := NewEnvironment(typeTable)

		for k, v := range item.vars {
			err = typeTable.addNode(k, v.datatype)
			assert.NoError(err)

			env.set(k, v.value)
		}

		eval := &Eval{}
		_, err = eval.Evaluate(item.node, env)
		assert.NoError(err)
		//assert.NotNil(result)
		//assert.EqualValues(item.result, result)
	}
}
