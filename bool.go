package pp

import (
	"reflect"
	"strconv"
)

func (p *PPrinter) PrintBool(v reflect.Value) {
	p.WriteString(strconv.FormatBool(v.Bool()))
}
