package rum

import (
	"reflect"

	"github.com/gin-gonic/gin"
)

type Operator interface {
	Output(*gin.Context) (interface{}, error)
}

type PathOperator interface {
	Path() string
}

type MethodOperator interface {
	Method() string
}

func DeepCopyOperator(op Operator) Operator {
	return NewOperatorFactory(op).NewOperator()
}

type OperatorFactory struct {
	Type     reflect.Type
	Operator Operator
}

func NewOperatorFactory(op Operator) *OperatorFactory {
	fact := &OperatorFactory{}

	fact.Type = deReflectType(reflect.TypeOf(op))
	fact.Operator = op

	return fact
}

func (fact *OperatorFactory) NewOperator() Operator {

	opc := reflect.New(fact.Type).Interface().(Operator)

	return opc

}
