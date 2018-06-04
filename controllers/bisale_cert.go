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
)

func GetCertList(c echo.Context) error {
	keyword := c.QueryParam("keyword")
	status := c.QueryParam("status")

	page, _ := strconv.ParseInt(c.QueryParam("page"), 10, 32)
	size, _ := strconv.ParseInt(c.QueryParam("size"), 10, 32)

	if size < 10 {
		size = 10
	}

	log, _ := common.GetLoggerWithTraceId(c)
	userService := common.GetBisaleUserServiceClient()

	ctx := context.Background()

	log.Info("---wawawawawaw---")
	log.Info("key", keyword, "status", status, int32(page), int32(size))
	res, err := userService.SelectUserKycByConditions(ctx, "", keyword, status, int32(page), int32(size))
	log.Info(res)
	if err != nil {
		log.Error(err)
		return Status(c, codes.ServiceError, nil)
	}
	return Status(c, codes.Success, res)
}

func GetCertDetailById(c echo.Context) error {
	id, _ := strconv.ParseInt(c.QueryParam("id"), 10, 32)

	log, traceId := common.GetLoggerWithTraceId(c)
	userService := common.GetBisaleUserServiceClient()
	ctx := context.Background()
	res, err := userService.SelectUserKycById(ctx, "", int32(id))
	if err != nil {
		log.Error(err)
		return Status(c, codes.ServiceError, nil)
	}

	storageService := common.GetStorageServiceClient()
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
			return Status(c, codes.ServiceError, nil)
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

func PostCertResult(c echo.Context) error {
	req := new(domain.PostCertRequest)
	if err := c.Bind(req); err != nil {
		return Status(c, codes.ServiceError, nil)
	}

	log, traceId := common.GetLoggerWithTraceId(c)
	userService := common.GetBisaleUserServiceClient()
	businessService := common.GetBisaleBusinessServiceClient()
	resp, err := userService.AuditUserKyc(context.Background(), "", req.Id, req.Status, req.Mark, req.UserId)
	messageService := common.GetMessageServiceClient()
	ctx := context.Background()
	if req.Status == "2" {
		err := businessService.EnableParticipant(context.Background(), "", req.UserId)
		if err != nil {
			log.Error(err)
		}
		messageService.SendMail(ctx, traceId, "bisale-admin", resp.Email, "template::mail::kyc-success", "{\"username\":"+"\""+resp.Email+"\"}", "zh-CN", 0)
	} else {
		messageService.SendMail(ctx, traceId, "bisale-admin", resp.Email, "template::mail::kyc-failed", "{\"username\":"+"\""+resp.Email+"\"}", "zh-CN", 0)
	}
	if err != nil {
		log.Error(err)
		return Status(c, codes.CertError, nil)
	}
	return Status(c, codes.Success, nil)

}

func GetCertListCount(c echo.Context) error {
	log, _ := common.GetLoggerWithTraceId(c)
	keyword := c.QueryParam("keyword")
	status := c.QueryParam("status")
	userService := common.GetBisaleUserServiceClient()
	res, err := userService.SelectUserKycCountByConditions(context.Background(), "", keyword, status)
	if err != nil {
		log.Error(err)
		return Status(c, codes.ServiceError, err)
	}

	return Status(c, codes.Success, res)
}
