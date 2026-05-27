package mp4

import (
	"io"
)

type Writer struct {
	writer  io.WriteSeeker
	biStack []*BoxInfo
}

func NewWriter(w io.WriteSeeker) *Writer { _ = "STUB: not implemented"; return nil }

func (w *Writer) Write(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (w *Writer) Seek(offset int64, whence int) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (w *Writer) StartBox(bi *BoxInfo) (*BoxInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w *Writer) EndBox() (*BoxInfo, error) { _ = "STUB: not implemented"; return nil, nil }

func (w *Writer) CopyBox(r io.ReadSeeker, bi *BoxInfo) error { _ = "STUB: not implemented"; return nil }
