package kvdatabase

// KVDatabase is a set of named KVTables.
type KVDatabase interface {
	Open() error
	Close() error
	GetTableNames() ([]string, error)
	CreateTable(string) (KVTable, error)
	GetTable(string) (KVTable, error)
	DeleteTable(string) error
}

// KVTable is a key/value map of strings to interface{} objects.
// It is intended to store Go objects, serialized as JSON strings.
type KVTable interface {
	Name() string
	Add(string, interface{}) error
	Get(string, interface{}) error
	Delete(string) error
	GetKeys() ([]string, error)
}
