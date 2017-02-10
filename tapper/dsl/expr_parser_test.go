package dsl

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//--------------------------

func Test10(t *testing.T) {
	assert := assert.New(t)
	assert.True(!false)
}

func Test11(t *testing.T) {
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
			node: nil,
		},
	}

	equals := func(a Node, b Node) bool {
		return a == b
	}

	for i, tc := range data {

		ep := &ExprParser{}
		node, err := ep.Parse(tc.toks)
		assert.NoError(err)
		assert.NotNil(node)

		for i := 0; i < len(toks); i++ {
			//log.Printf("%s\n", toks[i].String())
			assert.EqualValues(v[i], toks[i])
		}
	}

}
