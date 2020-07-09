package infrastructure

import (
	"context"
	"fmt"
	"path/filepath"
	"video-splitter/domain"
)

type Split struct {
	encoder Encoder
}

func (s Split) Split(ctx context.Context, input string, seqs []domain.Sequence, outputDir string) error {
	for i, seq := range seqs {
		suffix := fmt.Sprintf("_%03d", i+1)
		if err := s.encoder.Encode(input, seq, outputDir+"/"+insertSuffix(input, suffix)); err != nil {
			return err
		}
	}
	return nil
}

func insertSuffix(path, suffix string) string {
	ext := filepath.Ext(path)
	return path[:len(path)-len(ext)] + suffix + ext
}

func NewSplit(encoder Encoder) *Split {
	return &Split{encoder: encoder}
}
