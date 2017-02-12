package dsl

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTypeParser(t *testing.T) {
	assert := assert.New(t)

	tp := &TypeParser{}
	tabl, err := NewTypeTable()
	assert.NoError(err)

	for _, tc := range typeTestData {
		//log.Printf("========= %s ====== %v ==========", tc.name, tc.node)
		if tc.token == nil {
			continue
		}
		node, err := tp.Parse(tc.token, tabl)
		assert.NoError(err)
		assert.EqualValues(tc.node, node)
	}
}
