package controllers

import (
	"fmt"
	"context"
	"github.com/labstack/echo"
	"bisale/bisale-console-api/codes"
	"bisale/bisale-console-api/common"
)

func GetLoginSMSCode(c echo.Context) error {
	// log := common.GetLoggerWithTraceId(c)
	messageService := common.GetMessageServiceClient()
	ctx := context.Background()
	res, _ := messageService.Ping(ctx)
	fmt.Println(res)
	return Status(c, codes.Success, "success")
}

func PostLogin(c echo.Context) error {
	return Status(c, codes.Success, "success")
}
