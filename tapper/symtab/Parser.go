package symtab

import "log"
import "strings"

type Parser struct {
	symbolTable *SymbolTable
}

type DeclBlock map[Symbol]Decl

type Decl interface{}
type StructDecl map[Symbol]Decl
type StringDecl string

func (p *Parser) Parse(in *DeclBlock) {

	p.symbolTable = NewSymbolTable()
	p.symbolTable.Init()

	// collect the symbols that are declared
	for k, _ := range *in {
		p.symbolTable.add(k, nil)
	}

	for k, v := range *in {
		switch v.(type) {

		case map[string]interface{}:
			structDecl := convertDeclToStructDecl(&v)
			p.symbolTable.add(k, nil)
			s := p.parseStructDecl(k, structDecl)
			log.Printf("  XXX*** %s --- %s", k, s)
		case string:
			stringDecl := convertDeclToStringDecl(&v)
			p.symbolTable.add(k, nil)
			s := p.parseStringDecl(k, stringDecl)
			log.Printf("  XXX*** %s --- %s", k, s)

		default:
			panic(99)
		}
	}
}

func convertDeclToStructDecl(d *Decl) *StructDecl {
	s := StructDecl{}
	dd := (*d).(map[string]interface{})
	for k, v := range dd {
		s[Symbol(k)] = v.(Decl)
	}
	return &s
}

func convertDeclToStringDecl(d *Decl) *StringDecl {
	s := (*d).(string)
	stringDecl := StringDecl(s)
	return &stringDecl
}

func (p *Parser) parseStructDecl(name Symbol, structDecl *StructDecl) string {
	//log.Printf("PARSING Struct %s: %v", name, *structDecl)
	s := ""
	for k, v := range *structDecl {
		stringDecl := convertDeclToStringDecl(&v)
		result := p.parseStringDecl(k, stringDecl)
		s += string(k) + "/" + string(result) + " "
	}
	return s
}

func (p *Parser) parseStringDecl(name Symbol, stringDecl *StringDecl) string {
	//log.Printf("PARSING String %s: %v", name, *stringDecl)

	in := string(*stringDecl)

	in = strings.TrimSpace(in)

	s := ""

	switch {
	case p.symbolTable.has(Symbol(in)):
		s += "(" + strings.ToUpper(in) + ")"
	case strings.HasPrefix(in, "[map]"):
		s += "(MAP "
		i := len("[map]")
		rest := StringDecl(in[i:])
		s += p.parseStringDecl(name, &rest)
		s += ")"
	case strings.HasPrefix(in, "[]"):
		s += "([] "
		i := len("[]")
		rest := StringDecl(in[i:])
		s += p.parseStringDecl(name, &rest)
		s += ")"
	case strings.HasPrefix(in, "[4]"):
		s += "([4] "
		i := len("[4]")
		rest := StringDecl(in[i:])
		s += p.parseStringDecl(name, &rest)
		s += ")"
	default:
		log.Printf("Invalid declaration for %s: %s", name, in)
		panic(9)
	}
	return s
}
