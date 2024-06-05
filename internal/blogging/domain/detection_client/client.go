package detection_client

import "golang.org/x/net/context"

type IClientAdapter interface {
	DetectToxicComment(
		ctx context.Context,
		comment string,
	) (*PredictionComment, error)
}

type PredictionComment struct {
	IsToxicComment         bool
	Comment                string
	ToxicPredictionComment []int32
}
