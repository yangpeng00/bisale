package controllers

import (
	"context"
	"github.com/labstack/echo"
	"bisale/bisale-console-api/codes"
	"bisale/bisale-console-api/common"
	accountInputs "bisale/thrift-account/thrift/inputs"
)

func PostCreateMember(c echo.Context) error {
	log, traceId := common.GetLoggerWithTraceId(c)
	ctx := context.Background()

	accountService := common.GetAccountServiceClient()

	createMemberInput := accountInputs.CreateMemberInput{
		Account:  "koyeo",
		Email:    "koyeo@qq.com",
		Mobile:   "+86-18817392521",
		Password: "Helloshic",
		Status:   0,
	}

	createMemberOutput, err := accountService.CreateMember(ctx, traceId, &createMemberInput)

	if err != nil {
		log.Error(err)
		return err
	}
	return Status(c, codes.Success, createMemberOutput)
}
