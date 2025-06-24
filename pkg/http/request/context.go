package request

import (
	"context"
	"net/http"

	"github.com/flazhgrowth/fg-tamagochi/constant"
	"github.com/flazhgrowth/fg-tamagochi/pkg/db/entity"
	"github.com/flazhgrowth/fg-tamagochi/pkg/http/apierrors"
)

func (req *RequestImpl) GetContext() context.Context {
	return req.Context()
}

func (req *RequestImpl) NativeRequest() *http.Request {
	return req.Request
}

func (req *RequestImpl) GetAccountInfo() (accountInfo entity.AccountInfo, err error) {
	ctx := req.GetContext()
	accountInfo, ok := ctx.Value(constant.CtxKeyAccountInfo).(entity.AccountInfo)
	if !ok {
		return accountInfo, apierrors.ErrorUnauthorized()
	}

	return accountInfo, nil
}
