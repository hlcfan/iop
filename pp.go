package pp

import (
	"fmt"
	"io"
	"os"
	"reflect"
	"sync"
	"text/tabwriter"

	"github.com/hlcfan/pp/inspector"
)

var std = New()

type PPrinter struct {
	mutex      sync.Mutex
	out        *tabwriter.Writer
	inspectors []Inspectable
	maxDepth   int
	label      string
}

type Inspectable interface {
	Applicable(reflect.Type, reflect.Value) bool
	Inspect(inspector.Printable, reflect.Type, reflect.Value, int)
}

func SetOutput(out io.Writer) {
	std.SetOutput(out)
}

func Puts(variables ...interface{}) {
	std.Puts(variables...)
}

func Label(label string) *PPrinter {
	std.label = label

	return std
}

func New() *PPrinter {
	w := tabwriter.NewWriter(os.Stdout, 4, 4, 1, ' ', 0)
	return &PPrinter{
		out: w,
		// maxDepth: 2,
		inspectors: []Inspectable{
			inspector.NewSliceInspector(),
			inspector.NewMapInspector(),
			inspector.NewIntegerInspector(),
			inspector.NewTimeInspector(),
			inspector.NewStructInspector(),
			inspector.NewStringInspector(),
			inspector.NewBoolInspector(),
			inspector.NewInterfaceInspector(),
			// inspector.NewFallbackInspector(),
		},
	}
}

func (p *PPrinter) Puts(variables ...interface{}) {
	for _, variable := range variables {
		fmt.Fprint(std.out, std.label)
		v := reflect.ValueOf(variable)
		p.Inspect(v, 0)
	}

	p.out.Flush()
}

func (p *PPrinter) SetOutput(out io.Writer) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	w := tabwriter.NewWriter(out, 4, 4, 1, ' ', 0)
	p.out = w
}

func (p *PPrinter) Inspect(variable reflect.Value, level int) {
	if p.maxDepth > 0 && level > p.maxDepth {
		return
	}

	var inspector Inspectable

	v := variable
	t := reflect.TypeOf(v)

	for _, i := range p.inspectors {
		if i.Applicable(t, v) {
			inspector = i
			break
		}
	}

	if inspector == nil {
		return
	}

	inspector.Inspect(p, t, v, level)
}

func (p *PPrinter) Output() io.Writer {
	return p.out
}

func (p *PPrinter) Flush() {
	p.out.Flush()
}
