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
	e.Use(middlewares.FilterRequests)

	e.GET("/ping", controllers.Ping)

	// auth 路由
	auth := e.Group("/api/auth")
	auth.POST("/login", controllers.PostLogin)
	auth.GET("/login/sms/code", controllers.GetLoginSMSCode)
	auth.POST("/role", controllers.PostCreateRole)

	// member 路由
	member := e.Group("/api/member", middlewares.Auth)
	member.POST("", controllers.PostCreateMember)

	// bisale 业务路由
	bisale := e.Group("/api/bisale", middlewares.Auth)
	bisale.GET("/users", controllers.GetBisaleUsers)

	e.Logger.Fatal(e.Start(config.GetListenNetAddress()))
}
