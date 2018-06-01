package controllers

import (
	"net/http"
	"github.com/labstack/echo"
	"bisale/bisale-console-api/common"
)

func PostLogin(c echo.Context) error {
	return c.JSON(http.StatusOK, common.Status{

	})
}
