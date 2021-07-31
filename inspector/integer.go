package inspector

import (
	"fmt"
	"reflect"
)

type IntegerInspector struct{}

func NewIntegerInspector() *IntegerInspector {
	return &IntegerInspector{}
}

func (r *IntegerInspector) Applicable(t reflect.Type, v reflect.Value) bool {
	integerTypes := []reflect.Kind{
		reflect.Int,
		reflect.Int8,
		reflect.Int16,
		reflect.Int32,
		reflect.Int64,
	}

	for _, kind := range integerTypes {
		if v.Kind() == kind {
			return true
		}
	}

	return false
}

func (r *IntegerInspector) Inspect(ioP Printable, t reflect.Type, v reflect.Value, level int) {
	format := "%d\n"
	if level > 0 {
		format = "%d,\n"
	}

	fmt.Fprintf(ioP.Output(), format, v.Int())
}
