package inspector

import (
	"fmt"
	"reflect"
)

type flag uintptr

var (
	// flagRO indicates whether the value field of a reflect.Value
	// is read-only.
	flagRO flag

	// flagAddr indicates whether the address of the reflect.Value's
	// value may be taken.
	flagAddr flag
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
		fmt.Fprintf(ioP.Output(), "%s\t%s:", tabs, typeField.Name)

		// key := typeField.Name
		field := v.Field(j)
		ioP.Inspect(field, level+1)
	}

	var comma string
	if level > 0 {
		comma = ","
	}

	fmt.Fprintf(ioP.Output(), "%s}%s\n", tabs, comma)
}
