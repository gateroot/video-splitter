package infrastructure

type FakeDetector struct {
}

func (f FakeDetector) Detect(src string) []Transition {
	return []Transition{
		{Start: 15.0667, End: 20.3667},
		{Start: 25.0442, End: 30.4314},
	}
}

func (f FakeDetector) EndSec(src string) float64 {
	return 35.433333
}

func NewFakeDetector() *FakeDetector {
	return &FakeDetector{}
}
