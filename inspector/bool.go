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
	var tab string
	var comma string

	if level > 0 {
		tab = "\t"
		comma = ","
	}

	fmt.Fprintf(ioP.Output(), "%s%t%s\n", tab, v.Bool(), comma)
}
