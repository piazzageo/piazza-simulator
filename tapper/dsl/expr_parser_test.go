package dsl

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//--------------------------

func TestExprParser(t *testing.T) {
	assert := assert.New(t)

	type testcase struct {
		toks []*Token
		node Node
	}
	data := []testcase{
		testcase{
			toks: []*Token{ //"  a*(b + c )":
				&Token{Line: 1, Column: 4, Id: TokenSymbol, Text: "a"},
				&Token{Line: 1, Column: 7, Id: TokenSymbol, Text: "b"},
				&Token{Line: 1, Column: 11, Id: TokenSymbol, Text: "c"},
				&Token{Line: 1, Column: 9, Id: TokenAdd, Text: "+"},
				&Token{Line: 1, Column: 5, Id: TokenMultiply, Text: "*"},
			},
			node: NewNodeMultiply(
				NewNodeAdd(NewNodeSymbol("c"), NewNodeSymbol("b")),
				NewNodeSymbol("a")),
		},
	}

	for _, tc := range data {
		ep := &ExprParser{}
		node, err := ep.Parse(tc.toks)
		assert.NoError(err)
		assert.NotNil(node)

		//log.Printf("%v", tc.node)
		//log.Printf("%v", node)
		assert.Equal(tc.node, node)
	}
}
