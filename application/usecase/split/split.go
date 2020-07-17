package split

import (
	"context"
	"fmt"
	"github.com/gateroot/video-splitter/application/service"
	"github.com/gateroot/video-splitter/application/usecase"
)

type UseCaseHandler struct {
	fileChecker  usecase.FileChecker
	splitService service.SplitService
}

func (u UseCaseHandler) Handle(ctx context.Context, src, dst string) error {
	if !u.fileChecker.Exists(src) {
		return fmt.Errorf("input file not exist: %s\n", src)
	}

	if !u.fileChecker.IsDirectory(dst) {
		return fmt.Errorf("dst is not directory: %s\n", dst)
	}

	if err := u.splitService.Split(ctx, src, dst); err != nil {
		return fmt.Errorf("split file failed: %w", err)
	}

	return nil
}

func NewUseCaseHandler(fileChecker usecase.FileChecker, splitService service.SplitService) *UseCaseHandler {
	return &UseCaseHandler{fileChecker: fileChecker, splitService: splitService}
}
