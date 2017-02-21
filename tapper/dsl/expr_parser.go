package dsl

import (
	"fmt"
	"log"
	"strconv"
)

type ExprParser struct {
	tokens    []*Token
	typeTable *TypeTable
}

// TODO: <<, >>

// Parser converts tokens in RPN to a Node tree
func (ep *ExprParser) Parse(typeTable *TypeTable, toks []*Token) (ExprNode, error) {

	ep.tokens = toks
	ep.typeTable = typeTable

	ast, err := ep.buildTree()
	if err != nil {
		return nil, err
	}

	return ast, nil
}

func (ep *ExprParser) pop() *Token {
	n := len(ep.tokens)
	x := ep.tokens[n-1]
	ep.tokens = ep.tokens[:n-1]
	return x
}

func (ep *ExprParser) buildTree() (ExprNode, error) {

	var err error
	var left ExprNode
	var right ExprNode
	var out ExprNode
	tok := ep.pop()

	switch tok.Id {

	case TokenMultiply:
		left, err = ep.buildTree()
		if err != nil {
			return nil, err
		}
		right, err = ep.buildTree()
		if err != nil {
			return nil, err
		}

		out = NewExprNodeMultiply(left, right)

	case TokenAdd:
		left, err = ep.buildTree()
		if err != nil {
			return nil, err
		}
		right, err = ep.buildTree()
		if err != nil {
			return nil, err
		}
		out = NewExprNodeAdd(left, right)

	case TokenSubtract:
		left, err = ep.buildTree()
		if err != nil {
			return nil, err
		}
		right, err = ep.buildTree()
		if err != nil {
			return nil, err
		}
		out = NewExprNodeSubtract(left, right)

	case TokenSymbol:
		tnode := ep.typeTable.getStruct(StructName(tok.Text))
		if tnode == nil {
			return nil, fmt.Errorf("no declaration for symbol: %s", tok.Text)
		}
		baseType := InvalidType
		switch tnode.(type) {
		case *TypeNodeInt:
			baseType = IntType
		default:
			log.Printf("%s: %#T", tok.Text, tnode)
			panic(20)
		}
		out = NewExprNodeSymbolRef(tok.Text, baseType)

	case TokenInt:
		i, err := strconv.Atoi(tok.Text)
		if err != nil {
			return nil, err
		}
		out = NewExprNodeIntConstant(i)

	case TokenFloat:
		f, err := strconv.ParseFloat(tok.Text, 64)
		if err != nil {
			return nil, err
		}
		out = NewExprNodeFloatConstant(f)

	case TokenString:
		out = NewExprNodeStringConstant(tok.Text)

	case TokenBool:
		var b bool
		switch tok.Text {
		case "true":
			b = true
		case "false":
			b = false
		default:
			panic(23)
		}
		if err != nil {
			return nil, err
		}
		out = NewExprNodeBoolConstant(b)

	case TokenBitwiseXor:
		left, err = ep.buildTree()
		if err != nil {
			return nil, err
		}
		right, err = ep.buildTree()
		if err != nil {
			return nil, err
		}
		out = NewExprNodeBitwiseXor(left, right)

	case TokenBitwiseAnd:
		left, err = ep.buildTree()
		if err != nil {
			return nil, err
		}
		right, err = ep.buildTree()
		if err != nil {
			return nil, err
		}
		out = NewExprNodeBitwiseAnd(left, right)

	case TokenBitwiseOr:
		left, err = ep.buildTree()
		if err != nil {
			return nil, err
		}
		right, err = ep.buildTree()
		if err != nil {
			return nil, err
		}
		out = NewExprNodeBitwiseOr(left, right)

	case TokenLogicalOr:
		left, err = ep.buildTree()
		if err != nil {
			return nil, err
		}
		right, err = ep.buildTree()
		if err != nil {
			return nil, err
		}
		out = NewExprNodeLogicalOr(left, right)

	case TokenLogicalAnd:
		left, err = ep.buildTree()
		if err != nil {
			return nil, err
		}
		right, err = ep.buildTree()
		if err != nil {
			return nil, err
		}
		out = NewExprNodeLogicalAnd(left, right)

	case TokenEqualsEquals:
		left, err = ep.buildTree()
		if err != nil {
			return nil, err
		}
		right, err = ep.buildTree()
		if err != nil {
			return nil, err
		}
		out = NewExprNodeEqualsEquals(left, right)

	case TokenNotEquals:
		left, err = ep.buildTree()
		if err != nil {
			return nil, err
		}
		right, err = ep.buildTree()
		if err != nil {
			return nil, err
		}
		out = NewExprNodeNotEquals(left, right)

	case TokenGreaterOrEqualThan:
		left, err = ep.buildTree()
		if err != nil {
			return nil, err
		}
		right, err = ep.buildTree()
		if err != nil {
			return nil, err
		}
		out = NewExprNodeGreaterEqual(left, right)

	case TokenGreaterThan:
		left, err = ep.buildTree()
		if err != nil {
			return nil, err
		}
		right, err = ep.buildTree()
		if err != nil {
			return nil, err
		}
		out = NewExprNodeGreater(left, right)

	case TokenLessOrEqualThan:
		left, err = ep.buildTree()
		if err != nil {
			return nil, err
		}
		right, err = ep.buildTree()
		if err != nil {
			return nil, err
		}
		out = NewExprNodeLessEqual(left, right)

	case TokenLessThan:
		left, err = ep.buildTree()
		if err != nil {
			return nil, err
		}
		right, err = ep.buildTree()
		if err != nil {
			return nil, err
		}
		out = NewExprNodeLess(left, right)

	case TokenModulus:
		left, err = ep.buildTree()
		if err != nil {
			return nil, err
		}
		right, err = ep.buildTree()
		if err != nil {
			return nil, err
		}
		out = NewExprNodeModulus(left, right)

	case TokenDivide:
		left, err = ep.buildTree()
		if err != nil {
			return nil, err
		}
		right, err = ep.buildTree()
		if err != nil {
			return nil, err
		}
		out = NewExprNodeDivide(left, right)

	case TokenPeriod:
		left, err = ep.buildTree()
		if err != nil {
			return nil, err
		}
		right, err = ep.buildTree()
		if err != nil {
			return nil, err
		}
		out = NewExprNodeDivide(left, right)

	case TokenRightBracket:
		out = nil

	default:
		return nil, fmt.Errorf("expr_parser: Unknown token: %d (\"%s\")", tok.Id, tok.Text)
	}

	return out, nil
}
