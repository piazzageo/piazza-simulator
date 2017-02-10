package dsl

import "fmt"

type TypeParser struct {
}

func (p *TypeParser) Parse(toks []Token, typeTable *TypeTable) (TNode, error) {

	t0 := toks[0]
	t1ok := len(toks) > 1
	//t2ok := len(toks) > 2

	var out TNode

	switch t0.Id {

	case TokenSymbol:
		if t1ok {
			return nil, fmt.Errorf("extra token after %v\n\t%v", t0, toks[1])
		}
		if typeTable.isBuiltin(t0.Text) {
			out = typeTable.get(t0.Text).Node
		} else {
			out = &TNodeUserType{Name: t0.Text}
		}

	case TokenTypeSlice:
		if !t1ok {
			return nil, fmt.Errorf("no token after %v", t0)
		}
		next, err := p.Parse(toks[1:], typeTable)
		if err != nil {
			return nil, err
		}
		out = &TNodeSlice{ElemType: next}

	case TokenTypeMap:
		if !t1ok {
			return nil, fmt.Errorf("no token after %v", t0)
		}
		next, err := p.Parse(toks[1:], typeTable)
		if err != nil {
			return nil, err
		}
		out = &TNodeMap{KeyType: &TNodeString{}, ValueType: next}

	case TokenTypeArray:
		if !t1ok {
			return nil, fmt.Errorf("no token after %v", t0)
		}
		next, err := p.Parse(toks[1:], typeTable)
		if err != nil {
			return nil, err
		}
		out = &TNodeArray{ElemType: next, Len: t0.Value.(int)}

	default:
		return nil, fmt.Errorf("unhandled token: " + t0.String())
	}

	return out, nil
}
