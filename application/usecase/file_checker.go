//go:generate mockgen -source=$GOFILE -destination=mock_$GOFILE -package=$GOPACKAGE
package usecase

type FileChecker interface {
	Exists(filename string) bool
	IsDirectory(filename string) bool
}
