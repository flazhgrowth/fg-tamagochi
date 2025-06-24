package response

type BaseResponse struct {
	StatusCode int    `json:"-"`
	Code       string `json:"code"`
	Message    string `json:"message"`
	Data       any    `json:"data"`
}
