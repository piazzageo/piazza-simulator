package dsl

// shunting yard implementation adapted from https://github.com/mgenware/go-shunting-yard

import (
	"errors"
	"fmt"
)

type ExprTokenizer struct {
}

// TODO: <<, >>

// Tokenize returns the tokens for an expr, in RPN ("3,4,+")
func (ep *ExprTokenizer) Tokenize(s string) ([]*Token, error) {
	sc := &Scanner{}

	tokens, err := sc.Scan(s, false)
	if err != nil {
		return nil, err
	}

	//for _, v := range tokens {
	//	log.Printf("%s\n", v.String())
	//}

	toks, err := ep.makeRPN(tokens)
	if err != nil {
		return nil, err
	}

	//for _, tok := range toks {
	//log.Printf("%v\n", tok)
	//}

	return toks, nil
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
	priorities[TokenString] = 99
	priorities[TokenBool] = 99
	priorities[TokenFloat] = 99
	priorities[TokenInt] = 99
	priorities[TokenLeftBracket] = 99
	priorities[TokenRightBracket] = 99
	priorities[TokenBang] = 99

	priorities[TokenPeriod] = 6
	priorities[TokenMultiply] = 5
	priorities[TokenDivide] = 5
	priorities[TokenModulus] = 5
	priorities[TokenBitwiseAnd] = 5
	priorities[TokenAdd] = 4
	priorities[TokenSubtract] = 4
	priorities[TokenBitwiseOr] = 4
	priorities[TokenBitwiseXor] = 4
	priorities[TokenEqualsEquals] = 3
	priorities[TokenNotEquals] = 3
	priorities[TokenLessThan] = 3
	priorities[TokenLessOrEqualThan] = 3
	priorities[TokenGreaterThan] = 3
	priorities[TokenGreaterOrEqualThan] = 3
	priorities[TokenLogicalAnd] = 2
	priorities[TokenLogicalOr] = 1

	// if not set, associativity will be false(left-associative)
}

func (ep *ExprTokenizer) makeRPN(tokens []Token) ([]*Token, error) {
	var ret []*Token
	//log.Printf("==== %v", tokens)
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
					return nil, fmt.Errorf("expr_tokenizer: Unknown operator: %v", &token)
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

	//log.Printf("==== %v", ret)

	return ret, nil
}

//===========================================================================
