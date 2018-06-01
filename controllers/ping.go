package controllers

import (
	"github.com/labstack/echo"
)

func Ping(c echo.Context) error {

	return Response(c, Status{
		Code: 200,
	})
}
