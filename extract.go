package mp4

import (
	"io"
)

type BoxInfoWithPayload struct {
	Info    BoxInfo
	Payload IBox
}

func ExtractBoxWithPayload(r io.ReadSeeker, parent *BoxInfo, path BoxPath) ([]*BoxInfoWithPayload, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ExtractBoxesWithPayload(r io.ReadSeeker, parent *BoxInfo, paths []BoxPath) ([]*BoxInfoWithPayload, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ExtractBox(r io.ReadSeeker, parent *BoxInfo, path BoxPath) ([]*BoxInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ExtractBoxes(r io.ReadSeeker, parent *BoxInfo, paths []BoxPath) ([]*BoxInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func matchPath(paths []BoxPath, path BoxPath) (forwardMatch bool, match bool) {
	_ = "STUB: not implemented"
	return false, false
}
