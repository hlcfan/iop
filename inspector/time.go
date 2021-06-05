package inspector

import (
	"fmt"
	"reflect"
	"time"
)

type TimeInspector struct {
}

func NewTimeInspector() *TimeInspector {
	return &TimeInspector{}
}

func (r *TimeInspector) Applicable(t reflect.Type, v reflect.Value) bool {
	return v.Kind() == reflect.Struct &&
		v.Type().String() == "time.Time"
}

func (r *TimeInspector) Inspect(ioP IOP, t reflect.Type, v reflect.Value, level int) {
	// fmt.Printf("===Value : %#v\n", v.Type().String())
	// fmt.Printf("===Value : %#v\n", v.Interface())
	var tabs string
	// TODO: may use buffer
	for i := 0; i < level; i++ {
		tabs += "\t"
	}

	// var structType string
	// if level == 0 {
	// 	structType = v.Type().String()
	// }

	// fmt.Fprintf(ioP.Output(), "%s%v{\n", tabs, structType)
	fmt.Fprintf(ioP.Output(), "%s%v", tabs, v.Interface().(time.Time))
	// for j := 0; j < v.NumField(); j++ {
	// 	typeField := v.Type().Field(j)
	// 	valueField := v.Field(j)
	// 	fmt.Fprintf(ioP.Output(), "%s\t%s:", tabs, typeField.Name)
	// 	fmt.Printf("===Value Field: %#v\n", valueField)
	// 	ioP.Inspect(valueField.Interface(), level+1)
	// }

	var comma string
	if level > 0 {
		comma = ","
	}

	fmt.Fprintf(ioP.Output(), "%s\n", comma)
}
