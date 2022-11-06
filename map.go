package pp

import (
	"fmt"
	"reflect"
)

func (p *PPrinter) PrintMap(v reflect.Value, level int) {
	p.WriteString("{")
	p.writeNewline()

	for _, key := range v.MapKeys() {
		p.WriteString(p.nextLineIndent(level))
		v := v.MapIndex(key)
		p.WriteString(fmt.Sprintf("%v: ", key.Interface()))
		p.Inspect(v, level+1)
		p.WriteByte(',')
		p.writeNewline()
	}

	p.WriteString(p.currentLineIndent(level))
	p.WriteString("}")
}
