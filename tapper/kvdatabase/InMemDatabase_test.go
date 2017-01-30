package kvdatabase

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//--------------------------

func Test00(t *testing.T) {
	assert := assert.New(t)
	assert.True(!false)
}

func Test01Database(t *testing.T) {
	assert := assert.New(t)

	var err error

	db := InMemDatabase{}
	err = db.Open()
	assert.NoError(err)
	defer db.Close()

	table0, err := db.CreateTable("table0")
	assert.NoError(err)
	tmp, err := db.GetTable("table0")
	assert.NoError(err)
	assert.Equal("table0", table0.Name())
	assert.Equal("table0", tmp.Name())

	table1, err := db.CreateTable("table1")
	assert.NoError(err)
	tmp, err = db.GetTable("table1")
	assert.NoError(err)
	assert.Equal("table1", table1.Name())
	assert.Equal("table1", tmp.Name())
	assert.NotEqual(table1.Name(), table0.Name())

	names, err := db.GetTableNames()
	assert.Len(names, 2)
	ok1 := names[0] == table0.Name() && names[1] == table1.Name()
	ok2 := names[0] == table1.Name() && names[1] == table0.Name()
	assert.True((ok1 && !ok2) || (!ok1 && ok2))

	err = db.DeleteTable("tableX")
	assert.Error(err)
	err = db.DeleteTable("table0")
	assert.NoError(err)
	names, err = db.GetTableNames()
	assert.Len(names, 1)
	assert.Equal("table1", names[0])

	tmp, err = db.CreateTable("table1")
	assert.Error(err)

	err = db.DeleteTable("table1")
	assert.NoError(err)
	names, err = db.GetTableNames()
	assert.Len(names, 0)
}

func Test02Table(t *testing.T) {
	assert := assert.New(t)

	var err error

	db := InMemDatabase{}
	err = db.Open()
	assert.NoError(err)
	defer db.Close()

	table, err := db.CreateTable("table")
	assert.NoError(err)

	type T struct {
		S string
		I int
	}

	err = table.Add("key0", &T{S: "value0", I: 10})
	assert.NoError(err)
	err = table.Add("key1", &T{S: "value1", I: 11})
	assert.NoError(err)
	err = table.Add("key0", &T{S: "valueX", I: -999})
	assert.Error(err)

	v := T{}
	err = table.Get("key0", &v)
	assert.NoError(err)
	assert.EqualValues(T{S: "value0", I: 10}, v)
	err = table.Get("key1", &v)
	assert.NoError(err)
	assert.EqualValues(T{S: "value1", I: 11}, v)
	err = table.Get("key2", &v)
	assert.Error(err)

	ks, err := table.GetKeys()
	assert.Len(ks, 2)
	ok1 := ks[0] == "key0" && ks[1] == "key1"
	ok2 := ks[0] == "key1" && ks[1] == "key0"
	assert.True((ok1 && !ok2) || (!ok1 && ok2))

	err = table.Delete("key0")
	assert.NoError(err)
	err = table.Get("key0", &v)
	assert.Error(err)
	err = table.Get("key1", &v)
	assert.NoError(err)
}
