package file_service

import (
	"encoding/json"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"mime/multipart"
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
	err := gwmux.HandlePath(
		"POST", "/api/v1/image/upload",
		func(w http.ResponseWriter, req *http.Request, _ map[string]string) {
			err := req.ParseMultipartForm(10 << 20)
			if err != nil {
				return
			}

			// Get handler for filename, size and headers
			file, handler, err := req.FormFile("image")
			if err != nil {
				_, err := fmt.Fprintf(w, "Error Retrieving the File: %v", err)
				if err != nil {
					return
				}
				return
			}

			defer func(file multipart.File) {
				err := file.Close()
				if err != nil {

				}
			}(file)

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
				_, err := fmt.Fprintf(w, "catch an error when upload image to cloudinary: %v", err)
				if err != nil {
					return
				}
			} else {
				response := UploadImageResponse{
					Code:    0,
					Message: "Success",
					Data: &UploadImageResponse_Data{
						Url: *url,
					},
				}
				w.Header().Set("Content-Type", "application/json")
				err := json.NewEncoder(w).Encode(response)
				if err != nil {
					return
				}
			}
		})
	if err != nil {
		fmt.Printf("Error Retrieving the File: %v", err)
		return
	}
}
