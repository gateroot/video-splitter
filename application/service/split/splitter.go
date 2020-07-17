//go:generate mockgen -source=$GOFILE -destination=mock_$GOFILE -package=$GOPACKAGE
package split

import (
	"context"
	"github.com/gateroot/video-splitter/domain"
)

type Splitter interface {
	Split(ctx context.Context, input string, seqs []domain.Sequence, outputDir string) error
}
