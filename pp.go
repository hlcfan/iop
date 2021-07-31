package pp

import (
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
	Out        *tabwriter.Writer
	inspectors []Inspectable
	maxDepth   int
}

type Inspectable interface {
	Applicable(reflect.Type, reflect.Value) bool
	Inspect(inspector.Printable, reflect.Type, reflect.Value, int)
}

func SetOutput(out io.Writer) {
	std.SetOutput(out)
}

func Puts(variables ...interface{}) {
	for _, variable := range variables {
		v := reflect.ValueOf(variable)
		std.Inspect(v, 0)
	}

	std.Out.Flush()
}

func New() *PPrinter {
	w := tabwriter.NewWriter(os.Stdout, 4, 4, 1, ' ', tabwriter.Debug)
	return &PPrinter{
		Out: w,
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

func (p *PPrinter) SetOutput(out io.Writer) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	w := tabwriter.NewWriter(out, 4, 4, 1, ' ', 0)
	p.Out = w
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
	return p.Out
}

func (p *PPrinter) Flush() {
	p.Out.Flush()
}
