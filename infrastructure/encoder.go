//go:generate mockgen -source=$GOFILE -destination=mock_$GOFILE -package=$GOPACKAGE
package infrastructure

import "video-splitter/domain"

type Encoder interface {
	Encode(input string, seq domain.Sequence, output string) error
}
