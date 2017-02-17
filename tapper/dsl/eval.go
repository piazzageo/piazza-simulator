package dsl

type Eval struct {
}

func (e *Eval) Evaluate(expr ExprNode, env *Environment) (*ExprValue, error) {
	x := expr.Eval(env)
	return x, nil
}
