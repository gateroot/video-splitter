//go:generate mockgen -source=$GOFILE -destination=mock_$GOFILE -package=$GOPACKAGE
package usecase

import "context"

type FileReader interface {
	Exists(ctx context.Context, path string) (bool, error)
}
