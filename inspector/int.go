package inspector

import (
	"fmt"
	"reflect"
)

type IntegerInspector struct{}

func NewIntegerInspector() *IntegerInspector {
	return &IntegerInspector{}
}

func (r *IntegerInspector) Applicable(t reflect.Type, _ reflect.Value) bool {
	integerTypes := []reflect.Kind{
		reflect.Int,
		reflect.Int8,
		reflect.Int16,
		reflect.Int32,
		reflect.Int64,
	}

	for _, kind := range integerTypes {
		if t.Kind() == kind {
			return true
		}
	}

	return false
}

func (r *IntegerInspector) Inspect(ioP IOP, t reflect.Type, v reflect.Value, level int) {
	var tabs string
	for i := 0; i < level; i++ {
		tabs += "\t"
	}
	fmt.Fprintf(ioP.Output(), "%s%d\n", tabs, v.Int())
}
