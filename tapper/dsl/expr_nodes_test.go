package dsl

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//---------------------------------------------------------------------------

func TestExprNodes(t *testing.T) {
	assert := assert.New(t)

	multiply := NewExprNodeMultiply(NewExprNodeIntConstant(17), NewExprNodeIntConstant(19))
	assert.Equal("MULTIPLY(17, 19)", multiply.String())

	add := NewExprNodeAdd(NewExprNodeIntConstant(2), NewExprNodeIntConstant(-3))
	assert.Equal("ADD(2, -3)", add.String())

	symbolRef := NewExprNodeSymbolRef("sym", FloatType)
	assert.Equal("symF", symbolRef.String())

	intConstant := NewExprNodeIntConstant(2)
	assert.Equal("2", intConstant.String())

	floatConstant := NewExprNodeFloatConstant(0.5)
	assert.Equal("0.500000", floatConstant.String())

	boolConstant := NewExprNodeBoolConstant(true)
	assert.Equal("true", boolConstant.String())

	stringConstant := NewExprNodeStringConstant("ssttrriinngg")
	assert.Equal(`"ssttrriinngg"`, stringConstant.String())

	structRef := NewExprNodeStructRef("structo", "aa")
	assert.Equal("STRUCTREF(structo.aa)", structRef.String())

	arrayRef := NewExprNodeArrayRef("arrayo", NewExprNodeIntConstant(17))
	assert.Equal("ARRAYREF(arrayo,17)", arrayRef.String())

	mapRef := NewExprNodeMapRef("mappo", NewExprNodeBoolConstant(true))
	assert.Equal("MAPREF(mappo,true)", mapRef.String())

}

func TestNodeEquality(t *testing.T) {
	assert := assert.New(t)

	//eq := func(a, b Node) bool {
	//	return a == b
	//}

	a := NewExprNodeMapRef("mappy", NewExprNodeStringConstant("asdf"))
	b := NewExprNodeMapRef("mappy", NewExprNodeStringConstant("asdf"))
	c := NewExprNodeStringConstant("qwerty")
	d := NewExprNodeFloatConstant(2.1)
	assert.Equal(a, b)
	assert.NotEqual(a, c)
	assert.NotEqual(a, d)
	assert.NotEqual(c, d)
	assert.EqualValues(a, b)
}
