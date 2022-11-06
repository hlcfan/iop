package pp

import (
	"reflect"
	"strconv"
)

func (p *PPrinter) PrintUnsignedInteger(v reflect.Value) {
	p.WriteString(strconv.FormatUint(v.Uint(), 10))
}

func (p *PPrinter) PrintSignedInteger(v reflect.Value) {
	p.WriteString(strconv.FormatInt(v.Int(), 10))
}
