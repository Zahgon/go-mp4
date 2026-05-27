package mp4

import (
	"io"

	"github.com/abema/go-mp4/internal/bitio"
)

/*************************** btrt ****************************/

func BoxTypeBtrt() BoxType { _ = "STUB: not implemented"; return *new(BoxType) }

func init() {
	AddBoxDef(&Btrt{}, 0)
}

type Btrt struct {
	Box
	BufferSizeDB uint32 `mp4:"0,size=32"`
	MaxBitrate   uint32 `mp4:"1,size=32"`
	AvgBitrate   uint32 `mp4:"2,size=32"`
}

// GetType returns the BoxType
func (*Btrt) GetType() BoxType {
	_ = "STUB: not implemented"
	return

	/*************************** co64 ****************************/
	*new(BoxType)
}

func BoxTypeCo64() BoxType { _ = "STUB: not implemented"; return *new(BoxType) }

func init() {
	AddBoxDef(&Co64{}, 0)
}

type Co64 struct {
	FullBox     `mp4:"0,extend"`
	EntryCount  uint32   `mp4:"1,size=32"`
	ChunkOffset []uint64 `mp4:"2,size=64,len=dynamic"`
}

// GetType returns the BoxType
func (*Co64) GetType() BoxType {
	_ = "STUB: not implemented"
	return *

	// GetFieldLength returns length of dynamic field
	new(BoxType)
}

func (co64 *Co64) GetFieldLength(name string, ctx Context) uint {
	_ = "STUB: not implemented"
	return 0
}

/*************************** colr ****************************/

func BoxTypeColr() BoxType { _ = "STUB: not implemented"; return *new(BoxType) }

func init() {
	AddBoxDef(&Colr{})
}

type Colr struct {
	Box
	ColourType              [4]byte `mp4:"0,size=8,string"`
	ColourPrimaries         uint16  `mp4:"1,size=16,opt=dynamic"`
	TransferCharacteristics uint16  `mp4:"2,size=16,opt=dynamic"`
	MatrixCoefficients      uint16  `mp4:"3,size=16,opt=dynamic"`
	FullRangeFlag           bool    `mp4:"4,size=1,opt=dynamic"`
	Reserved                uint8   `mp4:"5,size=7,opt=dynamic"`
	Profile                 []byte  `mp4:"6,size=8,opt=dynamic"`
	Unknown                 []byte  `mp4:"7,size=8,opt=dynamic"`
}

func (colr *Colr) IsOptFieldEnabled(name string, ctx Context) bool {
	_ = "STUB: not implemented"
	return false
}

// GetType returns the BoxType
func (*Colr) GetType() BoxType {
	_ = "STUB: not implemented"
	return

	/*************************** cslg ****************************/
	*new(BoxType)
}

func BoxTypeCslg() BoxType { _ = "STUB: not implemented"; return *new(BoxType) }

func init() {
	AddBoxDef(&Cslg{}, 0, 1)
}

type Cslg struct {
	FullBox                        `mp4:"0,extend"`
	CompositionToDTSShiftV0        int32 `mp4:"1,size=32,ver=0"`
	LeastDecodeToDisplayDeltaV0    int32 `mp4:"2,size=32,ver=0"`
	GreatestDecodeToDisplayDeltaV0 int32 `mp4:"3,size=32,ver=0"`
	CompositionStartTimeV0         int32 `mp4:"4,size=32,ver=0"`
	CompositionEndTimeV0           int32 `mp4:"5,size=32,ver=0"`
	CompositionToDTSShiftV1        int64 `mp4:"6,size=64,nver=0"`
	LeastDecodeToDisplayDeltaV1    int64 `mp4:"7,size=64,nver=0"`
	GreatestDecodeToDisplayDeltaV1 int64 `mp4:"8,size=64,nver=0"`
	CompositionStartTimeV1         int64 `mp4:"9,size=64,nver=0"`
	CompositionEndTimeV1           int64 `mp4:"10,size=64,nver=0"`
}

// GetType returns the BoxType
func (*Cslg) GetType() BoxType { _ = "STUB: not implemented"; return *new(BoxType) }

func (cslg *Cslg) GetCompositionToDTSShift() int64 { _ = "STUB: not implemented"; return 0 }

func (cslg *Cslg) GetLeastDecodeToDisplayDelta() int64 { _ = "STUB: not implemented"; return 0 }

func (cslg *Cslg) GetGreatestDecodeToDisplayDelta() int64 { _ = "STUB: not implemented"; return 0 }

func (cslg *Cslg) GetCompositionStartTime() int64 { _ = "STUB: not implemented"; return 0 }

func (cslg *Cslg) GetCompositionEndTime() int64 { _ = "STUB: not implemented"; return 0 }

/*************************** ctts ****************************/

func BoxTypeCtts() BoxType { _ = "STUB: not implemented"; return *new(BoxType) }

func init() {
	AddBoxDef(&Ctts{}, 0, 1)
}

type Ctts struct {
	FullBox    `mp4:"0,extend"`
	EntryCount uint32      `mp4:"1,size=32"`
	Entries    []CttsEntry `mp4:"2,len=dynamic,size=64"`
}

type CttsEntry struct {
	SampleCount    uint32 `mp4:"0,size=32"`
	SampleOffsetV0 uint32 `mp4:"1,size=32,ver=0"`
	SampleOffsetV1 int32  `mp4:"2,size=32,ver=1"`
}

// GetType returns the BoxType
func (*Ctts) GetType() BoxType {
	_ = "STUB: not implemented"
	return *

	// GetFieldLength returns length of dynamic field
	new(BoxType)
}

func (ctts *Ctts) GetFieldLength(name string, ctx Context) uint {
	_ = "STUB: not implemented"
	return 0
}

func (ctts *Ctts) GetSampleOffset(index int) int64 { _ = "STUB: not implemented"; return 0 }

/*************************** dinf ****************************/

func BoxTypeDinf() BoxType { _ = "STUB: not implemented"; return *new(BoxType) }

func init() {
	AddBoxDef(&Dinf{})
}

// Dinf is ISOBMFF dinf box type
type Dinf struct {
	Box
}

// GetType returns the BoxType
func (*Dinf) GetType() BoxType {
	_ = "STUB: not implemented"
	return

	/*************************** dref ****************************/
	*new(BoxType)
}

func BoxTypeDref() BoxType { _ = "STUB: not implemented"; return *new(BoxType) }
func BoxTypeUrl() BoxType  { _ = "STUB: not implemented"; return *new(BoxType) }
func BoxTypeUrn() BoxType  { _ = "STUB: not implemented"; return *new(BoxType) }

func init() {
	AddBoxDef(&Dref{}, 0)
	AddBoxDef(&Url{}, 0)
	AddBoxDef(&Urn{}, 0)
}

// Dref is ISOBMFF dref box type
type Dref struct {
	FullBox    `mp4:"0,extend"`
	EntryCount uint32 `mp4:"1,size=32"`
}

// GetType returns the BoxType
func (*Dref) GetType() BoxType { _ = "STUB: not implemented"; return *new(BoxType) }

type Url struct {
	FullBox  `mp4:"0,extend"`
	Location string `mp4:"1,string,nopt=0x000001"`
}

func (*Url) GetType() BoxType { _ = "STUB: not implemented"; return *new(BoxType) }

const UrlSelfContained = 0x000001

type Urn struct {
	FullBox  `mp4:"0,extend"`
	Name     string `mp4:"1,string,nopt=0x000001"`
	Location string `mp4:"2,string,nopt=0x000001"`
}

func (*Urn) GetType() BoxType { _ = "STUB: not implemented"; return *new(BoxType) }

const UrnSelfContained = 0x000001

/*************************** edts ****************************/

func BoxTypeEdts() BoxType { _ = "STUB: not implemented"; return *new(BoxType) }

func init() {
	AddBoxDef(&Edts{})
}

// Edts is ISOBMFF edts box type
type Edts struct {
	Box
}

// GetType returns the BoxType
func (*Edts) GetType() BoxType {
	_ = "STUB: not implemented"
	return

	/*************************** elst ****************************/
	*new(BoxType)
}

func BoxTypeElst() BoxType { _ = "STUB: not implemented"; return *new(BoxType) }

func init() {
	AddBoxDef(&Elst{}, 0, 1)
}

// Elst is ISOBMFF elst box type
type Elst struct {
	FullBox    `mp4:"0,extend"`
	EntryCount uint32      `mp4:"1,size=32"`
	Entries    []ElstEntry `mp4:"2,len=dynamic,size=dynamic"`
}

type ElstEntry struct {
	SegmentDurationV0 uint32 `mp4:"0,size=32,ver=0"`
	MediaTimeV0       int32  `mp4:"1,size=32,ver=0"`
	SegmentDurationV1 uint64 `mp4:"2,size=64,ver=1"`
	MediaTimeV1       int64  `mp4:"3,size=64,ver=1"`
	MediaRateInteger  int16  `mp4:"4,size=16"`
	MediaRateFraction int16  `mp4:"5,size=16,const=0"`
}

// GetType returns the BoxType
func (*Elst) GetType() BoxType {
	_ = "STUB: not implemented"
	return *

	// GetFieldSize returns size of dynamic field
	new(BoxType)
}

func (elst *Elst) GetFieldSize(name string, ctx Context) uint { _ = "STUB: not implemented"; return 0 }

/* segmentDurationV0 */
/* mediaTimeV0       */
/* mediaRateInteger  */
/* mediaRateFraction */

/* segmentDurationV1 */
/* mediaTimeV1       */
/* mediaRateInteger  */
/* mediaRateFraction */

// GetFieldLength returns length of dynamic field
func (elst *Elst) GetFieldLength(name string, ctx Context) uint {
	_ = "STUB: not implemented"
	return 0
}

func (elst *Elst) GetSegmentDuration(index int) uint64 { _ = "STUB: not implemented"; return 0 }

func (elst *Elst) GetMediaTime(index int) int64 { _ = "STUB: not implemented"; return 0 }

/*************************** emsg ****************************/

func BoxTypeEmsg() BoxType { _ = "STUB: not implemented"; return *new(BoxType) }

func init() {
	AddBoxDef(&Emsg{}, 0, 1)
}

// Emsg is ISOBMFF emsg box type
type Emsg struct {
	FullBox               `mp4:"0,extend"`
	SchemeIdUri           string `mp4:"1,string"`
	Value                 string `mp4:"2,string"`
	Timescale             uint32 `mp4:"3,size=32"`
	PresentationTimeDelta uint32 `mp4:"4,size=32,ver=0"`
	PresentationTime      uint64 `mp4:"5,size=64,ver=1"`
	EventDuration         uint32 `mp4:"6,size=32"`
	Id                    uint32 `mp4:"7,size=32"`
	MessageData           []byte `mp4:"8,size=8,string"`
}

func (emsg *Emsg) OnReadField(name string, r bitio.ReadSeeker, leftBits uint64, ctx Context) (rbits uint64, override bool, err error) {
	_ = "STUB: not implemented"
	return 0, false, nil
}

func (emsg *Emsg) OnWriteField(name string, w bitio.Writer, ctx Context) (wbits uint64, override bool, err error) {
	_ = "STUB: not implemented"
	return 0, false, nil
}

// GetType returns the BoxType
func (*Emsg) GetType() BoxType {
	_ = "STUB: not implemented"
	return

	/*************************** fiel ****************************/
	*new(BoxType)
}

func BoxTypeFiel() BoxType { _ = "STUB: not implemented"; return *new(BoxType) }

func init() {
	AddBoxDef(&Fiel{})
}

type Fiel struct {
	Box
	FieldCount    uint8 `mp4:"0,size=8"`
	FieldOrdering uint8 `mp4:"1,size=8"`
}

func (Fiel) GetType() BoxType {
	_ = "STUB: not implemented"
	return

	/************************ free, skip *************************/
	*new(BoxType)
}

func BoxTypeFree() BoxType { _ = "STUB: not implemented"; return *new(BoxType) }
func BoxTypeSkip() BoxType { _ = "STUB: not implemented"; return *new(BoxType) }

func init() {
	AddBoxDef(&Free{})
	AddBoxDef(&Skip{})
}

type FreeSpace struct {
	Box
	Data []uint8 `mp4:"0,size=8"`
}

type Free FreeSpace

func (*Free) GetType() BoxType { _ = "STUB: not implemented"; return *new(BoxType) }

type Skip FreeSpace

func (*Skip) GetType() BoxType {
	_ = "STUB: not implemented"
	return

	/*************************** frma ****************************/
	*new(BoxType)
}

func BoxTypeFrma() BoxType { _ = "STUB: not implemented"; return *new(BoxType) }

func init() {
	AddBoxDef(&Frma{})
}

// Frma is ISOBMFF frma box type
type Frma struct {
	Box
	DataFormat [4]byte `mp4:"0,size=8,string"`
}

// GetType returns the BoxType
func (*Frma) GetType() BoxType {
	_ = "STUB: not implemented"
	return

	/*************************** ftyp ****************************/
	*new(BoxType)
}

func BoxTypeFtyp() BoxType { _ = "STUB: not implemented"; return *new(BoxType) }

func init() {
	AddBoxDef(&Ftyp{})
}

func BrandQT() [4]byte   { _ = "STUB: not implemented"; return nil }
func BrandISOM() [4]byte { _ = "STUB: not implemented"; return nil }
func BrandISO2() [4]byte { _ = "STUB: not implemented"; return nil }
func BrandISO3() [4]byte { _ = "STUB: not implemented"; return nil }
func BrandISO4() [4]byte { _ = "STUB: not implemented"; return nil }
func BrandISO5() [4]byte { _ = "STUB: not implemented"; return nil }
func BrandISO6() [4]byte { _ = "STUB: not implemented"; return nil }
func BrandISO7() [4]byte { _ = "STUB: not implemented"; return nil }
func BrandISO8() [4]byte { _ = "STUB: not implemented"; return nil }
func BrandISO9() [4]byte { _ = "STUB: not implemented"; return nil }
func BrandAVC1() [4]byte { _ = "STUB: not implemented"; return nil }
func BrandMP41() [4]byte { _ = "STUB: not implemented"; return nil }
func BrandMP71() [4]byte { _ = "STUB: not implemented"; return nil }

// Ftyp is ISOBMFF ftyp box type
type Ftyp struct {
	Box
	MajorBrand       [4]byte               `mp4:"0,size=8,string"`
	MinorVersion     uint32                `mp4:"1,size=32"`
	CompatibleBrands []CompatibleBrandElem `mp4:"2,size=32"` // reach to end of the box
}

type CompatibleBrandElem struct {
	CompatibleBrand [4]byte `mp4:"0,size=8,string"`
}

func (ftyp *Ftyp) AddCompatibleBrand(cb [4]byte) { _ = "STUB: not implemented"; return }

func (ftyp *Ftyp) RemoveCompatibleBrand(cb [4]byte) { _ = "STUB: not implemented"; return }

func (ftyp *Ftyp) HasCompatibleBrand(cb [4]byte) bool { _ = "STUB: not implemented"; return false }

// GetType returns the BoxType
func (*Ftyp) GetType() BoxType {
	_ = "STUB: not implemented"
	return

	/*************************** hdlr ****************************/
	*new(BoxType)
}

func BoxTypeHdlr() BoxType { _ = "STUB: not implemented"; return *new(BoxType) }

func init() {
	AddBoxDef(&Hdlr{}, 0)
}

// Hdlr is ISOBMFF hdlr box type
type Hdlr struct {
	FullBox `mp4:"0,extend"`
	// Predefined corresponds to component_type of QuickTime.
	// pre_defined of ISO-14496 has always zero,
	// however component_type has "mhlr" or "dhlr".
	PreDefined  uint32    `mp4:"1,size=32"`
	HandlerType [4]byte   `mp4:"2,size=8,string"`
	Reserved    [3]uint32 `mp4:"3,size=32,const=0"`
	Name        string    `mp4:"4,string"`
}

// GetType returns the BoxType
func (*Hdlr) GetType() BoxType { _ = "STUB: not implemented"; return *new(BoxType) }

func (hdlr *Hdlr) OnReadField(name string, r bitio.ReadSeeker, leftBits uint64, ctx Context) (rbits uint64, override bool, err error) {
	_ = "STUB: not implemented"
	return 0, false, nil
}

func (hdlr *Hdlr) OnReadName(r bitio.ReadSeeker, leftBits uint64, ctx Context) (rbits uint64, override bool, err error) {
	_ = "STUB: not implemented"
	return 0, false, nil
}

// Pascal-style String

// C-style String

/*************************** hvcC ****************************/

func BoxTypeHvcC() BoxType { _ = "STUB: not implemented"; return *new(BoxType) }

func init() {
	AddBoxDef(&HvcC{})
}

type HEVCNalu struct {
	BaseCustomFieldObject
	Length  uint16 `mp4:"0,size=16"`
	NALUnit []byte `mp4:"1,size=8,len=dynamic"`
}

func (s HEVCNalu) GetFieldLength(name string, ctx Context) uint {
	_ = "STUB: not implemented"
	return 0
}

type HEVCNaluArray struct {
	BaseCustomFieldObject
	Completeness bool       `mp4:"0,size=1"`
	Reserved     bool       `mp4:"1,size=1"`
	NaluType     uint8      `mp4:"2,size=6"`
	NumNalus     uint16     `mp4:"3,size=16"`
	Nalus        []HEVCNalu `mp4:"4,len=dynamic"`
}

func (a HEVCNaluArray) GetFieldLength(name string, ctx Context) uint {
	_ = "STUB: not implemented"
	return 0
}

type HvcC struct {
	Box
	ConfigurationVersion        uint8           `mp4:"0,size=8"`
	GeneralProfileSpace         uint8           `mp4:"1,size=2"`
	GeneralTierFlag             bool            `mp4:"2,size=1"`
	GeneralProfileIdc           uint8           `mp4:"3,size=5"`
	GeneralProfileCompatibility [32]bool        `mp4:"4,size=1"`
	GeneralConstraintIndicator  [6]uint8        `mp4:"5,size=8"`
	GeneralLevelIdc             uint8           `mp4:"6,size=8"`
	Reserved1                   uint8           `mp4:"7,size=4,const=15"`
	MinSpatialSegmentationIdc   uint16          `mp4:"8,size=12"`
	Reserved2                   uint8           `mp4:"9,size=6,const=63"`
	ParallelismType             uint8           `mp4:"10,size=2"`
	Reserved3                   uint8           `mp4:"11,size=6,const=63"`
	ChromaFormatIdc             uint8           `mp4:"12,size=2"`
	Reserved4                   uint8           `mp4:"13,size=5,const=31"`
	BitDepthLumaMinus8          uint8           `mp4:"14,size=3"`
	Reserved5                   uint8           `mp4:"15,size=5,const=31"`
	BitDepthChromaMinus8        uint8           `mp4:"16,size=3"`
	AvgFrameRate                uint16          `mp4:"17,size=16"`
	ConstantFrameRate           uint8           `mp4:"18,size=2"`
	NumTemporalLayers           uint8           `mp4:"19,size=2"`
	TemporalIdNested            uint8           `mp4:"20,size=2"`
	LengthSizeMinusOne          uint8           `mp4:"21,size=2"`
	NumOfNaluArrays             uint8           `mp4:"22,size=8"`
	NaluArrays                  []HEVCNaluArray `mp4:"23,len=dynamic"`
}

func (HvcC) GetType() BoxType { _ = "STUB: not implemented"; return *new(BoxType) }

func (hvcc HvcC) GetFieldLength(name string, ctx Context) uint { _ = "STUB: not implemented"; return 0 }

/*************************** mdat ****************************/

func BoxTypeMdat() BoxType { _ = "STUB: not implemented"; return *new(BoxType) }

func init() {
	AddBoxDef(&Mdat{})
}

// Mdat is ISOBMFF mdat box type
type Mdat struct {
	Box
	Data []byte `mp4:"0,size=8"`
}

// GetType returns the BoxType
func (*Mdat) GetType() BoxType {
	_ = "STUB: not implemented"
	return

	/*************************** mdhd ****************************/
	*new(BoxType)
}

func BoxTypeMdhd() BoxType { _ = "STUB: not implemented"; return *new(BoxType) }

func init() {
	AddBoxDef(&Mdhd{}, 0, 1)
}

// Mdhd is ISOBMFF mdhd box type
type Mdhd struct {
	FullBox            `mp4:"0,extend"`
	CreationTimeV0     uint32 `mp4:"1,size=32,ver=0"`
	ModificationTimeV0 uint32 `mp4:"2,size=32,ver=0"`
	CreationTimeV1     uint64 `mp4:"3,size=64,ver=1"`
	ModificationTimeV1 uint64 `mp4:"4,size=64,ver=1"`
	Timescale          uint32 `mp4:"5,size=32"`
	DurationV0         uint32 `mp4:"6,size=32,ver=0"`
	DurationV1         uint64 `mp4:"7,size=64,ver=1"`
	//
	Pad        bool    `mp4:"8,size=1,hidden"`
	Language   [3]byte `mp4:"9,size=5,iso639-2"` // ISO-639-2/T language code
	PreDefined uint16  `mp4:"10,size=16"`
}

// GetType returns the BoxType
func (*Mdhd) GetType() BoxType { _ = "STUB: not implemented"; return *new(BoxType) }

func (mdhd *Mdhd) GetCreationTime() uint64 { _ = "STUB: not implemented"; return 0 }

func (mdhd *Mdhd) GetModificationTime() uint64 { _ = "STUB: not implemented"; return 0 }

func (mdhd *Mdhd) GetDuration() uint64 { _ = "STUB: not implemented"; return 0 }

/*************************** mdia ****************************/

func BoxTypeMdia() BoxType { _ = "STUB: not implemented"; return *new(BoxType) }

func init() {
	AddBoxDef(&Mdia{})
}

// Mdia is ISOBMFF mdia box type
type Mdia struct {
	Box
}

// GetType returns the BoxType
func (*Mdia) GetType() BoxType {
	_ = "STUB: not implemented"
	return

	/*************************** mehd ****************************/
	*new(BoxType)
}

func BoxTypeMehd() BoxType { _ = "STUB: not implemented"; return *new(BoxType) }

func init() {
	AddBoxDef(&Mehd{}, 0, 1)
}

// Mehd is ISOBMFF mehd box type
type Mehd struct {
	FullBox            `mp4:"0,extend"`
	FragmentDurationV0 uint32 `mp4:"1,size=32,ver=0"`
	FragmentDurationV1 uint64 `mp4:"2,size=64,ver=1"`
}

// GetType returns the BoxType
func (*Mehd) GetType() BoxType { _ = "STUB: not implemented"; return *new(BoxType) }

func (mdhd *Mehd) GetFragmentDuration() uint64 { _ = "STUB: not implemented"; return 0 }

/*************************** meta ****************************/

func BoxTypeMeta() BoxType { _ = "STUB: not implemented"; return *new(BoxType) }

func init() {
	AddBoxDef(&Meta{}, 0)
}

// Meta is ISOBMFF meta box type
type Meta struct {
	FullBox `mp4:"0,extend"`
}

// GetType returns the BoxType
func (*Meta) GetType() BoxType { _ = "STUB: not implemented"; return *new(BoxType) }

func (meta *Meta) BeforeUnmarshal(r io.ReadSeeker, size uint64, ctx Context) (n uint64, override bool, err error) {
	_ = "STUB: not implemented"
	// for Apple Quick Time
	return 0, false, nil
}

/*************************** mfhd ****************************/

func BoxTypeMfhd() BoxType { _ = "STUB: not implemented"; return *new(BoxType) }

func init() {
	AddBoxDef(&Mfhd{}, 0)
}

// Mfhd is ISOBMFF mfhd box type
type Mfhd struct {
	FullBox        `mp4:"0,extend"`
	SequenceNumber uint32 `mp4:"1,size=32"`
}

// GetType returns the BoxType
func (*Mfhd) GetType() BoxType {
	_ = "STUB: not implemented"
	return

	/*************************** mfra ****************************/
	*new(BoxType)
}

func BoxTypeMfra() BoxType { _ = "STUB: not implemented"; return *new(BoxType) }

func init() {
	AddBoxDef(&Mfra{})
}

// Mfra is ISOBMFF mfra box type
type Mfra struct {
	Box
}

// GetType returns the BoxType
func (*Mfra) GetType() BoxType {
	_ = "STUB: not implemented"
	return

	/*************************** mfro ****************************/
	*new(BoxType)
}

func BoxTypeMfro() BoxType { _ = "STUB: not implemented"; return *new(BoxType) }

func init() {
	AddBoxDef(&Mfro{}, 0)
}

// Mfro is ISOBMFF mfro box type
type Mfro struct {
	FullBox `mp4:"0,extend"`
	Size    uint32 `mp4:"1,size=32"`
}

// GetType returns the BoxType
func (*Mfro) GetType() BoxType {
	_ = "STUB: not implemented"
	return

	/*************************** minf ****************************/
	*new(BoxType)
}

func BoxTypeMinf() BoxType { _ = "STUB: not implemented"; return *new(BoxType) }

func init() {
	AddBoxDef(&Minf{})
}

// Minf is ISOBMFF minf box type
type Minf struct {
	Box
}

// GetType returns the BoxType
func (*Minf) GetType() BoxType {
	_ = "STUB: not implemented"
	return

	/*************************** moof ****************************/
	*new(BoxType)
}

func BoxTypeMoof() BoxType { _ = "STUB: not implemented"; return *new(BoxType) }

func init() {
	AddBoxDef(&Moof{})
}

// Moof is ISOBMFF moof box type
type Moof struct {
	Box
}

// GetType returns the BoxType
func (*Moof) GetType() BoxType {
	_ = "STUB: not implemented"
	return

	/*************************** moov ****************************/
	*new(BoxType)
}

func BoxTypeMoov() BoxType { _ = "STUB: not implemented"; return *new(BoxType) }

func init() {
	AddBoxDef(&Moov{})
}

// Moov is ISOBMFF moov box type
type Moov struct {
	Box
}

// GetType returns the BoxType
func (*Moov) GetType() BoxType {
	_ = "STUB: not implemented"
	return

	/*************************** mvex ****************************/
	*new(BoxType)
}

func BoxTypeMvex() BoxType { _ = "STUB: not implemented"; return *new(BoxType) }

func init() {
	AddBoxDef(&Mvex{})
}

// Mvex is ISOBMFF mvex box type
type Mvex struct {
	Box
}

// GetType returns the BoxType
func (*Mvex) GetType() BoxType {
	_ = "STUB: not implemented"
	return

	/*************************** mvhd ****************************/
	*new(BoxType)
}

func BoxTypeMvhd() BoxType { _ = "STUB: not implemented"; return *new(BoxType) }

func init() {
	AddBoxDef(&Mvhd{}, 0, 1)
}

// Mvhd is ISOBMFF mvhd box type
type Mvhd struct {
	FullBox            `mp4:"0,extend"`
	CreationTimeV0     uint32    `mp4:"1,size=32,ver=0"`
	ModificationTimeV0 uint32    `mp4:"2,size=32,ver=0"`
	CreationTimeV1     uint64    `mp4:"3,size=64,ver=1"`
	ModificationTimeV1 uint64    `mp4:"4,size=64,ver=1"`
	Timescale          uint32    `mp4:"5,size=32"`
	DurationV0         uint32    `mp4:"6,size=32,ver=0"`
	DurationV1         uint64    `mp4:"7,size=64,ver=1"`
	Rate               int32     `mp4:"8,size=32"` // fixed-point 16.16 - template=0x00010000
	Volume             int16     `mp4:"9,size=16"` // template=0x0100
	Reserved           int16     `mp4:"10,size=16,const=0"`
	Reserved2          [2]uint32 `mp4:"11,size=32,const=0"`
	Matrix             [9]int32  `mp4:"12,size=32,hex"` // template={ 0x00010000,0,0,0,0x00010000,0,0,0,0x40000000 }
	PreDefined         [6]int32  `mp4:"13,size=32"`
	NextTrackID        uint32    `mp4:"14,size=32"`
}

// GetType returns the BoxType
func (*Mvhd) GetType() BoxType {
	_ = "STUB: not implemented"
	return *

	// StringifyField returns field value as string
	new(BoxType)
}

func (mvhd *Mvhd) StringifyField(name string, indent string, depth int, ctx Context) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

func (mvhd *Mvhd) GetCreationTime() uint64 { _ = "STUB: not implemented"; return 0 }

func (mvhd *Mvhd) GetModificationTime() uint64 { _ = "STUB: not implemented"; return 0 }

func (mvhd *Mvhd) GetDuration() uint64 { _ = "STUB: not implemented"; return 0 }

// GetRate returns value of rate as float64
func (mvhd *Mvhd) GetRate() float64 { _ = "STUB: not implemented"; return 0 }

// GetRateInt returns value of rate as int16
func (mvhd *Mvhd) GetRateInt() int16 { _ = "STUB: not implemented"; return 0 }

/*************************** saio ****************************/

func BoxTypeSaio() BoxType { _ = "STUB: not implemented"; return *new(BoxType) }

func init() {
	AddBoxDef(&Saio{}, 0, 1)
}

type Saio struct {
	FullBox              `mp4:"0,extend"`
	AuxInfoType          [4]byte  `mp4:"1,size=8,opt=0x000001,string"`
	AuxInfoTypeParameter uint32   `mp4:"2,size=32,opt=0x000001,hex"`
	EntryCount           uint32   `mp4:"3,size=32"`
	OffsetV0             []uint32 `mp4:"4,size=32,ver=0,len=dynamic"`
	OffsetV1             []uint64 `mp4:"5,size=64,nver=0,len=dynamic"`
}

func (saio *Saio) GetFieldLength(name string, ctx Context) uint {
	_ = "STUB: not implemented"
	return 0
}

func (*Saio) GetType() BoxType { _ = "STUB: not implemented"; return *new(BoxType) }

func (saio *Saio) GetOffset(index int) uint64 { _ = "STUB: not implemented"; return 0 }

/*************************** saiz ****************************/

func BoxTypeSaiz() BoxType { _ = "STUB: not implemented"; return *new(BoxType) }

func init() {
	AddBoxDef(&Saiz{}, 0)
}

type Saiz struct {
	FullBox               `mp4:"0,extend"`
	AuxInfoType           [4]byte `mp4:"1,size=8,opt=0x000001,string"`
	AuxInfoTypeParameter  uint32  `mp4:"2,size=32,opt=0x000001,hex"`
	DefaultSampleInfoSize uint8   `mp4:"3,size=8,dec"`
	SampleCount           uint32  `mp4:"4,size=32"`
	SampleInfoSize        []uint8 `mp4:"5,size=8,opt=dynamic,len=dynamic,dec"`
}

func (saiz *Saiz) IsOptFieldEnabled(name string, ctx Context) bool {
	_ = "STUB: not implemented"
	return false
}

func (saiz *Saiz) GetFieldLength(name string, ctx Context) uint {
	_ = "STUB: not implemented"
	return 0
}

func (*Saiz) GetType() BoxType {
	_ = "STUB: not implemented"
	return

	/*********************** SampleEntry *************************/
	*new(BoxType)
}

func BoxTypeMp4v() BoxType { _ = "STUB: not implemented"; return *new(BoxType) }
func BoxTypeAvc1() BoxType { _ = "STUB: not implemented"; return *new(BoxType) }
func BoxTypeEncv() BoxType { _ = "STUB: not implemented"; return *new(BoxType) }
func BoxTypeHev1() BoxType { _ = "STUB: not implemented"; return *new(BoxType) }
func BoxTypeHvc1() BoxType { _ = "STUB: not implemented"; return *new(BoxType) }
func BoxTypeMp4a() BoxType { _ = "STUB: not implemented"; return *new(BoxType) }
func BoxTypeEnca() BoxType { _ = "STUB: not implemented"; return *new(BoxType) }
func BoxTypeAvcC() BoxType { _ = "STUB: not implemented"; return *new(BoxType) }
func BoxTypePasp() BoxType { _ = "STUB: not implemented"; return *new(BoxType) }
func BoxTypeStpp() BoxType { _ = "STUB: not implemented"; return *new(BoxType) }
func BoxTypeSbtt() BoxType { _ = "STUB: not implemented"; return *new(BoxType) }

func init() {
	AddAnyTypeBoxDef(&VisualSampleEntry{}, BoxTypeMp4v())
	AddAnyTypeBoxDef(&VisualSampleEntry{}, BoxTypeAvc1())
	AddAnyTypeBoxDef(&VisualSampleEntry{}, BoxTypeEncv())
	AddAnyTypeBoxDef(&VisualSampleEntry{}, BoxTypeHev1())
	AddAnyTypeBoxDef(&VisualSampleEntry{}, BoxTypeHvc1())
	AddAnyTypeBoxDef(&AudioSampleEntry{}, BoxTypeMp4a())
	AddAnyTypeBoxDef(&AudioSampleEntry{}, BoxTypeEnca())
	AddAnyTypeBoxDef(&AVCDecoderConfiguration{}, BoxTypeAvcC())
	AddAnyTypeBoxDef(&PixelAspectRatioBox{}, BoxTypePasp())
	AddAnyTypeBoxDef(&XMLSubtitleSampleEntry{}, BoxTypeStpp())
	AddAnyTypeBoxDef(&TextSubtitleSampleEntry{}, BoxTypeSbtt())
}

type SampleEntry struct {
	AnyTypeBox
	Reserved           [6]uint8 `mp4:"0,size=8,const=0"`
	DataReferenceIndex uint16   `mp4:"1,size=16"`
}

type VisualSampleEntry struct {
	SampleEntry     `mp4:"0,extend"`
	PreDefined      uint16    `mp4:"1,size=16"`
	Reserved        uint16    `mp4:"2,size=16,const=0"`
	PreDefined2     [3]uint32 `mp4:"3,size=32"`
	Width           uint16    `mp4:"4,size=16"`
	Height          uint16    `mp4:"5,size=16"`
	Horizresolution uint32    `mp4:"6,size=32"`
	Vertresolution  uint32    `mp4:"7,size=32"`
	Reserved2       uint32    `mp4:"8,size=32,const=0"`
	FrameCount      uint16    `mp4:"9,size=16"`
	Compressorname  [32]byte  `mp4:"10,size=8"`
	Depth           uint16    `mp4:"11,size=16"`
	PreDefined3     int16     `mp4:"12,size=16"`
}

// StringifyField returns field value as string
func (vse *VisualSampleEntry) StringifyField(name string, indent string, depth int, ctx Context) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

type AudioSampleEntry struct {
	SampleEntry   `mp4:"0,extend,opt=dynamic"`
	EntryVersion  uint16    `mp4:"1,size=16,opt=dynamic"`
	Reserved      [3]uint16 `mp4:"2,size=16,opt=dynamic,const=0"`
	ChannelCount  uint16    `mp4:"3,size=16,opt=dynamic"`
	SampleSize    uint16    `mp4:"4,size=16,opt=dynamic"`
	PreDefined    uint16    `mp4:"5,size=16,opt=dynamic"`
	Reserved2     uint16    `mp4:"6,size=16,opt=dynamic,const=0"`
	SampleRate    uint32    `mp4:"7,size=32,opt=dynamic"` // fixed-point 16.16
	QuickTimeData []byte    `mp4:"8,size=8,opt=dynamic,len=dynamic"`
}

func (ase *AudioSampleEntry) IsOptFieldEnabled(name string, ctx Context) bool {
	_ = "STUB: not implemented"
	return false
}

func (ase *AudioSampleEntry) GetFieldLength(name string, ctx Context) uint {
	_ = "STUB: not implemented"
	return 0
}

// StringifyField returns field value as string
func (ase *AudioSampleEntry) StringifyField(name string, indent string, depth int, ctx Context) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

func (ase *AudioSampleEntry) GetSampleRate() float64 { _ = "STUB: not implemented"; return 0 }

func (ase *AudioSampleEntry) GetSampleRateInt() uint16 { _ = "STUB: not implemented"; return 0 }

const (
	AVCBaselineProfile uint8 = 66  // 0x42
	AVCMainProfile     uint8 = 77  // 0x4d
	AVCExtendedProfile uint8 = 88  // 0x58
	AVCHighProfile     uint8 = 100 // 0x64
	AVCHigh10Profile   uint8 = 110 // 0x6e
	AVCHigh422Profile  uint8 = 122 // 0x7a
)

type AVCDecoderConfiguration struct {
	AnyTypeBox
	ConfigurationVersion         uint8             `mp4:"0,size=8"`
	Profile                      uint8             `mp4:"1,size=8"`
	ProfileCompatibility         uint8             `mp4:"2,size=8"`
	Level                        uint8             `mp4:"3,size=8"`
	Reserved                     uint8             `mp4:"4,size=6,const=63"`
	LengthSizeMinusOne           uint8             `mp4:"5,size=2"`
	Reserved2                    uint8             `mp4:"6,size=3,const=7"`
	NumOfSequenceParameterSets   uint8             `mp4:"7,size=5"`
	SequenceParameterSets        []AVCParameterSet `mp4:"8,len=dynamic"`
	NumOfPictureParameterSets    uint8             `mp4:"9,size=8"`
	PictureParameterSets         []AVCParameterSet `mp4:"10,len=dynamic"`
	HighProfileFieldsEnabled     bool              `mp4:"11,hidden"`
	Reserved3                    uint8             `mp4:"12,size=6,opt=dynamic,const=63"`
	ChromaFormat                 uint8             `mp4:"13,size=2,opt=dynamic"`
	Reserved4                    uint8             `mp4:"14,size=5,opt=dynamic,const=31"`
	BitDepthLumaMinus8           uint8             `mp4:"15,size=3,opt=dynamic"`
	Reserved5                    uint8             `mp4:"16,size=5,opt=dynamic,const=31"`
	BitDepthChromaMinus8         uint8             `mp4:"17,size=3,opt=dynamic"`
	NumOfSequenceParameterSetExt uint8             `mp4:"18,size=8,opt=dynamic"`
	SequenceParameterSetsExt     []AVCParameterSet `mp4:"19,len=dynamic,opt=dynamic"`
}

func (avcc *AVCDecoderConfiguration) GetFieldLength(name string, ctx Context) uint {
	_ = "STUB: not implemented"
	return 0
}

func (avcc *AVCDecoderConfiguration) IsOptFieldEnabled(name string, ctx Context) bool {
	_ = "STUB: not implemented"
	return false
}

func (avcc *AVCDecoderConfiguration) OnReadField(name string, r bitio.ReadSeeker, leftBits uint64, ctx Context) (rbits uint64, override bool, err error) {
	_ = "STUB: not implemented"
	return 0, false, nil
}

func (avcc *AVCDecoderConfiguration) OnWriteField(name string, w bitio.Writer, ctx Context) (wbits uint64, override bool, err error) {
	_ = "STUB: not implemented"
	return 0, false, nil
}

type AVCParameterSet struct {
	BaseCustomFieldObject
	Length  uint16 `mp4:"0,size=16"`
	NALUnit []byte `mp4:"1,size=8,len=dynamic"`
}

func (s *AVCParameterSet) GetFieldLength(name string, ctx Context) uint {
	_ = "STUB: not implemented"
	return 0
}

type PixelAspectRatioBox struct {
	AnyTypeBox
	HSpacing uint32 `mp4:"0,size=32"`
	VSpacing uint32 `mp4:"1,size=32"`
}

type XMLSubtitleSampleEntry struct {
	SampleEntry        `mp4:"0,extend"`
	Namespace          string `mp4:"1,string"` // space-separated list
	SchemaLocation     string `mp4:"2,string"` // space-separated list, optional
	AuxiliaryMIMETypes string `mp4:"3,string"` // space-separated list, optional
}

func (xse *XMLSubtitleSampleEntry) GetNamespaceList() []string {
	_ = "STUB: not implemented"
	return nil
}

func (xse *XMLSubtitleSampleEntry) GetSchemaLocationList() []string {
	_ = "STUB: not implemented"
	return nil
}

func (xse *XMLSubtitleSampleEntry) GetAuxiliaryMIMETypesList() []string {
	_ = "STUB: not implemented"
	return nil
}

type TextSubtitleSampleEntry struct {
	SampleEntry     `mp4:"0,extend"`
	ContentEncoding string `mp4:"1,string"` // optional
	MIMEFormat      string `mp4:"2,string"`
}

/*************************** sbgp ****************************/

func BoxTypeSbgp() BoxType { _ = "STUB: not implemented"; return *new(BoxType) }

func init() {
	AddBoxDef(&Sbgp{}, 0, 1)
}

type Sbgp struct {
	FullBox               `mp4:"0,extend"`
	GroupingType          uint32      `mp4:"1,size=32"`
	GroupingTypeParameter uint32      `mp4:"2,size=32,ver=1"`
	EntryCount            uint32      `mp4:"3,size=32"`
	Entries               []SbgpEntry `mp4:"4,len=dynamic,size=64"`
}

type SbgpEntry struct {
	SampleCount           uint32 `mp4:"0,size=32"`
	GroupDescriptionIndex uint32 `mp4:"1,size=32"`
}

func (sbgp *Sbgp) GetFieldLength(name string, ctx Context) uint {
	_ = "STUB: not implemented"
	return 0
}

func (*Sbgp) GetType() BoxType {
	_ = "STUB: not implemented"
	return

	/*************************** schi ****************************/
	*new(BoxType)
}

func BoxTypeSchi() BoxType { _ = "STUB: not implemented"; return *new(BoxType) }

func init() {
	AddBoxDef(&Schi{})
}

type Schi struct {
	Box
}

func (*Schi) GetType() BoxType {
	_ = "STUB: not implemented"
	return

	/*************************** schm ****************************/
	*new(BoxType)
}

func BoxTypeSchm() BoxType { _ = "STUB: not implemented"; return *new(BoxType) }

func init() {
	AddBoxDef(&Schm{}, 0)
}

type Schm struct {
	FullBox       `mp4:"0,extend"`
	SchemeType    [4]byte `mp4:"1,size=8,string"`
	SchemeVersion uint32  `mp4:"2,size=32,hex"`
	SchemeUri     []byte  `mp4:"3,size=8,opt=0x000001,string"`
}

func (*Schm) GetType() BoxType {
	_ = "STUB: not implemented"
	return

	/*************************** sdtp ****************************/
	*new(BoxType)
}

func BoxTypeSdtp() BoxType { _ = "STUB: not implemented"; return *new(BoxType) }

func init() {
	AddBoxDef(&Sdtp{}, 0)
}

type Sdtp struct {
	FullBox `mp4:"0,extend"`
	Samples []SdtpSampleElem `mp4:"1,size=8"`
}

type SdtpSampleElem struct {
	IsLeading           uint8 `mp4:"0,size=2"`
	SampleDependsOn     uint8 `mp4:"1,size=2"`
	SampleIsDependedOn  uint8 `mp4:"2,size=2"`
	SampleHasRedundancy uint8 `mp4:"3,size=2"`
}

func (*Sdtp) GetType() BoxType {
	_ = "STUB: not implemented"
	return

	/*************************** sgpd ****************************/
	*new(BoxType)
}

func BoxTypeSgpd() BoxType { _ = "STUB: not implemented"; return *new(BoxType) }

func init() {
	AddBoxDef(&Sgpd{}, 1, 2) // version 0 is deprecated by ISO/IEC 14496-12
}

type Sgpd struct {
	FullBox                       `mp4:"0,extend"`
	GroupingType                  [4]byte                    `mp4:"1,size=8,string"`
	DefaultLength                 uint32                     `mp4:"2,size=32,ver=1"`
	DefaultSampleDescriptionIndex uint32                     `mp4:"3,size=32,ver=2"`
	EntryCount                    uint32                     `mp4:"4,size=32"`
	RollDistances                 []int16                    `mp4:"5,size=16,opt=dynamic"`
	RollDistancesL                []RollDistanceWithLength   `mp4:"6,size=16,opt=dynamic"`
	AlternativeStartupEntries     []AlternativeStartupEntry  `mp4:"7,size=dynamic,len=dynamic,opt=dynamic"`
	AlternativeStartupEntriesL    []AlternativeStartupEntryL `mp4:"8,len=dynamic,opt=dynamic"`
	VisualRandomAccessEntries     []VisualRandomAccessEntry  `mp4:"9,len=dynamic,opt=dynamic"`
	VisualRandomAccessEntriesL    []VisualRandomAccessEntryL `mp4:"10,len=dynamic,opt=dynamic"`
	TemporalLevelEntries          []TemporalLevelEntry       `mp4:"11,len=dynamic,opt=dynamic"`
	TemporalLevelEntriesL         []TemporalLevelEntryL      `mp4:"12,len=dynamic,opt=dynamic"`
	Unsupported                   []byte                     `mp4:"13,size=8,opt=dynamic"`
}

type RollDistanceWithLength struct {
	DescriptionLength uint32 `mp4:"0,size=32"`
	RollDistance      int16  `mp4:"1,size=16"`
}

type AlternativeStartupEntry struct {
	BaseCustomFieldObject
	RollCount         uint16                       `mp4:"0,size=16"`
	FirstOutputSample uint16                       `mp4:"1,size=16"`
	SampleOffset      []uint32                     `mp4:"2,size=32,len=dynamic"`
	Opts              []AlternativeStartupEntryOpt `mp4:"3,size=32"`
}

type AlternativeStartupEntryL struct {
	DescriptionLength       uint32 `mp4:"0,size=32"`
	AlternativeStartupEntry `mp4:"1,extend,size=dynamic"`
}

type AlternativeStartupEntryOpt struct {
	NumOutputSamples uint16 `mp4:"0,size=16"`
	NumTotalSamples  uint16 `mp4:"1,size=16"`
}

type VisualRandomAccessEntry struct {
	NumLeadingSamplesKnown bool  `mp4:"0,size=1"`
	NumLeadingSamples      uint8 `mp4:"1,size=7"`
}

type VisualRandomAccessEntryL struct {
	DescriptionLength       uint32 `mp4:"0,size=32"`
	VisualRandomAccessEntry `mp4:"1,extend"`
}

type TemporalLevelEntry struct {
	LevelIndependentlyDecodable bool  `mp4:"0,size=1"`
	Reserved                    uint8 `mp4:"1,size=7,const=0"`
}

type TemporalLevelEntryL struct {
	DescriptionLength  uint32 `mp4:"0,size=32"`
	TemporalLevelEntry `mp4:"1,extend"`
}

func (sgpd *Sgpd) GetFieldSize(name string, ctx Context) uint { _ = "STUB: not implemented"; return 0 }

func (sgpd *Sgpd) GetFieldLength(name string, ctx Context) uint {
	_ = "STUB: not implemented"
	return 0
}

func (sgpd *Sgpd) IsOptFieldEnabled(name string, ctx Context) bool {
	_ = "STUB: not implemented"
	return false
}

func (*Sgpd) GetType() BoxType { _ = "STUB: not implemented"; return *new(BoxType) }

func (entry *AlternativeStartupEntry) GetFieldLength(name string, ctx Context) uint {
	_ = "STUB: not implemented"
	return 0
}

func (entry *AlternativeStartupEntryL) GetFieldSize(name string, ctx Context) uint {
	_ = "STUB: not implemented"
	return 0
}

/*************************** sidx ****************************/

func BoxTypeSidx() BoxType { _ = "STUB: not implemented"; return *new(BoxType) }

func init() {
	AddBoxDef(&Sidx{}, 0, 1)
}

type Sidx struct {
	FullBox                    `mp4:"0,extend"`
	ReferenceID                uint32          `mp4:"1,size=32"`
	Timescale                  uint32          `mp4:"2,size=32"`
	EarliestPresentationTimeV0 uint32          `mp4:"3,size=32,ver=0"`
	FirstOffsetV0              uint32          `mp4:"4,size=32,ver=0"`
	EarliestPresentationTimeV1 uint64          `mp4:"5,size=64,nver=0"`
	FirstOffsetV1              uint64          `mp4:"6,size=64,nver=0"`
	Reserved                   uint16          `mp4:"7,size=16,const=0"`
	ReferenceCount             uint16          `mp4:"8,size=16"`
	References                 []SidxReference `mp4:"9,size=96,len=dynamic"`
}

type SidxReference struct {
	ReferenceType      bool   `mp4:"0,size=1"`
	ReferencedSize     uint32 `mp4:"1,size=31"`
	SubsegmentDuration uint32 `mp4:"2,size=32"`
	StartsWithSAP      bool   `mp4:"3,size=1"`
	SAPType            uint32 `mp4:"4,size=3"`
	SAPDeltaTime       uint32 `mp4:"5,size=28"`
}

func (*Sidx) GetType() BoxType { _ = "STUB: not implemented"; return *new(BoxType) }

func (sidx *Sidx) GetFieldLength(name string, ctx Context) uint {
	_ = "STUB: not implemented"
	return 0
}

func (sidx *Sidx) GetEarliestPresentationTime() uint64 { _ = "STUB: not implemented"; return 0 }

func (sidx *Sidx) GetFirstOffset() uint64 { _ = "STUB: not implemented"; return 0 }

/*************************** sinf ****************************/

func BoxTypeSinf() BoxType { _ = "STUB: not implemented"; return *new(BoxType) }

func init() {
	AddBoxDef(&Sinf{})
}

type Sinf struct {
	Box
}

func (*Sinf) GetType() BoxType {
	_ = "STUB: not implemented"
	return

	/*************************** smhd ****************************/
	*new(BoxType)
}

func BoxTypeSmhd() BoxType { _ = "STUB: not implemented"; return *new(BoxType) }

func init() {
	AddBoxDef(&Smhd{}, 0)
}

type Smhd struct {
	FullBox  `mp4:"0,extend"`
	Balance  int16  `mp4:"1,size=16"` // fixed-point 8.8 template=0
	Reserved uint16 `mp4:"2,size=16,const=0"`
}

func (*Smhd) GetType() BoxType {
	_ = "STUB: not implemented"
	return *

	// StringifyField returns field value as string
	new(BoxType)
}

func (smhd *Smhd) StringifyField(name string, indent string, depth int, ctx Context) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

// GetBalance returns value of width as float32
func (smhd *Smhd) GetBalance() float32 { _ = "STUB: not implemented"; return 0 }

// GetBalanceInt returns value of width as int8
func (smhd *Smhd) GetBalanceInt() int8 { _ = "STUB: not implemented"; return 0 }

/*************************** stbl ****************************/

func BoxTypeStbl() BoxType { _ = "STUB: not implemented"; return *new(BoxType) }

func init() {
	AddBoxDef(&Stbl{})
}

// Stbl is ISOBMFF stbl box type
type Stbl struct {
	Box
}

// GetType returns the BoxType
func (*Stbl) GetType() BoxType {
	_ = "STUB: not implemented"
	return

	/*************************** stco ****************************/
	*new(BoxType)
}

func BoxTypeStco() BoxType { _ = "STUB: not implemented"; return *new(BoxType) }

func init() {
	AddBoxDef(&Stco{}, 0)
}

// Stco is ISOBMFF stco box type
type Stco struct {
	FullBox     `mp4:"0,extend"`
	EntryCount  uint32   `mp4:"1,size=32"`
	ChunkOffset []uint32 `mp4:"2,size=32,len=dynamic"`
}

// GetType returns the BoxType
func (*Stco) GetType() BoxType {
	_ = "STUB: not implemented"
	return *

	// GetFieldLength returns length of dynamic field
	new(BoxType)
}

func (stco *Stco) GetFieldLength(name string, ctx Context) uint {
	_ = "STUB: not implemented"
	return 0
}

/*************************** stsc ****************************/

func BoxTypeStsc() BoxType { _ = "STUB: not implemented"; return *new(BoxType) }

func init() {
	AddBoxDef(&Stsc{}, 0)
}

// Stsc is ISOBMFF stsc box type
type Stsc struct {
	FullBox    `mp4:"0,extend"`
	EntryCount uint32      `mp4:"1,size=32"`
	Entries    []StscEntry `mp4:"2,len=dynamic,size=96"`
}

type StscEntry struct {
	FirstChunk             uint32 `mp4:"0,size=32"`
	SamplesPerChunk        uint32 `mp4:"1,size=32"`
	SampleDescriptionIndex uint32 `mp4:"2,size=32"`
}

// GetType returns the BoxType
func (*Stsc) GetType() BoxType {
	_ = "STUB: not implemented"
	return *

	// GetFieldLength returns length of dynamic field
	new(BoxType)
}

func (stsc *Stsc) GetFieldLength(name string, ctx Context) uint {
	_ = "STUB: not implemented"
	return 0
}

/*************************** stsd ****************************/

func BoxTypeStsd() BoxType { _ = "STUB: not implemented"; return *new(BoxType) }

func init() {
	AddBoxDef(&Stsd{}, 0)
}

// Stsd is ISOBMFF stsd box type
type Stsd struct {
	FullBox    `mp4:"0,extend"`
	EntryCount uint32 `mp4:"1,size=32"`
}

// GetType returns the BoxType
func (*Stsd) GetType() BoxType {
	_ = "STUB: not implemented"
	return

	/*************************** stss ****************************/
	*new(BoxType)
}

func BoxTypeStss() BoxType { _ = "STUB: not implemented"; return *new(BoxType) }

func init() {
	AddBoxDef(&Stss{}, 0)
}

type Stss struct {
	FullBox      `mp4:"0,extend"`
	EntryCount   uint32   `mp4:"1,size=32"`
	SampleNumber []uint32 `mp4:"2,len=dynamic,size=32"`
}

// GetType returns the BoxType
func (*Stss) GetType() BoxType {
	_ = "STUB: not implemented"
	return *

	// GetFieldLength returns length of dynamic field
	new(BoxType)
}

func (stss *Stss) GetFieldLength(name string, ctx Context) uint {
	_ = "STUB: not implemented"
	return 0
}

/*************************** stsz ****************************/

func BoxTypeStsz() BoxType { _ = "STUB: not implemented"; return *new(BoxType) }

func init() {
	AddBoxDef(&Stsz{}, 0)
}

// Stsz is ISOBMFF stsz box type
type Stsz struct {
	FullBox     `mp4:"0,extend"`
	SampleSize  uint32   `mp4:"1,size=32"`
	SampleCount uint32   `mp4:"2,size=32"`
	EntrySize   []uint32 `mp4:"3,size=32,len=dynamic"`
}

// GetType returns the BoxType
func (*Stsz) GetType() BoxType {
	_ = "STUB: not implemented"
	return *

	// GetFieldLength returns length of dynamic field
	new(BoxType)
}

func (stsz *Stsz) GetFieldLength(name string, ctx Context) uint {
	_ = "STUB: not implemented"
	return 0
}

/*************************** stts ****************************/

func BoxTypeStts() BoxType { _ = "STUB: not implemented"; return *new(BoxType) }

func init() {
	AddBoxDef(&Stts{}, 0)
}

// Stts is ISOBMFF stts box type
type Stts struct {
	FullBox    `mp4:"0,extend"`
	EntryCount uint32      `mp4:"1,size=32"`
	Entries    []SttsEntry `mp4:"2,len=dynamic,size=64"`
}

type SttsEntry struct {
	SampleCount uint32 `mp4:"0,size=32"`
	SampleDelta uint32 `mp4:"1,size=32"`
}

// GetType returns the BoxType
func (*Stts) GetType() BoxType {
	_ = "STUB: not implemented"
	return *

	// GetFieldLength returns length of dynamic field
	new(BoxType)
}

func (stts *Stts) GetFieldLength(name string, ctx Context) uint {
	_ = "STUB: not implemented"
	return 0
}

/*************************** styp ****************************/

func BoxTypeStyp() BoxType { _ = "STUB: not implemented"; return *new(BoxType) }

func init() {
	AddBoxDef(&Styp{})
}

type Styp struct {
	Box
	MajorBrand       [4]byte               `mp4:"0,size=8,string"`
	MinorVersion     uint32                `mp4:"1,size=32"`
	CompatibleBrands []CompatibleBrandElem `mp4:"2,size=32"` // reach to end of the box
}

func (*Styp) GetType() BoxType {
	_ = "STUB: not implemented"
	return

	/*************************** tfdt ****************************/
	*new(BoxType)
}

func BoxTypeTfdt() BoxType { _ = "STUB: not implemented"; return *new(BoxType) }

func init() {
	AddBoxDef(&Tfdt{}, 0, 1)
}

// Tfdt is ISOBMFF tfdt box type
type Tfdt struct {
	FullBox               `mp4:"0,extend"`
	BaseMediaDecodeTimeV0 uint32 `mp4:"1,size=32,ver=0"`
	BaseMediaDecodeTimeV1 uint64 `mp4:"2,size=64,ver=1"`
}

// GetType returns the BoxType
func (*Tfdt) GetType() BoxType { _ = "STUB: not implemented"; return *new(BoxType) }

func (tfdt *Tfdt) GetBaseMediaDecodeTime() uint64 { _ = "STUB: not implemented"; return 0 }

/*************************** tfhd ****************************/

func BoxTypeTfhd() BoxType { _ = "STUB: not implemented"; return *new(BoxType) }

func init() {
	AddBoxDef(&Tfhd{}, 0)
}

// Tfhd is ISOBMFF tfhd box type
type Tfhd struct {
	FullBox `mp4:"0,extend"`
	TrackID uint32 `mp4:"1,size=32"`

	// optional
	BaseDataOffset         uint64 `mp4:"2,size=64,opt=0x000001"`
	SampleDescriptionIndex uint32 `mp4:"3,size=32,opt=0x000002"`
	DefaultSampleDuration  uint32 `mp4:"4,size=32,opt=0x000008"`
	DefaultSampleSize      uint32 `mp4:"5,size=32,opt=0x000010"`
	DefaultSampleFlags     uint32 `mp4:"6,size=32,opt=0x000020,hex"`
}

const (
	TfhdBaseDataOffsetPresent         = 0x000001
	TfhdSampleDescriptionIndexPresent = 0x000002
	TfhdDefaultSampleDurationPresent  = 0x000008
	TfhdDefaultSampleSizePresent      = 0x000010
	TfhdDefaultSampleFlagsPresent     = 0x000020
	TfhdDurationIsEmpty               = 0x010000
	TfhdDefaultBaseIsMoof             = 0x020000
)

// GetType returns the BoxType
func (*Tfhd) GetType() BoxType {
	_ = "STUB: not implemented"
	return

	/*************************** tfra ****************************/
	*new(BoxType)
}

func BoxTypeTfra() BoxType { _ = "STUB: not implemented"; return *new(BoxType) }

func init() {
	AddBoxDef(&Tfra{}, 0, 1)
}

// Tfra is ISOBMFF tfra box type
type Tfra struct {
	FullBox               `mp4:"0,extend"`
	TrackID               uint32      `mp4:"1,size=32"`
	Reserved              uint32      `mp4:"2,size=26,const=0"`
	LengthSizeOfTrafNum   byte        `mp4:"3,size=2"`
	LengthSizeOfTrunNum   byte        `mp4:"4,size=2"`
	LengthSizeOfSampleNum byte        `mp4:"5,size=2"`
	NumberOfEntry         uint32      `mp4:"6,size=32"`
	Entries               []TfraEntry `mp4:"7,len=dynamic,size=dynamic"`
}

type TfraEntry struct {
	TimeV0       uint32 `mp4:"0,size=32,ver=0"`
	MoofOffsetV0 uint32 `mp4:"1,size=32,ver=0"`
	TimeV1       uint64 `mp4:"2,size=64,ver=1"`
	MoofOffsetV1 uint64 `mp4:"3,size=64,ver=1"`
	TrafNumber   uint32 `mp4:"4,size=dynamic"`
	TrunNumber   uint32 `mp4:"5,size=dynamic"`
	SampleNumber uint32 `mp4:"6,size=dynamic"`
}

// GetType returns the BoxType
func (*Tfra) GetType() BoxType {
	_ = "STUB: not implemented"
	return *

	// GetFieldSize returns size of dynamic field
	new(BoxType)
}

func (tfra *Tfra) GetFieldSize(name string, ctx Context) uint { _ = "STUB: not implemented"; return 0 }

/* TimeV0       */
/* MoofOffsetV0 */
/* TrafNumber   */
/* TrunNumber   */
/* SampleNumber */

/* TimeV1       */
/* MoofOffsetV1 */
/* TrafNumber   */
/* TrunNumber   */
/* SampleNumber */

// GetFieldLength returns length of dynamic field
func (tfra *Tfra) GetFieldLength(name string, ctx Context) uint {
	_ = "STUB: not implemented"
	return 0
}

func (tfra *Tfra) GetTime(index int) uint64 { _ = "STUB: not implemented"; return 0 }

func (tfra *Tfra) GetMoofOffset(index int) uint64 { _ = "STUB: not implemented"; return 0 }

/*************************** tkhd ****************************/

func BoxTypeTkhd() BoxType { _ = "STUB: not implemented"; return *new(BoxType) }

func init() {
	AddBoxDef(&Tkhd{}, 0, 1)
}

// Tkhd is ISOBMFF tkhd box type
type Tkhd struct {
	FullBox            `mp4:"0,extend"`
	CreationTimeV0     uint32 `mp4:"1,size=32,ver=0"`
	ModificationTimeV0 uint32 `mp4:"2,size=32,ver=0"`
	CreationTimeV1     uint64 `mp4:"3,size=64,ver=1"`
	ModificationTimeV1 uint64 `mp4:"4,size=64,ver=1"`
	TrackID            uint32 `mp4:"5,size=32"`
	Reserved0          uint32 `mp4:"6,size=32,const=0"`
	DurationV0         uint32 `mp4:"7,size=32,ver=0"`
	DurationV1         uint64 `mp4:"8,size=64,ver=1"`
	//
	Reserved1      [2]uint32 `mp4:"9,size=32,const=0"`
	Layer          int16     `mp4:"10,size=16"` // template=0
	AlternateGroup int16     `mp4:"11,size=16"` // template=0
	Volume         int16     `mp4:"12,size=16"` // template={if track_is_audio 0x0100 else 0}
	Reserved2      uint16    `mp4:"13,size=16,const=0"`
	Matrix         [9]int32  `mp4:"14,size=32,hex"` // template={ 0x00010000,0,0,0,0x00010000,0,0,0,0x40000000 };
	Width          uint32    `mp4:"15,size=32"`     // fixed-point 16.16
	Height         uint32    `mp4:"16,size=32"`     // fixed-point 16.16
}

// GetType returns the BoxType
func (*Tkhd) GetType() BoxType {
	_ = "STUB: not implemented"
	return *

	// StringifyField returns field value as string
	new(BoxType)
}

func (tkhd *Tkhd) StringifyField(name string, indent string, depth int, ctx Context) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

func (tkhd *Tkhd) GetCreationTime() uint64 { _ = "STUB: not implemented"; return 0 }

func (tkhd *Tkhd) GetModificationTime() uint64 { _ = "STUB: not implemented"; return 0 }

func (tkhd *Tkhd) GetDuration() uint64 { _ = "STUB: not implemented"; return 0 }

// GetWidth returns value of width as float64
func (tkhd *Tkhd) GetWidth() float64 { _ = "STUB: not implemented"; return 0 }

// GetWidthInt returns value of width as uint16
func (tkhd *Tkhd) GetWidthInt() uint16 { _ = "STUB: not implemented"; return 0 }

// GetHeight returns value of height as float64
func (tkhd *Tkhd) GetHeight() float64 { _ = "STUB: not implemented"; return 0 }

// GetHeightInt returns value of height as uint16
func (tkhd *Tkhd) GetHeightInt() uint16 { _ = "STUB: not implemented"; return 0 }

/*************************** traf ****************************/

func BoxTypeTraf() BoxType { _ = "STUB: not implemented"; return *new(BoxType) }

func init() {
	AddBoxDef(&Traf{})
}

// Traf is ISOBMFF traf box type
type Traf struct {
	Box
}

// GetType returns the BoxType
func (*Traf) GetType() BoxType {
	_ = "STUB: not implemented"
	return

	/*************************** trak ****************************/
	*new(BoxType)
}

func BoxTypeTrak() BoxType { _ = "STUB: not implemented"; return *new(BoxType) }

func init() {
	AddBoxDef(&Trak{})
}

// Trak is ISOBMFF trak box type
type Trak struct {
	Box
}

// GetType returns the BoxType
func (*Trak) GetType() BoxType {
	_ = "STUB: not implemented"
	return

	/*************************** trep ****************************/
	*new(BoxType)
}

func BoxTypeTrep() BoxType { _ = "STUB: not implemented"; return *new(BoxType) }

func init() {
	AddBoxDef(&Trep{}, 0)
}

// Trep is ISOBMFF trep box type
type Trep struct {
	FullBox `mp4:"0,extend"`
	TrackID uint32 `mp4:"1,size=32"`
}

// GetType returns the BoxType
func (*Trep) GetType() BoxType {
	_ = "STUB: not implemented"
	return

	/*************************** trex ****************************/
	*new(BoxType)
}

func BoxTypeTrex() BoxType { _ = "STUB: not implemented"; return *new(BoxType) }

func init() {
	AddBoxDef(&Trex{}, 0)
}

// Trex is ISOBMFF trex box type
type Trex struct {
	FullBox                       `mp4:"0,extend"`
	TrackID                       uint32 `mp4:"1,size=32"`
	DefaultSampleDescriptionIndex uint32 `mp4:"2,size=32"`
	DefaultSampleDuration         uint32 `mp4:"3,size=32"`
	DefaultSampleSize             uint32 `mp4:"4,size=32"`
	DefaultSampleFlags            uint32 `mp4:"5,size=32,hex"`
}

// GetType returns the BoxType
func (*Trex) GetType() BoxType {
	_ = "STUB: not implemented"
	return

	/*************************** trun ****************************/
	*new(BoxType)
}

func BoxTypeTrun() BoxType { _ = "STUB: not implemented"; return *new(BoxType) }

func init() {
	AddBoxDef(&Trun{}, 0, 1)
}

// Trun is ISOBMFF trun box type
type Trun struct {
	FullBox     `mp4:"0,extend"`
	SampleCount uint32 `mp4:"1,size=32"`

	// optional fields
	DataOffset       int32       `mp4:"2,size=32,opt=0x000001"`
	FirstSampleFlags uint32      `mp4:"3,size=32,opt=0x000004,hex"`
	Entries          []TrunEntry `mp4:"4,len=dynamic,size=dynamic"`
}

type TrunEntry struct {
	SampleDuration                uint32 `mp4:"0,size=32,opt=0x000100"`
	SampleSize                    uint32 `mp4:"1,size=32,opt=0x000200"`
	SampleFlags                   uint32 `mp4:"2,size=32,opt=0x000400,hex"`
	SampleCompositionTimeOffsetV0 uint32 `mp4:"3,size=32,opt=0x000800,ver=0"`
	SampleCompositionTimeOffsetV1 int32  `mp4:"4,size=32,opt=0x000800,nver=0"`
}

// GetType returns the BoxType
func (*Trun) GetType() BoxType {
	_ = "STUB: not implemented"
	return *

	// GetFieldSize returns size of dynamic field
	new(BoxType)
}

func (trun *Trun) GetFieldSize(name string, ctx Context) uint { _ = "STUB: not implemented"; return 0 }

// SampleDuration

// SampleSize

// SampleFlags

// SampleCompositionTimeOffsetV0 or V1

// GetFieldLength returns length of dynamic field
func (trun *Trun) GetFieldLength(name string, ctx Context) uint {
	_ = "STUB: not implemented"
	return 0
}

func (trun *Trun) GetSampleCompositionTimeOffset(index int) int64 {
	_ = "STUB: not implemented"
	return 0
}

/*************************** udta ****************************/

func BoxTypeUdta() BoxType { _ = "STUB: not implemented"; return *new(BoxType) }

func init() {
	AddBoxDef(&Udta{})
}

// Udta is ISOBMFF udta box type
type Udta struct {
	Box
}

// GetType returns the BoxType
func (*Udta) GetType() BoxType { _ = "STUB: not implemented"; return *new(BoxType) }

func isUnderUdta(ctx Context) bool {
	_ = "STUB: not implemented"
	return

	/*************************** vmhd ****************************/
	false
}

func BoxTypeVmhd() BoxType { _ = "STUB: not implemented"; return *new(BoxType) }

func init() {
	AddBoxDef(&Vmhd{}, 0)
}

// Vmhd is ISOBMFF vmhd box type
type Vmhd struct {
	FullBox      `mp4:"0,extend"`
	Graphicsmode uint16    `mp4:"1,size=16"` // template=0
	Opcolor      [3]uint16 `mp4:"2,size=16"` // template={0, 0, 0}
}

// GetType returns the BoxType
func (*Vmhd) GetType() BoxType {
	_ = "STUB: not implemented"
	return

	/*************************** wave ****************************/
	*new(BoxType)
}

func BoxTypeWave() BoxType { _ = "STUB: not implemented"; return *new(BoxType) }

func init() {
	AddBoxDef(&Wave{})
}

// Wave is QuickTime wave box
type Wave struct {
	Box
}

// GetType returns the BoxType
func (*Wave) GetType() BoxType { _ = "STUB: not implemented"; return *new(BoxType) }
