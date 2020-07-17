package infrastructure

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
	"github.com/gateroot/video-splitter/domain"
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
	dst := "hoge"

	gomock.InOrder(
		encoder.EXPECT().Encode("test.m4v", domain.Sequence{Start: 10.325, End: 15.253}, "hoge/test_001.m4v"),
		encoder.EXPECT().Encode("test.m4v", domain.Sequence{Start: 20.585, End: 25.952}, "hoge/test_002.m4v"),
	)

	ctx := context.Background()
	err := splitter.Split(ctx, input, seqs, dst)
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
