package response

type BaseResponse struct {
	StatusCode int    `json:"-"`
	Code       string `json:"code"`
	Message    string `json:"message"`
	Data       any    `json:"data"`
	ServerTime int64  `json:"servertime"`

	ErrorID string `json:"error_id,omitempty"`
}
