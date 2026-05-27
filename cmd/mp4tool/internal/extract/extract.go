package extract

import (
	"io"

	"github.com/abema/go-mp4"
)

const (
	blockSize        = 128 * 1024
	blockHistorySize = 4
)

func Main(args []string) int { _ = "STUB: not implemented"; return 0 }

func extract(r io.ReadSeeker, boxType mp4.BoxType) error { _ = "STUB: not implemented"; return nil }
