package dsl

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//--------------------------

func Test20(t *testing.T) {
	assert := assert.New(t)

	data := map[string][]*Token{
		"  a*(b + c )": []*Token{
			&Token{Line: 1, Column: 4, Id: TokenSymbol, Text: "a"},
			&Token{Line: 1, Column: 7, Id: TokenSymbol, Text: "b"},
			&Token{Line: 1, Column: 11, Id: TokenSymbol, Text: "c"},
			&Token{Line: 1, Column: 9, Id: TokenAdd, Text: "+"},
			&Token{Line: 1, Column: 5, Id: TokenMultiply, Text: "*"},
		},
	}

	for k, v := range data {

		ep := &ExprTokenizer{}
		toks, err := ep.Tokenize(k)
		assert.NoError(err)
		assert.NotNil(toks)

		assert.Len(toks, len(v))
		for i := 0; i < len(toks); i++ {
			//log.Printf("%s\n", toks[i].String())
			assert.EqualValues(v[i], toks[i])
		}
	}
}