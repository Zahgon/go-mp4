package edit

import (
	"math"
)

const UNoValue = math.MaxUint64

type Values struct {
	BaseMediaDecodeTime uint64
}

type Boxes []string

func (b Boxes) Exists(boxType string) bool { _ = "STUB: not implemented"; return false }

type Config struct {
	values    Values
	dropBoxes Boxes
}

var config Config

func Main(args []string) int { _ = "STUB: not implemented"; return 0 }

func editFile(inputPath, outputPath string) error { _ = "STUB: not implemented"; return nil }

// drop

// copy all data

// write header

// read payload

// edit some fields

// write payload

// expand all of offsprings

// rewrite box size
