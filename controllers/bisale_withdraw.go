package controllers

import (
	"github.com/labstack/echo"
	"strconv"
	"bisale/bisale-console-api/codes"
	"bisale/bisale-console-api/common"
	"context"
	"bisale/bisale-console-api/thrift/finance"
)

func GetWithdrawList(c echo.Context) error {
	page, _ := strconv.ParseInt(c.QueryParam("page"), 10, 32)
	size, _ := strconv.ParseInt(c.QueryParam("size"), 10, 32)

	if size < 10 {
		size = 10
	}

	log, _ := common.GetLoggerWithTraceId(c)
	withdrawService := common.GetBisaleWithdrawServiceClient()

	withdrawParams := new(finance.TDepositWithdrawParams)

	withdrawParams.UserName = c.QueryParam("keyword")
	withdrawParams.Status = c.QueryParam("status")
	withdrawParams.StartTime = c.QueryParam("startTime")
	withdrawParams.EndTime = c.QueryParam("endTime")
	withdrawParams.PageSize = int32(size)
	withdrawParams.StartPage = int32(page)

	log.Info("==gogogo==")
	log.Info(withdrawParams.StartTime, withdrawParams.EndTime, withdrawParams.PageSize, withdrawParams.StartPage)

	res, err := withdrawService.SelectDepositWithdrawByConditions(context.Background(), withdrawParams)

	log.Info(res)

	if err != nil {
		log.Error(err)
		return Status(c, codes.ServiceError, nil)
	}
	return Status(c, codes.Success, res)
}
