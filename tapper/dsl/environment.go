package dsl

type Environment struct {
	// symbol -> type
	tabl *TypeTable

	// symbol -> value
	data map[string]interface{}
}

func NewEnvironment(tabl *TypeTable) *Environment {
	env := &Environment{
		tabl: tabl,
		data: map[string]interface{}{},
	}
	return env
}

func (env *Environment) set(name string, value interface{}) {
	env.data[name] = value
}

func (env *Environment) get(name string) interface{} {
	return env.data[name]
}
