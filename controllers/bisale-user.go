package controllers

import (
	"github.com/labstack/echo"
	"bisale/bisale-console-api/common"
	"bisale/bisale-console-api/codes"
	"bisale/bisale-console-api/thrift/user"
	"context"
	"strconv"
	"encoding/json"
	"fmt"
)

type UserIdRequest struct {
	Id int32 `json:"id"`
}

type EmailRequest struct {
	Email string `json:"email"`
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

func GetWithdrawStatusById(c echo.Context) error {
	id, _ := strconv.ParseInt(c.QueryParam("id"), 10, 32)
	log, _ := common.GetLoggerWithTraceId(c)

	walletService, walletClient := common.GetWalletServiceClient()
	defer common.WalletServicePool.Put(walletClient)

	config := make(map[string]interface{})
	config["user_id"] = id
	config["type"] = 2
	configStr, _ := json.Marshal(config)

	result, err := walletService.Execute(context.Background(),"BlackList", "get", string(configStr))

	if err != nil {
		log.Error(err)
		return Status(c, codes.ServiceError, err)
	}

	return Status(c, codes.Success, result)

}

func PostWithdrawStatusById(c echo.Context) error {
	userRequest := new(UserIdRequest)
	c.Bind(userRequest)
	log, _ := common.GetLoggerWithTraceId(c)

	walletService, walletClient := common.GetWalletServiceClient()
	defer common.WalletServicePool.Put(walletClient)

	config := make(map[string]interface{})
	config["service_id"] = "exchange_dashboard"
	config["token"] = "h6u3w65nbs!@#5tertjjthsrtq4i68k58pr"
	config["currencies"] = "*"
	config["user_id"] = userRequest.Id
	config["type"] = 2
	configStr, _ := json.Marshal(config)

	fmt.Println(string(configStr))

	result, err := walletService.Execute(context.Background(),"BlackList", "update", string(configStr))

	if err != nil {
		log.Error(err)
		return Status(c, codes.ServiceError, err)
	}

	return Status(c, codes.Success, result)

}

func DeleteWithdrawStatusById(c echo.Context) error {
	id, _ := strconv.ParseInt(c.QueryParam("id"), 10, 32)
	log, _ := common.GetLoggerWithTraceId(c)

	walletService, walletClient := common.GetWalletServiceClient()
	defer common.WalletServicePool.Put(walletClient)

	config := make(map[string]interface{})
	config["service_id"] = "exchange_dashboard"
	config["token"] = "h6u3w65nbs!@#5tertjjthsrtq4i68k58pr"
	config["currencies"] = "*"
	config["user_id"] = id
	config["type"] = 2
	configStr, _ := json.Marshal(config)

	result, err := walletService.Execute(context.Background(),"BlackList", "remove", string(configStr))

	if err != nil {
		log.Error(err)
		return Status(c, codes.ServiceError, err)
	}

	return Status(c, codes.Success, result)
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

func PostCaptchaCountChange(c echo.Context) error {
	req := new(EmailRequest)
	if err := c.Bind(req); err != nil {
		return Status(c, codes.ServiceError, err)
	}

	log, traceId := common.GetLoggerWithTraceId(c)
	captchaService, captchaClient := common.GetCaptchaServiceClient()

	defer common.CaptchaServicePool.Put(captchaClient)

	err := captchaService.ClearCount(context.Background(), traceId, req.Email + "login")
	if err != nil {
		log.Error(err)
		return Status(c, codes.ServiceError, err)
	}

	return Status(c, codes.Success, nil)

}
