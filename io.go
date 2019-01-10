package ChartoGopher

import "io"

type Writer interface {
	write(chart) (int, error)
}

type Reader interface {
	read(io.Reader)
}
