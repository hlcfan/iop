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
	format := "%s\n"
	if level > 0 {
		format = "%s,\n"
	}

	fmt.Fprintf(ioP.Output(), format, v.String())
}
