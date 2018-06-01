package controllers

import (
	"net/http"
	"github.com/labstack/echo"
	"bisale/bisale-console-api/common"
)

func PostCreateMember(c echo.Context) error {
	return c.JSON(http.StatusOK, common.Status{

	})
}
