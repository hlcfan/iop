package pp

import (
	"reflect"
)

func (p *PPrinter) PrintMap(v reflect.Value, level int) {
	p.WriteString("{")
	p.writeNewline()

	for _, key := range v.MapKeys() {
		p.WriteString(p.nextLineIndent(level))
		v := v.MapIndex(key)
		p.Inspect(key, level)
		p.WriteString(p.styler.PrintCharacter(": "))
		p.Inspect(v, level+1)
		p.WriteString(p.styler.PrintCharacter(","))
		p.writeNewline()
	}

	p.WriteString(p.currentLineIndent(level))
	p.WriteString("}")
}
