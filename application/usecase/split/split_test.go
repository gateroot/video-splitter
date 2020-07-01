package split

import (
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
	"video-splitter/application/usecase"
)

func TestUseCaseHandler_Handle(t *testing.T) {
	setup := func() (*usecase.MockFileReader, *usecase.MockSplitService, usecase.SplitUseCase) {
		ctrl := gomock.NewController(t)
		reader := usecase.NewMockFileReader(ctrl)
		service := usecase.NewMockSplitService(ctrl)
		handler := NewUseCaseHandler(reader, service)
		return reader, service, handler
	}
	src := "/tmp/test.mp4"
	dst := "/tmp/split"

	t.Run("file exists", func(t *testing.T) {
		reader, service, handler := setup()
		exists := reader.EXPECT().Exists(gomock.Any(), src).Times(1).Return(true, nil)
		service.EXPECT().Split(gomock.Any(), src, dst).After(exists).Times(1).Return(nil)

		ctx := context.Background()
		err := handler.Handle(ctx, src, dst)

		assert.NoError(t, err)
	})

	t.Run("file not exist", func(t *testing.T) {
		reader, service, handler := setup()
		reader.EXPECT().Exists(gomock.Any(), src).Times(1).Return(false, nil)
		service.EXPECT().Split(gomock.Any(), src, dst).Times(0).Return(nil)

		ctx := context.Background()
		err := handler.Handle(ctx, src, dst)

		assert.NoError(t, err)
	})

	t.Run("check file exists failed", func(t *testing.T) {
		reader, service, handler := setup()
		reader.EXPECT().Exists(gomock.Any(), src).Times(1).Return(false, errors.New("something wrong"))
		service.EXPECT().Split(gomock.Any(), src, dst).Times(0).Return(nil)

		ctx := context.Background()
		err := handler.Handle(ctx, src, dst)

		assert.Error(t, err)
	})
}
