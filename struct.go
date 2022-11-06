package pp

import (
	"fmt"
	"reflect"
	"time"
)

func (p *PPrinter) PrintStruct(v reflect.Value, level int) {
	if v.Type().String() == "time.Time" {
		p.printTime(v, level)
		return
	}

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

func (p *PPrinter) printTime(v reflect.Value, level int) {
	s := fmt.Sprintf("%v", v.Interface().(time.Time))
	p.WriteString(s)
}
