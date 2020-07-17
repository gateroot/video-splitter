package infrastructure

import (
	"context"
	"github.com/gateroot/video-splitter/application/service/split"
	"github.com/gateroot/video-splitter/domain"
)

type BlackDetector struct {
	detector Detector
}

func (b BlackDetector) Detect(ctx context.Context, src string) (split.Analysis, error) {
	seqs := b.analyze(src)
	return split.Analysis{Sequences: seqs}, nil
}

func (b BlackDetector) analyze(input string) []domain.Sequence {
	transitions := b.detector.Detect(input)
	gen := generator()
	seqs := make([]domain.Sequence, 0)
	var transition Transition
	for _, transition = range transitions {
		seqs = append(seqs, gen(transition))
	}
	endSec := b.detector.EndSec(input)
	seqs = append(seqs, domain.Sequence{
		Start: transition.End,
		End:   endSec,
	})
	return seqs
}

func generator() func(transition Transition) domain.Sequence {
	cur := float64(0)
	return func(transition Transition) domain.Sequence {
		seq := domain.Sequence{
			Start: cur,
			End:   transition.Start,
		}
		cur = transition.End
		return seq
	}
}

func NewBlackDetector(detector Detector) *BlackDetector {
	return &BlackDetector{detector: detector}
}
