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

func (r *StructInspector) Inspect(ioP IOP, t reflect.Type, v reflect.Value) {
	fmt.Fprint(ioP.Output(), "{\n")
	for j := 0; j < v.NumField(); j++ {
		typeField := v.Type().Field(j)
		valueField := v.Field(j)
		// fmt.Fprintf(ioP.Output(), "\t\t\t%s:\t%v,\n", typeField.Name, valueField.Interface())
		fmt.Fprintf(ioP.Output(), "\t\t\t%s:\t", typeField.Name)
		ioP.Inspect(valueField.Interface())
	}
	fmt.Fprintln(ioP.Output(), "\t\t},")
	fmt.Fprintln(ioP.Output(), "}")
}
