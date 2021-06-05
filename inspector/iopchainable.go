package inspector

import "io"

type IOP interface {
	Inspect(variable interface{}, level int)
	Output() io.Writer
}
