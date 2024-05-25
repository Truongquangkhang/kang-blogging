package image

import (
	"bytes"
	"golang.org/x/net/context"
	"image"
	"image/jpeg"
	"kang-blogging/internal/blogging/infra/genproto/blogging"
	"log"
	"os"
)

func (g GrpcService) UploadImage(
	ctx context.Context,
	request *blogging.UploadImageRequest,
) (*blogging.UploadImageResponse, error) {
	serveFrames(request.Image)
	return &blogging.UploadImageResponse{
		Code:    0,
		Message: "Success",
		Data: &blogging.UploadImageResponse_Data{
			Url: "khang",
		},
	}, nil
}

func serveFrames(imgByte []byte) {

	img, _, err := image.Decode(bytes.NewReader(imgByte))
	if err != nil {
		log.Fatalln(err)
	}

	out, _ := os.Create("img.jpeg")
	defer out.Close()

	var opts jpeg.Options
	opts.Quality = 1

	err = jpeg.Encode(out, img, &opts)
	//jpeg.Encode(out, img, nil)
	if err != nil {
		log.Println(err)
	}

}
