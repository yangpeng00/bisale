package controllers

import (
	"github.com/labstack/echo"
	"bisale/bisale-console-api/common"
	"context"
	"bisale/bisale-console-api/codes"
	"bisale/bisale-console-api/thrift/finance"
)

type OrderRequest struct {
	Page int32 `query:"page"`
	Size int32 `query:"size"`
	UserId int32 `query:"userId"`
	Email string `query:"email"`
	Currency string `query:"currency"`
	Status string `query:"status"`
	StartTime string `query:"startTime"`
	EndTime string `query:"endTime"`
}

type OrderResult struct {
	list []*finance.TDepositWithdrawResult_
	count int32
}

func GetDepositOrder(c echo.Context) error {
	withdrawService, withdrawClient := common.GetBisaleWithdrawServiceClient()
	defer common.BisaleWithdrawServicePool.Put(withdrawClient)

	log, traceId := common.GetLoggerWithTraceId(c)

	request := new(OrderRequest)
	if err := c.Bind(request); err != nil {
		log.Error(err)
		return Status(c, codes.ServiceError, err)
	}

	params := new(finance.TDepositWithdrawParams)
	params.TraceId = traceId
	params.PageSize = request.Size
	params.StartPage = request.Page
	params.UserId = request.UserId
	params.Email = request.Email
	params.Currency = request.Currency
	params.Status = request.Status
	params.StartTime = request.StartTime
	params.EndTime = request.EndTime

	listResult, err := withdrawService.SelectSlaveDepositListByConditions(context.Background(), params)

	if err != nil {
		log.Error(err)
		return Status(c, codes.ServiceError, err)
	}

	countResult, err := withdrawService.SelectSlaveDepositCountByConditions(context.Background(), params)

	if err != nil {
		log.Error(err)
		return Status(c, codes.ServiceError, err)
	}

	r := new(OrderResult)
	r.list = listResult
	r.count = countResult

	return Status(c, codes.Success, r)
}

func GetWithdrawOrder(c echo.Context) error {
	withdrawService, withdrawClient := common.GetBisaleWithdrawServiceClient()
	defer common.BisaleWithdrawServicePool.Put(withdrawClient)

	log, traceId := common.GetLoggerWithTraceId(c)

	request := new(OrderRequest)
	if err := c.Bind(request); err != nil {
		log.Error(err)
		return Status(c, codes.ServiceError, err)
	}

	params := new(finance.TDepositWithdrawParams)
	params.TraceId = traceId
	params.PageSize = request.Size
	params.StartPage = request.Page
	params.UserId = request.UserId
	params.Email = request.Email
	params.Currency = request.Currency
	params.Status = request.Status
	params.StartTime = request.StartTime
	params.EndTime = request.EndTime

	countResult, err := withdrawService.SelectSlaveWithdrawLCountByConditions(context.Background(), params)

	if err != nil {
		log.Error(err)
		return Status(c, codes.Success, err)
	}

	listResult, err := withdrawService.SelectSlaveWithdrawListByConditions(context.Background(), params)

	if err != nil {
		log.Error(err)
		return Status(c, codes.Success, err)
	}

	r := new(OrderResult)
	r.list = listResult
	r.count = countResult

	return Status(c, codes.Success, r)
}
