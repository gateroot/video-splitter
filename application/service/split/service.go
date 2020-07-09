package split

import (
	"context"
	"fmt"
)

type Service struct {
	detector SceneDetector
	splitter Splitter
}

func (s Service) Split(ctx context.Context, src, dst string) error {
	ana, err := s.detector.Detect(ctx, src)
	if err != nil {
		return fmt.Errorf("scene detect failed: %w", err)
	}

	if ana.Sequences == nil || len(ana.Sequences) == 0 {
		return nil
	}

	err = s.splitter.Split(ctx, src, ana.Sequences, dst)
	if err != nil {
		return fmt.Errorf("split failed: %w", err)
	}

	return nil
}

func NewSplitService(detector SceneDetector, splitter Splitter) *Service {
	return &Service{detector: detector, splitter: splitter}
}
