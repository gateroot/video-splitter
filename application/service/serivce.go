//go:generate mockgen -source=$GOFILE -destination=../usecase/mock_split_service.go -package=usecase
package service

import "context"

type SplitService interface {
	Split(ctx context.Context, src, dst string) error
}
