package image

import (
	"golang.org/x/net/context"
	"kang-blogging/internal/blogging/infra/genproto/blogging"
)

func (g GrpcService) UploadImage(
	ctx context.Context,
	request *blogging.UploadImageRequest,
) (*blogging.UploadImageResponse, error) {
	return &blogging.UploadImageResponse{
		Code:    0,
		Message: "Success",
		Data: &blogging.UploadImageResponse_Data{
			Url: "url_receive_from_cloudinary",
		},
	}, nil
}
