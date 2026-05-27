package mp4

/*********************** WebVTT Sample Entry ****************************/

func BoxTypeVttC() BoxType { _ = "STUB: not implemented"; return *new(BoxType) }
func BoxTypeVlab() BoxType { _ = "STUB: not implemented"; return *new(BoxType) }
func BoxTypeWvtt() BoxType { _ = "STUB: not implemented"; return *new(BoxType) }

func init() {
	AddBoxDef(&WebVTTConfigurationBox{})
	AddBoxDef(&WebVTTSourceLabelBox{})
	AddAnyTypeBoxDef(&WVTTSampleEntry{}, BoxTypeWvtt())
}

type WebVTTConfigurationBox struct {
	Box
	Config string `mp4:"0,boxstring"`
}

func (WebVTTConfigurationBox) GetType() BoxType { _ = "STUB: not implemented"; return *new(BoxType) }

type WebVTTSourceLabelBox struct {
	Box
	SourceLabel string `mp4:"0,boxstring"`
}

func (WebVTTSourceLabelBox) GetType() BoxType { _ = "STUB: not implemented"; return *new(BoxType) }

type WVTTSampleEntry struct {
	SampleEntry `mp4:"0,extend"`
}

/*********************** WebVTT Sample Format ****************************/

func BoxTypeVttc() BoxType { _ = "STUB: not implemented"; return *new(BoxType) }
func BoxTypeVsid() BoxType { _ = "STUB: not implemented"; return *new(BoxType) }
func BoxTypeCtim() BoxType { _ = "STUB: not implemented"; return *new(BoxType) }
func BoxTypeIden() BoxType { _ = "STUB: not implemented"; return *new(BoxType) }
func BoxTypeSttg() BoxType { _ = "STUB: not implemented"; return *new(BoxType) }
func BoxTypePayl() BoxType { _ = "STUB: not implemented"; return *new(BoxType) }
func BoxTypeVtte() BoxType { _ = "STUB: not implemented"; return *new(BoxType) }
func BoxTypeVtta() BoxType { _ = "STUB: not implemented"; return *new(BoxType) }

func init() {
	AddBoxDef(&VTTCueBox{})
	AddBoxDef(&CueSourceIDBox{})
	AddBoxDef(&CueTimeBox{})
	AddBoxDef(&CueIDBox{})
	AddBoxDef(&CueSettingsBox{})
	AddBoxDef(&CuePayloadBox{})
	AddBoxDef(&VTTEmptyCueBox{})
	AddBoxDef(&VTTAdditionalTextBox{})
}

type VTTCueBox struct {
	Box
}

func (VTTCueBox) GetType() BoxType { _ = "STUB: not implemented"; return *new(BoxType) }

type CueSourceIDBox struct {
	Box
	SourceId uint32 `mp4:"0,size=32"`
}

func (CueSourceIDBox) GetType() BoxType { _ = "STUB: not implemented"; return *new(BoxType) }

type CueTimeBox struct {
	Box
	CueCurrentTime string `mp4:"0,boxstring"`
}

func (CueTimeBox) GetType() BoxType { _ = "STUB: not implemented"; return *new(BoxType) }

type CueIDBox struct {
	Box
	CueId string `mp4:"0,boxstring"`
}

func (CueIDBox) GetType() BoxType { _ = "STUB: not implemented"; return *new(BoxType) }

type CueSettingsBox struct {
	Box
	Settings string `mp4:"0,boxstring"`
}

func (CueSettingsBox) GetType() BoxType { _ = "STUB: not implemented"; return *new(BoxType) }

type CuePayloadBox struct {
	Box
	CueText string `mp4:"0,boxstring"`
}

func (CuePayloadBox) GetType() BoxType { _ = "STUB: not implemented"; return *new(BoxType) }

type VTTEmptyCueBox struct {
	Box
}

func (VTTEmptyCueBox) GetType() BoxType { _ = "STUB: not implemented"; return *new(BoxType) }

type VTTAdditionalTextBox struct {
	Box
	CueAdditionalText string `mp4:"0,boxstring"`
}

func (VTTAdditionalTextBox) GetType() BoxType { _ = "STUB: not implemented"; return *new(BoxType) }
