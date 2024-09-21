package middlewares

import (
	"errors"
	"fmt"
	"socialmood/internal/jwt"
	"strings"

	"github.com/google/uuid"
	"github.com/savsgio/atreugo/v11"
	"github.com/valyala/fasthttp"
)

var (
	ErrAuthorizationRequired = errors.New("authorization required")
)

type AccessToken struct {
	Sub string
	Iat string
}

func AuthMiddleware(jwtService *jwt.JWT) func(ctx *atreugo.RequestCtx) error {
	return func(ctx *atreugo.RequestCtx) error {
		if string(ctx.Method()) == "OPTIONS" {
			fmt.Println(ctx.Path())
			return ctx.Next()
		}

		switch string(ctx.Path()) {
		case "/api/user/register":
			return ctx.Next()
		case "/api/auth/login":
			return ctx.Next()
		}

		authHeader := string(ctx.Request.Header.Peek("Authorization"))
		if len(authHeader) == 0 {
			return ctx.ErrorResponse(ErrAuthorizationRequired, fasthttp.StatusUnauthorized)
		}

		splitAuthHeader := strings.Split(authHeader, " ")

		accessToken := splitAuthHeader[1]

		if splitAuthHeader[0] != "Bearer" || len(accessToken) == 0 {
			return ctx.ErrorResponse(ErrAuthorizationRequired, fasthttp.StatusUnauthorized)
		}

		userID, err := jwtService.ParseToken(accessToken)
		if err != nil {
			return ctx.ErrorResponse(ErrAuthorizationRequired, fasthttp.StatusUnauthorized)
		}

		parsedUserID, err := uuid.Parse(userID)
		if err != nil {
			return ctx.ErrorResponse(ErrAuthorizationRequired, fasthttp.StatusUnauthorized)
		}

		ctx.SetUserValue("user", parsedUserID)
		return ctx.Next()
	}
}
