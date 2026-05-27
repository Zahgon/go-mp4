package mp4

/*************************** pssh ****************************/

func BoxTypePssh() BoxType { _ = "STUB: not implemented"; return *new(BoxType) }

func init() {
	AddBoxDef(&Pssh{}, 0, 1)
}

// Pssh is ISOBMFF pssh box type
type Pssh struct {
	FullBox  `mp4:"0,extend"`
	SystemID [16]byte  `mp4:"1,size=8,uuid"`
	KIDCount uint32    `mp4:"2,size=32,nver=0"`
	KIDs     []PsshKID `mp4:"3,nver=0,len=dynamic,size=128"`
	DataSize int32     `mp4:"4,size=32"`
	Data     []byte    `mp4:"5,size=8,len=dynamic"`
}

type PsshKID struct {
	KID [16]byte `mp4:"0,size=8,uuid"`
}

// GetFieldLength returns length of dynamic field
func (pssh *Pssh) GetFieldLength(name string, ctx Context) uint {
	_ = "STUB: not implemented"
	return 0
}

// StringifyField returns field value as string
func (pssh *Pssh) StringifyField(name string, indent string, depth int, ctx Context) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

// GetType returns the BoxType
func (*Pssh) GetType() BoxType {
	_ = "STUB: not implemented"
	return

	/*************************** tenc ****************************/
	*new(BoxType)
}

func BoxTypeTenc() BoxType { _ = "STUB: not implemented"; return *new(BoxType) }

func init() {
	AddBoxDef(&Tenc{}, 0, 1)
}

// Tenc is ISOBMFF tenc box type
type Tenc struct {
	FullBox                `mp4:"0,extend"`
	Reserved               uint8    `mp4:"1,size=8,dec"`
	DefaultCryptByteBlock  uint8    `mp4:"2,size=4,dec"` // always 0 on version 0
	DefaultSkipByteBlock   uint8    `mp4:"3,size=4,dec"` // always 0 on version 0
	DefaultIsProtected     uint8    `mp4:"4,size=8,dec"`
	DefaultPerSampleIVSize uint8    `mp4:"5,size=8,dec"`
	DefaultKID             [16]byte `mp4:"6,size=8,uuid"`
	DefaultConstantIVSize  uint8    `mp4:"7,size=8,opt=dynamic,dec"`
	DefaultConstantIV      []byte   `mp4:"8,size=8,opt=dynamic,len=dynamic"`
}

func (tenc *Tenc) IsOptFieldEnabled(name string, ctx Context) bool {
	_ = "STUB: not implemented"
	return false
}

func (tenc *Tenc) GetFieldLength(name string, ctx Context) uint {
	_ = "STUB: not implemented"
	return 0
}

// GetType returns the BoxType
func (*Tenc) GetType() BoxType { _ = "STUB: not implemented"; return *new(BoxType) }
