package inspector

import (
	"fmt"
	"reflect"
)

type SliceInspector struct {
}

func NewSliceInspector() *SliceInspector {
	return &SliceInspector{}
}

func (r *SliceInspector) Applicable(t reflect.Type, v reflect.Value) bool {
	return v.Kind() == reflect.Slice
}

func (r *SliceInspector) Inspect(ioP IOP, t reflect.Type, v reflect.Value, level int) {
	var tabs string
	for i := 0; i < level; i++ {
		tabs += "\t"
	}
	fmt.Fprintf(ioP.Output(), "%s%s{\n", tabs, t)
	for i := 0; i < v.Len(); i++ {
		ele := v.Index(i)
		// tt := reflect.TypeOf(ele)
		// vv := reflect.ValueOf(ele)
		// fmt.Printf("===Ele: %#v\n", ele)
		// fmt.Fprintln(ioP.Output(), "\t\t{")
		// Interate each struct field
		ioP.Inspect(ele.Interface(), level)
		// fmt.Fprintln(ioP.Output(), "\t\t},")
	}

	fmt.Fprintf(ioP.Output(), "%s}\n", tabs)
}
