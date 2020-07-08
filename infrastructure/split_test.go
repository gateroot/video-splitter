package infrastructure

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
	"video-splitter/domain"
)

func TestSplit_Split(t *testing.T) {
	ctrl := gomock.NewController(t)
	encoder := NewMockEncoder(ctrl)
	splitter := NewSplit(encoder)

	input := "test.m4v"
	seqs := []domain.Sequence{
		{Start: 10.325, End: 15.253},
		{Start: 20.585, End: 25.952},
	}

	gomock.InOrder(
		encoder.EXPECT().Encode("test.m4v", domain.Sequence{Start: 10.325, End: 15.253}, "test_001.m4v"),
		encoder.EXPECT().Encode("test.m4v", domain.Sequence{Start: 20.585, End: 25.952}, "test_002.m4v"),
	)

	err := splitter.Split(input, seqs)
	assert.NoError(t, err)
}

func Test_insertSuffix(t *testing.T) {
	type args struct {
		path   string
		suffix string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "", args: args{path: "test.m4v", suffix: "_001"}, want: "test_001.m4v"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := insertSuffix(tt.args.path, tt.args.suffix); got != tt.want {
				t.Errorf("insertSuffix() = %v, want %v", got, tt.want)
			}
		})
	}
}
