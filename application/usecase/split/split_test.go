package split

import (
	"context"
	"testing"
)

func UseCaseHandler_Test(t *testing.T) {
	reader := NewMockFileReader()
	service := NewMockSplitService()
	handler := NewUseCaseHandler(reader, service)

	ctx := context.Background()
	src := "/tmp/test.mp4"
	dst := "/tmp/split"
	handler.Handle(ctx, src, dst)


}
