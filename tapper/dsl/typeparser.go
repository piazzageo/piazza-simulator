package dsl

import (
	"encoding/json"
	"fmt"
)

type TypeParser struct {
	typeTable *TypeTable
}

func NewTypeParser() (*TypeParser, error) {
	typeTable := NewTypeTable()
	err := typeTable.Init()
	if err != nil {
		return nil, err
	}

	tp := &TypeParser{
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
func (p *TypeParser) ParseJson(s string) error {
	declBlock := &DeclBlock{}
	err := json.Unmarshal([]byte(s), declBlock)
	if err != nil {
		return err
	}

	return p.Parse(declBlock)
}

// Parse takes a declaration block expressed as a DeclBlock object and parses it.
func (p *TypeParser) Parse(block *DeclBlock) error {
	var err error

	err = p.declareSymbols(block)
	if err != nil {
		return err
	}

	err = p.parseBlock(block)
	if err != nil {
		return err
	}

	return nil
}

// collect the symbols that are declared,
// put them into the symbol table
func (p *TypeParser) declareSymbols(block *DeclBlock) error {
	var err error

	for name, decl := range *block {
		switch decl.(type) {
		case string:
			err = p.typeTable.add(name)
			if err != nil {
				return err
			}
		case map[string]interface{}:
			err = p.typeTable.add(name)
			if err != nil {
				return err
			}
			for fieldName, _ := range decl.(map[string]interface{}) {
				err = p.typeTable.add(name + "." + fieldName)
				if err != nil {
					return err
				}
			}
		default:
			return fmt.Errorf("bad decl type: %T", decl)
		}
	}

	return nil
}

func (p *TypeParser) parseBlock(block *DeclBlock) error {
	var err error
	var tnode TNode

	for name, decl := range *block {

		switch decl.(type) {

		case map[string]interface{}:
			structDecl := decl.(map[string]interface{})
			for fieldName, xfieldDecl := range structDecl {
				fieldDecl := xfieldDecl.(string)
				tnode, err = p.parseDecl(name+"."+fieldName, fieldDecl)
				if err != nil {
					return err
				}
				p.typeTable.set(name+"."+fieldName, tnode)
			}

		case string:
			stringDecl := decl.(string)
			tnode, err = p.parseDecl(name, stringDecl)
			if err != nil {
				return err
			}
			p.typeTable.set(name, tnode)
			//log.Printf("$$ %s $$ %v", name, tnode)

		default:
			return fmt.Errorf("unknown type declation: %v", decl)
		}
	}

	return nil
}

// given a string like "[map][]float", return the TNode tree for it
func (p *TypeParser) parseDecl(name string, stringDecl string) (TNode, error) {

	scanner := Scanner{}

	toks, err := scanner.Scan(stringDecl)
	if err != nil {
		return nil, err
	}

	tnode, err := p.parseDeclToken(toks)
	if err != nil {
		return nil, err
	}

	//log.Printf("$$ %s $$ %s $$ %v", name, stringDecl, tnode)
	return tnode, nil
}

func (p *TypeParser) parseDeclToken(toks []Token) (TNode, error) {

	t0 := toks[0]
	t1ok := len(toks) > 1
	//t2ok := len(toks) > 2

	var out TNode

	switch t0.Id {

	case TokenSymbol:
		if t1ok {
			return nil, fmt.Errorf("extra token after %v\n\t%v", t0, toks[1])
		}
		if p.typeTable.isBuiltin(t0.Text) {
			out = p.typeTable.get(t0.Text).Node
		} else {
			out = &TNodeUserType{Name: t0.Text}
		}

	case TokenTypeSlice:
		if !t1ok {
			return nil, fmt.Errorf("no token after %v", t0)
		}
		next, err := p.parseDeclToken(toks[1:])
		if err != nil {
			return nil, err
		}
		out = &TNodeSlice{ElemType: next}

	case TokenTypeMap:
		if !t1ok {
			return nil, fmt.Errorf("no token after %v", t0)
		}
		next, err := p.parseDeclToken(toks[1:])
		if err != nil {
			return nil, err
		}
		out = &TNodeMap{KeyType: &TNodeString{}, ValueType: next}

	case TokenTypeArray:
		if !t1ok {
			return nil, fmt.Errorf("no token after %v", t0)
		}
		next, err := p.parseDeclToken(toks[1:])
		if err != nil {
			return nil, err
		}
		out = &TNodeArray{ElemType: next, Len: t0.Value.(int)}

	default:
		return nil, fmt.Errorf("unhandled token: " + t0.String())
	}

	return out, nil
}
