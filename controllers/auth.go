package controllers

import (
	"net/http"
	"github.com/labstack/echo"
)

func PostLogin(c echo.Context) error {
	return c.JSON(http.StatusOK, Status{
		Code: 200,
	})
}
