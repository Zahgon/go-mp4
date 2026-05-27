package mp4

import (
	"errors"
	"io"
	"math"
	"reflect"

	"github.com/abema/go-mp4/internal/bitio"
)

const (
	anyVersion = math.MaxUint8
)

var ErrUnsupportedBoxVersion = errors.New("unsupported box version")

func readerHasSize(reader bitio.ReadSeeker, size uint64) bool {
	_ = "STUB: not implemented"
	return false
}

type marshaller struct {
	writer bitio.Writer
	wbits  uint64
	src    IImmutableBox
	ctx    Context
}

func Marshal(w io.Writer, src IImmutableBox, ctx Context) (n uint64, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (m *marshaller) marshal(v reflect.Value, fi *fieldInstance) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *marshaller) marshalPtr(v reflect.Value, fi *fieldInstance) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *marshaller) marshalStruct(v reflect.Value, fs []*field) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *marshaller) marshalArray(v reflect.Value, fi *fieldInstance) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *marshaller) marshalSlice(v reflect.Value, fi *fieldInstance) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *marshaller) marshalInt(v reflect.Value, fi *fieldInstance) error {
	_ = "STUB: not implemented"
	return nil
}

// set sign bit

func (m *marshaller) marshalUint(v reflect.Value, fi *fieldInstance) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *marshaller) marshalBool(v reflect.Value, fi *fieldInstance) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *marshaller) marshalString(v reflect.Value, fi *fieldInstance) error {
	_ = "STUB: not implemented"
	return nil
}

// null character

func (m *marshaller) writeUvarint(u uint64) error { _ = "STUB: not implemented"; return nil }

type unmarshaller struct {
	reader bitio.ReadSeeker
	dst    IBox
	size   uint64
	rbits  uint64
	ctx    Context
}

func UnmarshalAny(r io.ReadSeeker, boxType BoxType, payloadSize uint64, ctx Context) (box IBox, n uint64, err error) {
	_ = "STUB: not implemented"
	return *new(IBox), 0, nil
}

func Unmarshal(r io.ReadSeeker, payloadSize uint64, dst IBox, ctx Context) (n uint64, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (u *unmarshaller) unmarshal(v reflect.Value, fi *fieldInstance) error {
	_ = "STUB: not implemented"
	return nil
}

func (u *unmarshaller) unmarshalPtr(v reflect.Value, fi *fieldInstance) error {
	_ = "STUB: not implemented"
	return nil
}

func (u *unmarshaller) unmarshalStructInternal(v reflect.Value, fi *fieldInstance) error {
	_ = "STUB: not implemented"
	return nil
}

func (u *unmarshaller) unmarshalStruct(v reflect.Value, fs []*field) error {
	_ = "STUB: not implemented"
	return nil
}

func (u *unmarshaller) unmarshalArray(v reflect.Value, fi *fieldInstance) error {
	_ = "STUB: not implemented"
	return nil
}

func (u *unmarshaller) unmarshalSlice(v reflect.Value, fi *fieldInstance) error {
	_ = "STUB: not implemented"
	return nil
}

func (u *unmarshaller) unmarshalInt(v reflect.Value, fi *fieldInstance) error {
	_ = "STUB: not implemented"
	return nil
}

func (u *unmarshaller) unmarshalUint(v reflect.Value, fi *fieldInstance) error {
	_ = "STUB: not implemented"
	return nil
}

func (u *unmarshaller) unmarshalBool(v reflect.Value, fi *fieldInstance) error {
	_ = "STUB: not implemented"
	return nil
}

func (u *unmarshaller) unmarshalString(v reflect.Value, fi *fieldInstance) error {
	_ = "STUB: not implemented"
	return nil
}

func (u *unmarshaller) unmarshalStringC(v reflect.Value) error {
	_ = "STUB: not implemented"
	return nil
}

// null character

func (u *unmarshaller) unmarshalStringCP(v reflect.Value, fi *fieldInstance) error {
	_ = "STUB: not implemented"
	return nil
}

func (u *unmarshaller) tryReadPString(v reflect.Value, fi *fieldInstance) (ok bool, err error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (u *unmarshaller) readUvarint() (uint64, error) { _ = "STUB: not implemented"; return 0, nil }
