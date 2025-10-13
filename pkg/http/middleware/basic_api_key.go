package middleware

import (
	"net/http"

	"github.com/flazhgrowth/fg-gotools/hash/sha256"
	"github.com/flazhgrowth/fg-tamagochi/pkg/http/apierrors"
	"github.com/flazhgrowth/fg-tamagochi/pkg/http/response"
	"github.com/flazhgrowth/fg-tamagochi/pkg/vault"
)

// BasicAPIKeyMiddleware adds a security check, by checking header value of X-API-Key
/*
	This middleware simply comparing the value of sha256(x-api-key) == secret.apikey
*/
func BasicAPIKeyMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apikey := w.Header().Get("X-API-Key")
		resp := response.New(w)
		secret := vault.GetVault().GetStringWithDefault("secret.apikey", "")
		if secret == "" {
			resp.Respond(nil, apierrors.ErrorBadRequest("apikey is not set"))
			return
		}

		if apikey == "" {
			resp.Respond(nil, apierrors.ErrorUnauthorized("invalid api key"))
			return
		}
		hashedApikey := sha256.Hash(apikey)
		if hashedApikey != secret {
			resp.Respond(nil, apierrors.ErrorUnauthorized("invalid api key"))
			return
		}

		next.ServeHTTP(w, r)
	})
}
