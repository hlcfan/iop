package inspector

import (
	"fmt"
	"io"
	"reflect"
)

type MapInspector struct {
}

func NewMapInspector() *MapInspector {
	return &MapInspector{}
}

func (r *MapInspector) Applicable(t reflect.Type, v reflect.Value) bool {
	return v.Kind() == reflect.Map
}

func (r *MapInspector) Inspect(out io.Writer, t reflect.Type, v reflect.Value) {
	fmt.Println("===================")
	// fmt.Println("===Ele type: ", t)
	// fmt.Println("===Ele type: ", t.Elem())
	fmt.Fprintf(out, "%s {\n", t)
	for _, e := range v.MapKeys() {
		v := v.MapIndex(e)
		fmt.Fprintf(out, "\t\t\t%s:\t%v,\n", e, v)
		// switch t := v.Interface().(type) {
		// case int:
		// 	fmt.Fprintln(out, e, t)
		// case string:
		// 	fmt.Fprintln(out, e, t)
		// case bool:
		// 	fmt.Fprintln(out, e, t)
		// default:
		// 	fmt.Fprintln(out, "not found")
		// }
	}

	fmt.Fprintln(out, "}")

	// fmt.Printf("%#v\n", v)
	fmt.Println("===================")
}
