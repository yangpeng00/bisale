package main

import (
	"github.com/labstack/echo"
	"bisale/bisale-console-api/config"
	"github.com/labstack/echo/middleware"
	"bisale/bisale-console-api/controllers"
)

func main() {

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(serverHeader)

	e.GET("/ping", controllers.Ping)

	api := e.Group("/api")

	api.POST("/login", controllers.PostLogin)
	api.POST("/member", controllers.PostCreateMember)
	api.POST("/role", controllers.PostCreateRole)

	e.Logger.Fatal(e.Start(config.GetListenNetAddress()))
}

func serverHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderServer, "Echo/3.0")
		//c.Response().
		if err := next(c); err != nil {
			c.Error(err)
		}
		return nil
	}
}
