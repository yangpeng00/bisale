package controllers

import (
	"context"
	"github.com/labstack/echo"
	"bisale/bisale-console-api/codes"
	"bisale/bisale-console-api/common"
	"bisale/bisale-console-api/config"
	"bisale/bisale-console-api/thrift/message"
	accountInputs "bisale/bisale-console-api/thrift/inputs"
)

type LoginForm struct {
	Username string `json:"username" validate:"required"`
	Code   string `json:"code" validate:"required"`
	Key  string `json:"key"`
}

func GetLoginCodeIdentify(mobile string) string {
	return "login::" + mobile
}

func PostLoginSMSCode(c echo.Context) error {
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
		}
		return Status(c, codes.ServiceError, err)
	}

	times++
	expired, _ := captchaService.GetTodayLeftSeconds(ctx, traceId)
	err = captchaService.SetCount(ctx, traceId, identify, times, expired)

	if err != nil {
		return Status(c, codes.ServiceError, err)
	}

	return Status(c, codes.Success, map[string]string{
		"token": numberCode.Token,
	})
}

func PostLogin(c echo.Context) error {

	loginForm := new(LoginForm)

	if err := c.Bind(loginForm); err != nil {
		return Status(c, codes.FormIsEmpty, err)
	}
	if err := c.Validate(loginForm); err != nil {
		return Status(c, codes.ValidateError, err)
	}

	log, traceId := common.GetLoggerWithTraceId(c)
	ctx := context.Background()

	captchaService := common.GetCaptchaServiceClient()

	correct, err := captchaService.ValidateNumberCode(ctx, traceId, loginForm.Username, loginForm.Code, loginForm.Key)

	if err != nil {
		log.Error(err)
		return Status(c, codes.ServiceError, err)
	}

	if !correct {
		return Status(c, codes.SMSCodeError, err)
	}

	accountService := common.GetAccountServiceClient()
	token, err := accountService.GenerateJWTToken(ctx, traceId, &accountInputs.JWTInput{MemberId: "123"}, config.Config.JWTToken, 12)

	if err != nil {
		log.Error(err)
		return Status(c, codes.ServiceError, err)
	}

	return Status(c, codes.Success, map[string]string{
		"token": token,
	})
}
