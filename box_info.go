package mp4

import (
	"io"
)

type Context struct {
	// IsQuickTimeCompatible represents whether ftyp.compatible_brands contains "qt  ".
	IsQuickTimeCompatible bool

	// QuickTimeKeysMetaEntryCount the expected number of items under the ilst box as observed from the keys box
	QuickTimeKeysMetaEntryCount int

	// UnderWave represents whether current box is under the wave box.
	UnderWave bool

	// UnderIlst represents whether current box is under the ilst box.
	UnderIlst bool

	// UnderIlstMeta represents whether current box is under the metadata box under the ilst box.
	UnderIlstMeta bool

	// UnderIlstFreeMeta represents whether current box is under "----" box.
	UnderIlstFreeMeta bool

	// UnderUdta represents whether current box is under the udta box.
	UnderUdta bool
}

// BoxInfo has common infomations of box
type BoxInfo struct {
	// Offset specifies an offset of the box in a file.
	Offset uint64

	// Size specifies size(bytes) of box.
	Size uint64

	// HeaderSize specifies size(bytes) of common fields which are defined as "Box" class member at ISO/IEC 14496-12.
	HeaderSize uint64

	// Type specifies box type which is represented by 4 characters.
	Type BoxType

	// ExtendToEOF is set true when Box.size is zero. It means that end of box equals to end of file.
	ExtendToEOF bool

	// Context would be set by ReadBoxStructure, not ReadBoxInfo.
	Context
}

func (bi *BoxInfo) IsSupportedType() bool { _ = "STUB: not implemented"; return false }

const (
	SmallHeaderSize = 8
	LargeHeaderSize = 16
)

// EncodeBoxInfo encodes the common header fields of a box as defined in ISO/IEC 14496-12.
// If bi.ExtendToEOF is true, the size field is set to zero (indicating the box extends to EOF).
// If the size fits in a uint32 and HeaderSize is not LargeHeaderSize, an 8-byte small header is used.
// Otherwise, a 16-byte large header is used with size field set to 1 and the actual size in the extended field.
func EncodeBoxInfo(bi *BoxInfo) []byte { _ = "STUB: not implemented"; return nil }

// WriteBoxInfo writes common fields which are defined as "Box" class member at ISO/IEC 14496-12.
// This function ignores bi.Offset and returns BoxInfo which contains real Offset and recalculated Size/HeaderSize.
func WriteBoxInfo(w io.WriteSeeker, bi *BoxInfo) (*BoxInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ReadBoxInfo reads common fields which are defined as "Box" class member at ISO/IEC 14496-12.
func ReadBoxInfo(r io.ReadSeeker) (*BoxInfo, error) { _ = "STUB: not implemented"; return nil, nil }

// read 8 bytes

// pick size and type

// box extends to end of file

// read more 8 bytes

func (bi *BoxInfo) SeekToStart(s io.Seeker) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (bi *BoxInfo) SeekToPayload(s io.Seeker) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (bi *BoxInfo) SeekToEnd(s io.Seeker) (int64, error) { _ = "STUB: not implemented"; return 0, nil }
