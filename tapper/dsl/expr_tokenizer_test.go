package dsl

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

//---------------------------------------------------------------------------

type exprTest struct {
	tag        string
	declBlock  *DeclBlock
	exprText   string
	exprTokens []*Token
	exprNode   ExprNode
	env        *Environment
	result     *ExprValue
}

var exprTestData = []exprTest{
	exprTest{
		tag: "Test 0",
		declBlock: &DeclBlock{
			"a": "int",
			"b": "int",
			"c": "int",
		},
		exprText: "  a*(b + c )",
		exprTokens: []*Token{
			&Token{Line: 1, Column: 4, Id: TokenSymbol, Text: "a"},
			&Token{Line: 1, Column: 7, Id: TokenSymbol, Text: "b"},
			&Token{Line: 1, Column: 11, Id: TokenSymbol, Text: "c"},
			&Token{Line: 1, Column: 9, Id: TokenAdd, Text: "+"},
			&Token{Line: 1, Column: 5, Id: TokenMultiply, Text: "*"},
		},
		exprNode: NewExprNodeMultiply(
			NewExprNodeAdd(NewExprNodeSymbolRef("c", IntType), NewExprNodeSymbolRef("b", IntType)),
			NewExprNodeSymbolRef("a", IntType)),
	},
	exprTest{
		tag: "Test 1",
		declBlock: &DeclBlock{
			"a": "int",
			"b": "int",
			"c": "int",
			"d": "int",
			"e": "int",
		},
		exprText: "a/b%c^d-e",
		exprTokens: []*Token{
			&Token{Line: 1, Column: 2, Id: TokenSymbol, Text: "a"},
			&Token{Line: 1, Column: 4, Id: TokenSymbol, Text: "b"},
			&Token{Line: 1, Column: 3, Id: TokenDivide, Text: "/"},
			&Token{Line: 1, Column: 6, Id: TokenSymbol, Text: "c"},
			&Token{Line: 1, Column: 5, Id: TokenModulus, Text: "%"},
			&Token{Line: 1, Column: 8, Id: TokenSymbol, Text: "d"},
			&Token{Line: 1, Column: 7, Id: TokenBitwiseXor, Text: "^"},
			&Token{Line: 1, Column: 10, Id: TokenSymbol, Text: "e"},
			&Token{Line: 1, Column: 9, Id: TokenSubtract, Text: "-"},
		},
		exprNode: NewExprNodeSubtract(
			NewExprNodeSymbolRef("e", IntType),
			NewExprNodeBitwiseXor(
				NewExprNodeSymbolRef("d", IntType),
				NewExprNodeModulus(
					NewExprNodeSymbolRef("c", IntType),
					NewExprNodeDivide(
						NewExprNodeSymbolRef("b", IntType), NewExprNodeSymbolRef("a", IntType))))),
	},
	exprTest{
		tag: "Test 2",
		declBlock: &DeclBlock{
			"a": "int",
			"b": "int",
			"c": "int",
			"d": "int",
			"e": "int",
			"f": "int",
			"g": "int",
			"h": "int",
		},
		exprText: "a < b && c <= d && e > f && g >= h",
		exprTokens: []*Token{
			&Token{Line: 1, Column: 2, Id: TokenSymbol, Text: "a"},
			&Token{Line: 1, Column: 6, Id: TokenSymbol, Text: "b"},
			&Token{Line: 1, Column: 4, Id: TokenLessThan, Text: "<"},
			&Token{Line: 1, Column: 11, Id: TokenSymbol, Text: "c"},
			&Token{Line: 1, Column: 16, Id: TokenSymbol, Text: "d"},
			&Token{Line: 1, Column: 13, Id: TokenLessOrEqualThan, Text: "<="},
			&Token{Line: 1, Column: 8, Id: TokenLogicalAnd, Text: "&&"},
			&Token{Line: 1, Column: 21, Id: TokenSymbol, Text: "e"},
			&Token{Line: 1, Column: 25, Id: TokenSymbol, Text: "f"},
			&Token{Line: 1, Column: 23, Id: TokenGreaterThan, Text: ">"},
			&Token{Line: 1, Column: 18, Id: TokenLogicalAnd, Text: "&&"},
			&Token{Line: 1, Column: 30, Id: TokenSymbol, Text: "g"},
			&Token{Line: 1, Column: 35, Id: TokenSymbol, Text: "h"},
			&Token{Line: 1, Column: 32, Id: TokenGreaterOrEqualThan, Text: ">="},
			&Token{Line: 1, Column: 27, Id: TokenLogicalAnd, Text: "&&"},
		},
	},
	exprTest{
		tag: "Test 3",
		declBlock: &DeclBlock{
			"foo": map[string]string{"z": "int"},
		},
		exprText: "z.foo + 1",
		exprTokens: []*Token{
			&Token{Line: 1, Column: 2, Id: TokenSymbol, Text: "z"},
			&Token{Line: 1, Column: 6, Id: TokenSymbol, Text: "foo"},
			&Token{Line: 1, Column: 3, Id: TokenPeriod, Text: "."},
			&Token{Line: 1, Column: 10, Id: TokenInt, Text: "1"},
			&Token{Line: 1, Column: 8, Id: TokenAdd, Text: "+"},
		},
	},
	exprTest{
		tag: "Test 4",
		declBlock: &DeclBlock{
			"a": "int",
			"b": "int",
			"c": "int",
		},
		exprText: "q[2] + 3",
		exprTokens: []*Token{
			&Token{Line: 1, Column: 2, Id: TokenSymbol, Text: "q"},
			&Token{Line: 1, Column: 3, Id: TokenLeftBracket, Text: "["},
			&Token{Line: 1, Column: 4, Id: TokenInt, Text: "2"},
			&Token{Line: 1, Column: 5, Id: TokenRightBracket, Text: "]"},
			&Token{Line: 1, Column: 9, Id: TokenInt, Text: "3"},
			&Token{Line: 1, Column: 7, Id: TokenAdd, Text: "+"},
		},
	},
	exprTest{
		tag: "Test 5",
		declBlock: &DeclBlock{
			"a": "int",
			"b": "int",
			"c": "int",
		},
		exprText: `m["k"] + 0.5`,
		exprTokens: []*Token{
			&Token{Line: 1, Column: 2, Id: TokenSymbol, Text: "m"},
			&Token{Line: 1, Column: 3, Id: TokenLeftBracket, Text: "["},
			&Token{Line: 1, Column: 6, Id: TokenString, Text: `"k"`},
			&Token{Line: 1, Column: 7, Id: TokenRightBracket, Text: "]"},
			&Token{Line: 1, Column: 13, Id: TokenFloat, Text: "0.5"},
			&Token{Line: 1, Column: 9, Id: TokenAdd, Text: "+"},
		},
	},
	exprTest{
		tag: "Test 6",
		declBlock: &DeclBlock{
			"t": "bool",
		},
		exprText: "t == true",
		exprTokens: []*Token{
			&Token{Line: 1, Column: 2, Id: TokenSymbol, Text: "t"},
			&Token{Line: 1, Column: 10, Id: TokenBool, Text: "true"},
			&Token{Line: 1, Column: 4, Id: TokenEqualsEquals, Text: "=="},
		},
	},
	exprTest{
		tag: "Test 7",
		declBlock: &DeclBlock{
			"a": "int",
			"b": "int",
			"c": "int",
		},
		exprText: `f != false`,
		exprTokens: []*Token{
			&Token{Line: 1, Column: 2, Id: TokenSymbol, Text: "f"},
			&Token{Line: 1, Column: 11, Id: TokenBool, Text: "false"},
			&Token{Line: 1, Column: 4, Id: TokenNotEquals, Text: "!="},
		},
	},
	exprTest{
		tag: "Test 8",
		declBlock: &DeclBlock{
			"a": "int",
			"b": "int",
			"c": "int",
		},
		exprText: `a && b || c`,
		exprTokens: []*Token{
			&Token{Line: 1, Column: 2, Id: TokenSymbol, Text: "a"},
			&Token{Line: 1, Column: 7, Id: TokenSymbol, Text: "b"},
			&Token{Line: 1, Column: 4, Id: TokenLogicalAnd, Text: "&&"},
			&Token{Line: 1, Column: 12, Id: TokenSymbol, Text: "c"},
			&Token{Line: 1, Column: 9, Id: TokenLogicalOr, Text: "||"},
		},
	},
}

//---------------------------------------------------------------------------

func TestExprTokenizer(t *testing.T) {
	assert := assert.New(t)

	var err error

	for index, tc := range exprTestData {
		var typeTable *TypeTable
		var exprTokens []*Token

		log.Printf("------------------ %s", tc.tag)

		// decl text --> type table
		{
			assert.NotNil(tc.declBlock)

			typeTokenizer := TypeTokenizer{}
			typeTable, err = typeTokenizer.Parse(tc.declBlock)
			assert.NoError(err)
			assert.NotNil(typeTable)
		}

		// expr text -> expr tokens
		{
			assert.NotEmpty(tc.exprText)

			tokenizer := &ExprTokenizer{}
			exprTokens, err = tokenizer.Tokenize(tc.exprText)
			assert.NoError(err)
			assert.NotNil(exprTokens)
		}

		// verify expr tokens
		{
			assert.NotNil(tc.exprTokens)

			assert.Len(exprTokens, len(tc.exprTokens))
			for i := 0; i < len(exprTokens); i++ {
				//log.Printf("%s\n", tokens[i].String())
				assert.EqualValues(tc.exprTokens[i], exprTokens[i])
			}
		}

		// expr tokens -> expr nodes
		{
			assert.NotNil(tc.exprTokens)
			//assert.NotNil(tc.exprNode)

			parser := &ExprParser{}
			exprNode, err := parser.Parse(typeTable, tc.exprTokens)
			assert.NoError(err)
			assert.NotNil(exprNode)

			// verify expr nodes
			//		assert.Equal(tc.exprNode.String(), exprNode.String())
		}

		if index == 999 {
			break
		}
	}
}
