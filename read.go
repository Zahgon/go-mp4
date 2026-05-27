package mp4

import (
	"io"
)

func checkPayloadSize(r io.ReadSeeker, bi *BoxInfo) error { _ = "STUB: not implemented"; return nil }

type BoxPath []BoxType

func (lhs BoxPath) compareWith(rhs BoxPath) (forwardMatch bool, match bool) {
	_ = "STUB: not implemented"
	return false, false
}

type ReadHandle struct {
	Params      []interface{}
	BoxInfo     BoxInfo
	Path        BoxPath
	ReadPayload func() (box IBox, n uint64, err error)
	ReadData    func(io.Writer) (n uint64, err error)
	Expand      func(params ...interface{}) (vals []interface{}, err error)
}

type ReadHandler func(handle *ReadHandle) (val interface{}, err error)

func ReadBoxStructure(r io.ReadSeeker, handler ReadHandler, params ...interface{}) ([]interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ReadBoxStructureFromInternal(r io.ReadSeeker, bi *BoxInfo, handler ReadHandler, params ...interface{}) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func readBoxStructureFromInternal(r io.ReadSeeker, bi *BoxInfo, path BoxPath, handler ReadHandler, params []interface{}) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// check comatible-brands

// parse numbered ilst items after keys box by saving EntryCount field to context

func readBoxStructure(r io.ReadSeeker, totalSize uint64, isRoot bool, path BoxPath, ctx Context, handler ReadHandler, params []interface{}) ([]interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// preserve keys entry count on context for subsequent ilst number item box
