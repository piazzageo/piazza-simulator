package dsl

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//--------------------------

func Test50(t *testing.T) {
	assert := assert.New(t)
	assert.True(!false)
}

func Test51(t *testing.T) {
	assert := assert.New(t)

	const src = `  a*(b + c )`
	et := &ExprTokenizer{}
	toks, err := et.Tokenize(src)
	assert.NoError(err)
	assert.NotNil(toks)

	ep := &ExprParser{}
	node, err := ep.Parse(toks)
	assert.NoError(err)
	assert.NotNil(node)

	//for _, v := range toks {
	//log.Printf("%s\n", v.String())
	//}
}
