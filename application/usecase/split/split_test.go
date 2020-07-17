package split

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
	"video-splitter/application/usecase"
)

func TestUseCaseHandler_Handle(t *testing.T) {
	setup := func() (*usecase.MockFileChecker, *usecase.MockSplitService, usecase.SplitUseCase) {
		ctrl := gomock.NewController(t)
		checker := usecase.NewMockFileChecker(ctrl)
		service := usecase.NewMockSplitService(ctrl)
		handler := NewUseCaseHandler(checker, service)
		return checker, service, handler
	}
	src := "/tmp/test.mp4"
	dst := "/tmp/split"

	t.Run("file exists", func(t *testing.T) {
		reader, service, handler := setup()
		gomock.InOrder(
			reader.EXPECT().Exists(src).Times(1).Return(true),
			reader.EXPECT().IsDirectory(dst).Times(1).Return(true),
			service.EXPECT().Split(gomock.Any(), src, dst).Times(1).Return(nil),
		)

		ctx := context.Background()
		err := handler.Handle(ctx, src, dst)

		assert.NoError(t, err)
	})

	t.Run("file not exist", func(t *testing.T) {
		reader, service, handler := setup()
		gomock.InOrder(
			reader.EXPECT().Exists(src).Times(1).Return(false),
			reader.EXPECT().IsDirectory(dst).Times(0),
			service.EXPECT().Split(gomock.Any(), src, dst).Times(0),
		)

		ctx := context.Background()
		err := handler.Handle(ctx, src, dst)

		assert.Error(t, err)
	})

	t.Run("dst is not directory", func(t *testing.T) {
		reader, service, handler := setup()
		gomock.InOrder(
			reader.EXPECT().Exists(src).Times(1).Return(true),
			reader.EXPECT().IsDirectory(dst).Times(1).Return(false),
			service.EXPECT().Split(gomock.Any(), src, dst).Times(0),
		)

		ctx := context.Background()
		err := handler.Handle(ctx, src, dst)

		assert.Error(t, err)
	})
}
