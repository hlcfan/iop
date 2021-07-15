package pp

import (
	"io"
	"os"
	"reflect"
	"sync"

	"github.com/hlcfan/pp/inspector"
)

var std = New()

type PPrinter struct {
	mutex      sync.Mutex
	Out        io.Writer
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

// func Puts(variable interface{}) {
// 	fmt.Fprintf("===Variable: %#v\n", variable)
// }

func Inspect(variable interface{}) {
	v := reflect.ValueOf(variable)
	std.Inspect(v, 0)
}

func New() *PPrinter {
	return &PPrinter{
		Out: os.Stdout,
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
	p.Out = out
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
