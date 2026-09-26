package response

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/flazhgrowth/fg-tamagochi/pkg/http/apierrors"
)

func (resp *ResponseImpl) RespondJSON(data any, err error, statusCode ...int) {
	baseResp := BaseResponse{
		StatusCode: http.StatusOK,
		Code:       "success",
		Message:    "Success",
		Data:       data,
		ServerTime: time.Now().Unix(),
	}
	if len(statusCode) > 0 {
		baseResp.StatusCode = statusCode[0]
	}
	if err != nil {
		ogError, ok := err.(apierrors.HTTPError)
		if !ok {
			ogError = apierrors.ErrorInternalServerError()
		}

		baseResp.Code = ogError.Code
		baseResp.Message = ogError.Message
		baseResp.StatusCode = int(ogError.StatusCode)
		baseResp.Data = nil
		baseResp.ErrorID = ogError.ID
	}

	resp.
		writeBasicHeader().
		write(baseResp)
}

func (resp *ResponseImpl) writeBasicHeader() *ResponseImpl {
	resp.Header().Set("Content-Type", "application/json")
	return resp
}

func (resp *ResponseImpl) write(baseResp BaseResponse) {
	resp.WriteHeader(baseResp.StatusCode)
	respData, _ := json.MarshalIndent(baseResp, "", "    ")
	resp.Write(respData)
}
