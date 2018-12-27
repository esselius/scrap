package paasio

import "io"

type wrajter struct {
	io.Writer
}

type rejder struct {
	io.Reader
}

type rejdwrajter struct {
	wrajter
	rejder
}

func (w wrajter) WriteCount() (n int64, nops int) {
	return 0, 0
}

func (r rejder) ReadCount() (n int64, nops int) {
	return 0, 0
}

func NewWriteCounter(writer io.Writer) WriteCounter {
	return wrajter{}
}

func NewReadCounter(reader io.Reader) ReadCounter {
	return rejder{reader}
}

func NewReadWriteCounter(readWriter io.ReadWriter) ReadWriteCounter {
	return rejdwrajter{}
}
