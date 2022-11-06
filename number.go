package pp

import (
	"reflect"
	"strconv"
)

func (p *PPrinter) PrintUnsignedInteger(v reflect.Value) {
	p.WriteString(p.styler.PrintNumber(strconv.FormatUint(v.Uint(), 10)))
}

func (p *PPrinter) PrintSignedInteger(v reflect.Value) {
	p.WriteString(p.styler.PrintNumber(strconv.FormatInt(v.Int(), 10)))
}
