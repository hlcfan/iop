package inspector

import "io"

type IOP interface {
	Inspect(variable interface{})
	Output() io.Writer
}
