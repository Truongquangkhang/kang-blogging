package toxicity_detection_client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"golang.org/x/net/context"
	"io"
	"kang-blogging/internal/blogging/domain/detection_client"
	"kang-blogging/internal/common/server/httpheader"
	"net/http"
	"os"
)

type ToxicityDetectionClient struct {
	URL string
}

func NewToxicityDetectionClient() *ToxicityDetectionClient {
	url := os.Getenv("TOXICITY_DETECTION_URL")
	return &ToxicityDetectionClient{
		URL: url,
	}
}

type DetectContentResponse struct {
	Text        string  `json:"text"`
	Predictions []int32 `json:"predictions"`
}

const (
	DETECT_CONTENT = "/api/v1/detect-content/"
)

//func (c *ToxicityDetectionClient) DetectToxicComment(
//	ctx context.Context,
//	comment string,
//) (*detection_client.PredictionComment, error) {
//	countLetter := strings.Count(comment, " ") + 1
//	prediction := make([]int32, countLetter)
//	isToxicComment := false
//	for index, c := range strings.Split(comment, " ") {
//		if c == "ngu" {
//			isToxicComment = true
//			prediction[index] = 1
//		}
//	}
//	return &detection_client.PredictionComment{
//		IsToxicComment:         isToxicComment,
//		Comment:                comment,
//		ToxicPredictionComment: prediction,
//	}, nil
//}

func (c *ToxicityDetectionClient) DetectToxicComment(
	ctx context.Context,
	comment string,
) (*detection_client.PredictionComment, error) {
	body := []byte(fmt.Sprintf(`{
		"content": "%s"
	}`, comment))
	response, err := http.Post(
		fmt.Sprintf("%s%s", c.URL, DETECT_CONTENT),
		httpheader.CONTENT_TYPE_APPLICATION_JSON, bytes.NewBuffer(body),
	)
	if err != nil {
		return nil, err
	}

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}
	var detectContentResponse DetectContentResponse
	err = json.Unmarshal(responseData, &detectContentResponse)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response body: %v", err)
	}
	isToxicComment := false
	for _, c := range detectContentResponse.Predictions {
		if c == 1 {
			isToxicComment = true
			break
		}
	}
	return &detection_client.PredictionComment{
		IsToxicComment:         isToxicComment,
		Comment:                detectContentResponse.Text,
		ToxicPredictionComment: detectContentResponse.Predictions,
	}, nil
}
