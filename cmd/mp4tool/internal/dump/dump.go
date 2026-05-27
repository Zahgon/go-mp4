package dump

import (
	"io"
	"os"

	"golang.org/x/term"
)

const (
	indentSize       = 2
	blockSize        = 128 * 1024
	blockHistorySize = 4
)

var terminalWidth = 180

func init() {
	if width, _, err := term.GetSize(int(os.Stdin.Fd())); err == nil {
		terminalWidth = width
	}
}

func Main(args []string) int { _ = "STUB: not implemented"; return 0 }

type mp4dump struct {
	full    map[string]struct{}
	showAll bool
	offset  bool
	hex     bool
}

func (m *mp4dump) dumpFile(fpath string) error { _ = "STUB: not implemented"; return nil }

func (m *mp4dump) dump(r io.ReadSeeker) error { _ = "STUB: not implemented"; return nil }

// supported box type

// unsupported box type

func printIndent(w io.Writer, depth int) { _ = "STUB: not implemented"; return }
