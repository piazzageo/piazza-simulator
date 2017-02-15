package dsl

import (
	"encoding/json"
	"fmt"
)

type TypeTokenizer struct {
	typeTable *TypeTable
}

func NewTypeTokenizer() (*TypeTokenizer, error) {
	typeTable, err := NewTypeTable()
	if err != nil {
		return nil, err
	}

	tp := &TypeTokenizer{
		typeTable: typeTable,
	}

	return tp, nil
}

// A DeclBlock is a JSON object that is a map from symbol names
// to symbol types. The symbol type can be either a
// "struct decl"" or a "simple decl".
//
// A simple decl looks like these:
//     int
//     []bool
//     [map][4]float
//
// A struct decl looks like these:
//     { "a" : "[map]int", "b" : "[map]int" }
//     { "user": "string" }
type DeclBlock map[string]interface{}

// ParseJson takes a declaration block expressed as JSON string and parses it.
func (p *TypeTokenizer) ParseJson(s string) (*TypeTable, error) {
	declBlock := &DeclBlock{}
	err := json.Unmarshal([]byte(s), declBlock)
	if err != nil {
		return nil, err
	}

	return p.Parse(declBlock)
}

// Parse takes a declaration block expressed as a DeclBlock object and parses it.
func (p *TypeTokenizer) Parse(block *DeclBlock) (*TypeTable, error) {
	var err error

	err = p.parseBlock(block)
	if err != nil {
		return nil, err
	}

	return p.typeTable, nil
}

func (p *TypeTokenizer) parseBlock(block *DeclBlock) error {
	var err error
	var tnode TypeNode

	for name, decl := range *block {

		switch decl.(type) {

		case map[string]interface{}:
			structNode := NewTypeNodeStruct()

			structDecl := decl.(map[string]interface{})
			for fieldName, xfieldDecl := range structDecl {
				fieldDecl := xfieldDecl.(string)
				tnode, err = p.parseDecl(name+"."+fieldName, fieldDecl)
				if err != nil {
					return err
				}
				structNode.Fields[fieldName] = NewTypeNodeField(fieldName, tnode)
			}
			err = p.typeTable.addNode(name, structNode)
			if err != nil {
				return err
			}

		case string:
			stringDecl := decl.(string)
			tnode, err = p.parseDecl(name, stringDecl)
			if err != nil {
				return err
			}
			err = p.typeTable.addNode(name, tnode)
			if err != nil {
				return err
			}
			//log.Printf("$$ %s $$ %v", name, tnode)

		default:
			return fmt.Errorf("unknown type declation: %v", decl)
		}
	}

	return nil
}

func (p *TypeTokenizer) parseDecl(name string, decl string) (TypeNode, error) {

	scanner := Scanner{}

	toks, err := scanner.Scan(decl)
	if err != nil {
		return nil, err
	}

	tnode, err := parseTheTokens(toks, p.typeTable)
	if err != nil {
		return nil, err
	}

	//log.Printf("$$ %s $$ %s $$ %v", name, stringDecl, tnode)
	return tnode, nil
}

//---------------------------------------------------------------------------

func parseTheTokens(toks []Token, typeTable *TypeTable) (TypeNode, error) {

	t0 := toks[0]
	t1ok := len(toks) > 1
	//t2ok := len(toks) > 2

	var out TypeNode

	switch t0.Id {

	case TokenSymbol:
		if t1ok {
			return nil, fmt.Errorf("extra token after %v\n\t%v", t0, toks[1])
		}
		if typeTable.isBuiltin(t0.Text) {
			out = typeTable.getNode(t0.Text)
		} else {
			out = NewTypeNodeName(t0.Text)
		}

	case TokenTypeSlice:
		if !t1ok {
			return nil, fmt.Errorf("no token after %v", t0)
		}
		next, err := parseTheTokens(toks[1:], typeTable)
		if err != nil {
			return nil, err
		}
		out = NewTypeNodeSlice(next)

	case TokenTypeMap:
		if !t1ok {
			return nil, fmt.Errorf("no token after %v", t0)
		}
		next, err := parseTheTokens(toks[1:], typeTable)
		if err != nil {
			return nil, err
		}
		out = NewTypeNodeMap(NewTypeNodeString(), next)

	case TokenTypeArray:
		if !t1ok {
			return nil, fmt.Errorf("no token after %v", t0)
		}
		next, err := parseTheTokens(toks[1:], typeTable)
		if err != nil {
			return nil, err
		}
		out = NewTypeNodeArray(next, t0.Value.(int))

	default:
		return nil, fmt.Errorf("unhandled token: " + t0.String())
	}

	return out, nil
}
