package di

import (
	"video-splitter/application/service/split"
	"video-splitter/application/usecase"
	split2 "video-splitter/application/usecase/split"
	"video-splitter/infrastructure"
	detector2 "video-splitter/infrastructure/detector"
	encoder2 "video-splitter/infrastructure/encoder"
)

func InjectSplitUseCase() usecase.SplitUseCase {
	// infrastructure
	encoder := encoder2.NewFFmpegEncoder()
	detector := detector2.NewFFmpegDetector()
	blackDetector := infrastructure.NewBlackDetector(detector)
	splitter := infrastructure.NewSplit(encoder)
	checker := infrastructure.NewFileChecker()

	// service
	service := split.NewSplitService(blackDetector, splitter)

	// usecase
	handler := split2.NewUseCaseHandler(checker, service)

	return handler
}
