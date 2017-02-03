package symtab

import (
	"testing"

	"encoding/json"
)

func Test20(t *testing.T) {
	//assert := assert.New(t)

	// this is a DeclBlock, containing 3 interfaces (2 StructDecls and 1 StringDecl)
	s := `
{ 
    "Point": {
        "x": "float",
        "y": "float"
    },
    "MyType1": "int",
    "MyType2": "[map]int",
    "MyType3": "[map]Point",
    "MyType4": "[]Point",
    "MyType5": "[4]Point",
    "Message": {
        "alpha": "string",
        "beta": "Point",
        "gamma": "MyType2"
    }
}`

	declBlock := &DeclBlock{}

	err := json.Unmarshal([]byte(s), declBlock)
	if err != nil {
		panic(err)
	}

	p := &Parser{}
	p.Parse(declBlock)
}
