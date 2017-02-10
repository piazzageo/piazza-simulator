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

	const src = `  a*(b + c )`
	et := &ExprTokenizer{}
	toks, err := et.Tokenize(src)
	assert.NoError(err)
	assert.NotNil(toks)

	ep := &ExprParser{}
	node, err := ep.Parse(toks)
	assert.NoError(err)
	assert.NotNil(node)
}
