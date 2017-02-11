package dsl

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type dslTestItem struct {
	decl string
	expr string
}

var dslTestData = []dslTestItem{
	dslTestItem{
		decl: `{ "Point": { "x": "float32", "y": "float32" } }`,
		expr: "a * (b + c )",
	},
}

func TestDsl(t *testing.T) {
	assert := assert.New(t)

	for _, item := range dslTestData {
		dsl, err := NewDsl()
		assert.NoError(err)

		tId, err := dsl.ParseDeclaration(item.decl)
		assert.NoError(err)
		assert.NotEqual(InvalidId, tId)

		eId, err := dsl.ParseExpression(item.expr)
		assert.NoError(err)
		assert.NotEqual(InvalidId, eId)
	}
}
