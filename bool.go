package pp

import (
	"reflect"
	"strconv"
)

func (p *PPrinter) PrintBool(v reflect.Value) {
	p.WriteString(p.styler.PrintBool(strconv.FormatBool(v.Bool())))
}
