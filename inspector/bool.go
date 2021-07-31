package inspector

import (
	"fmt"
	"reflect"
)

type BoolInspector struct{}

func NewBoolInspector() *BoolInspector {
	return &BoolInspector{}
}

func (r *BoolInspector) Applicable(t reflect.Type, v reflect.Value) bool {
	return v.Kind() == reflect.Bool
}

func (r *BoolInspector) Inspect(ioP Printable, t reflect.Type, v reflect.Value, level int) {
	format := "%t\n"
	if level > 0 {
		format = "%t,\n"
	}

	fmt.Fprintf(ioP.Output(), format, v.Bool())
}
