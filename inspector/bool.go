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

func (r *BoolInspector) Inspect(ioP IOP, t reflect.Type, v reflect.Value, level int) {
	var tabs string
	for i := 0; i < level; i++ {
		tabs += "\t"
	}
	fmt.Fprintf(ioP.Output(), "%s%t,\n", tabs, v.Bool())
}
