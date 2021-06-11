package inspector

import (
	"io"
	"reflect"
)

type IOP interface {
	Inspect(variable reflect.Value, level int)
	Output() io.Writer
}
