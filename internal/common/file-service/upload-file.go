package file_service

import (
	"encoding/json"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"kang-blogging/internal/common/server/httperr"
	"net/http"
	"os"
	"time"
)

type UploadImageResponse_Data struct {
	Url string `json:"url"`
}

// UploadImageResponse struct for the response
type UploadImageResponse struct {
	Code    int                       `json:"code"`
	Message string                    `json:"message"`
	Data    *UploadImageResponse_Data `json:"data"`
}

func RegisterFileServiceHandler(gwmux *runtime.ServeMux) {
	gwmux.HandlePath(
		"POST", "/api/v1/image/upload",
		func(w http.ResponseWriter, req *http.Request, _ map[string]string) {
			req.ParseMultipartForm(10 << 20)

			// Get handler for filename, size and headers
			file, handler, err := req.FormFile("image")
			if err != nil {
				fmt.Fprintf(w, "Error Retrieving the File: %v", err)
				return
			}
			defer file.Close()

			cloudinary := Cloudinary{
				Url:       os.Getenv("CLOUDINARY_URL"),
				CloudName: os.Getenv("CLOUDINARY_CLOUD_NAME"),
				ApiKey:    os.Getenv("CLOUDINARY_API_KEY"),
				ApiSecret: os.Getenv("CLOUDINARY_API_SECRET"),
			}
			url, err := cloudinary.UploadImage(
				file, handler.Filename, os.Getenv("CLOUDINARY_EAGER"), time.Now().Unix(),
			)
			if err != nil {
				httperr.RespondWithError(err, w, req)
			} else {
				response := UploadImageResponse{
					Code:    0,
					Message: "Success",
					Data: &UploadImageResponse_Data{
						Url: *url,
					},
				}
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(response)
			}
		})
}
