package inspector

import (
	"fmt"
	"reflect"
)

type InterfaceInspector struct{}

func NewInterfaceInspector() *InterfaceInspector {
	return &InterfaceInspector{}
}

func (r *InterfaceInspector) Applicable(t reflect.Type, v reflect.Value) bool {
	return v.Kind() == reflect.Interface
}

func (r *InterfaceInspector) Inspect(ioP Printable, t reflect.Type, v reflect.Value, level int) {
	fmt.Fprintf(ioP.Output(), "\t%v,\n", v)
}
