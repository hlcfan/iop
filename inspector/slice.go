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

func (r *SliceInspector) Inspect(ioP Printable, t reflect.Type, v reflect.Value, level int) {
	var tabs string
	for i := 0; i < level; i++ {
		tabs += "\t"
	}

	// fmt.Printf("===Kind: %s", v.Type())
	fmt.Fprintf(ioP.Output(), "%s%s {\n", "", v.Type())
	for i := 0; i < v.Len(); i++ {
		ele := v.Index(i)
		// tt := reflect.TypeOf(ele)
		// vv := reflect.ValueOf(ele)
		// fmt.Printf("===Ele: %#v\n", ele)
		fmt.Fprintf(ioP.Output(), "%s\t", tabs)
		// Interate each struct field
		// fmt.Println("===Level: ", level)
		ioP.Inspect(ele, level+1)
		// fmt.Fprintln(ioP.Output(), "\t\t},")
	}

	var comma string
	if level > 0 {
		comma = ","
	}

	fmt.Fprintf(ioP.Output(), "%s}%s\n", tabs, comma)
}
