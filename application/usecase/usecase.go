package usecase

import "context"

type SplitUseCase interface {
	Handle(ctx context.Context, src, dst string) error
}

type SplitAutoUseCase interface {
	Handle(ctx context.Context, srcDir, dstDir string) error
}
