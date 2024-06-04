package toxicity_detection_client

import (
	"golang.org/x/net/context"
	"kang-blogging/internal/blogging/domain/detection_client"
	"os"
	"strings"
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

func (c *ToxicityDetectionClient) DetectToxicComment(
	ctx context.Context,
	comment string,
) (*detection_client.PredictionComment, error) {
	countLetter := strings.Count(comment, " ") + 1
	prediction := make([]int32, countLetter)
	isToxicComment := false
	for index, c := range strings.Split(comment, " ") {
		if c == "ngu" {
			isToxicComment = true
			prediction[index] = 1
		}
	}
	return &detection_client.PredictionComment{
		IsToxicComment:         isToxicComment,
		Comment:                comment,
		ToxicPredictionComment: prediction,
	}, nil
}
