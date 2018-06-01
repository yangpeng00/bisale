package main

import (
	"github.com/labstack/echo"
	"bisale/bisale-console-api/config"
	"github.com/labstack/echo/middleware"
	"bisale/bisale-console-api/controllers"
)

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

func main() {

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(serverHeader)

	e.GET("/ping", controllers.Ping)
	e.POST("/api/login", controllers.PostLogin)
	e.POST("/api/member", controllers.PostCreateMember)
	e.POST("/api/role", controllers.PostCreateRole)

	e.Logger.Fatal(e.Start(config.GetListenNetAddress()))
}
