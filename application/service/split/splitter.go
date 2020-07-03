//go:generate mockgen -source=$GOFILE -destination=mock_$GOFILE -package=$GOPACKAGE
package split

import "context"

type Splitter interface {
	Split(ctx context.Context, src, dst string, analysis Analysis) error
}
