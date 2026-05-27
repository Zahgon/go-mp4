package mp4

import (
	"errors"
	"reflect"
)

var ErrBoxInfoNotFound = errors.New("box info not found")

// BoxType is mpeg box type
type BoxType [4]byte

func StrToBoxType(code string) BoxType { _ = "STUB: not implemented"; return *new(BoxType) }

// Uint32ToBoxType returns a new BoxType from the provied uint32
func Uint32ToBoxType(i uint32) BoxType { _ = "STUB: not implemented"; return *new(BoxType) }

func (boxType BoxType) String() string { _ = "STUB: not implemented"; return "" }

func isASCII(c byte) bool { _ = "STUB: not implemented"; return false }

func isPrintable(c byte) bool { _ = "STUB: not implemented"; return false }

func (lhs BoxType) MatchWith(rhs BoxType) bool { _ = "STUB: not implemented"; return false }

var boxTypeAny = BoxType{0x00, 0x00, 0x00, 0x00}

func BoxTypeAny() BoxType { _ = "STUB: not implemented"; return *new(BoxType) }

type boxDef struct {
	dataType reflect.Type
	versions []uint8
	isTarget func(Context) bool
	fields   []*field
}

var boxMap = make(map[BoxType][]boxDef, 64)

func AddBoxDef(payload IBox, versions ...uint8) { _ = "STUB: not implemented"; return }

func AddBoxDefEx(payload IBox, isTarget func(Context) bool, versions ...uint8) {
	_ = "STUB: not implemented"
	return
}

func AddAnyTypeBoxDef(payload IAnyType, boxType BoxType, versions ...uint8) {
	_ = "STUB: not implemented"
	return
}

func AddAnyTypeBoxDefEx(payload IAnyType, boxType BoxType, isTarget func(Context) bool, versions ...uint8) {
	_ = "STUB: not implemented"
	return
}

var itemBoxFields = buildFields(&Item{})

func (boxType BoxType) getBoxDef(ctx Context) *boxDef { _ = "STUB: not implemented"; return nil }

func (boxType BoxType) IsSupported(ctx Context) bool { _ = "STUB: not implemented"; return false }

func (boxType BoxType) New(ctx Context) (IBox, error) {
	_ = "STUB: not implemented"
	return *new(IBox), nil
}

func (boxType BoxType) GetSupportedVersions(ctx Context) ([]uint8, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (boxType BoxType) IsSupportedVersion(ver uint8, ctx Context) bool {
	_ = "STUB: not implemented"
	return false
}
