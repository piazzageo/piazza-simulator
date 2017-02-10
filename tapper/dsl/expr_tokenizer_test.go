package dsl

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//--------------------------

func Test20(t *testing.T) {
	assert := assert.New(t)

	const src = `  a*(b + c )`

	ep := &ExprTokenizer{}

	toks, err := ep.Tokenize(src)
	assert.NoError(err)
	assert.NotNil(toks)

	//for _, v := range toks {
	//log.Printf("%s\n", v.String())
	//}
}
