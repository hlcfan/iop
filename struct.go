package pp

import "reflect"

func (p *PPrinter) PrintStruct(v reflect.Value, level int) {
	p.WriteString(p.styler.PrintIdentifier(v.Type().Name()))
	p.WriteString(" {")
	p.writeNewline()

	for j := 0; j < v.NumField(); j++ {
		p.WriteString(p.nextLineIndent(level))
		typeField := v.Type().Field(j)
		p.WriteString(typeField.Name)
		p.WriteString(p.styler.PrintCharacter(": "))

		field := v.Field(j)
		p.Inspect(field, level+1)
		p.WriteString(p.styler.PrintCharacter(","))
		p.writeNewline()
	}

	p.WriteString(p.currentLineIndent(level))
	p.WriteByte('}')
}
