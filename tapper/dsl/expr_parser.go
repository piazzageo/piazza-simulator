package dsl

import "fmt"

type ExprParser struct {
	tokens []*Token
}

// TODO: <<, >>

// Parser converts tokens in RPN to a Node tree
func (ep *ExprParser) Parse(toks []*Token) (Node, error) {

	ep.tokens = toks

	ast, err := ep.buildTree()
	if err != nil {
		return nil, err
	}

	//log.Printf("%v\n", ast)

	return ast, nil
}

func (ep *ExprParser) pop() *Token {
	n := len(ep.tokens)
	x := ep.tokens[n-1]
	ep.tokens = ep.tokens[:n-1]
	return x
}

func (ep *ExprParser) buildTree() (Node, error) {

	var err error
	var left Node
	var right Node
	var out Node
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

		out = NewNodeMultiply(left, right)

	case TokenAdd:
		left, err = ep.buildTree()
		if err != nil {
			return nil, err
		}
		right, err = ep.buildTree()
		if err != nil {
			return nil, err
		}
		out = NewNodeAdd(left, right)

	case TokenSymbol:
		out = NewNodeSymbol(tok.Text)

	default:
		return nil, fmt.Errorf("Unknown token building ast: %d (\"%s\")", tok.Id, tok.Text)
	}

	return out, nil
}
