package split

import (
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
	"github.com/gateroot/video-splitter/application/service"
	"github.com/gateroot/video-splitter/domain"
)

func TestService_Split(t *testing.T) {
	setup := func() (service.SplitService, *MockSceneDetector, *MockSplitter) {
		ctrl := gomock.NewController(t)
		detector := NewMockSceneDetector(ctrl)
		splitter := NewMockSplitter(ctrl)
		s := NewSplitService(detector, splitter)
		return s, detector, splitter
	}

	src := "/tmp"
	dst := "/tmp/split"

	t.Run("ok", func(t *testing.T) {
		s, detector, splitter := setup()
		anl := Analysis{
			Sequences: []domain.Sequence{
				{Start: 0, End: 1000},
			},
		}
		gomock.InOrder(
			detector.EXPECT().Detect(gomock.Any(), src).Times(1).Return(anl, nil),
			splitter.EXPECT().Split(gomock.Any(), src, anl.Sequences, dst).Times(1).Return(nil),
		)

		ctx := context.Background()
		err := s.Split(ctx, src, dst)

		assert.NoError(t, err)
	})

	t.Run("empty analysis", func(t *testing.T) {
		s, detector, splitter := setup()
		anl := Analysis{
			Sequences: []domain.Sequence{},
		}
		gomock.InOrder(
			detector.EXPECT().Detect(gomock.Any(), src).Times(1).Return(anl, nil),
			splitter.EXPECT().Split(gomock.Any(), src, dst, anl).Times(0),
		)

		ctx := context.Background()
		err := s.Split(ctx, src, dst)

		assert.NoError(t, err)
	})

	t.Run("analyze failed", func(t *testing.T) {
		s, detector, splitter := setup()
		gomock.InOrder(
			detector.EXPECT().Detect(gomock.Any(), src).Times(1).Return(Analysis{}, errors.New("something wrong")),
			splitter.EXPECT().Split(gomock.Any(), src, dst, gomock.Any()).Times(0),
		)

		ctx := context.Background()
		err := s.Split(ctx, src, dst)

		assert.Error(t, err)
	})

	t.Run("split failed", func(t *testing.T) {
		s, detector, splitter := setup()
		anl := Analysis{
			Sequences: []domain.Sequence{
				{Start: 0, End: 1000},
			},
		}
		gomock.InOrder(
			detector.EXPECT().Detect(gomock.Any(), src).Times(1).Return(anl, nil),
			splitter.EXPECT().Split(gomock.Any(), src, gomock.Any(), dst).Times(1).Return(errors.New("something wrong")),
		)

		ctx := context.Background()
		err := s.Split(ctx, src, dst)

		assert.Error(t, err)
	})
}
