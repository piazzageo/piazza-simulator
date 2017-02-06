package symtab

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

	const src = `  as > 10 | b +3 < c`

	err := ParseExpr(src)
	assert.NoError(err)
	//assert.NotNil(toks)

	//for _, v := range toks {
	//log.Printf("%s\n", v.String())
	//}
}
