package inspector

import (
	"fmt"
	"reflect"
)

type StringInspector struct{}

func NewStringInspector() *StringInspector {
	return &StringInspector{}
}

func (r *StringInspector) Applicable(t reflect.Type, v reflect.Value) bool {
	return v.Kind() == reflect.String
}

func (r *StringInspector) Inspect(ioP IOP, t reflect.Type, v reflect.Value, level int) {
	var tabs string
	for i := 0; i < level; i++ {
		tabs += "\t"
	}
	fmt.Fprintf(ioP.Output(), "%s%s,\n", tabs, v.String())
}
