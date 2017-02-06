package dsl

import (
	"log"
	"regexp"
	"strconv"
	"strings"
)

type Parser struct {
	symbolTable *SymbolTable
}

type DeclBlock map[Symbol]Decl

type Decl interface{}
type StructDecl map[Symbol]Decl
type StringDecl string

var arrayRegexp *regexp.Regexp

func init() {
	arrayRegexp = regexp.MustCompile(`^\[(\d+)\]`)
}

func (p *Parser) Parse(in *DeclBlock) {

	p.symbolTable = NewSymbolTable()
	p.symbolTable.Init()

	// collect the symbols that are declared,
	// put them into the symbol table
	for k, _ := range *in {
		p.symbolTable.add(k, nil)
	}

	for k, v := range *in {
		switch v.(type) {

		case map[string]interface{}:
			structDecl := convertDeclToStructDecl(&v)
			dslType := p.parseStructDecl(k, structDecl)
			p.symbolTable.add(k, dslType)
			log.Printf("%s: %s", k, dslType)

		case string:
			stringDecl := convertDeclToStringDecl(&v)
			dslType := p.parseStringDecl(k, stringDecl)
			p.symbolTable.add(k, dslType)
			log.Printf("%s: %s", k, dslType)

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

func (p *Parser) parseStructDecl(name Symbol, structDecl *StructDecl) DslType {
	dslType := &StructDslType{
		Fields: map[Symbol]DslType{},
	}

	for k, v := range *structDecl {
		stringDecl := convertDeclToStringDecl(&v)
		fieldType := p.parseStringDecl(k, stringDecl)
		dslType.Fields[k] = fieldType
	}
	return dslType
}

func (p *Parser) parseStringDecl(name Symbol, stringDecl *StringDecl) DslType {

	in := string(*stringDecl)
	in = strings.TrimSpace(in)

	var dslType DslType

	arrayMatch, arrayLen := matchArrayPrefix(in)

	switch {

	case p.symbolTable.has(Symbol(in)):
		dslType = &SymbolDslType{Symbol: Symbol(in)}

	case strings.HasPrefix(in, "[map]"):
		i := len("[map]")
		rest := StringDecl(in[i:])
		valueType := p.parseStringDecl(name, &rest)
		keyType := &SymbolDslType{Symbol: "string"}
		dslType = &MapDslType{KeyType: keyType, ValueType: valueType}

	case strings.HasPrefix(in, "[]"):

		i := len("[]")
		rest := StringDecl(in[i:])
		elemType := p.parseStringDecl(name, &rest)
		dslType = &SliceDslType{ElemType: elemType}

	case arrayMatch:
		i := strings.Index(in, "]")
		rest := StringDecl(in[i+1:])
		elemType := p.parseStringDecl(name, &rest)
		dslType = &ArrayDslType{
			ElemType: elemType,
			Len:      arrayLen,
		}

	default:
		log.Printf("Invalid declaration for %s: %s", name, in)
		panic(9)
	}

	return dslType
}

func matchArrayPrefix(s string) (bool, int) {
	ok := arrayRegexp.Match([]byte(s))
	if !ok {
		return false, -1
	}
	sub := arrayRegexp.FindSubmatch([]byte(s))

	siz, err := strconv.Atoi(string(sub[1]))
	if err != nil {
		panic(err)
	}
	return true, siz
}
