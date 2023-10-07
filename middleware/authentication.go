package middleware

import (
	"context"
	"github.com/golang-jwt/jwt/v5"
	"github.com/samber/do"
	"github.com/t-kuni/go-graphql-template/const/app"
	"github.com/t-kuni/go-graphql-template/domain/model"
	"net/http"
)

type Authentication struct {
	Middleware func(http.Handler) http.Handler
}

func NewAuthentication(i *do.Injector) (*Authentication, error) {
	return &Authentication{
		Middleware: func(next http.Handler) http.Handler {
			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				tokenStr := r.Header.Get("Authorization")

				token, _, err := jwt.NewParser().ParseUnverified(tokenStr, jwt.MapClaims{})
				if err != nil {
					// TODO: エラーログ出力
					http.Error(w, "Unauthorized", http.StatusForbidden)
					return
				}

				var user model.User
				if claims, ok := token.Claims.(jwt.MapClaims); ok {
					user.ID = claims["sub"].(string)
				} else {
					// TODO: エラーログ出力
					http.Error(w, "Unauthorized", http.StatusForbidden)
					return
				}

				ctx := context.WithValue(r.Context(), app.CtxKeyUser, user)

				r = r.WithContext(ctx)
				next.ServeHTTP(w, r)
			})
		},
	}, nil
}
