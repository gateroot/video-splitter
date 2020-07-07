package detector

import (
	"bufio"
	"os/exec"
	"regexp"
	"strconv"
	"video-splitter/infrastructure"
)

type FFmpegDetector struct {
}

func (d FFmpegDetector) Detect(input string) []infrastructure.Transition {
	cmd := exec.Command("ffmpeg", "-i", input, "-vf", "blackdetect=0.5:.95:.1, lutrgb=r=negval:g=negval:b=negval, blackdetect=0.5:.95:.1", "-f", "null", "-")

	stderr, _ := cmd.StderrPipe()
	defer stderr.Close()

	_ = cmd.Start()

	s := bufio.NewScanner(stderr)
	var end float64
	transitions := make([]infrastructure.Transition, 0)
	for s.Scan() {
		line := s.Text()
		r, _ := regexp.Compile("black_start:(.*) black_end:(.*) black_duration:(.*)")
		if !r.MatchString(line) {
			continue
		}
		group := r.FindStringSubmatch(line)
		start, _ := strconv.ParseFloat(group[1], 64)
		end, _ = strconv.ParseFloat(group[2], 64)
		transitions = append(transitions, infrastructure.Transition{
			Start: start,
			End:   end,
		})
	}
	return transitions
}

func (d FFmpegDetector) EndSec(input string) float64 {
	cmd := exec.Command("ffprobe", "-v", "error", "-show_entries", "format=duration", "-of", "default=noprint_wrappers=1:nokey=1", input)
	stdout, _ := cmd.StdoutPipe()
	defer stdout.Close()

	_ = cmd.Start()

	s := bufio.NewScanner(stdout)
	var endSec float64
	for s.Scan() {
		line := s.Text()
		f, err := strconv.ParseFloat(line, 64)
		if err != nil {
			panic(err)
		}
		endSec = f
	}
	return endSec
}

func NewFFmpegDetector() *FFmpegDetector {
	return &FFmpegDetector{}
}
