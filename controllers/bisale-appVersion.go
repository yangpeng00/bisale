package controllers

import (
	"github.com/labstack/echo"
	"bisale/bisale-console-api/codes"
	"bisale/bisale-console-api/common"
	"context"
	"bisale/bisale-console-api/thrift/content"
	"strconv"
)

type AppVersionResult struct {
	List []*content.TAppVersion `json:"list"`
	Count int32 `json:"count"`
	LangTypes []string `json:"langTypes"`
	SourceTypes []string `json:"sourceType"`
	StatusTypes []string `json:"statusTypes"`
}

func GetAppVersion(c echo.Context) error {
	log, _ := common.GetLoggerWithTraceId(c)
	appVersionService, appVersionClient := common.GetBisaleAppVersionServiceClient()
	defer common.BisaleAppVersionServicePool.Put(appVersionClient)

	page, _ := strconv.ParseInt(c.QueryParam("page"), 10, 32)
	size, _ := strconv.ParseInt(c.QueryParam("size"), 10, 32)
	version := c.QueryParam("version")
	status := c.QueryParam("status")

	list, err := appVersionService.SelectAppVersions(context.Background(), version, status, int32(page), int32(size))
	count, err := appVersionService.SelectAppVersionCount(context.Background(), version, status)
	langTypes, err := appVersionService.SelectLangTypes(context.Background())
	sourceTypes, err := appVersionService.SelectSourceTypes(context.Background())
	statusTypes, err := appVersionService.SelectStatusTypes(context.Background())
	if err != nil {
		log.Error(err)
		return Status(c, codes.Success, err)
	}
	result := new(AppVersionResult)
	result.List = list
	result.Count = count
	result.LangTypes = langTypes
	result.SourceTypes = sourceTypes
	result.StatusTypes = statusTypes
	return Status(c, codes.Success, result)
}

func PostAppVersion(c echo.Context) error {
	log, _ := common.GetLoggerWithTraceId(c)
	appVersionService, appVersionClient := common.GetBisaleAppVersionServiceClient()
	defer common.BisaleAppVersionServicePool.Put(appVersionClient)

	params := new(content.TAppVersion)
	c.Bind(params)

	_, err := appVersionService.AddAppVersion(context.Background(), params)
	if err != nil {
		log.Error(err)
		return Status(c, codes.ServiceError, err)
	}
	return Status(c, codes.Success, nil)
}

func PutAppVersion(c echo.Context) error {
	log, _ := common.GetLoggerWithTraceId(c)
	appVersionService, appVersionClient := common.GetBisaleAppVersionServiceClient()
	defer common.BisaleAppVersionServicePool.Put(appVersionClient)

	params := new(content.TAppVersion)
	c.Bind(params)

	_, err := appVersionService.EditAppVersion(context.Background(), params)
	if err != nil {
		log.Error(err)
		return Status(c, codes.ServiceError, err)
	}
	return Status(c, codes.Success, nil)
}
