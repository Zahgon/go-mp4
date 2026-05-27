package probe

import (
	"io"
)

func Main(args []string) int { _ = "STUB: not implemented"; return 0 }

type report struct {
	MajorBrand       string   `yaml:"major_brand"`
	MinorVersion     uint32   `yaml:"minor_version"`
	CompatibleBrands []string `yaml:"compatible_brands"`
	FastStart        bool     `yaml:"fast_start"`
	Timescale        uint32   `yaml:"timescale"`
	Duration         uint64   `yaml:"duration"`
	DurationSeconds  float32  `yaml:"duration_seconds"`
	Tracks           []*track `yaml:"tracks"`
}

type track struct {
	TrackID         uint32  `yaml:"track_id"`
	Timescale       uint32  `yaml:"timescale"`
	Duration        uint64  `yaml:"duration"`
	DurationSeconds float32 `yaml:"duration_seconds"`
	Codec           string  `yaml:"codec"`
	Encrypted       bool    `yaml:"encrypted"`
	Width           uint16  `json:",omitempty" yaml:"width,omitempty"`
	Height          uint16  `json:",omitempty" yaml:"height,omitempty"`
	SampleNum       int     `json:",omitempty" yaml:"sample_num,omitempty"`
	ChunkNum        int     `json:",omitempty" yaml:"chunk_num,omitempty"`
	IDRFrameNum     int     `json:",omitempty" yaml:"idr_frame_num,omitempty"`
	Bitrate         uint64  `json:",omitempty" yaml:"bitrate,omitempty"`
	MaxBitrate      uint64  `json:",omitempty" yaml:"max_bitrate,omitempty"`
}

func buildReport(r io.ReadSeeker) (*report, error) { _ = "STUB: not implemented"; return nil, nil }
