package symtab

// shunting yard implementation adapted from https://github.com/mgenware/go-shunting-yard

import (
	"errors"
	"fmt"
	"log"
	"strings"
	"text/scanner"
)

type Expression struct {
}

func ParseExpr(s string) error {
	tokens, err := Scan(s)
	if err != nil {
		return err
	}

	for _, v := range tokens {
		log.Printf("%s\n", v.String())
	}

	asts, err := Parse(tokens)
	if err != nil {
		return err
	}

	for _, ast := range asts {
		log.Printf("%s\n", ast.String())
	}

	return nil
}

//===========================================================================

// precedence of operators
var priorities map[string]int

// associativities of operators
var associativities map[string]bool

func init() {
	priorities = make(map[string]int, 0)
	associativities = make(map[string]bool, 0)

	priorities["+"] = 0
	priorities["-"] = 0
	priorities["*"] = -1
	priorities["/"] = -1
	priorities["^"] = -2
	priorities["<"] = -3 // CHECK
	priorities[">"] = -3 // CHECK
	priorities["|"] = -4 // CHECK

	// if not set, associativity will be false(left-associative)
}

type AST struct {
	Token interface{}
}

func (ast *AST) String() string {
	return fmt.Sprintf("%v", ast.Token)
}

func Parse(tokens []Token) ([]*AST, error) {
	var ret []*AST

	var operators []string
	for _, token := range tokens {
		if token.TokenId == -2 || token.TokenId == -3 {
			operandToken := &AST{Token: token}
			ret = append(ret, operandToken)
		} else {
			// check parentheses
			if token.Text == "(" {
				operators = append(operators, token.Text)
			} else if token.Text == ")" {
				foundLeftParenthesis := false
				// pop until "(" is fouund
				for len(operators) > 0 {
					oper := operators[len(operators)-1]
					operators = operators[:len(operators)-1]

					if oper == "(" {
						foundLeftParenthesis = true
						break
					} else {
						ret = append(ret, &AST{Token: oper})
					}
				}
				if !foundLeftParenthesis {
					return nil, errors.New("Mismatched parentheses found")
				}
			} else {
				// operator priority and associativity
				priority, ok := priorities[token.Text]
				if !ok {
					return nil, fmt.Errorf("Unknown operator: %v", token)
				}
				rightAssociative := associativities[token.Text]

				for len(operators) > 0 {
					top := operators[len(operators)-1]

					if top == "(" {
						break
					}

					prevPriority := priorities[top]

					if (rightAssociative && priority < prevPriority) || (!rightAssociative && priority <= prevPriority) {
						// pop current operator
						operators = operators[:len(operators)-1]
						ret = append(ret, &AST{Token: top})
					} else {
						break
					}
				} // end of for len(operators) > 0

				operators = append(operators, token.Text)
			} // end of if token == "("
		} // end of if isOperand(token)
	} // end of for _, token := range tokens

	// process remaining operators
	for len(operators) > 0 {
		// pop
		operator := operators[len(operators)-1]
		operators = operators[:len(operators)-1]

		if operator == "(" {
			return nil, errors.New("Mismatched parentheses found")
		}
		ret = append(ret, &AST{Token: operator})
	}

	return ret, nil
}

//===========================================================================

type Token struct {
	Line    int
	Column  int
	Text    string
	TokenId int
}

func (t *Token) String() string {
	return fmt.Sprintf("[%d:%d] %d %s", t.Line, t.Column, t.TokenId, t.Text)
}

func Scan(input string) ([]Token, error) {

	const src = `
	// This is scanned code.
	if a > 10 {
		someParsable = text
	}`

	tokens := []Token{}

	var s scanner.Scanner
	s.Init(strings.NewReader(input))
	var tok rune
	for {
		tok = s.Scan()
		if tok == scanner.EOF {
			break
		}

		//value := strings.TrimSpace(s.TokenText())
		//	fmt.Println("At position", s.Pos(), ":", s.TokenText(), "==", tok)

		token := Token{
			Line:    s.Pos().Line,
			Column:  s.Pos().Column,
			Text:    s.TokenText(),
			TokenId: int(tok),
		}
		tokens = append(tokens, token)
	}

	return tokens, nil
}
