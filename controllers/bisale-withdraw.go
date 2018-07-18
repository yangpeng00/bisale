package controllers

import (
	"github.com/labstack/echo"
	"strconv"
	"bisale/bisale-console-api/codes"
	"bisale/bisale-console-api/common"
	"context"
	"bisale/bisale-console-api/thrift/finance"
)

type PostWithdrawResultRequest struct {
	Id int32 `json:"id"`
	UserId int32 `json:"userId"`
	Status string `json:"status"`
	Mark string `json:"mark"`
}

func GetWithdrawList(c echo.Context) error {
	page, _ := strconv.ParseInt(c.QueryParam("page"), 10, 32)
	size, _ := strconv.ParseInt(c.QueryParam("size"), 10, 32)

	if size < 10 {
		size = 10
	}

	log, traceId := common.GetLoggerWithTraceId(c)
	withdrawService, withdrawClient := common.GetBisaleWithdrawServiceClient()
	defer common.BisaleWithdrawServicePool.Put(withdrawClient)

	withdrawParams := new(finance.TDepositWithdrawParams)

	withdrawParams.TraceId = traceId
	withdrawParams.UserName = c.QueryParam("keyword")
	withdrawParams.Status = c.QueryParam("status")
	withdrawParams.StartTime = c.QueryParam("startedAt")
	withdrawParams.EndTime = c.QueryParam("endedAt")
	withdrawParams.PageSize = int32(size)
	withdrawParams.StartPage = int32(page)

	res, err := withdrawService.SelectDepositWithdrawByConditions(context.Background(), withdrawParams)

	log.Info(res)

	if err != nil {
		log.Error(err)
		return Status(c, codes.ServiceError, nil)
	}
	return Status(c, codes.Success, res)
}

func GetWithdrawListCount(c echo.Context) error {
	log, traceId := common.GetLoggerWithTraceId(c)
	withdrawService, withdrawClient := common.GetBisaleWithdrawServiceClient()
	defer common.BisaleWithdrawServicePool.Put(withdrawClient)

	withdrawParams := new(finance.TDepositWithdrawParams)

	withdrawParams.TraceId = traceId
	withdrawParams.UserName = c.QueryParam("keyword")
	withdrawParams.Status = c.QueryParam("status")
	withdrawParams.StartTime = c.QueryParam("startedAt")
	withdrawParams.EndTime = c.QueryParam("endedAt")

	res, err := withdrawService.SelectDepositWithdrawCountByConditions(context.Background(), withdrawParams)

	log.Info(res)

	if err != nil {
		log.Error(err)
		return Status(c, codes.ServiceError, nil)
	}
	return Status(c, codes.Success, res)

}

func PostWithdrawResult(c echo.Context) error {
	log, traceId := common.GetLoggerWithTraceId(c)
	userService, userClient := common.GetBisaleWithdrawServiceClient()
	defer common.BisaleWithdrawServicePool.Put(userClient)

	req := new(PostWithdrawResultRequest)

	err := c.Bind(req)
	if err != nil {
		log.Error(err)
		return Status(c, codes.ServiceError, err)
	}

	res, err :=userService.AuditDepositWithdraw(context.Background(), traceId, req.Status, req.Mark, req.Id)
	if err != nil || res == false {
		log.Error(err)
		return Status(c, codes.ServiceError, err)
	}

	return Status(c, codes.Success, res)

}
