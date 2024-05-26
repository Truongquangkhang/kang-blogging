package file_service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"kang-blogging/internal/common/utils"
	"mime/multipart"
	"net/http"
	"strconv"
)

type Cloudinary struct {
	Url       string
	CloudName string
	ApiKey    string
	ApiSecret string
}

type UploadImageWithEagerResponse struct {
	Eager []struct {
		Url string `json:"url"`
	} `json:"eager"`
	Url string `json:"url"`
}

const PATH_UPLOAD_IMAGE = "/image/upload"

func (c *Cloudinary) UploadImage(file multipart.File, fileName string, eager string, timestamp int64) (*string, error) {
	params := map[string]string{
		"api_key":   c.ApiKey,
		"eager":     eager,
		"public_id": fileName,
		"timestamp": strconv.FormatInt(timestamp, 10),
		"signature": c.genAuthorization(fileName, eager, timestamp),
	}
	var b bytes.Buffer
	writer := multipart.NewWriter(&b)
	for key, val := range params {
		_ = writer.WriteField(key, val)
	}
	part, err := writer.CreateFormFile("file", fileName)
	if err != nil {
		return nil, fmt.Errorf("error writing to form file: %v", err)
	}

	_, err = io.Copy(part, file)
	if err != nil {
		return nil, fmt.Errorf("error copy file: %v", err)
	}
	writer.Close()

	uploadURL := fmt.Sprintf("%s%s%s", c.Url, c.CloudName, PATH_UPLOAD_IMAGE)
	request, err := http.NewRequest("POST", uploadURL, &b)
	if err != nil {
		return nil, fmt.Errorf("error creating upload request: %v", err)
	}

	request.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return nil, fmt.Errorf("error sending upload request: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Error Reading Response: %v", err)
	}

	var result UploadImageWithEagerResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	if len(result.Eager) > 0 {
		return &result.Eager[0].Url, nil
	}
	return &result.Url, nil
}

func (c *Cloudinary) genAuthorization(fileName string, eager string, timestamp int64) string {
	signature := fmt.Sprintf(
		"eager=%s&public_id=%s&timestamp=%s%s", eager, fileName, strconv.FormatInt(timestamp, 10), c.ApiSecret,
	)
	return utils.EncodeSHA1(signature)
}
