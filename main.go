package main

import (
	"net/http"
	"github.com/labstack/echo"
	"bisale/bisale-console-api/config"
	"github.com/labstack/echo/middleware"
	"bisale/bisale-console-api/controllers"
	"bisale/bisale-console-api/common"
)

func main() {

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, common.Status{
			Code: 200,
		}.String())
	})

	e.POST("/api/login", controllers.PostLogin)
	e.POST("/api/member", controllers.PostCreateMember)
	e.POST("/api/role", controllers.PostCreateRole)

	e.Logger.Fatal(e.Start(config.GetListenNetAddress()))
}
