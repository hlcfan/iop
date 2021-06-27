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

func (r *TimeInspector) Inspect(ioP Printable, t reflect.Type, v reflect.Value, level int) {
	// fmt.Printf("===Value : %#v\n", v.Type().String())
	// fmt.Printf("===Value : %#v\n", v.Interface())

	fmt.Fprintf(ioP.Output(), "\t%v", v.Interface().(time.Time))

	var comma string
	if level > 0 {
		comma = ","
	}

	fmt.Fprintf(ioP.Output(), "%s\n", comma)
}
