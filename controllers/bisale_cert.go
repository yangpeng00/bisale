package controllers

import (
	"github.com/labstack/echo"
	"bisale/bisale-console-api/codes"
	"bisale/bisale-console-api/common"
	"context"
	"strconv"
	"github.com/labstack/gommon/log"
	"bisale/bisale-console-api/domain"
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
	userService := common.GetBisaleServiceClient()

	ctx := context.Background()
	res, err := userService.SelectUserKycByConditions(ctx, keyword, status, int32(page), int32(size))
	if err != nil {
		log.Error(err)
		return Status(c, codes.InteralServerError, nil)
	}
	return Status(c, codes.Success, res)
}

func GetCertDetailById(c echo.Context) error {
	id, _ := strconv.ParseInt(c.QueryParam("id"), 10, 32)

	log, _ := common.GetLoggerWithTraceId(c)
	userService := common.GetBisaleServiceClient()

	res, err := userService.SelectUserKycById(context.Background(), int32(id))
	if err != nil {
		log.Error(err)
		return Status(c, codes.InteralServerError, nil)
	}
	return Status(c, codes.Success, res)
}

func PostCertResult(c echo.Context) error {
	req := new(domain.PostCertRequest)
	if err := c.Bind(req); err != nil {
		return Status(c, codes.InteralServerError, nil)
	}

	log, _ := common.GetLoggerWithTraceId(c)
	userService := common.GetBisaleServiceClient()

	resp, err := userService.AuditUserKyc(context.Background(), req.Id, req.Status, req.Mark)
	if resp == 0 || err != nil {
		log.Error(err)
		return Status(c, codes.CertError, nil)
	}
	return Status(c, codes.Success, nil)

}

func GetCertListCount(c echo.Context) error {
	keyword := c.QueryParam("keyword")
	status := c.QueryParam("status")

	userService := common.GetBisaleServiceClient()
	res, err := userService.SelectCountByConditions(context.Background(), keyword, status)
	if err != nil {
		log.Error(err)
		return Status(c, codes.InteralServerError, nil)
	}

	return Status(c, codes.Success, res)
}