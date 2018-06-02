package controllers

import (
	"errors"
	"github.com/labstack/echo"
	"bisale/bisale-console-api/codes"
	"bisale/bisale-console-api/common"
)

func GetBisaleUsers(c echo.Context) error {
	log := common.GetLoggerWithTraceId(c)
	log.Error(errors.New("some thing wrong"))
	return Status(c, codes.ValidateError, errors.New("some thing wrong"))
	return Status(c, codes.Success, "")
}
