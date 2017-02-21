package dsl

import (
	"fmt"
	"strconv"
)

//---------------------------------------------------------------------

type ExprValue struct {
	Type  BaseType
	Value interface{}
}

type ExprNode interface {
	String() string
	Eval(*Environment) *ExprValue
}

//---------------------------------------------------------------------

type ExprNodeMultiply struct {
	left  ExprNode
	right ExprNode
}

func NewExprNodeMultiply(left ExprNode, right ExprNode) *ExprNodeMultiply {
	n := &ExprNodeMultiply{
		left:  left,
		right: right,
	}
	return n
}

func (n *ExprNodeMultiply) Eval(env *Environment) *ExprValue {
	l := n.left.Eval(env)
	r := n.right.Eval(env)

	if l.Type != r.Type {
		panic(19)
	}
	switch l.Type {
	case IntType:
		return &ExprValue{Type: IntType, Value: l.Value.(int) * r.Value.(int)}
	case FloatType:
		return &ExprValue{Type: FloatType, Value: l.Value.(float64) * r.Value.(float64)}
	}
	panic(19)
}

func (n *ExprNodeMultiply) String() string {
	return fmt.Sprintf("MULTIPLY(%v, %v)", n.left, n.right)
}

//---------------------------------------------------------------------

type ExprNodeDivide struct {
	left  ExprNode
	right ExprNode
}

func NewExprNodeDivide(left ExprNode, right ExprNode) *ExprNodeDivide {
	n := &ExprNodeDivide{
		left:  left,
		right: right,
	}
	return n
}

func (n *ExprNodeDivide) Eval(env *Environment) *ExprValue {
	l := n.left.Eval(env)
	r := n.right.Eval(env)

	if l.Type != r.Type {
		panic(19)
	}
	switch l.Type {
	case IntType:
		return &ExprValue{Type: IntType, Value: l.Value.(int) / r.Value.(int)}
	case FloatType:
		return &ExprValue{Type: FloatType, Value: l.Value.(float64) / r.Value.(float64)}
	}
	panic(19)
}

func (n *ExprNodeDivide) String() string {
	return fmt.Sprintf("DIVIDE(%v, %v)", n.left, n.right)
}

//---------------------------------------------------------------------

type ExprNodeAdd struct {
	left  ExprNode
	right ExprNode
}

func NewExprNodeAdd(left ExprNode, right ExprNode) *ExprNodeAdd {
	n := &ExprNodeAdd{
		left:  left,
		right: right,
	}
	return n
}

func (n *ExprNodeAdd) Eval(env *Environment) *ExprValue {
	l := n.left.Eval(env)
	r := n.right.Eval(env)

	if l.Type != r.Type {
		panic(19)
	}
	switch l.Type {
	case IntType:
		return &ExprValue{Type: IntType, Value: l.Value.(int) + r.Value.(int)}
	case FloatType:
		return &ExprValue{Type: FloatType, Value: l.Value.(float64) + r.Value.(float64)}
	}
	panic(19)
}

func (n *ExprNodeAdd) String() string {
	return fmt.Sprintf("ADD(%v, %v)", n.left, n.right)
}

//---------------------------------------------------------------------

type ExprNodeSubtract struct {
	left  ExprNode
	right ExprNode
}

func NewExprNodeSubtract(left ExprNode, right ExprNode) *ExprNodeSubtract {
	n := &ExprNodeSubtract{
		left:  left,
		right: right,
	}
	return n
}

func (n *ExprNodeSubtract) Eval(env *Environment) *ExprValue {
	l := n.left.Eval(env)
	r := n.right.Eval(env)

	if l.Type != r.Type {
		panic(19)
	}
	switch l.Type {
	case IntType:
		return &ExprValue{Type: IntType, Value: l.Value.(int) - r.Value.(int)}
	case FloatType:
		return &ExprValue{Type: FloatType, Value: l.Value.(float64) - r.Value.(float64)}
	}
	panic(19)
}

func (n *ExprNodeSubtract) String() string {
	return fmt.Sprintf("SUBTRACT(%v, %v)", n.left, n.right)
}

//---------------------------------------------------------------------

type ExprNodeBitwiseXor struct {
	left  ExprNode
	right ExprNode
}

func NewExprNodeBitwiseXor(left ExprNode, right ExprNode) *ExprNodeBitwiseXor {
	n := &ExprNodeBitwiseXor{
		left:  left,
		right: right,
	}
	return n
}

func (n *ExprNodeBitwiseXor) Eval(env *Environment) *ExprValue {
	l := n.left.Eval(env)
	r := n.right.Eval(env)

	if l.Type != r.Type {
		panic(19)
	}
	switch l.Type {
	case IntType:
		return &ExprValue{Type: IntType, Value: l.Value.(int) ^ r.Value.(int)}
	}
	panic(19)
}

func (n *ExprNodeBitwiseXor) String() string {
	return fmt.Sprintf("^(%v, %v)", n.left, n.right)
}

//---------------------------------------------------------------------

type ExprNodeBitwiseAnd struct {
	left  ExprNode
	right ExprNode
}

func NewExprNodeBitwiseAnd(left ExprNode, right ExprNode) *ExprNodeBitwiseAnd {
	n := &ExprNodeBitwiseAnd{
		left:  left,
		right: right,
	}
	return n
}

func (n *ExprNodeBitwiseAnd) Eval(env *Environment) *ExprValue {
	l := n.left.Eval(env)
	r := n.right.Eval(env)

	if l.Type != r.Type {
		panic(19)
	}
	switch l.Type {
	case IntType:
		return &ExprValue{Type: IntType, Value: l.Value.(int) & r.Value.(int)}
	}
	panic(19)
}

func (n *ExprNodeBitwiseAnd) String() string {
	return fmt.Sprintf("&(%v, %v)", n.left, n.right)
}

//---------------------------------------------------------------------

type ExprNodeBitwiseOr struct {
	left  ExprNode
	right ExprNode
}

func NewExprNodeBitwiseOr(left ExprNode, right ExprNode) *ExprNodeBitwiseOr {
	n := &ExprNodeBitwiseOr{
		left:  left,
		right: right,
	}
	return n
}

func (n *ExprNodeBitwiseOr) Eval(env *Environment) *ExprValue {
	l := n.left.Eval(env)
	r := n.right.Eval(env)

	if l.Type != r.Type {
		panic(19)
	}
	switch l.Type {
	case IntType:
		return &ExprValue{Type: IntType, Value: l.Value.(int) | r.Value.(int)}
	}
	panic(19)
}

func (n *ExprNodeBitwiseOr) String() string {
	return fmt.Sprintf("|(%v, %v)", n.left, n.right)
}

//---------------------------------------------------------------------

type ExprNodeLogicalAnd struct {
	left  ExprNode
	right ExprNode
}

func NewExprNodeLogicalAnd(left ExprNode, right ExprNode) *ExprNodeLogicalAnd {
	n := &ExprNodeLogicalAnd{
		left:  left,
		right: right,
	}
	return n
}

func (n *ExprNodeLogicalAnd) Eval(env *Environment) *ExprValue {
	l := n.left.Eval(env)
	r := n.right.Eval(env)

	if l.Type != r.Type {
		panic(19)
	}
	switch l.Type {
	case BoolType:
		return &ExprValue{Type: IntType, Value: l.Value.(bool) && r.Value.(bool)}
	}
	panic(19)
}

func (n *ExprNodeLogicalAnd) String() string {
	return fmt.Sprintf("&&(%v, %v)", n.left, n.right)
}

//---------------------------------------------------------------------

type ExprNodeLogicalOr struct {
	left  ExprNode
	right ExprNode
}

func NewExprNodeLogicalOr(left ExprNode, right ExprNode) *ExprNodeLogicalOr {
	n := &ExprNodeLogicalOr{
		left:  left,
		right: right,
	}
	return n
}

func (n *ExprNodeLogicalOr) Eval(env *Environment) *ExprValue {
	l := n.left.Eval(env)
	r := n.right.Eval(env)

	if l.Type != r.Type {
		panic(19)
	}
	switch l.Type {
	case BoolType:
		return &ExprValue{Type: IntType, Value: l.Value.(bool) || r.Value.(bool)}
	}
	panic(19)
}

func (n *ExprNodeLogicalOr) String() string {
	return fmt.Sprintf("&(%v, %v)", n.left, n.right)
}

//---------------------------------------------------------------------

type ExprNodeModulus struct {
	left  ExprNode
	right ExprNode
}

func NewExprNodeModulus(left ExprNode, right ExprNode) *ExprNodeModulus {
	n := &ExprNodeModulus{
		left:  left,
		right: right,
	}
	return n
}

func (n *ExprNodeModulus) Eval(env *Environment) *ExprValue {
	l := n.left.Eval(env)
	r := n.right.Eval(env)

	if l.Type != r.Type {
		panic(19)
	}
	switch l.Type {
	case IntType:
		return &ExprValue{Type: IntType, Value: l.Value.(int) % r.Value.(int)}
	}
	panic(19)
}

func (n *ExprNodeModulus) String() string {
	return fmt.Sprintf("%%(%v, %v)", n.left, n.right)
}

//---------------------------------------------------------------------

type ExprNodeGreaterEqual struct {
	left  ExprNode
	right ExprNode
}

func NewExprNodeGreaterEqual(left ExprNode, right ExprNode) *ExprNodeGreaterEqual {
	n := &ExprNodeGreaterEqual{
		left:  left,
		right: right,
	}
	return n
}

func (n *ExprNodeGreaterEqual) Eval(env *Environment) *ExprValue {
	l := n.left.Eval(env)
	r := n.right.Eval(env)

	if l.Type != r.Type {
		panic(19)
	}
	switch l.Type {
	case IntType:
		return &ExprValue{Type: IntType, Value: l.Value.(int) >= r.Value.(int)}
	case FloatType:
		return &ExprValue{Type: FloatType, Value: l.Value.(float64) >= r.Value.(float64)}
	case StringType:
		return &ExprValue{Type: StringType, Value: l.Value.(string) >= r.Value.(string)}
	}
	panic(19)
}

func (n *ExprNodeGreaterEqual) String() string {
	return fmt.Sprintf(">=(%v, %v)", n.left, n.right)
}

//---------------------------------------------------------------------

type ExprNodeGreater struct {
	left  ExprNode
	right ExprNode
}

func NewExprNodeGreater(left ExprNode, right ExprNode) *ExprNodeGreater {
	n := &ExprNodeGreater{
		left:  left,
		right: right,
	}
	return n
}

func (n *ExprNodeGreater) Eval(env *Environment) *ExprValue {
	l := n.left.Eval(env)
	r := n.right.Eval(env)

	if l.Type != r.Type {
		panic(19)
	}
	switch l.Type {
	case IntType:
		return &ExprValue{Type: IntType, Value: l.Value.(int) > r.Value.(int)}
	case FloatType:
		return &ExprValue{Type: FloatType, Value: l.Value.(float64) > r.Value.(float64)}
	case StringType:
		return &ExprValue{Type: StringType, Value: l.Value.(string) > r.Value.(string)}
	}
	panic(19)
}

func (n *ExprNodeGreater) String() string {
	return fmt.Sprintf(">(%v, %v)", n.left, n.right)
}

//---------------------------------------------------------------------

type ExprNodeLessEqual struct {
	left  ExprNode
	right ExprNode
}

func NewExprNodeLessEqual(left ExprNode, right ExprNode) *ExprNodeLessEqual {
	n := &ExprNodeLessEqual{
		left:  left,
		right: right,
	}
	return n
}

func (n *ExprNodeLessEqual) Eval(env *Environment) *ExprValue {
	l := n.left.Eval(env)
	r := n.right.Eval(env)

	if l.Type != r.Type {
		panic(19)
	}
	switch l.Type {
	case IntType:
		return &ExprValue{Type: IntType, Value: l.Value.(int) <= r.Value.(int)}
	case FloatType:
		return &ExprValue{Type: FloatType, Value: l.Value.(float64) <= r.Value.(float64)}
	case StringType:
		return &ExprValue{Type: StringType, Value: l.Value.(string) <= r.Value.(string)}
	}
	panic(19)
}

func (n *ExprNodeLessEqual) String() string {
	return fmt.Sprintf("<=(%v, %v)", n.left, n.right)
}

//---------------------------------------------------------------------

type ExprNodeLess struct {
	left  ExprNode
	right ExprNode
}

func NewExprNodeLess(left ExprNode, right ExprNode) *ExprNodeLess {
	n := &ExprNodeLess{
		left:  left,
		right: right,
	}
	return n
}

func (n *ExprNodeLess) Eval(env *Environment) *ExprValue {
	l := n.left.Eval(env)
	r := n.right.Eval(env)

	if l.Type != r.Type {
		panic(19)
	}
	switch l.Type {
	case IntType:
		return &ExprValue{Type: IntType, Value: l.Value.(int) < r.Value.(int)}
	case FloatType:
		return &ExprValue{Type: FloatType, Value: l.Value.(float64) < r.Value.(float64)}
	case StringType:
		return &ExprValue{Type: StringType, Value: l.Value.(string) < r.Value.(string)}
	}
	panic(19)
}

func (n *ExprNodeLess) String() string {
	return fmt.Sprintf("<(%v, %v)", n.left, n.right)
}

//---------------------------------------------------------------------

type ExprNodeEqualsEquals struct {
	left  ExprNode
	right ExprNode
}

func NewExprNodeEqualsEquals(left ExprNode, right ExprNode) *ExprNodeEqualsEquals {
	n := &ExprNodeEqualsEquals{
		left:  left,
		right: right,
	}
	return n
}

func (n *ExprNodeEqualsEquals) Eval(env *Environment) *ExprValue {
	l := n.left.Eval(env)
	r := n.right.Eval(env)

	if l.Type != r.Type {
		panic(19)
	}
	switch l.Type {
	case IntType:
		return &ExprValue{Type: IntType, Value: l.Value.(int) == r.Value.(int)}
	case FloatType:
		return &ExprValue{Type: FloatType, Value: l.Value.(float64) == r.Value.(float64)}
	case StringType:
		return &ExprValue{Type: StringType, Value: l.Value.(string) == r.Value.(string)}
	case BoolType:
		return &ExprValue{Type: StringType, Value: l.Value.(bool) == r.Value.(bool)}
	}
	panic(19)
}

func (n *ExprNodeEqualsEquals) String() string {
	return fmt.Sprintf("==(%v, %v)", n.left, n.right)
}

//---------------------------------------------------------------------

type ExprNodeNotEquals struct {
	left  ExprNode
	right ExprNode
}

func NewExprNodeNotEquals(left ExprNode, right ExprNode) *ExprNodeNotEquals {
	n := &ExprNodeNotEquals{
		left:  left,
		right: right,
	}
	return n
}

func (n *ExprNodeNotEquals) Eval(env *Environment) *ExprValue {
	l := n.left.Eval(env)
	r := n.right.Eval(env)

	if l.Type != r.Type {
		panic(19)
	}
	switch l.Type {
	case IntType:
		return &ExprValue{Type: IntType, Value: l.Value.(int) != r.Value.(int)}
	case FloatType:
		return &ExprValue{Type: FloatType, Value: l.Value.(float64) != r.Value.(float64)}
	case StringType:
		return &ExprValue{Type: StringType, Value: l.Value.(string) != r.Value.(string)}
	case BoolType:
		return &ExprValue{Type: StringType, Value: l.Value.(bool) != r.Value.(bool)}
	}
	panic(19)
}

func (n *ExprNodeNotEquals) String() string {
	return fmt.Sprintf("!=(%v, %v)", n.left, n.right)
}

//---------------------------------------------------------------------------

type ExprNodeSymbolRef struct {
	typ  BaseType
	name string
}

func NewExprNodeSymbolRef(name string, typ BaseType) *ExprNodeSymbolRef {
	n := &ExprNodeSymbolRef{
		name: name,
		typ:  typ,
	}
	return n
}

func (n *ExprNodeSymbolRef) Eval(env *Environment) *ExprValue {
	value := env.get(n.name)
	if n.typ != value.Type {
		panic(19)
	}
	return value
}

func (n *ExprNodeSymbolRef) String() string {
	t := ""
	switch n.typ {
	case IntType:
		t = "I"
	case FloatType:
		t = "F"
	case BoolType:
		t = "B"
	case StringType:
		t = "S"
	default:
		panic(21)
	}
	return fmt.Sprintf("%s%s", n.name, t)
}

//---------------------------------------------------------------------------

type ExprNodeIntConstant struct {
	value *ExprValue
}

func NewExprNodeIntConstant(value int) *ExprNodeIntConstant {
	n := &ExprNodeIntConstant{
		value: &ExprValue{Type: IntType, Value: value},
	}
	return n
}

func (n *ExprNodeIntConstant) Eval(env *Environment) *ExprValue {
	return n.value
}

func (n *ExprNodeIntConstant) String() string {
	return fmt.Sprintf("%d", n.value.Value.(int))
}

//---------------------------------------------------------------------------

type ExprNodeFloatConstant struct {
	value *ExprValue
}

func NewExprNodeFloatConstant(value float64) *ExprNodeFloatConstant {
	n := &ExprNodeFloatConstant{
		value: &ExprValue{Type: FloatType, Value: value},
	}
	return n
}

func (n *ExprNodeFloatConstant) Eval(env *Environment) *ExprValue {
	return n.value
}

func (n *ExprNodeFloatConstant) String() string {
	return fmt.Sprintf("%f", n.value.Value.(float64))
}

//---------------------------------------------------------------------------

type ExprNodeBoolConstant struct {
	value *ExprValue
}

func NewExprNodeBoolConstant(value bool) *ExprNodeBoolConstant {
	n := &ExprNodeBoolConstant{
		value: &ExprValue{Type: BoolType, Value: value},
	}
	return n
}

func (n *ExprNodeBoolConstant) Eval(env *Environment) *ExprValue {
	return n.value
}

func (n *ExprNodeBoolConstant) String() string {
	return fmt.Sprintf("%t", n.value.Value.(bool))
}

//---------------------------------------------------------------------------

type ExprNodeStringConstant struct {
	value *ExprValue
}

func NewExprNodeStringConstant(value string) *ExprNodeStringConstant {
	n := &ExprNodeStringConstant{
		value: &ExprValue{Type: StringType, Value: value},
	}
	return n
}

func (n *ExprNodeStringConstant) Eval(env *Environment) *ExprValue {
	return n.value
}

func (n *ExprNodeStringConstant) String() string {
	return fmt.Sprintf(`"%s"`, n.value.Value.(string))
}

//---------------------------------------------------------------------------

// "message.timestamp"
type ExprNodeStructRef struct {
	symbol string
	field  string
}

func NewExprNodeStructRef(symbol string, field string) *ExprNodeStructRef {
	n := &ExprNodeStructRef{
		symbol: symbol,
		field:  field,
	}
	return n
}

func (n *ExprNodeStructRef) Eval(env *Environment) *ExprValue {
	return env.get(n.symbol + "." + n.field)
}

func (n *ExprNodeStructRef) String() string {
	return fmt.Sprintf("STRUCTREF(%s.%s)", n.symbol, n.field)
}

//---------------------------------------------------------------------------

type ExprNodeArrayRef struct {
	symbol string
	index  ExprNode
}

func NewExprNodeArrayRef(symbol string, index ExprNode) *ExprNodeArrayRef {
	n := &ExprNodeArrayRef{
		symbol: symbol,
		index:  index,
	}
	return n
}

func (n *ExprNodeArrayRef) Eval(env *Environment) *ExprValue {
	idx := n.index.Eval(env)
	if idx.Type != IntType {
		panic(19)
	}
	idxstr := strconv.Itoa(idx.Value.(int))
	return env.get(n.symbol + "#" + idxstr)
}

func (n *ExprNodeArrayRef) String() string {
	return fmt.Sprintf("ARRAYREF(%s,%v)", n.symbol, n.index)
}

//---------------------------------------------------------------------------

type ExprNodeMapRef struct {
	symbol string
	key    ExprNode
}

func NewExprNodeMapRef(symbol string, key ExprNode) *ExprNodeMapRef {
	n := &ExprNodeMapRef{
		symbol: symbol,
		key:    key,
	}
	return n
}

func (n *ExprNodeMapRef) String() string {
	return fmt.Sprintf("MAPREF(%s,%v)", n.symbol, n.key)
}

//---------------------------------------------------------------------------
