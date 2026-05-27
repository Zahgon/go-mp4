package bitio

import "io"

type Reader interface {
	io.Reader

	// alignment:
	//  |-1-byte-block-|--------------|--------------|--------------|
	//  |<-offset->|<-------------------width---------------------->|
	ReadBits(width uint) (data []byte, err error)

	ReadBit() (bit bool, err error)
}

type ReadSeeker interface {
	Reader
	io.Seeker
}

type reader struct {
	reader io.Reader
	octet  byte
	width  uint
}

func NewReader(r io.Reader) Reader { _ = "STUB: not implemented"; return *new(Reader) }

func (r *reader) Read(p []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

func (r *reader) ReadBits(size uint) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (r *reader) ReadBit() (bool, error) { _ = "STUB: not implemented"; return false, nil }

type readSeeker struct {
	reader
	seeker io.Seeker
}

func NewReadSeeker(r io.ReadSeeker) ReadSeeker { _ = "STUB: not implemented"; return *new(ReadSeeker) }

func (r *readSeeker) Seek(offset int64, whence int) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}
