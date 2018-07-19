package controllers

import (
	"context"
	"strconv"
	"github.com/labstack/echo"
	"bisale/bisale-console-api/codes"
	"bisale/bisale-console-api/common"
	"bisale/bisale-console-api/config"
	"bisale/bisale-console-api/domain"
	"strings"
	"github.com/sirupsen/logrus"
	"fmt"
	"bisale/bisale-console-api/thrift/user"
	"encoding/json"
)

func GetCertList(c echo.Context) error {
	log, traceId := common.GetLoggerWithTraceId(c)
	userService, userClient := common.GetBisaleUserKycServiceClient()
	defer common.BisaleUserKycServicePool.Put(userClient)

	userId, _ := strconv.ParseInt(c.QueryParam("userId"), 10, 32)
	page, _ := strconv.ParseInt(c.QueryParam("page"), 10, 32)
	size, _ := strconv.ParseInt(c.QueryParam("size"), 10, 32)

	userParams := new(user.TUserKycParams)
	userParams.TraceId = traceId
	userParams.StartPage = int32(page)
	userParams.PageSize = int32(size)
	userParams.UserId = int32(userId)
	userParams.UserName = c.QueryParam("username")
	userParams.Mobile = c.QueryParam("mobile")
	userParams.Email = c.QueryParam("email")
	userParams.Status = c.QueryParam("status")

	xxx, _ := json.Marshal(userParams)
	fmt.Println("======")
	fmt.Println(string(xxx))

	ctx := context.Background()

	res, err := userService.SelectUserKycByConditions(ctx, userParams)
	log.Info(res)
	if err != nil {
		log.Error(err)
		return Status(c, codes.ServiceError, err)
	}
	return Status(c, codes.Success, res)
}

func GetCertDetailById(c echo.Context) error {
	id, _ := strconv.ParseInt(c.QueryParam("id"), 10, 32)

	log, traceId := common.GetLoggerWithTraceId(c)

	userService, userClient := common.GetBisaleUserKycServiceClient()
	defer common.BisaleUserKycServicePool.Put(userClient)

	ctx := context.Background()
	res, err := userService.SelectUserKycById(ctx, traceId, int32(id))
	if err != nil {
		log.Error(err)
		return Status(c, codes.ServiceError, err)
	}

	storageService, storageClient := common.GetStorageServiceClient()
	defer common.StorageServicePool.Put(storageClient)
	if res == nil {
		return Status(c, codes.MemberNotExist, "")
	}
	if strings.HasPrefix(res.IdPicFront, "U/") ||
		strings.HasPrefix(res.IdPicBack, "U/") ||
		strings.HasPrefix(res.IdPicHold, "U/") ||
		strings.HasPrefix(res.PassportPicFront, "U/") ||
		strings.HasPrefix(res.PassportPicInfo, "U/") ||
		strings.HasPrefix(res.PassportPicHold, "U/") {

		images, err := storageService.GetProcessUrls(ctx, traceId, config.Config.KYCBucket, map[string]string{
			"IdPicFront":       res.IdPicFront,
			"IdPicBack":        res.IdPicBack,
			"IdPicHold":        res.IdPicHold,
			"PassportPicFront": res.PassportPicFront,
			"PassportPicInfo":  res.PassportPicInfo,
			"PassportPicHold":  res.PassportPicHold,
		}, "", 60)

		if err != nil {
			log.Error(err)
			return Status(c, codes.ServiceError, err)
		}

		res.IdPicFront = images["IdPicFront"]
		res.IdPicBack = images["IdPicBack"]
		res.IdPicHold = images["IdPicHold"]
		res.PassportPicFront = images["PassportPicFront"]
		res.PassportPicInfo = images["PassportPicInfo"]
		res.PassportPicHold = images["PassportPicHold"]

	} else {
		res.IdPicFront = config.Config.OldKYCHost + res.IdPicFront
		res.IdPicBack = config.Config.OldKYCHost + res.IdPicBack
		res.IdPicHold = config.Config.OldKYCHost + res.IdPicHold
		res.PassportPicFront = config.Config.OldKYCHost + res.PassportPicFront
		res.PassportPicInfo = config.Config.OldKYCHost + res.PassportPicInfo
		res.PassportPicHold = config.Config.OldKYCHost + res.PassportPicHold
	}

	return Status(c, codes.Success, res)
}

func splitByLine(arg string) string {
	arr := strings.Split(arg, "-")
	length := len(arr)
	return arr[length - 1]
}

func PostCertResult(c echo.Context) error {
	req := new(domain.PostCertRequest)
	if err := c.Bind(req); err != nil {
		return Status(c, codes.ServiceError, err)
	}

	log, traceId := common.GetLoggerWithTraceId(c)

	userService, userClient := common.GetBisaleUserKycServiceClient()
	defer common.BisaleUserKycServicePool.Put(userClient)

	businessService, businessClient := common.GetBisaleBusinessServiceClient()
	defer common.BisaleBusinessServicePool.Put(businessClient)

	resp, err := userService.AuditUserKyc(context.Background(), traceId, req.Id, req.Status, req.Mark, req.UserId)
	messageService, messageClient := common.GetMessageServiceClient()
	defer common.MessageServicePool.Put(messageClient)

	ctx := context.Background()
	log.Info(fmt.Sprintf("The request status is %s", req.Status))
	if req.Status == "2" {
		err := businessService.EnableParticipant(context.Background(), traceId, req.UserId)
		if err != nil {
			log.Error(err)
			log.WithFields(logrus.Fields{
				"user_id": req.UserId,
				"err":     err.Error(),
			}).Error("邀请糖果发送失败")
			return Status(c, codes.ServiceError, err)
		}
		log.WithFields(logrus.Fields{
			"user_id": req.UserId,
		}).Info("邀请糖果发送成功")
		if resp != nil {
			if resp.Email != "" {
				err := messageService.SendMail(ctx, traceId, config.Config.KycSuccessMail.AppId, resp.Email, config.Config.KycSuccessMail.TemplateId, "{\"username\":"+"\""+resp.Email+"\"}", "zh-CN", 0)
				if err != nil {
					log.WithFields(logrus.Fields{
						"user_id": req.UserId,
					}).Error("KYC邮件发送失败", err)
				} else {
					log.WithFields(logrus.Fields{
						"user_id": req.UserId,
					}).Info("KYC邮件发送成功")
				}
			} else {
				err := messageService.SendSMS(ctx, traceId, config.Config.KycFailedSMS.AppId, resp.Mobile, config.Config.KycSuccessSMS.TemplateId, "{\"username\":"+"\""+splitByLine(resp.Mobile)+"\"}", "zh-CN", 0)
				if err != nil {
					log.WithFields(logrus.Fields{
						"user_id": req.UserId,
					}).Error("KYC短信发送失败", err)
				} else {
					log.WithFields(logrus.Fields{
						"user_id": req.UserId,
					}).Info("KYC短信发送成功")
				}
			}
		} else {
			log.WithFields(logrus.Fields{
				"user_id": req.UserId,
			}).Error("KYC审核服务返回数据错误，邮件未发送")
		}
	} else {
		if resp != nil {
			if resp.Email != "" {
				err := messageService.SendMail(ctx, traceId, config.Config.KycFailedMail.AppId, resp.Email, config.Config.KycFailedMail.TemplateId, "{\"username\":"+"\""+resp.Email+"\",\"reason\":\""+req.Mark+"\"}", "zh-CN", 0)
				if err != nil {
					log.WithFields(logrus.Fields{
						"user_id": req.UserId,
					}).Error("KYC邮件发送失败", err)
				} else {
					log.WithFields(logrus.Fields{
						"user_id": req.UserId,
					}).Info("KYC邮件发送成功")
				}
			} else {
				err := messageService.SendSMS(ctx, traceId, config.Config.KycFailedSMS.AppId, resp.Mobile, config.Config.KycFailedSMS.TemplateId, "{\"username\":"+"\""+splitByLine(resp.Mobile)+"\",\"reason\":\""+req.Mark+"\"}", "zh-CN", 0)
				if err != nil {
					log.WithFields(logrus.Fields{
						"user_id": req.UserId,
					}).Error("KYC短信发送失败", err)
				} else {
					log.WithFields(logrus.Fields{
						"user_id": req.UserId,
					}).Info("KYC短信发送成功")
				}
			}
		} else {
			log.WithFields(logrus.Fields{
				"user_id": req.UserId,
			}).Error("KYC审核服务返回数据错误，邮件未发送")
		}
	}
	if err != nil {
		log.Error(err)
		return Status(c, codes.CertError, err)
	}
	return Status(c, codes.Success, nil)

}

func GetCertListCount(c echo.Context) error {
	log, traceId := common.GetLoggerWithTraceId(c)

	userId, _ := strconv.ParseInt(c.QueryParam("userId"), 10, 32)
	userParams := new(user.TUserKycParams)
	userParams.TraceId = traceId
	userParams.UserId = int32(userId)
	userParams.UserName = c.QueryParam("username")
	userParams.Mobile = c.QueryParam("mobile")
	userParams.Email = c.QueryParam("email")
	userParams.Status = c.QueryParam("status")

	userService, userClient := common.GetBisaleUserKycServiceClient()
	defer common.BisaleUserKycServicePool.Put(userClient)

	res, err := userService.SelectUserKycCountByConditions(context.Background(), userParams)
	if err != nil {
		log.Error(err)
		return Status(c, codes.ServiceError, err)
	}

	return Status(c, codes.Success, res)
}
