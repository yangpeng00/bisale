package controllers

import (
	"context"
	"bisale/bisale-console-api/codes"
	"bisale/bisale-console-api/common"
	"github.com/labstack/echo"
	"bisale/bisale-console-api/thrift/engine"
	"bisale/bisale-console-api/thrift/finance"
	"strconv"
	"bisale/bisale-console-api/thrift/balanceAccount"
)

type TradeDetailResult struct {
	List []*engine.TTradeDetail `json:"list"`
	Count int32 `json:"count"`
}

type EngineAccountResult struct {
	List []*engine.TEngineAccountCheckingItem `json:"list"`
	Count int32 `json:"count"`
}

func GetReconciliationExchangeDetail(c echo.Context) error {
	log, _ := common.GetLoggerWithTraceId(c)
	tradeDetailService, tradeDetailClient := common.GetBisaleTradeDetailServiceClient()
	defer common.BisaleTradeDetailServicePool.Put(tradeDetailClient)

	params := new(engine.TTradeDetailsParams)
	pageSize, _ := strconv.ParseInt(c.QueryParam("size"), 10, 32)
	params.PageSize = int32(pageSize)
	startPage, _ := strconv.ParseInt(c.QueryParam("page"), 10, 32)
	params.StartPage = int32(startPage)

	buyerId, _ := strconv.ParseInt(c.QueryParam("buyerId"), 10, 32)
	params.BuyerId = int32(buyerId)
	params.BuyerEmail = c.QueryParam("buyerEmail")

	sellerId, _ := strconv.ParseInt(c.QueryParam("sellerId"), 10, 32)
	params.SellerId = int32(sellerId)
	params.BuyerEmail = c.QueryParam("buyerEmail")

	params.EndTime = c.QueryParam("endTime")
	params.StartTime = c.QueryParam("startTime")
	params.CurrencyPair = c.QueryParam("currencyPair")

	count, err := tradeDetailService.SelectTradeDetailsCountBy(context.Background(), params)
	list, err := tradeDetailService.SelectTradeDetailsBy(context.Background(), params)
	if err != nil {
		log.Error(err)
		return Status(c, codes.ServiceError, err)
	}
	
	result := new(TradeDetailResult)
	result.List = list
	result.Count = count

	return Status(c, codes.Success, result)
}

func GetReconciliation(c echo.Context) error {
	log, _ := common.GetLoggerWithTraceId(c)
	reconciliationService, reconciliationClient := common.GetBisaleAccountStatementServiceClient()
	defer common.BisaleAccountStatementPool.Put(reconciliationClient)

	startTime := c.QueryParam("startTime")
	endTime := c.QueryParam("endTime")

	params := new(finance.TAccountStatementParams)
	params.StartTime = startTime
	params.EndTime = endTime

	result, err := reconciliationService.SelectAccountStatement(context.Background(), params)

	if err != nil {
		log.Error(err)
		return Status(c, codes.ServiceError, err)
	}
	return Status(c, codes.Success, result )

}

func GetTransferRequest(c echo.Context) error {
	log, _ := common.GetLoggerWithTraceId(c)
	transferRequestService, reconciliationClient := common.GetBisaleTransferRequestServiceClient()
	defer common.BisaleTransferRequestServicePool.Put(reconciliationClient)

	params := new(balanceAccount.TTransferRequestParams)
	pageSize, _ := strconv.ParseInt(c.QueryParam("size"), 10, 32)
	params.PageSize = int32(pageSize)
	startPage, _ := strconv.ParseInt(c.QueryParam("page"), 10, 32)
	params.Page = int32(startPage)

	params.EndTime = c.QueryParam("endTime")
	params.StartTime = c.QueryParam("startTime")

	status, _ := strconv.ParseInt(c.QueryParam("status"), 10, 32)
	params.Status = int32(status)
	transferType, _ := strconv.ParseInt(c.QueryParam("type"), 10, 32)
	params.Type = int32(transferType)
	source, _ := strconv.ParseInt(c.QueryParam("source"), 10, 32)
	params.Source = int32(source)
	userId, _ := strconv.ParseInt(c.QueryParam("userId"), 10, 32)
	params.UserId = int32(userId)

	params.Email = c.QueryParam("email")
	params.Currency = c.QueryParam("currency")

	result, err := transferRequestService.GetTransferRequestList(context.Background(), params)

	if err != nil {
		log.Error(err)
		return Status(c, codes.ServiceError, err)
	}
	return Status(c, codes.Success, result)
}

func GetEngineAccount(c echo.Context) error {
	log, traceId := common.GetLoggerWithTraceId(c)
	engineAccountService, engineAccountClient := common.GetBisaleAccountTransferServiceClient()
	defer common.BisaleAccountTransferServicePool.Put(engineAccountClient)

	params := new(engine.TEngineAccountCheckingParams)
	pageSize, _ := strconv.ParseInt(c.QueryParam("size"), 10, 32)
	params.PageSize = int32(pageSize)
	startPage, _ := strconv.ParseInt(c.QueryParam("page"), 10, 32)
	params.StartPage = int32(startPage)
	userId, _ := strconv.ParseInt(c.QueryParam("userId"), 10, 32)
	params.UserId = int32(userId)
	params.Email = c.QueryParam("email")
	params.Currency = c.QueryParam("currency")
	params.StartTime = c.QueryParam("startTime")
	params.EndTime = c.QueryParam("endTime")
	params.Type = c.QueryParam("type")
	params.TraceId = traceId

	list, err := engineAccountService.SelectEngineAccountDataBy(context.Background(), params)
	count, err := engineAccountService.SelectEngineAccountCountBy(context.Background(), params)
	result := new(EngineAccountResult)
	result.List = list
	result.Count = count
	if err != nil {
		log.Error(err)
		return Status(c, codes.ServiceError, err)
	}
	return Status(c, codes.Success, result)
}

func GetBlockchainDeposit(c echo.Context) error {
	log, _ := common.GetLoggerWithTraceId(c)
	balanceAccountService, balanceAccountClient := common.GetBisaleBalanceAccountServiceClient()
	defer common.BisaleBalanceAccountServicePool.Put(balanceAccountClient)

	params := new(balanceAccount.TChainDepositWithdrawParams)
	params.TxId =  c.QueryParam("txId")
	params.Currency = c.QueryParam("currency")
	params.StartTime = c.QueryParam("startTime")
	params.EndTime = c.QueryParam("endTime")
	pageSize, _ := strconv.ParseInt(c.QueryParam("size"), 10, 32)
	params.PageSize = int32(pageSize)
	startPage, _ := strconv.ParseInt(c.QueryParam("page"), 10, 32)
	params.Page = int32(startPage)
	if c.QueryParam("orderId") != "" {
		orderId, _ := strconv.ParseInt(c.QueryParam("orderId"), 10, 32)
		params.OrderId = int32(orderId)
	}
	if c.QueryParam("checkExec") != "" {
		checkExec, _ := strconv.ParseInt(c.QueryParam("checkExec"), 10, 8)
		params.CheckExec = int8(checkExec)
	}
	if c.QueryParam("status") != "" {
		status, _ := strconv.ParseInt(c.QueryParam("status"), 10, 8)
		params.Status = int8(status)
	}

	result, err := balanceAccountService.GetChainDeposit(context.Background(), params)
	if err != nil {
		log.Error(err)
		return Status(c, codes.ServiceError, err)
	}
	return Status(c, codes.Success, result)
}

func GetBlockchainWithdraw(c echo.Context) error {
	log, _ := common.GetLoggerWithTraceId(c)
	balanceAccountService, balanceAccountClient := common.GetBisaleBalanceAccountServiceClient()
	defer common.BisaleBalanceAccountServicePool.Put(balanceAccountClient)

	params := new(balanceAccount.TChainDepositWithdrawParams)
	params.TxId =  c.QueryParam("txId")
	params.Currency = c.QueryParam("currency")
	params.StartTime = c.QueryParam("startTime")
	params.EndTime = c.QueryParam("endTime")
	pageSize, _ := strconv.ParseInt(c.QueryParam("size"), 10, 32)
	params.PageSize = int32(pageSize)
	startPage, _ := strconv.ParseInt(c.QueryParam("page"), 10, 32)
	params.Page = int32(startPage)
	orderId, _ := strconv.ParseInt(c.QueryParam("orderId"), 10, 32)
	params.OrderId = int32(orderId)
	checkExec, _ := strconv.ParseInt(c.QueryParam("checkExec"), 10, 8)
	params.CheckExec = int8(checkExec)
	status, _ := strconv.ParseInt(c.QueryParam("status"), 10, 8)
	params.Status = int8(status)

	result, err := balanceAccountService.GetChainWithdraw(context.Background(), params)
	if err != nil {
		log.Error(err)
		return Status(c, codes.ServiceError, err)
	}
	return Status(c, codes.Success, result)
}

