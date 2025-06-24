package request

import (
	"context"
	"net/http"

	"github.com/flazhgrowth/fg-tamagotchi/pkg/db/entity"
)

type RequestImpl struct {
	*http.Request

	// HTTP Headers
	generalHeaders        HTTPGeneralHeaders
	securityHeaders       HTTPSecurityHeaders
	contentRelatedHeaders HTTPContentRelatedHeaders

	// User account
	userAccount UserAccount
}

type Request interface {
	// GetContext get request context
	GetContext() context.Context

	// DecodeBody decode body bytes to struct. The argument of this method accepts struct pointer
	DecodeBody(dest any) error

	// DecodeQueryParam decode request param to struct. The argument of this method accepts struct pointer
	DecodeQueryParam(dest any) error

	// GeneralHeaders gets general headers. This includes Host, UserAgent, Accept, AcceptEncoding, Referer, Connection
	GeneralHeaders() HTTPGeneralHeaders

	/*
		SecurityHeaders gets security headers. This includes Authorization, ProxyAuthorization, Cookie. There's IsAuth also.
		This is a helper to decide if Authorization header is available and valid (valid would be the value starts with prefix "Bearer").

		Authorization value also does not include the "Bearer" prefix anymore
	*/
	SecurityHeaders() HTTPSecurityHeaders

	ContentHeaders() HTTPContentRelatedHeaders

	// GetAccountInfo gets account info from the context. Will return apierrors.ErrorUnauthorized() on failed get
	GetAccountInfo() (accountInfo entity.AccountInfo, err error)

	// GetNetHTTPHeaders gets native HTTP Headers
	GetNetHTTPHeaders() http.Header

	// NativeRequest returns the *http.Request native
	NativeRequest() *http.Request

	// URLParam gets url param value based on given key
	URLParam(key string) ParamsValue
}

func New(r *http.Request) Request {
	req := &RequestImpl{
		Request: r,
	}
	return req.
		fetchGeneralHeaders().
		fetchContentRelatedHeaders().
		fetchUserAccountData()
}
