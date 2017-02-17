package dsl

// shunting yard implementation adapted from https://github.com/mgenware/go-shunting-yard

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"text/scanner"
)

//===========================================================================
var arrayTypeRegexp *regexp.Regexp

func init() {
	arrayTypeRegexp = regexp.MustCompile(`^\[(\d+)\]`)
}

//===========================================================================

type Scanner struct{}

func (s *Scanner) Scan(str string, asDecl bool) ([]Token, error) {
	tokens, err := s.scan(str)
	if err != nil {
		return nil, err
	}

	if asDecl {
		tokens, err = declPeepholer(tokens)
		if err != nil {
			return nil, err
		}
	} else {
		tokens, err = exprPeepholer(tokens)
		if err != nil {
			return nil, err
		}
	}
	return tokens, nil
}

//===========================================================================

func (s *Scanner) scan(input string) ([]Token, error) {

	tokens := []Token{}

	var sx scanner.Scanner
	sx.Init(strings.NewReader(input))

	var tok rune

	for {
		tok = sx.Scan()
		if tok == scanner.EOF {
			break
		}

		tokenid := convertId(tok)
		if tokenid == TokenInvalid {
			return nil, fmt.Errorf("Unknown token type %v (%s)", tok, sx.TokenText())
		}
		token := Token{
			Line:   sx.Pos().Line,
			Column: sx.Pos().Column,
			Text:   sx.TokenText(),
			Id:     tokenid,
		}
		//log.Printf("TOK: %s", token.String())
		tokens = append(tokens, token)
	}

	return tokens, nil
}

func declPeepholer(tokens []Token) ([]Token, error) {

	result := []Token{}
	push := func(t Token) {
		result = append(result, t)
	}

	// combine these two tokens:
	//   ||, &&
	//   []
	// combine these three tokens:
	//   [map]
	//   [int]

	// the last index in use
	last := len(tokens) - 1

	for i := 0; i <= last; i += 0 { // TODO
		atLeastTwoLeft := i <= last-1
		atLeastThreeLeft := i <= last-2

		if atLeastThreeLeft {
			s := tokens[i].Text + tokens[i+1].Text + tokens[i+2].Text
			arrayMatch, arrayLen := matchArrayTypePrefix(s)

			switch {
			case s == "[map]":
				t := Token{
					Line:   tokens[i].Line,
					Column: tokens[i].Column,
					Text:   s,
					Id:     TokenTypeMap,
				}
				push(t)
				i += 3
				continue

			case arrayMatch: // [123]
				t := Token{
					Line:   tokens[i].Line,
					Column: tokens[i].Column,
					Text:   s,
					Id:     TokenTypeArray,
					Value:  arrayLen,
				}
				push(t)
				i += 3
				continue
			}
		}

		if atLeastTwoLeft {
			s := tokens[i].Text + tokens[i+1].Text

			switch s {
			case "[]":
				t := Token{
					Line:   tokens[i].Line,
					Column: tokens[i].Column,
					Text:   s,
					Id:     TokenTypeSlice,
				}
				push(t)
				i += 2
				continue
			}
		}

		// no peephole match
		push(tokens[i])
		i++
	}

	return result, nil
}

func exprPeepholer(tokens []Token) ([]Token, error) {
	result := []Token{}
	push := func(t Token) {
		result = append(result, t)
	}

	// combine these two tokens:
	//   ||, &&

	// the last index in use
	last := len(tokens) - 1

	for i := 0; i <= last; i += 0 { // TODO
		atLeastTwoLeft := i <= last-1

		switch {
		case tokens[i].Id == TokenSymbol &&
			(tokens[i].Text == "true" || tokens[i].Text == "false"):
			t := Token{
				Line:   tokens[i].Line,
				Column: tokens[i].Column,
				Text:   tokens[i].Text,
				Id:     TokenBool,
			}
			push(t)
			i += 1
			continue
		}

		if atLeastTwoLeft {
			t1 := tokens[i].Id
			t2 := tokens[i+1].Id

			switch {

			case t1 == TokenGreaterThan && t2 == TokenEquals:
				t := Token{
					Line:   tokens[i].Line,
					Column: tokens[i].Column,
					Text:   ">=",
					Id:     TokenGreaterOrEqualThan,
				}
				push(t)
				i += 2
				continue

			case t1 == TokenLessThan && t2 == TokenEquals:
				t := Token{
					Line:   tokens[i].Line,
					Column: tokens[i].Column,
					Text:   "<=",
					Id:     TokenLessOrEqualThan,
				}
				push(t)
				i += 2
				continue

			case t1 == TokenBitwiseOr && t2 == TokenBitwiseOr:
				t := Token{
					Line:   tokens[i].Line,
					Column: tokens[i].Column,
					Text:   "||",
					Id:     TokenLogicalOr,
				}
				push(t)
				i += 2
				continue

			case t1 == TokenBitwiseAnd && t2 == TokenBitwiseAnd:
				t := Token{
					Line:   tokens[i].Line,
					Column: tokens[i].Column,
					Text:   "&&",
					Id:     TokenLogicalAnd,
				}
				push(t)
				i += 2
				continue

			case t1 == TokenEquals && t2 == TokenEquals:
				t := Token{
					Line:   tokens[i].Line,
					Column: tokens[i].Column,
					Text:   "==",
					Id:     TokenEqualsEquals,
				}
				push(t)
				i += 2
				continue

			case t1 == TokenBang && t2 == TokenEquals:
				t := Token{
					Line:   tokens[i].Line,
					Column: tokens[i].Column,
					Text:   "!=",
					Id:     TokenNotEquals,
				}
				push(t)
				i += 2
				continue
			}
		}

		// no peephole match
		push(tokens[i])
		i++
	}

	return result, nil
}

func matchArrayTypePrefix(s string) (bool, int) {
	ok := arrayTypeRegexp.Match([]byte(s))
	if !ok {
		return false, -1
	}
	sub := arrayTypeRegexp.FindSubmatch([]byte(s))

	siz, err := strconv.Atoi(string(sub[1]))
	if err != nil || siz < 1 {
		panic(err)
	}
	return true, siz
}
