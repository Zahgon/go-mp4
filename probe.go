package mp4

import (
	"io"

	"github.com/abema/go-mp4/internal/bitio"
)

type ProbeInfo struct {
	MajorBrand       [4]byte
	MinorVersion     uint32
	CompatibleBrands [][4]byte
	FastStart        bool
	Timescale        uint32
	Duration         uint64
	Tracks           Tracks
	Segments         Segments
}

// Deprecated: replace with ProbeInfo
type FraProbeInfo = ProbeInfo

type Tracks []*Track

// Deprecated: replace with Track
type TrackInfo = Track

type Track struct {
	TrackID   uint32
	Timescale uint32
	Duration  uint64
	Codec     Codec
	Encrypted bool
	EditList  EditList
	Samples   Samples
	Chunks    Chunks
	AVC       *AVCDecConfigInfo
	MP4A      *MP4AInfo
}

type Codec int

const (
	CodecUnknown Codec = iota
	CodecAVC1
	CodecMP4A
)

type EditList []*EditListEntry

type EditListEntry struct {
	MediaTime       int64
	SegmentDuration uint64
}

type Samples []*Sample

type Sample struct {
	Size                  uint32
	TimeDelta             uint32
	CompositionTimeOffset int64
}

type Chunks []*Chunk

type Chunk struct {
	DataOffset      uint64
	SamplesPerChunk uint32
}

type AVCDecConfigInfo struct {
	ConfigurationVersion uint8
	Profile              uint8
	ProfileCompatibility uint8
	Level                uint8
	LengthSize           uint16
	Width                uint16
	Height               uint16
}

type MP4AInfo struct {
	OTI          uint8
	AudOTI       uint8
	ChannelCount uint16
}

type Segments []*Segment

// Deprecated: replace with Segment
type SegmentInfo = Segment

type Segment struct {
	TrackID               uint32
	MoofOffset            uint64
	BaseMediaDecodeTime   uint64
	DefaultSampleDuration uint32
	SampleCount           uint32
	Duration              uint32
	CompositionTimeOffset int32
	Size                  uint32
}

// Probe probes MP4 file
func Probe(r io.ReadSeeker) (*ProbeInfo, error) { _ = "STUB: not implemented"; return nil, nil }

// ProbeFra probes fragmented MP4 file
// Deprecated: replace with Probe
func ProbeFra(r io.ReadSeeker) (*FraProbeInfo, error) { _ = "STUB: not implemented"; return nil, nil }

func probeTrak(r io.ReadSeeker, bi *BoxInfo) (*Track, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func detectAACProfile(esds *Esds) (oti, audOTI uint8, err error) {
	_ = "STUB: not implemented"
	return 0, 0, nil
}

// audio object type

// sampling frequency index

func findDescriptorByTag(dscrs []Descriptor, tag int8) *Descriptor {
	_ = "STUB: not implemented"
	return nil
}

func getAudioObjectType(r bitio.Reader) (byte, int, error) {
	_ = "STUB: not implemented"
	return 0, 0, nil
}

func probeMoof(r io.ReadSeeker, bi *BoxInfo) (*Segment, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func FindIDRFrames(r io.ReadSeeker, trackInfo *TrackInfo) ([]int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (samples Samples) GetBitrate(timescale uint32) uint64 { _ = "STUB: not implemented"; return 0 }

func (samples Samples) GetMaxBitrate(timescale uint32, timeDelta uint64) uint64 {
	_ = "STUB: not implemented"
	return 0
}

func (segments Segments) GetBitrate(trackID uint32, timescale uint32) uint64 {
	_ = "STUB: not implemented"
	return 0
}

func (segments Segments) GetMaxBitrate(trackID uint32, timescale uint32) uint64 {
	_ = "STUB: not implemented"
	return 0
}
