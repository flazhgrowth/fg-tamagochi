package middleware

import (
	"fmt"
	"net/http"

	"github.com/flazhgrowth/fg-tamagochi/pkg/http/apierrors"
	"github.com/flazhgrowth/fg-tamagochi/pkg/http/response"
	"github.com/flazhgrowth/fg-tamagochi/pkg/vault"
	"github.com/flazhgrowth/fg-tamagopkg/hash/sha256"
)

// BasicAPIKeyMiddleware adds a security check, by checking header value of X-API-Key
/*
	This middleware simply comparing the value of sha256(x-api-key) == secret.apikey
*/
func BasicAPIKeyMiddleware(key string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			apikey := r.Header.Get("X-API-Key")
			resp := response.New(w)
			secret := vault.GetVault().GetStringWithDefault(fmt.Sprintf("secret.apikey.%s", key), "")
			if secret == "" {
				resp.RespondJSON(nil, apierrors.ErrorBadRequest("apikey is not set").WithCode("invalid_api_key"))
				return
			}

			if apikey == "" {
				resp.RespondJSON(nil, apierrors.ErrorUnauthorized("invalid api key").WithCode("invalid_api_key"))
				return
			}
			hashedApikey := sha256.Hash(apikey, sha256.HexEncoder)
			if hashedApikey != secret {
				resp.RespondJSON(nil, apierrors.ErrorUnauthorized("invalid api key").WithCode("invalid_api_key"))
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
