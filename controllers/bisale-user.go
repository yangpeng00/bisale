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
	"crypto/hmac"
	"crypto/sha1"
	"bisale/bisale-console-api/domain"
	"github.com/parnurzeal/gorequest"
	"encoding/hex"
	"encoding/base64"
	"github.com/satori/go.uuid"
	"bisale/bisale-console-api/config"
)

type UserIdRequest struct {
	Id int32 `json:"id"`
}

type EmailRequest struct {
	Email string `json:"email"`
}

// util
func Sha1(strMessage string) string {
	ctx := sha1.New()
	ctx.Write([]byte(strMessage))
	return base64.StdEncoding.EncodeToString(ctx.Sum(nil))
}

func HmacSha1Signature(strMessage string, strSecret string) string {
	key := []byte(strSecret)
	h := hmac.New(sha1.New, key)
	h.Write([]byte(strMessage))
	return hex.EncodeToString(h.Sum(nil))
	//return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

// router
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

func GetDepositAddressById(c echo.Context) error {
	id, _ := strconv.ParseInt(c.QueryParam("id"), 10, 32)
	log, _ := common.GetLoggerWithTraceId(c)

	walletService, walletClient := common.GetWalletServiceClient()
	defer common.WalletServicePool.Put(walletClient)

	config := make(map[string]interface{})
	config["user_id"] = id
	config["currency"] = "*"
	config["address_type"] = "1"
	configStr, _ := json.Marshal(config)
	fmt.Println(string(configStr))

	result, err := walletService.Execute(context.Background(),"Address", "getDepositAddress", string(configStr))

	if err != nil {
		log.Error(err)
		return Status(c, codes.ServiceError, err)
	}

	return Status(c, codes.Success, result)
}

func GetAccountStatusById(c echo.Context) error {
	log, _ := common.GetLoggerWithTraceId(c)

	account := c.QueryParam("account")
	key := c.QueryParam("key")

	accountInfo := new(domain.AccountInfo)
	accountInfo.MsgType = "GetAccountInfoRequest"

	crId, _ :=  uuid.NewV4()
	accountInfo.CRID = crId.String()

	accountInfo.Account = account
	accountInfo.Date = "20180709"

	serialMessage := accountInfo.Serialize()
	//hmac ,use sha1

	hashKey := Sha1(key)
	accountInfo.SIG = HmacSha1Signature(serialMessage, hashKey)

	accountInfoStr, err := json.Marshal(accountInfo)

	if err != nil {
		log.Error(err)
		return Status(c, codes.ServiceError, err)
	}

	_, str, error := gorequest.New().
		Post(config.Config.EngineUrl).
		Send(string(accountInfoStr)).
		Set("Accept", "application/json").
		End()

	if error != nil {
		return Status(c, codes.ServiceError, error)
	}

	return Status(c, codes.Success, str)

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
