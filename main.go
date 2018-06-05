package main

import (
	"github.com/labstack/echo"
	"bisale/bisale-console-api/config"
	"github.com/labstack/echo/middleware"
	"bisale/bisale-console-api/controllers"
	"bisale/bisale-console-api/middlewares"
	"github.com/sirupsen/logrus"
	"github.com/go-playground/validator"
)

type Validator struct {
	validator *validator.Validate
}

func (v *Validator) Validate(i interface{}) error {
	return v.validator.Struct(i)
}

func main() {

	e := echo.New()
	e.Validator = &Validator{validator: validator.New()}
	e.Logger = middlewares.LogrusLogger{logrus.StandardLogger()}
	e.Use(middlewares.LogrusHook())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.Use(middlewares.RequestHash)
	e.Use(middlewares.FilterIteratedRequests)

	e.GET("/ping", controllers.Ping)

	// auth 路由
	auth := e.Group("/api/auth")
	auth.POST("/login", controllers.PostLogin)
	auth.POST("/login/sms/code", controllers.PostLoginSMSCode)
	auth.POST("/role", controllers.PostCreateRole)

	// member 路由
	member := e.Group("/api/member", middlewares.Auth)
	member.POST("", controllers.PostCreateMember)

	// bisale 业务路由
	bisale := e.Group("/api/bisale", middlewares.Auth)

	bisale.GET("/cert/list", controllers.GetCertList)
	bisale.GET("/cert/list/count", controllers.GetCertListCount)
	bisale.GET("/cert/detail", controllers.GetCertDetailById)
	bisale.POST("/cert/result", controllers.PostCertResult, middlewares.FilterRequestsStrict)

	bisale.GET("/withdraw/list", controllers.GetWithdrawList)

	e.Logger.Fatal(e.Start(config.GetListenNetAddress()))
}
