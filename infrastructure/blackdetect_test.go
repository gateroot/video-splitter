package infrastructure

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
	"github.com/gateroot/video-splitter/application/service/split"
	"github.com/gateroot/video-splitter/domain"
)

func TestBlackDetector_Detect(t *testing.T) {
	detector := NewFakeDetector()
	blackDetector := NewBlackDetector(detector)

	ctx := context.Background()
	anl, err := blackDetector.Detect(ctx, "test.mov")

	assert.NoError(t, err)
	want := split.Analysis{
		Sequences: []domain.Sequence{
			{Start: 0, End: 15.0667},
			{Start: 20.3667, End: 25.0442},
			{Start: 30.4314, End: 35.433333},
		},
	}
	assert.Equal(t, want, anl)
}
