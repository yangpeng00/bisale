package controllers

import (
	"net/http"
	"github.com/labstack/echo"
)

func PostCreateRole(c echo.Context) error {
	return c.String(http.StatusOK, "success")
}
