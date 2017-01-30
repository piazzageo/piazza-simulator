package kvdatabase

import "fmt"
import "encoding/json"

//---------------------------------------------------------------------

type InMemDatabase struct {
	tables map[string]*InMemTable
}

type InMemTable struct {
	name string
	data map[string][]byte
}

func init() {
	var _ KVDatabase = (*InMemDatabase)(nil)
	var _ KVTable = (*InMemTable)(nil)
}

//---------------------------------------------------------------------

func (db *InMemDatabase) Open() error {
	db.tables = map[string]*InMemTable{}
	return nil
}

func (db *InMemDatabase) Close() error {
	db.tables = nil
	return nil
}

func (db *InMemDatabase) GetTableNames() ([]string, error) {
	names := make([]string, len(db.tables))
	i := 0
	for k, _ := range db.tables {
		names[i] = k
		i++
	}
	return names, nil
}

func (db *InMemDatabase) CreateTable(name string) (KVTable, error) {
	_, ok := db.tables[name]
	if ok {
		return nil, fmt.Errorf("table %s already exists", name)
	}

	table := &InMemTable{
		name: name,
		data: map[string][]byte{},
	}
	db.tables[name] = table
	return table, nil
}

func (db *InMemDatabase) GetTable(name string) (KVTable, error) {
	v, ok := db.tables[name]
	if !ok {
		return nil, fmt.Errorf("table %s does not exist", name)
	}
	return v, nil
}

func (db *InMemDatabase) DeleteTable(name string) error {
	_, ok := db.tables[name]
	if !ok {
		return fmt.Errorf("table %s does not exist", name)
	}
	delete(db.tables, name)
	return nil
}

//---------------------------------------------------------------------

func (table *InMemTable) Name() string {
	return table.name
}

func (table *InMemTable) Add(key string, object interface{}) error {
	_, ok := table.data[key]
	if ok {
		return fmt.Errorf("key %s already exists", key)
	}

	value, err := json.Marshal(object)
	if err != nil {
		return err
	}

	table.data[key] = value
	return nil
}

func (table *InMemTable) Get(key string, object interface{}) error {
	v, ok := table.data[key]
	if !ok {
		return fmt.Errorf("key %s not found", key)
	}

	err := json.Unmarshal(v, object)
	if err != nil {
		return err
	}

	return nil
}

func (table *InMemTable) Delete(key string) error {
	_, ok := table.data[key]
	if !ok {
		return fmt.Errorf("key %s not found", key)
	}
	delete(table.data, key)
	return nil

}

func (table *InMemTable) GetKeys() ([]string, error) {
	keys := make([]string, len(table.data))
	i := 0
	for k, _ := range table.data {
		keys[i] = k
		i++
	}
	return keys, nil
}
