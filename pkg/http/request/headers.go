package request

import (
	"net/http"
	"strings"
)

func (req *RequestImpl) GeneralHeaders() HTTPGeneralHeaders {
	return req.generalHeaders
}

func (req *RequestImpl) SecurityHeaders() HTTPSecurityHeaders {
	return req.fetchSecurityHeaders().securityHeaders
}

func (req *RequestImpl) ContentHeaders() HTTPContentRelatedHeaders {
	return req.contentRelatedHeaders
}

func (req *RequestImpl) GetNetHTTPHeaders() http.Header {
	return req.Header
}

func (req *RequestImpl) fetchGeneralHeaders() *RequestImpl {
	req.generalHeaders = HTTPGeneralHeaders{
		Host:           req.Host,
		UserAgent:      req.UserAgent(),
		Accept:         req.Header.Get("accept"),
		AcceptEncoding: req.Header.Get("accept-encoding"),
		RemoteAddr:     req.RemoteAddr,
		RequestURI:     req.RequestURI,
		Referer:        req.Referer(),
		Connection:     req.Header.Get("connection"),
	}

	return req
}

func (req *RequestImpl) fetchSecurityHeaders() *RequestImpl {
	authHeader := req.Header.Get("authorization")
	req.securityHeaders = HTTPSecurityHeaders{
		ProxyAuthorization: req.Header.Get("proxy-authorization"),
		Cookie:             req.Header.Get("cookie"),
	}
	req.securityHeaders.IsAuth = (authHeader != "" && strings.HasPrefix(authHeader, "Bearer"))
	if req.securityHeaders.IsAuth {
		req.securityHeaders.Authorization = strings.ReplaceAll(authHeader, "Bearer ", "")
	}

	return req
}

func (req *RequestImpl) fetchContentRelatedHeaders() *RequestImpl {
	req.contentRelatedHeaders = HTTPContentRelatedHeaders{
		ContentType:   req.Header.Get("content-type"),
		ContentLength: req.ContentLength,
	}

	return req
}
