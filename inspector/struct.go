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

	var structType string
	indentation := "\t"
	if level == 0 {
		structType = v.Type().String()
		indentation = ""
	}

	fmt.Fprintf(ioP.Output(), "%s%v{\n", indentation, structType)
	for j := 0; j < v.NumField(); j++ {
		typeField := v.Type().Field(j)
		valueField := v.Field(j)
		fmt.Fprintf(ioP.Output(), "%s\t%s:", tabs, typeField.Name)
		// fmt.Printf("===Value Field: %#v\n", valueField)
		// fmt.Fprintf(ioP.Output(), "\t")
		ioP.Inspect(valueField.Interface(), level+1)
	}

	var comma string
	if level > 0 {
		comma = ","
	}

	fmt.Fprintf(ioP.Output(), "%s}%s\n", tabs, comma)
}
