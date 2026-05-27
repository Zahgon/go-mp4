package mp4

import (
	"bytes"
	"io"
	"reflect"
)

type stringifier struct {
	buf    *bytes.Buffer
	src    IImmutableBox
	indent string
	ctx    Context
}

func Stringify(src IImmutableBox, ctx Context) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func StringifyWithIndent(src IImmutableBox, indent string, ctx Context) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (m *stringifier) stringify(v reflect.Value, fi *fieldInstance, depth int) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *stringifier) stringifyPtr(v reflect.Value, fi *fieldInstance, depth int) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *stringifier) stringifyStruct(v reflect.Value, fs []*field, depth int, extended bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *stringifier) stringifyArray(v reflect.Value, fi *fieldInstance, depth int) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *stringifier) stringifySlice(v reflect.Value, fi *fieldInstance, depth int) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *stringifier) stringifyInt(v reflect.Value, fi *fieldInstance, depth int) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *stringifier) stringifyUint(v reflect.Value, fi *fieldInstance, depth int) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *stringifier) stringifyBool(v reflect.Value, depth int) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *stringifier) stringifyString(v reflect.Value, depth int) error {
	_ = "STUB: not implemented"
	return nil
}

func writeIndent(w io.Writer, indent string, depth int) { _ = "STUB: not implemented"; return }
