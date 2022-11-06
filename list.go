package pp

import "reflect"

func (p *PPrinter) PrintList(v reflect.Value, level int) {
	p.WriteByte('[')
	p.writeNewline()

	for i := 0; i < v.Len(); i++ {
		p.WriteString(p.nextLineIndent(level))
		p.Inspect(v.Index(i), level+1)
		p.WriteByte(',')
		p.writeNewline()
	}

	p.WriteString(p.currentLineIndent(level))
	p.WriteByte(']')
}
