//go:generate mockgen -source=$GOFILE -destination=mock_$GOFILE -package=$GOPACKAGE
package split

import "context"

type SceneDetector interface {
	Detect(ctx context.Context, src string) (Analysis, error)
}
