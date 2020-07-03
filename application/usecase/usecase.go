package usecase

import "context"

type SplitUseCase interface {
	Handle(ctx context.Context, src, dst string) error
}
