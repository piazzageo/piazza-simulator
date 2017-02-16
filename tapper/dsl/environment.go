package dsl

type Environment struct {
	// symbol -> type
	tabl *TypeTable

	// symbol -> value
	data map[string]*ExprValue
}

func NewEnvironment(tabl *TypeTable) *Environment {
	env := &Environment{
		tabl: tabl,
		data: map[string]*ExprValue{},
	}
	return env
}

func (env *Environment) set(name string, value *ExprValue) {
	env.data[name] = value
}

func (env *Environment) get(name string) *ExprValue {
	return env.data[name]
}
