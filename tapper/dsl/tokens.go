package dsl

import "fmt"

//---------------------------------------------------------------------------

type TokenId int

const (
	TokenInvalid TokenId = iota
	TokenBool
	TokenSymbol
	TokenInt
	TokenFloat
	TokenString
	TokenEquals
	TokenAdd
	TokenMultiply
	TokenBang
	TokenBitwiseOr
	TokenBitwiseAnd
	TokenLeftParen
	TokenRightParen
	TokenLeftBracket
	TokenRightBracket
	TokenPeriod
	TokenSubtract
	TokenDivide
	TokenBitwiseXor
	TokenModulus
	TokenGreaterThan
	TokenLessThan

	// derived, for exprs
	TokenEqualsEquals
	TokenNotEquals
	TokenLogicalAnd
	TokenLogicalOr

	// derived, for decls
	TokenTypeMap
	TokenTypeArray
	TokenTypeSlice

	// not yet used
	TokenGreaterOrEqualThan
	TokenLessOrEqualThan
)

type Token struct {
	Line   int
	Column int
	Text   string
	Id     TokenId
	Value  interface{}
}

func (t *Token) String() string {
	s := fmt.Sprintf("[%d:%d] id=%d text=\"%s\"", t.Line, t.Column, t.Id, t.Text)
	if t.Value != nil {
		s += fmt.Sprintf(" value=<%v>", t.Value)
	}
	return s
}

func convertId(r rune) TokenId {
	switch r {
	case -2:
		return TokenSymbol
	case -3:
		return TokenInt
	case -4:
		return TokenFloat
	case -6:
		return TokenString
	case 33:
		return TokenBang
	case 37:
		return TokenModulus
	case 38:
		return TokenBitwiseAnd
	case 40:
		return TokenLeftParen
	case 41:
		return TokenRightParen
	case 42:
		return TokenMultiply
	case 43:
		return TokenAdd
	case 45:
		return TokenSubtract
	case 46:
		return TokenPeriod
	case 47:
		return TokenDivide
	case 60:
		return TokenLessThan
	case 61:
		return TokenEquals
	case 62:
		return TokenGreaterThan
	case 91:
		return TokenLeftBracket
	case 93:
		return TokenRightBracket
	case 94:
		return TokenBitwiseXor
	case 124:
		return TokenBitwiseOr
	default:
		return TokenInvalid
	}
}
