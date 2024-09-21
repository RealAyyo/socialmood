package server

import (
	"net"
	controllers2 "socialmood/api/controllers"
	"socialmood/api/middlewares"
	"socialmood/internal/config"
	"socialmood/internal/jwt"

	"github.com/valyala/fasthttp"

	"github.com/savsgio/atreugo/v11"
)

func NewServer(
	conf *config.Config,
	userController *controllers2.UserController,
	authController *controllers2.AuthController,
) error {
	addr := net.JoinHostPort(conf.Http.Host, conf.Http.Port)
	atreugoConf := atreugo.Config{
		Addr: addr,
	}

	jwtService := jwt.New(&conf.JWT)

	server := atreugo.New(atreugoConf)
	server.UseAfter(func(ctx *atreugo.RequestCtx) error {
		ctx.Response.Header.Set("Access-Control-Allow-Origin", "*")
		ctx.Response.Header.Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, DELETE, PUT")
		ctx.Response.Header.Set("Access-Control-Allow-Headers", "Origin, Content-Type, Accept, Authorization")

		if string(ctx.Method()) == "OPTIONS" {
			ctx.Response.Header.Set("Access-Control-Max-Age", "86400")
			ctx.SetStatusCode(fasthttp.StatusOK)
			return nil
		}
		return ctx.Next()
	})
	server.UseBefore(middlewares.AuthMiddleware(jwtService))
	apiV1 := server.NewGroupPath("/api")

	// AUTH
	authPath := apiV1.NewGroupPath("/auth")
	authPath.POST("/login", authController.Login)

	// user
	userPath := apiV1.NewGroupPath("/user")
	userPath.POST("/register", userController.Register)
	userPath.GET("/get/{id}", userController.GetById)
	userPath.POST("/Search", userController.Search)

	if err := server.ListenAndServe(); err != nil {
		return err
	}

	return nil
}
