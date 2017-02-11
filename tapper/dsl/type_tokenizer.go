package dsl

import (
	"encoding/json"
	"fmt"
)

type TypeTokenizer struct {
	typeTable *TypeTable
}

func NewTypeTokenizer() (*TypeTokenizer, error) {
	typeTable := NewTypeTable()
	err := typeTable.Init()
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

	err = p.declareSymbols(block)
	if err != nil {
		return nil, err
	}

	err = p.parseBlock(block)
	if err != nil {
		return nil, err
	}

	return p.typeTable, nil
}

// collect the symbols that are declared,
// put them into the symbol table
func (p *TypeTokenizer) declareSymbols(block *DeclBlock) error {
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
			fs := map[string]bool{}
			for fieldName, _ := range decl.(map[string]interface{}) {
				err = p.typeTable.add(name + "." + fieldName)
				if err != nil {
					return err
				}
				fs[fieldName] = true
			}
			err = p.typeTable.setNode(name, NewNodeStructType(fs))
		default:
			return fmt.Errorf("bad decl type: %T", decl)
		}
	}

	return nil
}

func (p *TypeTokenizer) parseBlock(block *DeclBlock) error {
	var err error
	var tnode Node

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
				p.typeTable.setNode(name+"."+fieldName, tnode)
			}

		case string:
			stringDecl := decl.(string)
			tnode, err = p.parseDecl(name, stringDecl)
			if err != nil {
				return err
			}
			p.typeTable.setNode(name, tnode)
			//log.Printf("$$ %s $$ %v", name, tnode)

		default:
			return fmt.Errorf("unknown type declation: %v", decl)
		}
	}

	return nil
}

func (p *TypeTokenizer) parseDecl(name string, decl string) (Node, error) {

	scanner := Scanner{}

	toks, err := scanner.Scan(decl)
	if err != nil {
		return nil, err
	}

	err = p.typeTable.setToken(name, toks)
	if err != nil {
		return nil, err
	}

	tp := &TypeParser{}

	tnode, err := tp.Parse(toks, p.typeTable)
	if err != nil {
		return nil, err
	}

	//log.Printf("$$ %s $$ %s $$ %v", name, stringDecl, tnode)
	return tnode, nil
}
