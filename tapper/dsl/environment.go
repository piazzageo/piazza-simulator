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

func (env *Environment) setInt(name string, value int) {
	env.data[name] = value
}

func (env *Environment) getInt(name string) int {
	return env.data[name].(int)
}

func (env *Environment) setFloat(name string, value float64) {
	env.data[name] = value
}

func (env *Environment) getFloat(name string) float64 {
	return env.data[name].(float64)
}
