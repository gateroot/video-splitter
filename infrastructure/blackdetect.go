package infrastructure

import (
	"bufio"
	"context"
	"os/exec"
	"regexp"
	"strconv"
	"video-splitter/application/service/split"
	"video-splitter/domain"
)

type BlackDetector struct {
}

func (b BlackDetector) Detect(ctx context.Context, src string) (split.Analysis, error) {
	seqs := analyze(src)
	return split.Analysis{Sequences: seqs}, nil
}

func analyze(input string) []domain.Sequence {
	cmd := exec.Command("ffmpeg", "-i", input, "-vf", "blackdetect=0.5:.95:.1, lutrgb=r=negval:g=negval:b=negval, blackdetect=0.5:.95:.1", "-f", "null", "-")

	stderr, _ := cmd.StderrPipe()
	defer stderr.Close()

	_ = cmd.Start()

	s := bufio.NewScanner(stderr)
	gen := generator()
	var bend float64
	seqs := make([]domain.Sequence, 0)
	for s.Scan() {
		line := s.Text()
		r, _ := regexp.Compile("black_start:(.*) black_end:(.*) black_duration:5.3")
		if !r.MatchString(line) {
			continue
		}
		group := r.FindStringSubmatch(line)
		bstart, _ := strconv.ParseFloat(group[1], 64)
		bend, _ = strconv.ParseFloat(group[2], 64)
		seqs = append(seqs, gen(bstart, bend))
	}

	return seqs
}

func generator() func(start, end float64) domain.Sequence {
	cur := float64(0)
	return func(start, end float64) domain.Sequence {
		seq := domain.Sequence{
			Start: cur,
			End:   end,
		}
		cur = end
		return seq
	}
}

func NewBlackDetector() *BlackDetector {
	return &BlackDetector{}
}
