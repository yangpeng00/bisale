package main

import (
	"github.com/labstack/echo"
	"bisale/bisale-console-api/config"
	"github.com/labstack/echo/middleware"
	"bisale/bisale-console-api/controllers"
	"bisale/bisale-console-api/middlewares"
	"github.com/sirupsen/logrus"
)

func main() {

	e := echo.New()
	e.Logger = middlewares.LogrusLogger{logrus.StandardLogger()}
	e.Use(middlewares.LogrusHook())
	e.Use(middleware.Recover())
	e.Use(middlewares.FilterRequests)

	e.GET("/ping", controllers.Ping)

	api := e.Group("/api")
	api.POST("/login", controllers.PostLogin)
	api.POST("/member", controllers.PostCreateMember)
	api.POST("/role", controllers.PostCreateRole)

	bisale := e.Group("/api/bisale")
	bisale.GET("/users", controllers.GetBisaleUsers)

	e.Logger.Fatal(e.Start(config.GetListenNetAddress()))
}
