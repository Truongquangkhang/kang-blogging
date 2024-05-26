package image

import (
	"context"
	"github.com/sirupsen/logrus"
	"kang-blogging/internal/common/decorator"
)

type UploadImageParams struct {
}

type UploadImageResult struct {
}

type UploadImageHandler decorator.UsecaseHandler[UploadImageParams, UploadImageResult]

type uploadImageHandler struct {
}

func NewUploadImageHandler(
	logger *logrus.Entry,
	metrics decorator.MetricsClient,
) UploadImageHandler {
	return decorator.ApplyUsecaseDecorators[UploadImageParams, UploadImageResult](
		uploadImageHandler{},
		logger,
		metrics,
	)
}

func (g uploadImageHandler) Handle(ctx context.Context, param UploadImageParams) (UploadImageResult, error) {
	err := param.Validate()
	if err != nil {
		return UploadImageResult{}, err
	}

	return UploadImageResult{}, err
}

func (p *UploadImageParams) Validate() error {
	return nil
}
