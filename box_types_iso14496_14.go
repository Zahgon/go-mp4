package mp4

/*************************** esds ****************************/

// https://developer.apple.com/library/content/documentation/QuickTime/QTFF/QTFFChap3/qtff3.html

func BoxTypeEsds() BoxType { _ = "STUB: not implemented"; return *new(BoxType) }

func init() {
	AddBoxDef(&Esds{}, 0)
}

const (
	ESDescrTag            = 0x03
	DecoderConfigDescrTag = 0x04
	DecSpecificInfoTag    = 0x05
	SLConfigDescrTag      = 0x06
)

// Esds is ES descripter box
type Esds struct {
	FullBox     `mp4:"0,extend"`
	Descriptors []Descriptor `mp4:"1,array"`
}

// GetType returns the BoxType
func (*Esds) GetType() BoxType { _ = "STUB: not implemented"; return *new(BoxType) }

type Descriptor struct {
	BaseCustomFieldObject
	Tag                     int8                     `mp4:"0,size=8"` // must be 0x03
	Size                    uint32                   `mp4:"1,varint"`
	ESDescriptor            *ESDescriptor            `mp4:"2,extend,opt=dynamic"`
	DecoderConfigDescriptor *DecoderConfigDescriptor `mp4:"3,extend,opt=dynamic"`
	Data                    []byte                   `mp4:"4,size=8,opt=dynamic,len=dynamic"`
}

// GetFieldLength returns length of dynamic field
func (ds *Descriptor) GetFieldLength(name string, ctx Context) uint {
	_ = "STUB: not implemented"
	return 0
}

func (ds *Descriptor) IsOptFieldEnabled(name string, ctx Context) bool {
	_ = "STUB: not implemented"
	return false
}

// StringifyField returns field value as string
func (ds *Descriptor) StringifyField(name string, indent string, depth int, ctx Context) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

type ESDescriptor struct {
	BaseCustomFieldObject
	ESID                 uint16 `mp4:"0,size=16"`
	StreamDependenceFlag bool   `mp4:"1,size=1"`
	UrlFlag              bool   `mp4:"2,size=1"`
	OcrStreamFlag        bool   `mp4:"3,size=1"`
	StreamPriority       int8   `mp4:"4,size=5"`
	DependsOnESID        uint16 `mp4:"5,size=16,opt=dynamic"`
	URLLength            uint8  `mp4:"6,size=8,opt=dynamic"`
	URLString            []byte `mp4:"7,size=8,len=dynamic,opt=dynamic,string"`
	OCRESID              uint16 `mp4:"8,size=16,opt=dynamic"`
}

func (esds *ESDescriptor) GetFieldLength(name string, ctx Context) uint {
	_ = "STUB: not implemented"
	return 0
}

func (esds *ESDescriptor) IsOptFieldEnabled(name string, ctx Context) bool {
	_ = "STUB: not implemented"
	return false
}

type DecoderConfigDescriptor struct {
	BaseCustomFieldObject
	ObjectTypeIndication byte   `mp4:"0,size=8"`
	StreamType           int8   `mp4:"1,size=6"`
	UpStream             bool   `mp4:"2,size=1"`
	Reserved             bool   `mp4:"3,size=1"`
	BufferSizeDB         uint32 `mp4:"4,size=24"`
	MaxBitrate           uint32 `mp4:"5,size=32"`
	AvgBitrate           uint32 `mp4:"6,size=32"`
}
