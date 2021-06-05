package inspector

import (
	"fmt"
	"reflect"
)

type StructInspector struct {
}

func NewStructInspector() *StructInspector {
	return &StructInspector{}
}

func (r *StructInspector) Applicable(t reflect.Type, v reflect.Value) bool {
	return v.Kind() == reflect.Struct
}

func (r *StructInspector) Inspect(ioP IOP, t reflect.Type, v reflect.Value, level int) {
	var tabs string
	// TODO: may use buffer
	for i := 0; i < level; i++ {
		tabs += "\t"
	}

	fmt.Fprintf(ioP.Output(), "%s{\n", tabs)
	for j := 0; j < v.NumField(); j++ {
		typeField := v.Type().Field(j)
		valueField := v.Field(j)
		fmt.Fprintf(ioP.Output(), "%s\t%s:", tabs, typeField.Name)
		ioP.Inspect(valueField.Interface(), level+1)
	}
	fmt.Fprintf(ioP.Output(), "%s},\n", tabs)
}
