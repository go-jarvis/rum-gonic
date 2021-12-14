package rum

import "reflect"

func NewOperatorFactory(op Operator) *OperatorFactory {
	fact := &OperatorFactory{}

	// get real operator refelct type
	fact.Type = deReflectType(reflect.TypeOf(op))
	fact.Operator = op

	return fact
}

// OperatorFactory to create new operator
type OperatorFactory struct {
	Type     reflect.Type
	Operator Operator
}

// New create a new operator
func (fact *OperatorFactory) New() Operator {
	opc := reflect.New(fact.Type).Interface().(Operator)
	return opc
}
