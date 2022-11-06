package pp

import (
	"bufio"
	"io"
	"os"
	"reflect"
	"strings"
	"sync"

	"github.com/hlcfan/pp/style"
)

var once sync.Once
var printer *PPrinter

const (
	defaultPrefix = ""
	defaultIndent = " "
)

type Styler interface {
	PrintIdentifier(string) string
	PrintNumber(string) string
	PrintString(string) string
	PrintBool(string) string
	PrintCharacter(string) string
}

type PPrinter struct {
	*bufio.Writer
	mutex      sync.Mutex
	maxDepth   int
	indent     string
	prefix     string
	label      string
	putNewline bool
	styler     Styler
}

func Puts(variable any) {
	p := New(os.Stdout, defaultPrefix, defaultIndent)

	p.Inspect(reflect.ValueOf(variable), 0)
	p.writeNewline()

	p.Flush()
}

func PutsWithLabel(variable any, label string) {
	p := New(os.Stdout, defaultPrefix, defaultIndent)
	p.label = label

	p.WriteString(p.label + ": ")
	p.Inspect(reflect.ValueOf(variable), 0)
	p.writeNewline()

	p.Flush()
}

func SetOutput(w io.Writer) {
	printer.SetOutput(bufio.NewWriter(w))
}

func New(w io.Writer, prefix, indent string) *PPrinter {
	once.Do(func() {
		printer = &PPrinter{
			maxDepth: 5,
			prefix:   prefix,
			indent:   indent,
			Writer:   bufio.NewWriter(w),
			styler:   style.Rebecca,
		}
	})

	return printer
}

func (p *PPrinter) SetOutput(w io.Writer) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	p.Writer = bufio.NewWriter(w)
}

func (p *PPrinter) Inspect(v reflect.Value, level int) {
	if p.maxDepth > 0 && level > p.maxDepth {
		return
	}

	if !v.IsValid() {
		p.WriteString("nil")
		return
	}

	p.writePrefix()

	switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		p.PrintSignedInteger(v)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		p.PrintUnsignedInteger(v)
	case reflect.String:
		p.PrintString(v)
	case reflect.Bool:
		p.PrintBool(v)
	case reflect.Slice:
		p.PrintList(v, level)
	case reflect.Struct:
		p.PrintStruct(v, level)
	case reflect.Map:
		p.PrintMap(v, level)
	case reflect.Interface, reflect.Ptr:
		p.Inspect(v.Elem(), level)
	default:
		p.WriteString(v.String())
	}
}

func (p *PPrinter) writeNewline() {
	p.WriteByte('\n')
	p.writePrefix()
}

func (p *PPrinter) writePrefix() {
	if len(p.prefix) > 0 {
		p.WriteString(p.prefix)
	}
}

func (p *PPrinter) currentLineIndent(level int) string {
	return strings.Repeat(p.indent, level)
}

func (p *PPrinter) nextLineIndent(level int) string {
	return strings.Repeat(p.indent, level+1)
}
