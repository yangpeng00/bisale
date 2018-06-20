package controllers

import (
	"fmt"
	"context"
	"github.com/labstack/echo"
	"bisale/bisale-console-api/codes"
	"bisale/bisale-console-api/common"
	accountInputs "bisale/bisale-console-api/thrift/inputs"
)

type CreateMemberForm struct {
	Mobile   string `json:"mobile" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func PostCreateMember(c echo.Context) error {

	log, traceId := common.GetLoggerWithTraceId(c)
	ctx := context.Background()

	createMemberForm := new(CreateMemberForm)

	if err := c.Bind(createMemberForm); err != nil {
		return Status(c, codes.FormIsEmpty, err)
	}

	fmt.Println(createMemberForm)

	if err := c.Validate(createMemberForm); err != nil {
		return Status(c, codes.ValidateError, err)
	}

	accountService,accountClient := common.GetAccountServiceClient()
	defer common.AccountServicePool.Put(accountClient)

	createMemberInput := accountInputs.CreateMemberInput{
		Mobile:   createMemberForm.Mobile,
		Password: createMemberForm.Password,
		Status:   0,
	}

	createMemberOutput, err := accountService.CreateMember(ctx, traceId, &createMemberInput)

	if err != nil {
		log.Error(err)
		return Status(c, codes.ServiceError, err)
	}
	return Status(c, codes.Success, createMemberOutput)
}
