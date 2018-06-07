package controllers

import (
	"github.com/labstack/echo"
	"bisale/bisale-console-api/common"
	"bisale/bisale-console-api/codes"
	"bisale/bisale-console-api/thrift/user"
	"context"
	"strconv"
)

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

	res, err := userService.SelectUserByConditions(context.Background(), userParams)
	if (err != nil) {
		log.Error(err)
		return Status(c, codes.ServiceError, err)
	}

	return Status(c, codes.Success, res)
}

func GetUserListCount(c echo.Context) error {
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

	res, err := userService.SelectUserCountByConditions(context.Background(), userParams)
	if (err != nil) {
		log.Error(err)
		return Status(c, codes.ServiceError, err)
	}

	return Status(c, codes.Success, res)
}
