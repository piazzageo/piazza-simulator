package dsl

import "testing"

type dslTestItem struct {
	decl   string
	expr   string
	env    *Environment
	result interface{}
}

var dslTestData = []dslTestItem{
	dslTestItem{
		decl: `{ "Point": { "x": "float", "y": "float" } }`,
		expr: "a * (b + c )",
		env: &Environment{
			values: map[string]*ExprValue{
				"a": &ExprValue{Type: IntType, Value: 2},
				"b": &ExprValue{Type: IntType, Value: 3},
				"c": &ExprValue{Type: IntType, Value: 4},
			},
		},
		result: 314,
	},
}

func TestDsl(t *testing.T) {
	/*	assert := assert.New(t)

		for _, item := range dslTestData {
			d, err := NewDsl()
			assert.NoError(err)

			tId, err := d.ParseDeclaration(item.decl)
			assert.NoError(err)
			assert.NotEqual(InvalidId, tId)

			eId, err := d.ParseExpression(tId, item.expr)
			assert.NoError(err)
			assert.NotEqual(InvalidId, eId)

			result, err := d.Evaluate(eId, tId, item.env)
			assert.NoError(err)
			assert.NotNil(result)
			assert.EqualValues(item.result, result)
		}
	*/
}
