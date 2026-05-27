package divide

import (
	"os"
)

const (
	videoDirName     = "video"
	audioDirName     = "audio"
	encVideoDirName  = "video_enc"
	encAudioDirName  = "audio_enc"
	initMP4FileName  = "init.mp4"
	playlistFileName = "playlist.m3u8"
)

func segmentFileName(i int) string { _ = "STUB: not implemented"; return "" }

func Main(args []string) int { _ = "STUB: not implemented"; return 0 }

type childInfo map[uint32]uint64

type segment struct {
	duration float64
}

type trackType int

const (
	trackVideo trackType = iota
	trackAudio
	trackEncVideo
	trackEncAudio
)

type track struct {
	id          uint32
	trackType   trackType
	timescale   uint32
	bandwidth   uint64
	height      uint16
	width       uint16
	segments    []segment
	outputDir   string
	initFile    *os.File
	segmentFile *os.File
}

func divide(inputFilePath, outputDir string) error { _ = "STUB: not implemented"; return nil }

// generate track map

// get trackID from Tkhd box

// get timescale from Mdhd box

// initialization segment

// get trackID from Tkhd box

// already writeAll is true in Moov box

// copy all data of payload

// rewrite headers

// media segment

// extract Tfdt-box

// extract Trun-box

// close last segment file and create next

// Mdat box

// skip

func outputMasterPlaylist(filePath string, trackTypeMap map[trackType]*track) error {
	_ = "STUB: not implemented"
	return nil
}

// FIXME: hard coding

func outputMediaPlaylist(filePath string, segments []segment) error {
	_ = "STUB: not implemented"
	return nil
}
