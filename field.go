package mp4

import (
	"reflect"
)

type (
	stringType uint8
	fieldFlag  uint16
)

const (
	stringType_C stringType = iota
	stringType_C_P

	fieldString        fieldFlag = 1 << iota // 0
	fieldExtend                              // 1
	fieldDec                                 // 2
	fieldHex                                 // 3
	fieldISO639_2                            // 4
	fieldUUID                                // 5
	fieldHidden                              // 6
	fieldOptDynamic                          // 7
	fieldVarint                              // 8
	fieldSizeDynamic                         // 9
	fieldLengthDynamic                       // 10
	fieldBoxString                           // 11 - non-null-terminated string (14496-30)
)

type field struct {
	children []*field
	name     string
	cnst     string
	order    int
	optFlag  uint32
	nOptFlag uint32
	size     uint
	length   uint
	flags    fieldFlag
	strType  stringType
	version  uint8
	nVersion uint8
}

func (f *field) set(flag fieldFlag) { _ = "STUB: not implemented"; return }

func (f *field) is(flag fieldFlag) bool { _ = "STUB: not implemented"; return false }

func buildFields(box IImmutableBox) []*field { _ = "STUB: not implemented"; return nil }

func buildFieldsStruct(t reflect.Type) []*field { _ = "STUB: not implemented"; return nil }

func buildFieldsAny(t reflect.Type) []*field { _ = "STUB: not implemented"; return nil }

func buildField(fieldName string, tag string) *field { _ = "STUB: not implemented"; return nil }

func parseFieldTag(str string) map[string]string { _ = "STUB: not implemented"; return nil }

type fieldInstance struct {
	field
	cfo ICustomFieldObject
}

func resolveFieldInstance(f *field, box IImmutableBox, parent reflect.Value, ctx Context) *fieldInstance {
	_ = "STUB: not implemented"
	return nil
}

func isTargetField(box IImmutableBox, fi *fieldInstance, ctx Context) bool {
	_ = "STUB: not implemented"
	return false
}
