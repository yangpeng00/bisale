package controllers

import (
	"github.com/labstack/echo"
	"bisale/bisale-console-api/codes"
)

func PostLogin(c echo.Context) error {
	return Status(c, codes.Success, "success")

}
