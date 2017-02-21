package dsl

type Environment struct {
	values map[string]*ExprValue
}

func NewEnvironment(typeTable *TypeTable) *Environment {
	env := &Environment{
		values: map[string]*ExprValue{},
	}
	return env
}

func (env *Environment) set(name string, value *ExprValue) {
	env.values[name] = value
}

func (env *Environment) get(name string) *ExprValue {
	return env.values[name]
}
