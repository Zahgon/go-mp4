package mp4

/*************************** Opus ****************************/

// https://opus-codec.org/docs/opus_in_isobmff.html

func BoxTypeOpus() BoxType { _ = "STUB: not implemented"; return *new(BoxType) }

func init() {
	AddAnyTypeBoxDef(&AudioSampleEntry{}, BoxTypeOpus())
}

/*************************** dOps ****************************/

// https://opus-codec.org/docs/opus_in_isobmff.html

func BoxTypeDOps() BoxType { _ = "STUB: not implemented"; return *new(BoxType) }

func init() {
	AddBoxDef(&DOps{})
}

type DOps struct {
	Box
	Version              uint8   `mp4:"0,size=8"`
	OutputChannelCount   uint8   `mp4:"1,size=8"`
	PreSkip              uint16  `mp4:"2,size=16"`
	InputSampleRate      uint32  `mp4:"3,size=32"`
	OutputGain           int16   `mp4:"4,size=16"`
	ChannelMappingFamily uint8   `mp4:"5,size=8"`
	StreamCount          uint8   `mp4:"6,opt=dynamic,size=8"`
	CoupledCount         uint8   `mp4:"7,opt=dynamic,size=8"`
	ChannelMapping       []uint8 `mp4:"8,opt=dynamic,size=8,len=dynamic"`
}

func (DOps) GetType() BoxType { _ = "STUB: not implemented"; return *new(BoxType) }

func (dops DOps) IsOptFieldEnabled(name string, ctx Context) bool {
	_ = "STUB: not implemented"
	return false
}

func (ops DOps) GetFieldLength(name string, ctx Context) uint { _ = "STUB: not implemented"; return 0 }
