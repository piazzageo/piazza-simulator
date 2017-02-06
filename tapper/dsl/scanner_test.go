package dsl

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

//--------------------------

func Test40(t *testing.T) {
	assert := assert.New(t)
	assert.True(!false)
}

func Test41(t *testing.T) {
	assert := assert.New(t)

	s := &Scanner{}

	type data struct {
		source string
		tokens []string
	}

	table := []data{
		{
			source: "  as > 10 | b +3 < c",
			tokens: []string{"as", ">", "10", "|", "b", "+", "3", "<", "c"},
		},
		{
			source: `[map] int32`,
			tokens: []string{"[map]", "int32"},
		},
	}

	for _, testcase := range table {

		tokens, err := s.Scan(testcase.source)
		assert.NoError(err)
		assert.NotNil(tokens)

		assert.Len(tokens, len(testcase.tokens))
		for i, _ := range testcase.tokens {
			assert.Equal(testcase.tokens[i], tokens[i].Text)
			log.Printf("%s", testcase.tokens[i])
		}
		log.Printf("---")

	}

}
