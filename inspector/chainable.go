package inspector

import (
	"io"
	"reflect"
)

type Printable interface {
	Inspect(variable reflect.Value, level int)
	Output() io.Writer
}
