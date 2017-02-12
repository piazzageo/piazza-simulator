package dsl

type Eval struct {
}

func (e *Eval) Evaluate(expr Node, env *Environment) (interface{}, error) {
	return nil, nil
}

//func (n *NodeSymbol) Eval(env *Environment) (interface{}, error) {
//	val := &Value{data: env.data[n.Value()], datatype: NewNodeNumberType(FlavorS32)}
//	return fmt.Sprintf("(* %v %v)", n.left, n.right)
//}
