package infrastructure

type Detector interface {
	Detect(input string) []Transition
	EndSec(input string) float64
}
