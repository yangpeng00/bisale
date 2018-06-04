package controllers

import (
	"fmt"
	"context"
	"github.com/labstack/echo"
	"bisale/bisale-console-api/codes"
	"bisale/bisale-console-api/common"
	"bisale/bisale-console-api/thrift/message"
)

func GetLoginCodeIdentify(mobile string) string {
	return "login::" + mobile
}

func GetLoginSMSCode(c echo.Context) error {
	log, traceId := common.GetLoggerWithTraceId(c)

	messageService := common.GetMessageServiceClient()
	captchaService := common.GetCaptchaServiceClient()

	ctx := context.Background()
	mobile := c.Param("mobile")
	identify := GetLoginCodeIdentify(mobile)
	times, _ := captchaService.GetCount(ctx, traceId, identify)
	if times > 10 {
		return Status(c, codes.SendCodeLock60Seconds, "")
	}
	numberCode, err := captchaService.GenerateNumberCode(ctx, traceId, mobile, 6, 600)
	if err != nil {
		log.Error(err)
		return err
	}
	appId := "bisale-console-api"
	template := "template::sms::login-code"
	err = messageService.SendSMS(ctx, traceId, appId, mobile, template, "{\"code\":"+numberCode.Value+"}", "zh-CN", 60)
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
	times++
	expired, _ := captchaService.GetTodayLeftSeconds(ctx, traceId)
	err = captchaService.SetCount(ctx, traceId, identify, times, expired)
	if err != nil {
		return err
	}
	return Status(c, codes.Success, map[string]string{
		"token": numberCode.Token,
	})
}

func PostLogin(c echo.Context) error {
	//mobile := c.FormValue("mobile")
	//code := c.FormValue("code")
	//token := c.FormValue("token")
	log, traceId := common.GetLoggerWithTraceId(c)
	ctx := context.Background()
	// messageService := common.GetMessageServiceClient()
	// captchaService := common.GetCaptchaServiceClient()
	accountService := common.GetAccountServiceClient()
	token, err := accountService.GenerateJWTToken(ctx, traceId, "welcome to bad", 12)
	if err != nil {
		log.Error(err)
		return err
	}
	return Status(c, codes.Success, map[string]string{
		"token": token,
	})
}
