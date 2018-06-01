package controllers

import (
	"net/http"
	"github.com/labstack/echo"
)

func PostCreateMember(c echo.Context) error {
	return c.JSON(http.StatusOK, Status{

	})
}
