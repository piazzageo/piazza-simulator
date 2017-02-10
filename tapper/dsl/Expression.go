package dsl

// shunting yard implementation adapted from https://github.com/mgenware/go-shunting-yard

import (
	"errors"
	"fmt"
	"log"
)

type Expression struct {
}

// TODO: <<, >>

func ParseExpr(s string) error {
	sc := &Scanner{}

	tokens, err := sc.Scan(s)
	if err != nil {
		return err
	}

	//for _, v := range tokens {
	//	log.Printf("%s\n", v.String())
	//}

	toks, err := Parse(tokens)
	if err != nil {
		return err
	}

	for _, tok := range toks {
		log.Printf("%v\n", ast)
	}

	ast, err := buildTree(toks)
	if err != nil {
		return err
	}

	ast.Column = 0
	return nil
}

func pop(toks []*Token) (*Token, []*Token) {
	n := len(toks)
	x := toks[n-1]
	toks2 := toks[:n-1]
	return x, toks2
}

func buildTree(toks []*Token) (*Token, []*Token, error) {

	log.Printf("%v", toks)

	tok, toks := pop(toks)

	switch tok.Id {
	case TokenMultiply:

	default:
		return nil, nil, fmt.Errorf("Unknown token building ast: %d (\"%s\")", tok.Id, tok.Text)

	}

	return nil, nil, nil
}

//===========================================================================

// precedence of operators, with Token.Id as key
var priorities map[TokenId]int

// associativities of operators
var associativities map[string]bool

func init() {
	priorities = make(map[TokenId]int, 0)
	associativities = make(map[string]bool, 0)

	priorities[TokenSymbol] = 99
	priorities[TokenNumber] = 99
	priorities[TokenMultiply] = 5
	priorities[TokenDivide] = 5
	priorities[TokenMod] = 5
	priorities[TokenBitwiseAnd] = 5
	priorities[TokenAdd] = 4
	priorities[TokenSubtract] = 4
	priorities[TokenBitwiseOr] = 4
	priorities[TokenExponent] = 4
	priorities[TokenEquals] = 3
	priorities[TokenNotEquals] = 3
	priorities[TokenLessThan] = 3
	priorities[TokenLessOrEqualThan] = 3
	priorities[TokenGreaterThan] = 3
	priorities[TokenGreaterOrEqualThan] = 3
	priorities[TokenLogicalAnd] = 2
	priorities[TokenLogicalOr] = 1

	// if not set, associativity will be false(left-associative)
}

func Parse(tokens []Token) ([]*Token, error) {
	var ret []*Token

	var operators []Token
	for _, token := range tokens {
		if token.Id == -2 || token.Id == -3 {
			operandToken := &token
			ret = append(ret, operandToken)
		} else {
			// check parentheses
			if token.Id == TokenLeftParen {
				operators = append(operators, token)
			} else if token.Id == TokenRightParen {
				foundLeftParenthesis := false
				// pop until "(" is fouund
				for len(operators) > 0 {
					oper := operators[len(operators)-1]
					operators = operators[:len(operators)-1]

					if oper.Id == TokenLeftParen {
						foundLeftParenthesis = true
						break
					} else {
						ret = append(ret, &oper)
					}
				}
				if !foundLeftParenthesis {
					return nil, errors.New("Mismatched parentheses found")
				}
			} else {
				// operator priority and associativity
				priority, ok := priorities[token.Id]
				if !ok {
					return nil, fmt.Errorf("Unknown operator: %v", &token)
				}
				rightAssociative := associativities[token.Text]

				for len(operators) > 0 {
					top := operators[len(operators)-1]

					if top.Id == TokenLeftParen {
						break
					}

					prevPriority := priorities[top.Id]

					if (rightAssociative && priority < prevPriority) || (!rightAssociative && priority <= prevPriority) {
						// pop current operator
						operators = operators[:len(operators)-1]
						ret = append(ret, &top)
					} else {
						break
					}
				} // end of for len(operators) > 0

				operators = append(operators, token)
			} // end of if token == "("
		} // end of if isOperand(token)
	} // end of for _, token := range tokens

	// process remaining operators
	for len(operators) > 0 {
		// pop
		operator := operators[len(operators)-1]
		operators = operators[:len(operators)-1]

		if operator.Id == TokenLeftParen {
			return nil, errors.New("Mismatched parentheses found")
		}
		ret = append(ret, &operator)
	}

	return ret, nil
}

//===========================================================================
