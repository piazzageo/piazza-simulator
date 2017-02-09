package dsl

import (
	"encoding/json"
	"fmt"
	"log"
)

type TypeParser struct {
	typeTable *TypeTable
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

	p.typeTable = NewTypeTable()
	p.typeTable.Init()

	// collect the symbols that are declared,
	// put them into the symbol table
	for name, decl := range *block {
		switch decl.(type) {
		case string:
			p.typeTable.insert(name)
		case map[string]interface{}:
			p.typeTable.insert(name)
			for fieldName, _ := range decl.(map[string]interface{}) {
				p.typeTable.insert(name + "." + fieldName)
			}
		default:
			return fmt.Errorf("bad decl type: %T", decl)
		}
	}

	err = p.tokenize(block)
	if err != nil {
		return err
	}

	err = p.parseAllTokens()
	return nil
}

func (p *TypeParser) parseAllTokens() error {

	tt := p.typeTable
	for _, v := range tt.Types {
		err := p.parseTokens(v)
		if err != nil {
			return err
		}
	}
	return nil
}

func (p *TypeParser) parseTokens(tte *TypeTableEntry) error {
	log.Printf("=== %v", tte)
	if tte.Tokens == nil {
		return nil
	}

	out, err := p.parsePiece(tte.Tokens)
	if err != nil {
		return err
	}
	tte.Node = out
	return nil
}

func (p *TypeParser) parsePiece(toks []Token) (TNode, error) {

	t0 := toks[0]
	t1ok := len(toks) > 1
	//t2ok := len(toks) > 2

	var out TNode

	switch t0.Id {
	case TokenSymbol:
		if t1ok {
			return nil, fmt.Errorf("extra token after %v\n\t%v", t0, toks[1])
		}
		out = &TNodeSymbol{Symbol: t0.Text}
	case TokenTypeSlice:
		if !t1ok {
			return nil, fmt.Errorf("no token after %v", t0)
		}
		next, err := p.parsePiece(toks[1:])
		if err != nil {
			return nil, err
		}
		out = &TNodeSlice{ElemType: next}
	default:
		return nil, fmt.Errorf("unhandled token: " + t0.String())
	}

	return out, nil
}

func (p *TypeParser) tokenize(block *DeclBlock) error {
	var err error

	for name, decl := range *block {

		switch decl.(type) {

		case map[string]interface{}:
			structDecl := decl.(map[string]interface{})
			for fieldName, xfieldDecl := range structDecl {
				fieldDecl := xfieldDecl.(string)
				err = p.parseStringDecl(name+"."+fieldName, &fieldDecl)
				if err != nil {
					return err
				}
			}
			p.typeTable.get(name).Tokens = nil
		case string:
			stringDecl := decl.(string)
			err = p.parseStringDecl(name, &stringDecl)
			if err != nil {
				return err
			}
			//log.Printf(">> %#v", toks)

		default:
			return fmt.Errorf("unknown type declation: %v", decl)
		}
	}

	return nil
}

func (p *TypeParser) parseStringDecl(name string, stringDecl *string) error {

	scanner := Scanner{}

	toks, err := scanner.Scan(*stringDecl)
	if err != nil {
		return err
	}

	p.typeTable.get(name).Tokens = toks

	return nil
}
