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
	fmt.Fprintf(ioP.Output(), "\t%t,\n", v.Bool())
}
