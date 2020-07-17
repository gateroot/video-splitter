package encoder

import (
	"github.com/gateroot/video-splitter/domain"
	"os/exec"
	"strconv"
)

type FFmpegEncoder struct {
}

func (f FFmpegEncoder) Encode(input string, seq domain.Sequence, output string) error {
	start := strconv.FormatFloat(seq.Start, 'f', -1, 64)
	end := strconv.FormatFloat(seq.End, 'f', -1, 64)
	cmd := exec.Command("ffmpeg", "-i", input, "-ss", start, "-to", end, "-vcodec copy", "-acodec copy", output)

	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}

func NewFFmpegEncoder() *FFmpegEncoder {
	return &FFmpegEncoder{}
}
