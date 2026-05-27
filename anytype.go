package mp4

type IAnyType interface {
	IBox
	SetType(BoxType)
}

type AnyTypeBox struct {
	Box
	Type BoxType
}

func (e *AnyTypeBox) GetType() BoxType { _ = "STUB: not implemented"; return *new(BoxType) }

func (e *AnyTypeBox) SetType(boxType BoxType) { _ = "STUB: not implemented"; return }
