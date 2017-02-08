package dsl

import (
	"encoding/json"
	"fmt"
	"log"
)

type TypeParser struct {
	symbolTable *SymbolTable
}

// A DeclBlock is a JSON object that is a map from symbol names
// to symbol types. The symbol type can be either a
// "struct decl"" or a "simple decl".
//
// A simple decl looks like these: int, []bool, [map][4]float
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

	p.symbolTable = NewSymbolTable()
	p.symbolTable.Init()

	// collect the symbols that are declared,
	// put them into the symbol table
	for name, _ := range *block {
		p.symbolTable.add(Symbol(name), nil)
	}

	table := map[string][]Token{}

	scanner := &Scanner{}

	var toks []Token

	for name, decl := range *block {

		switch decl.(type) {

		case map[string]interface{}:
			structDecl := decl.(map[string]interface{})
			for fieldName, xfieldDecl := range structDecl {
				fieldDecl := xfieldDecl.(string)
				toks, err = p.parseStringDecl(fieldName, &fieldDecl)
				if err != nil {
					return err
				}
				table[fieldName] = toks
				//p.symbolTable.add(name+"."+fieldName, dslType)
			}
			p.symbolTable.add(Symbol(name), nil)

		case string:
			stringDecl := decl.(string)
			toks, err := scanner.Scan(string(stringDecl))
			if err != nil {
				return err
			}
			log.Printf(">> %#v", toks)
			toks, err = p.parseStringDecl(name, &stringDecl)
			//p.symbolTable.add(name, dslType)
			if err != nil {
				return err
			}

		default:
			return fmt.Errorf("unknown type declation: %v", decl)
		}
	}
	return nil
}

func (p *TypeParser) parseStringDecl(name string, stringDecl *string) ([]Token, error) {

	scanner := Scanner{}

	toks, err := scanner.Scan(*stringDecl)
	if err != nil {
		return nil, err
	}
	log.Printf(">> %#v", toks)
	//	toks, err = p.parseStringDecl(name, stringDecl)
	//	if err != nil {
	//		return nil, err
	//	}
	//	p.symbolTable.add(name, dslType)

	return toks, nil
}

/*
	in := string(*stringDecl)
	in = strings.TrimSpace(in)

	var dslType Type

	arrayMatch, arrayLen := matchArrayTypePrefix(in)

	switch {

	case p.symbolTable.has(Symbol(in)):
		dslType = &SymbolType{Symbol: Symbol(in)}

	case strings.HasPrefix(in, "[map]"):
		i := len("[map]")
		rest := in[i:]
		valueType := p.parseStringDecl(name, &rest)
		keyType := &SymbolType{Symbol: "string"}
		dslType = &MapType{KeyType: keyType, ValueType: valueType}

	case strings.HasPrefix(in, "[]"):

		i := len("[]")
		rest := in[i:]
		elemType := p.parseStringDecl(name, &rest)
		dslType = &SliceType{ElemType: elemType}

	case arrayMatch:
		i := strings.Index(in, "]")
		rest := in[i+1:]
		elemType := p.parseStringDecl(name, &rest)
		dslType = &ArrayType{
			ElemType: elemType,
			Len:      arrayLen,
		}

	default:
		log.Printf("Invalid declaration for %s: %s", name, in)
		panic(9)
	}

	return dslType
*/
