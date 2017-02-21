package dsl

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//--------------------------

func TestScannerPrefixMatch(t *testing.T) {
	assert := assert.New(t)

	ok, siz := matchArrayTypePrefix("[]")
	assert.False(ok)
	ok, siz = matchArrayTypePrefix(" []")
	assert.False(ok)
	ok, siz = matchArrayTypePrefix("[}")
	assert.False(ok)
	ok, siz = matchArrayTypePrefix("[3a]")
	assert.False(ok)
	ok, siz = matchArrayTypePrefix("[3]")
	assert.True(ok)
	assert.Equal(3, siz)
	ok, siz = matchArrayTypePrefix("[32]")
	assert.True(ok)
	assert.Equal(32, siz)
}

func TestScanner(t *testing.T) {
	assert := assert.New(t)

	s := &Scanner{}

	type data struct {
		asDecl bool
		source string
		tokens []string
	}

	table := []data{
		{
			asDecl: false,
			source: "  as > 10 | b +3 < c",
			tokens: []string{"as", ">", "10", "|", "b", "+", "3", "<", "c"},
		},
		{
			asDecl: true,
			source: `[map] int`,
			tokens: []string{"[map]", "int"},
		},
		{
			asDecl: true,
			source: `[] int`,
			tokens: []string{"[]", "int"},
		},
		{
			asDecl: true,
			source: `[4] int`,
			tokens: []string{"[4]", "int"},
		},
	}

	for _, testcase := range table {
		tokens, err := s.Scan(testcase.source, testcase.asDecl)
		assert.NoError(err)
		assert.NotNil(tokens)

		assert.Len(tokens, len(testcase.tokens))
		for i := range testcase.tokens {
			assert.Equal(testcase.tokens[i], tokens[i].Text)
		}
	}
}
