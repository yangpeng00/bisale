package controllers

import (
	"fmt"
	"context"
	"github.com/labstack/echo"
	"bisale/bisale-console-api/codes"
	"bisale/bisale-console-api/common"
	"bisale/bisale-console-api/thrift/message"
)

func GetLoginSMSCode(c echo.Context) error {
	log, traceId := common.GetLoggerWithTraceId(c)
	messageService := common.GetMessageServiceClient()
	ctx := context.Background()
	err := messageService.SendSMS(
		ctx,
		traceId,
		"bisale-console-api",
		"+86-18817392521",
		"template::sms::login-code",
		`{"code":123456}`,
		"zh-CN",
		60,
	)
	if err != nil {
		log.Error(err)
		if status, ok := err.(*message.Status); ok {
			if status.Code == 30060 {
				return Status(c, codes.SendCodeLock60Seconds, "")
			}
		} else {
			fmt.Println(ok)
		}
		return err
	}
	return Status(c, codes.Success, "success")
}

func PostLogin(c echo.Context) error {
	return Status(c, codes.Success, "success")
}
