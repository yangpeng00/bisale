package controllers

import (
	"github.com/labstack/echo"
	"bisale/bisale-console-api/common"
	"bisale/bisale-console-api/codes"
	"bisale/bisale-console-api/thrift/user"
	"context"
	"strconv"
)

type UserIdRequest struct {
	Id int32 `json:"id"`
}

func GetUserList(c echo.Context) error {
	page, _ := strconv.ParseInt(c.QueryParam("page"), 10, 32)
	size, _ := strconv.ParseInt(c.QueryParam("size"), 10, 32)

	if size < 10 {
		size = 10
	}

	log, _ := common.GetLoggerWithTraceId(c)
	userService, userClient := common.GetBisaleUserServiceClient()
	defer common.BisaleUserServicePool.Put(userClient)

	userParams := new(user.TUserParams)
	userParams.StartPage = int32(page)
	userParams.PageSize = int32(size)
	userParams.UserName = c.QueryParam("keyword")
	userParams.AccountStatus = c.QueryParam("accountStatus")
	userParams.LoginStatus = c.QueryParam("loginStatus")
	userParams.KycStatus = c.QueryParam("kycStatus")

	res, err := userService.SelectUserByConditions(context.Background(), userParams)
	if (err != nil) {
		log.Error(err)
		return Status(c, codes.ServiceError, err)
	}

	return Status(c, codes.Success, res)
}

func GetUserListCount(c echo.Context) error {

	log, _ := common.GetLoggerWithTraceId(c)
	userService, userClient := common.GetBisaleUserServiceClient()
	defer common.BisaleUserServicePool.Put(userClient)

	userParams := new(user.TUserParams)
	userParams.UserName = c.QueryParam("keyword")
	userParams.AccountStatus = c.QueryParam("accountStatus")
	userParams.LoginStatus = c.QueryParam("loginStatus")
	userParams.KycStatus = c.QueryParam("kycStatus")

	res, err := userService.SelectUserCountByConditions(context.Background(), userParams)
	if err != nil {
		log.Error(err)
		return Status(c, codes.ServiceError, err)
	}

	return Status(c, codes.Success, res)
}

func GetUserDetailById(c echo.Context) error {
	id, _ := strconv.ParseInt(c.QueryParam("id"), 10, 32)


	log, traceId := common.GetLoggerWithTraceId(c)
	userService, userClient := common.GetBisaleUserServiceClient()
	defer common.BisaleUserServicePool.Put(userClient)

	userInfo, err := userService.SelectUserBaseInfoByUserId(context.Background(), traceId, int32(id))

	if err != nil {
		log.Error(err)
		return Status(c, codes.ServiceError, err)
	}

	return Status(c, codes.Success, userInfo)
}

func GetGoogleStatusById(c echo.Context) error {
	id, _ := strconv.ParseInt(c.QueryParam("id"), 10, 32)


	log, traceId := common.GetLoggerWithTraceId(c)
	userService, userClient := common.GetBisaleUserServiceClient()
	defer common.BisaleUserServicePool.Put(userClient)

	googleStatus, err := userService.SelectUserGoogleStatus(context.Background(), traceId, int32(id))

	if err != nil {
		log.Error(err)
		return Status(c, codes.ServiceError, err)
	}

	return Status(c, codes.Success, googleStatus)
}

func PostUserStatusChange(c echo.Context) error {
	req := new(UserIdRequest)
	if err := c.Bind(req); err != nil {
		return Status(c, codes.ServiceError, err)
	}

	log, traceId := common.GetLoggerWithTraceId(c)
	userService, userClient := common.GetBisaleUserServiceClient()
	defer common.BisaleUserServicePool.Put(userClient)

	_, err := userService.UpdateUserStatusByUserId(context.Background(), traceId, req.Id)

	if err != nil {
		log.Error(err)
		return Status(c, codes.ServiceError, err)
	}

	return Status(c, codes.Success, nil)
}

func PostGoogleCodeChange(c echo.Context) error {
	req := new(UserIdRequest)
	if err := c.Bind(req); err != nil {
		return Status(c, codes.ServiceError, err)
	}

	log, traceId := common.GetLoggerWithTraceId(c)
	userService, userClient := common.GetBisaleUserServiceClient()
	defer common.BisaleUserServicePool.Put(userClient)

	_, err := userService.ResetGoogleCode(context.Background(), traceId, req.Id)

	if err != nil {
		log.Error(err)
		return Status(c, codes.ServiceError, err)
	}

	return Status(c, codes.Success, nil)

}
