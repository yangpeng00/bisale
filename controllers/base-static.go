package controllers

import (
	"github.com/labstack/echo"
	"bisale/bisale-console-api/common"
	"context"
	"bisale/bisale-console-api/codes"
)

type StaticResponse struct {
	KycCount int32 `json:"kycCount"`
	UserCount int32 `json:"userCount"`
	RegisterDailyCount interface{} `json:"registerDailyCount"`
	UserKycDailyCount interface{} `json:"userKycDailyCount"`
}

func GetStatic(c echo.Context) error {
	log, traceId := common.GetLoggerWithTraceId(c)

	userService, userClient := common.GetBisaleUserServiceClient()
	defer common.BisaleUserServicePool.Put(userClient)

	userKycService, userKycClient := common.GetBisaleUserKycServiceClient()
	defer common.BisaleUserKycServicePool.Put(userKycClient)

	//userCount, err := userService.SelectSlaveAllUserCount(context.Background(), traceId)
	//if err != nil {
	//	log.Error(err)
	//	return Status(c, codes.ServiceError, err)
	//}
	kycCount, err := userKycService.SelectSlaveAllUserKycCount(context.Background(), traceId)
	if err != nil {
		log.Error(err)
		return Status(c, codes.ServiceError, err)
	}
	registerCountDay, err := userService.SelectSlaveRegisterCountDay(context.Background(), traceId, 7)
	if err != nil {
		log.Error(err)
		return Status(c, codes.ServiceError, err)
	}
	kycCountDay, err := userKycService.SelectSlaveUserKycCountDay(context.Background(), traceId, 7)
	if err != nil {
		log.Error(err)
		return Status(c, codes.ServiceError, err)
	}
	staticResponse := new(StaticResponse)
	//staticResponse.UserCount = userCount
	staticResponse.KycCount = kycCount
	staticResponse.RegisterDailyCount = registerCountDay
	staticResponse.UserKycDailyCount = kycCountDay

	if err != nil {
		log.Error(err)
		return Status(c, codes.ServiceError, err)
	}

	return Status(c, codes.Success, staticResponse)
}

func GetTop5Award(c echo.Context) error {
	log, traceId := common.GetLoggerWithTraceId(c)

	userService, userClient := common.GetBisaleUserServiceClient()
	defer common.BisaleUserServicePool.Put(userClient)

	top5Award, err := userService.SelectSlaveInviteFriendsAwardTop(context.Background(), traceId, 5)

	if err != nil {
		log.Error(err)
		return Status(c, codes.ServiceError, err)
	}

	return Status(c, codes.Success, top5Award)
}
