package httpresp

type BaseResponse struct {
	Code    int32  `json:"code"`
	Message string `json:"message"`
}
