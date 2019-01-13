package ChartoGopher

import "io"

type Writer interface {
	Write(Chart) (int, error)
}

type Reader interface {
	read(io.Reader) Chart
}
