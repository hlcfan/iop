package inspector

import (
	"fmt"
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

func (r *MapInspector) Inspect(ioP Printable, t reflect.Type, v reflect.Value, level int) {
	var tabs, indentation string
	// TODO: may use buffer
	for i := 0; i < level; i++ {
		tabs += "\t"
	}

	if level > 0 {
		indentation = ""
	}

	fmt.Fprintf(ioP.Output(), "%s%s {\n", indentation, v.Type())
	for _, key := range v.MapKeys() {
		v := v.MapIndex(key)
		fmt.Fprintf(ioP.Output(), "%s\t%v:\t", tabs, key.Interface())
		ioP.Inspect(v, level+1)
	}

	var comma string
	if level > 0 {
		comma = ","
	}

	fmt.Fprintf(ioP.Output(), "%s}%s\n", tabs, comma)
}
