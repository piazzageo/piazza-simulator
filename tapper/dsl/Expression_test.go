package dsl

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//--------------------------

func Test30(t *testing.T) {
	assert := assert.New(t)
	assert.True(!false)
}

func Test31(t *testing.T) {
	assert := assert.New(t)

	const src = `  a*(b + c )`

	err := ParseExpr(src)
	assert.NoError(err)
	//assert.NotNil(toks)

	//for _, v := range toks {
	//log.Printf("%s\n", v.String())
	//}
}
