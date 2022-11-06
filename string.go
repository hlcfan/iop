package pp

import (
	"reflect"
	"strconv"
)

func (p *PPrinter) PrintString(v reflect.Value) {
	p.WriteString(strconv.Quote(v.String()))
}
