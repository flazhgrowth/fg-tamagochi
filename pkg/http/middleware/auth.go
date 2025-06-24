package middleware

import (
	"context"
	"net/http"
	"time"

	"github.com/flazhgrowth/fg-gotools/jwt"
	"github.com/flazhgrowth/fg-tamagotchi/constant"
	"github.com/flazhgrowth/fg-tamagotchi/pkg/db/entity"
	"github.com/flazhgrowth/fg-tamagotchi/pkg/http/apierrors"
	"github.com/flazhgrowth/fg-tamagotchi/pkg/http/request"
	"github.com/flazhgrowth/fg-tamagotchi/pkg/http/response"
	"github.com/flazhgrowth/fg-tamagotchi/pkg/vault"
	"github.com/rs/zerolog/log"
)

// BasicBearerAuthMiddleware
/*
	BasicBearerAuthMiddleware only checks whether a route is supposedly guarded with auth or not.
	Further checks and validations for this required to be implemented on another middleware
*/
func BasicBearerAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		req := request.New(r)
		resp := response.New(w)
		secHeaders := req.SecurityHeaders()
		if !secHeaders.IsAuth {
			resp.Respond(nil, apierrors.ErrorUnauthorized())
			return
		}

		token := jwt.NewJWT()
		claims, err := token.ValidateToken(secHeaders.Authorization, vault.GetVault().GetStringWithDefault("tokens.secret", ""))
		if err != nil {
			log.Error().Msgf("auth-middleware: %s", err.Error())
			resp.Respond(nil, apierrors.ErrorUnauthorized(jwt.ErrInvalidToken.Error()))
			return
		}
		if claims.ExpiresAt.Before(time.Now()) || claims.ID == "" {
			log.Error().Msgf("auth-middleware: claims.ExpiresAt.Before(time.Now()) = %v", claims.ExpiresAt.Before(time.Now()))
			resp.Respond(nil, apierrors.ErrorUnauthorized(jwt.ErrInvalidToken.Error()))
			return
		}
		ctx := r.Context()
		ctx = context.WithValue(ctx, constant.CtxKeyAccountInfo, entity.AccountInfo{
			ID:       claims.ID,
			Username: claims.Username,
			Email:    claims.Email,
		})

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
