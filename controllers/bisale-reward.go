package controllers

import (
	"github.com/labstack/echo"
	"bisale/bisale-console-api/common"
	"bisale/bisale-console-api/thrift/finance"
	"context"
	"bisale/bisale-console-api/codes"
	"bisale/bisale-console-api/thrift/user"
	"strconv"
	"fmt"
	"bisale/bisale-console-api/thrift/system"
)

type RewardParam struct {
	TraceId string `query:"traceId" json:"traceId"`
	StartDate string `query:"startDate" json:"startDate"`
	EndDate string `query:"endDate" json:"endDate"`
	UserId int32 `query:"userId" json:"userId"`
	Mobile string `query:"mobile" json:"mobile"`
	Email string `query:"email" json:"email"`
	Page int32 `query:"page" json:"page"`
	Size int32 `query:"size" json:"size"`
	Sort string `query:"sort" json:"sort"`
	CurrentDate string `query:"currentDate" json:"currentDate"`
}

type UserParams struct {
	UserId int32 `query:"userId" json:"userId"`
	Email string `query:"email" json:"email"`
	Mobile string `query:"mobile" json:"mobile"`
	IsSpecial int32 `query:"isSpecial" json:"isSpecial"`
	Page int32 `query:"page" json:"page"`
	Size int32 `query:"size" json:"size"`
	IsTrade string `query:"isTrade" json:"isTrade"`
	IsDividend string `query:"isDividend" json:"isDividend"`
}

type SystemParams struct {
	Payload string `query:"payload" json:"payload"`
}

type AwardResult struct {
	List []*finance.TTradeAwardDaysListResult_ `json:"list"`
	TotalCount int32 `json:"totalCount"`
	TotalAward string `json:"totalAward"`
}

type AwardDetailResult struct {
	List []*finance.TTradeAwardDetailListResult_ `json:"list"`
	TotalCount int32 `json:"totalCount"`
	TotalAward string `json:"totalAward"`
	CurrentDate string `json:"currentDate"`
}

type BonusResult struct {
	List []*finance.TDividedDaysListResult_ `json:"list"`
	CurrencyList []*finance.TDividedCurrencyResult_ `json:"currencyList"`
	TotalCount int32 `json:"totalCount"`
	TotalAward string `json:"totalAward"`
}

type BonusDetailResult struct {
	List []*finance.TDividedDaysDetailListResult_ `json:"list"`
	CurrencyList []*finance.TDividedCurrencyResult_ `json:"currencyList"`
	TotalCount int32 `json:"totalCount"`
	TotalAward string `json:"totalAward"`
	CurrentDate string `json:"currentDate"`
}

type AttributeResult struct {
	List []*user.TUserAttributeResult_ `json:"list"`
	Count int32 `json:"count"`
	Config []*system.TSystemConfigResult_ `json:"config"`
}

func GetExchangeDetail(c echo.Context) error {
	withdrawService, withdrawClient := common.GetBisaleWithdrawServiceClient()
	defer common.BisaleWithdrawServicePool.Put(withdrawClient)

	log, traceId := common.GetLoggerWithTraceId(c)

	params := new(RewardParam)
	if err := c.Bind(params); err != nil {
		log.Error(err)
		return Status(c, codes.ServiceError, err)
	}

	query := new(finance.TTradeAwardParams)
	query.TraceId = traceId
	query.PageSize = params.Size
	query.StartPage = params.Page
	query.UserId = params.UserId
	query.Email = params.Email
	query.Mobile = params.Mobile
	query.Sort = params.Sort
	query.StartDate = params.StartDate
	query.EndDate = params.EndDate
	query.CurrentDate = params.CurrentDate

	resultList, err := withdrawService.SelectTradeAwardDaysDetailList(context.Background(), query)
	count, err := withdrawService.SelectTradeAwardDaysDetailCount(context.Background(), query)
	if err != nil {
		log.Error(err)
		return Status(c, codes.ServiceError, err)
	}
	result := new(AwardDetailResult)
	result.List = resultList
	result.TotalCount = count.RecordNumber
	result.CurrentDate = count.CurrentDate
	result.TotalAward = count.TotalAward

	return Status(c, codes.Success, result)
}

func GetExchangeList(c echo.Context) error {
	withdrawService, withdrawClient := common.GetBisaleWithdrawServiceClient()
	defer common.BisaleWithdrawServicePool.Put(withdrawClient)

	log, traceId := common.GetLoggerWithTraceId(c)

	params := new(RewardParam)
	if err := c.Bind(params); err != nil {
		log.Error(err)
		return Status(c, codes.ServiceError, err)
	}

	query := new(finance.TTradeAwardParams)
	query.TraceId = traceId
	query.PageSize = params.Size
	query.StartPage = params.Page
	query.UserId = params.UserId
	query.Email = params.Email
	query.Mobile = params.Mobile
	query.Sort = params.Sort
	query.StartDate = params.StartDate
	query.EndDate = params.EndDate
	query.CurrentDate = params.CurrentDate

	resultList, err := withdrawService.SelectTradeAwardDaysList(context.Background(), query)
	count, err := withdrawService.SelectTradeAwardDaysCount(context.Background(), query)
	if err != nil {
		log.Error(err)
		return Status(c, codes.ServiceError, err)
	}
	result := new(AwardResult)
	result.List = resultList
	result.TotalCount = count.RecordNumber
	result.TotalAward = count.TotalNumber

	return Status(c, codes.Success, result)
}

func GetBonusList(c echo.Context) error {
	withdrawService, withdrawClient := common.GetBisaleWithdrawServiceClient()
	defer common.BisaleWithdrawServicePool.Put(withdrawClient)

	log, traceId := common.GetLoggerWithTraceId(c)


	params := new(RewardParam)
	if err := c.Bind(params); err != nil {
		log.Error(err)
		return Status(c, codes.ServiceError, err)
	}

	query := new(finance.TTradeAwardParams)
	query.TraceId = traceId
	query.PageSize = params.Size
	query.StartPage = params.Page
	query.UserId = params.UserId
	query.Email = params.Email
	query.Mobile = params.Mobile
	query.StartDate = params.StartDate
	query.EndDate = params.EndDate
	query.CurrentDate = params.CurrentDate

	resultList, err := withdrawService.SelectDividedDaysList(context.Background(), query)

	count, err := withdrawService.SelectDividedDaysCount(context.Background(), query)

	if err != nil {
		log.Error(err)
		return Status(c, codes.ServiceError, err)
	}

	result := new(BonusResult)
	result.List = resultList
	result.CurrencyList = count.CurrencyList
	result.TotalCount = count.RecordNumber

	return Status(c, codes.Success, result)
}

func GetBonusDetail(c echo.Context) error {
	withdrawService, withdrawClient := common.GetBisaleWithdrawServiceClient()
	defer common.BisaleWithdrawServicePool.Put(withdrawClient)

	log, traceId := common.GetLoggerWithTraceId(c)

	params := new(RewardParam)
	if err := c.Bind(params); err != nil {
		log.Error(err)
		return Status(c, codes.ServiceError, err)
	}

	query := new(finance.TTradeAwardParams)
	query.TraceId = traceId
	query.PageSize = params.Size
	query.StartPage = params.Page
	query.UserId = params.UserId
	query.Email = params.Email
	query.Mobile = params.Mobile
	query.StartDate = params.StartDate
	query.EndDate = params.EndDate
	query.CurrentDate = params.CurrentDate

	resultList, err := withdrawService.SelectDividedDaysDetailList(context.Background(), query)
	count, err := withdrawService.SelectDividedDaysDetailCount(context.Background(), query)
	totalCount, err := withdrawService.SelectDividedDaysCount(context.Background(), query)
	if err != nil {
		log.Error(err)
		return Status(c, codes.ServiceError, err)
	}

	fmt.Println(totalCount)

	result := new(BonusDetailResult)
	result.List = resultList
	result.TotalCount = count.RecordNumber
	result.CurrentDate = count.CurrentDate
	result.CurrencyList = totalCount.CurrencyList

	return Status(c, codes.Success, result)
}

func GetUserAttribute(c echo.Context) error {
	userService, userClient := common.GetBisaleUserServiceClient()
	defer common.BisaleUserServicePool.Put(userClient)

	systemService, systemClient := common.GetBisaleSystemServiceClient()
	defer common.BisaleSystemServicePool.Put(systemClient)

	log, traceId := common.GetLoggerWithTraceId(c)

	query := new(UserParams)
	if err := c.Bind(query); err != nil {
		log.Error(err)
		return Status(c, codes.ServiceError, err)
	}

	params := new(user.TUserAttributeParam)
	params.IsSpecial = query.IsSpecial
	params.UserId = query.UserId
	params.Email = query.Email
	params.Mobile = query.Mobile
	params.StartPage = query.Page
	params.PageSize = query.Size

	list, err := userService.SelectUserAttributeListByConditions(context.Background(), traceId, params)
	count, err := userService.SelectUserAttributeCountByConditions(context.Background(), traceId, params)
	config, err := systemService.SelectSystemConfigByKeyWords(context.Background(), traceId, "kycLimit,noKycLimit,tradeRate,dividendRate,vipRate")
	if err != nil {
		log.Error(err)
		return Status(c, codes.ServiceError, err)
	}

	result := new(AttributeResult)
	result.List = list
	result.Count = count
	result.Config = config

	return Status(c, codes.Success, result)
}

func GetUserAttributeLog(c echo.Context) error {
	userService, userClient := common.GetBisaleUserServiceClient()
	defer common.BisaleUserServicePool.Put(userClient)

	log, traceId := common.GetLoggerWithTraceId(c)

	userId, err := strconv.ParseInt(c.QueryParam("userId"),10, 32)
	if err != nil {
		return Status(c, codes.ServiceError, err)
	}

	result, err := userService.SelectUserAttributeLogByUserId(context.Background(), traceId, int32(userId))
	if err != nil {
		log.Error(err)
		return Status(c, codes.ServiceError, err)
	}

	return Status(c, codes.Success, result)
}

func PostUserAttribute(c echo.Context) error {
	userService, userClient := common.GetBisaleUserServiceClient()
	defer common.BisaleUserServicePool.Put(userClient)

	log, traceId := common.GetLoggerWithTraceId(c)

	query := new(UserParams)
	if err := c.Bind(query); err != nil {
		log.Error(err)
		return Status(c, codes.ServiceError, err)
	}

	params := new(user.TUserAttributeParam)
	params.IsSpecial = query.IsSpecial
	params.IsTrade = query.IsTrade
	params.UserId = query.UserId
	params.Email = query.Email
	params.Mobile = query.Mobile
	params.StartPage = query.Page
	params.PageSize = query.Size

	_, err := userService.UpdateUserAttributeByUserId(context.Background(), traceId, params)
	if err != nil {
		log.Error(err)
		return Status(c, codes.ServiceError, err)
	}

	return Status(c, codes.Success, nil)
}

func PostSystemConfig(c echo.Context) error {
	systemService, systemClient := common.GetBisaleSystemServiceClient()
	defer common.BisaleSystemServicePool.Put(systemClient)

	log, traceId := common.GetLoggerWithTraceId(c)

	params := new(SystemParams)
	if err := c.Bind(params); err != nil {
		log.Error(err)
		return Status(c, codes.ServiceError, err)
	}

	_, error := systemService.UpdateSystemConfigByJsonValue(context.Background(), traceId, params.Payload)
	if error != nil {
		log.Error(error)
		return Status(c, codes.ServiceError, error)
	}

	return Status(c, codes.Success, nil)
}
