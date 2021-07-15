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

func (r *StringInspector) Inspect(ioP Printable, t reflect.Type, v reflect.Value, level int) {
	var tab string
	var comma string

	if level > 0 {
		tab = "\t"
		comma = ","
	}

	fmt.Fprintf(ioP.Output(), "%s%s%s\n", tab, v.String(), comma)
}
