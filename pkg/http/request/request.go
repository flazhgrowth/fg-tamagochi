package request

import (
	"context"
	"net/http"

	"github.com/flazhgrowth/fg-tamagochi/pkg/db/entity"
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
	/*
		Any error returned from this method will return apierrors.HTTPError type, in which the struct
		already implement the Error() string method, so it already fulfilled error interface.

		This also indicates that there's no need to return apierrors on api layer
	*/
	DecodeBody(dest any) error

	// DecodeQueryParam decode request param to struct. The argument of this method accepts struct pointer
	/*
		Any error returned from this method will return apierrors.HTTPError type, in which the struct
		already implement the Error() string method, so it already fulfilled error interface.

		This also indicates that there's no need to return apierrors on api layer
	*/
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

	// DecodeURLParam decode url param into struct with tag of path and pathtype for its type
	/*
		Any error returned from this method will return apierrors.HTTPError type, in which the struct
		already implement the Error() string method, so it already fulfilled error interface.

		This also indicates that there's no need to return apierrors on api layer.

		Please note that, it is still recommended to use URLParam method instead of using DecodeURLParam, due to efficiency.
	*/
	DecodeURLParam(dest any) error
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
