package inspector

import (
	"fmt"
	"io"
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

func (r *SliceInspector) Inspect(out io.Writer, t reflect.Type, v reflect.Value) {
	// fmt.Println("===Ele type: ", t)
	// fmt.Println("===Ele type: ", t.Elem())
	fmt.Fprintf(out, "%s {\n", t)
	for i := 0; i < v.Len(); i++ {
		ele := v.Index(i)
		fmt.Fprintln(out, "\t\t{")
		for j := 0; j < ele.NumField(); j++ {
			valueField := ele.Field(j)
			typeField := ele.Type().Field(j)
			fmt.Fprintf(out, "\t\t\t%s:\t%v,\n", typeField.Name, valueField.Interface())
		}
		fmt.Fprintln(out, "\t\t},")
		// fmt.Printf("===Index: %#v\n", val)
		// fmt.Printf("\t%#v\n", ele)
	}

	fmt.Fprintln(out, "}")

	// fmt.Printf("%#v\n", v)
	fmt.Println("===================")
}
