package split

import (
	"context"
	"fmt"
	"video-splitter/application/service"
	"video-splitter/application/usecase"
)

type UseCaseHandler struct {
	fileReader   usecase.FileReader
	splitService service.SplitService
}

func (u UseCaseHandler) Handle(ctx context.Context, src, dst string) error {
	b, err := u.fileReader.Exists(ctx, src)
	if err != nil {
		return fmt.Errorf("check file exists failed: %w", err)
	}

	if b {
		u.splitService.Split(ctx, src, dst)
	}

	return nil
}

func NewUseCaseHandler(fileReader usecase.FileReader, splitService service.SplitService) *UseCaseHandler {
	return &UseCaseHandler{fileReader: fileReader, splitService: splitService}
}
