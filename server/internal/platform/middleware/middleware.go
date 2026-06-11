package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/findardi/Wadi/server/internal/platform/response"
	"github.com/findardi/Wadi/server/internal/platform/token"
)

type TokenVerifier interface {
	VerifyToken(tokenString string) (*token.JwtClaims, error)
}

type ctxKey string

const claimsKey ctxKey = "auth_claims"

type Middleware struct {
	verifier TokenVerifier
}

func New(verifier TokenVerifier) *Middleware {
	return &Middleware{
		verifier: verifier,
	}
}

// RequireAuth
func (m *Middleware) RequireAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Authorization")
		parts := strings.SplitN(header, " ", 2)

		if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") {
			response.Error(w, http.StatusUnauthorized, "missing or invalid authorization", nil)
			return
		}

		claims, err := m.verifier.VerifyToken(parts[1])
		if err != nil {
			response.Error(w, http.StatusUnauthorized, "invalid or expired token", nil)
			return
		}

		ctx := context.WithValue(r.Context(), claimsKey, claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// ClaimsFromContext
func ClaimsFromContext(ctx context.Context) (*token.JwtClaims, bool) {
	claims, ok := ctx.Value(claimsKey).(*token.JwtClaims)
	return claims, ok
}
