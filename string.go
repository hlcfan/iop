package pp

import (
	"reflect"
	"strconv"
)

func (p *PPrinter) PrintString(v reflect.Value) {
	p.WriteString(p.styler.PrintString(strconv.Quote(v.String())))
}
