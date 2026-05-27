package bitio

import (
	"io"
)

type Writer interface {
	io.Writer

	// alignment:
	//  |-1-byte-block-|--------------|--------------|--------------|
	//  |<-offset->|<-------------------width---------------------->|
	WriteBits(data []byte, width uint) error

	WriteBit(bit bool) error
}

type writer struct {
	writer io.Writer
	octet  byte
	width  uint
}

func NewWriter(w io.Writer) Writer { _ = "STUB: not implemented"; return *new(Writer) }

func (w *writer) Write(p []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

func (w *writer) WriteBits(data []byte, width uint) error { _ = "STUB: not implemented"; return nil }

func (w *writer) WriteBit(bit bool) error { _ = "STUB: not implemented"; return nil }
