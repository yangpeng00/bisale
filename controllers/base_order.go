package controllers

import (
	"github.com/labstack/echo"
	"bisale/bisale-console-api/common"
	"context"
	"bisale/bisale-console-api/codes"
	"bisale/bisale-console-api/thrift/finance"
	"bisale/bisale-console-api/thrift/engine"
	"encoding/json"
	"time"
)

type OrderRequest struct {
	Page      int32  `query:"page"`
	Size      int32  `query:"size"`
	UserId    int32  `query:"userId"`
	Email     string `query:"email"`
	Mobile    string `query:"mobile"`
	Symbol    string `query:"symbol"`
	Status    string `query:"status"`
	Type      int32  `query:"type"`
	StartTime string `query:"startTime"`
	EndTime   string `query:"endTime"`
}

type ExchangeRequest struct {
	Page      int32  `query:"page"`
	Size      int32  `query:"size"`
	UserId    int32  `query:"userId"`
	Email     string `query:"email"`
	Mobile    string `query:"mobile"`
	Side      string `query:"side"`
	Status    string `query:"status"`
	Symbol    string `query:"symbol"`
	StartTime string `query:"StartTime"`
	EndTime   string `query:"endTime"`
}

type ExchangeDetailRequest struct {
	Page    int32  `query:"page"`
	Size    int32  `query:"size"`
	OrderId string `query:"orderId"`
}

type OrderResult struct {
	list  string
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
	params.Mobile = request.Mobile
	params.Currency = request.Symbol
	params.Status = request.Status
	params.StartTime = request.StartTime
	params.EndTime = request.EndTime
	params.Type = request.Type

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

	r := make(map[string]interface{})
	r["list"] = listResult
	r["count"] = countResult

	walletService, walletClient := common.GetWalletServiceClient()
	defer common.WalletServicePool.Put(walletClient)

	config := make(map[string]interface{})
	config["lang"] = "zh-CN"
	configStr, _ := json.Marshal(config)

	currencyInfo, err := walletService.Execute(context.Background(), "Currency", "getConfigs", string(configStr))
	if err != nil {
		log.Error(err)
		return Status(c, codes.ServiceError, err)
	}
	r["symbolList"] = currencyInfo

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
	params.Mobile = request.Mobile
	params.Currency = request.Symbol
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

	res := make(map[string]interface{})
	res["list"] = listResult
	res["count"] = countResult

	walletService, walletClient := common.GetWalletServiceClient()
	defer common.WalletServicePool.Put(walletClient)

	config := make(map[string]interface{})
	config["lang"] = "zh-CN"
	configStr, _ := json.Marshal(config)

	currencyInfo, err := walletService.Execute(context.Background(), "Currency", "getConfigs", string(configStr))
	if err != nil {
		log.Error(err)
		return Status(c, codes.ServiceError, err)
	}
	res["symbolList"] = currencyInfo

	return Status(c, codes.Success, res)
}

func GetExchangeOrder(c echo.Context) error {
	log, traceId := common.GetLoggerWithTraceId(c)

	request := new(ExchangeRequest)
	c.Bind(request)

	orderService, orderClient := common.GetBisaleOrderServiceClient()
	defer common.BisaleOrderServicePool.Put(orderClient)

	params := new(engine.TOrdersParams)
	params.TraceId = traceId
	params.UserId = request.UserId
	params.PageSize = request.Size
	params.StartPage = request.Page
	params.Status = request.Status
	params.StartTime = request.StartTime
	params.EndTime = request.EndTime
	params.Symbol = request.Symbol
	params.Side = request.Side
	params.Email = request.Email
	params.Mobile = request.Mobile

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	list, err := orderService.SelectEngineOrdersListByConditions(ctx, params)
	if err != nil {
		log.Error(err)
		return Status(c, codes.ServiceError, err)
	}

	count, err := orderService.SelectEngineOrdersCountByConditions(ctx, params)
	if err != nil {
		log.Error(err)
		return Status(c, codes.ServiceError, err)
	}

	symbolList, err := orderService.SelectSymbolsList(ctx)
	if err != nil {
		log.Error(err)
		return Status(c, codes.ServiceError, err)
	}

	res := make(map[string]interface{})
	res["list"] = list
	res["count"] = count
	res["symbolList"] = symbolList

	return Status(c, codes.Success, res)
}

func GetExchangeOrderDetail(c echo.Context) error {
	log, traceId := common.GetLoggerWithTraceId(c)

	request := new(ExchangeDetailRequest)
	c.Bind(request)

	orderService, orderClient := common.GetBisaleOrderServiceClient()
	defer common.BisaleOrderServicePool.Put(orderClient)

	params := new(engine.TOrdersDetailParams)
	params.TraceId = traceId
	params.StartPage = request.Page
	params.PageSize = request.Size
	params.OrderId = request.OrderId

	count, err := orderService.SelectEngineOrdersDetailCountByOrderId(context.Background(), params)

	if err != nil {
		log.Error(err)
		return Status(c, codes.ServiceError, err)
	}

	list, err := orderService.SelectEngineOrdersDetailListByOrderId(context.Background(), params)

	if err != nil {
		log.Error(err)
		return Status(c, codes.ServiceError, err)
	}

	res := make(map[string]interface{})
	res["list"] = list
	res["count"] = count

	return Status(c, codes.Success, res)
}
