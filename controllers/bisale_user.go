package controllers

import (
	"github.com/labstack/echo"
	"bisale/bisale-console-api/codes"
	"bisale/bisale-console-api/common"
	"context"
)

func GetBisaleUsers(c echo.Context) error {
	log, _ := common.GetLoggerWithTraceId(c)
	userService := common.GetBisaleServiceClient()
	ctx := context.Background()
	res, err := userService.SelectUserKycByConditions(ctx, "", "", 1, 100)
	if err != nil {
		log.Error(err)
	}
	return Status(c, codes.Success, res)
}
