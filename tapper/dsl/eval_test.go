package dsl

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type evalTestItem struct {
	node   Node
	env    *Environment
	result interface{}
}

var evalTestData = []evalTestItem{
	evalTestItem{ // "a * (b + c )"
		node: NewNodeMultiply(
			NewNodeAdd(NewNodeSymbol("c"), NewNodeSymbol("b")),
			NewNodeSymbol("a")),
		env: &Environment{
			data: map[string]interface{}{"a": 2, "b": 3, "c": 4},
		},
		result: 14,
	},
}

func TestEval(t *testing.T) {
	assert := assert.New(t)

	var err error

	typeTable, err := NewTypeTable()
	assert.NoError(err)

	for _, s := range []string{"a", "b", "c"} {
		err = typeTable.add(s)
		assert.NoError(err)
		typeTable.setNode(s, NewNodeIntType())
		assert.NoError(err)
	}

	/*for _, item := range evalTestData {
		eval := &Eval{}
		result, err := eval.Evaluate(item.node, typeTable, item.env)
		assert.NoError(err)
		assert.NotNil(result)
		assert.EqualValues(item.result, result)
	}*/
}
