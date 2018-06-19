package controllers

import (
	"github.com/labstack/echo"
	"bisale/bisale-console-api/codes"
	"bisale/bisale-console-api/common"
	"context"
	"bisale/bisale-console-api/config"
	"github.com/satori/go.uuid"
	"strconv"
	"encoding/json"
	"bisale/bisale-console-api/thrift/content"
)

type Banner struct {
	Id int32 `json:"id"`
	AppId string `json:"appId"`
	TemplateId int32 `json:"templateId"`
	Title string `json:"title"`
	Status string `json:"status"`
	Sort int32 `json:"sort"`
	Lang string `json:"lang"`
	LinkAddress string `json:"linkAddress"`
	PhotoAddress string `json:"photoAddress"`
	UpAt string `json:"createdAt"`
	DownAt string `json:"endAt"`
}

type Notice struct {
	Id int32 `json:"id"`
	Channel int32 `json:"channel"`
	Lang string `json:"lang"`
	Title string `json:"title"`
	LinkUrl string `json:"linkUrl"`
	Status int32 `json:"status"`
	Sort int32 `json:"sort"`
	UpTime string `json:"upTime"`
	DownTime string `json:"downTime"`
}

type StatusChange struct {
	Id int32 `json:"id"`
	Status int32 `json:"status"`
}

func GetImageMeta(c echo.Context) error {
	log, traceId := common.GetLoggerWithTraceId(c)
	storageService, storageClient := common.GetStorageServiceClient()
	defer common.StorageServicePool.Put(storageClient)

	size, _ := strconv.ParseInt(config.Config.BannerSize, 10, 32)
	dir, _ := uuid.NewV4()
	file, _ := uuid.NewV4()

	token, err := storageService.GeneratePolicyToken(context.Background(), traceId, config.Config.BannerBucket, dir.String(), int32(size), config.Config.BannerExpired)

	if err != nil {
		log.Error(err)
		return Status(c, codes.ServiceError, err)
	}

	var meta map[string]interface{}
	jsonError := json.Unmarshal([]byte(token), &meta)
	if jsonError != nil {
		log.Error(jsonError)
		return Status(c, codes.ServiceError, jsonError)
	}

	json.Unmarshal([]byte(token), &meta)
	meta["file"] = file.String()
	log.Info(meta)

	return Status(c, codes.Success, meta)
}

func GetImageUrl(c echo.Context) error {
	log, traceId := common.GetLoggerWithTraceId(c)
	storageService, storageClient := common.GetStorageServiceClient()
	defer common.StorageServicePool.Put(storageClient)

	path := c.QueryParam("path")

	url, err := storageService.GetUrl(context.Background(), traceId, config.Config.BannerBucket, path, 60)

	if (err != nil) {
		log.Error(err)
		return Status(c, codes.ServiceError, err)
	}

	return Status(c, codes.Success, url)

}

func GetBannerList(c echo.Context) error {
	page, _ := strconv.ParseInt(c.QueryParam("page"), 10, 32)
	size, _ := strconv.ParseInt(c.QueryParam("size"), 10, 32)

	if size < 10 {
		size =  10
	}

	log, traceId := common.GetLoggerWithTraceId(c)
	contentService, contentClient := common.GetBisaleContentServiceClient()
	defer common.BisaleContentServicePool.Put(contentClient)

	bannerList, err := contentService.SelectBannerList(context.Background(), traceId, int32(page), int32(size))

	if err != nil {
		log.Error(err)
		return Status(c, codes.ServiceError, err)
	}

	return Status(c, codes.Success, bannerList)
}

func GetBannerListCount(c echo.Context) error {

	log, traceId := common.GetLoggerWithTraceId(c)
	contentService, contentClient := common.GetBisaleContentServiceClient()
	defer common.BisaleContentServicePool.Put(contentClient)

	result, err := contentService.SelectBannerCount(context.Background(), traceId)

	if err != nil {
		log.Error(err)
		return Status(c, codes.ServiceError, err)
	}

	return Status(c, codes.Success, result)
}

func GetBanner(c echo.Context) error {

	log, traceId := common.GetLoggerWithTraceId(c)
	contentService, contentClient := common.GetBisaleContentServiceClient()
	defer common.BisaleContentServicePool.Put(contentClient)

	id, _ :=  strconv.ParseInt(c.QueryParam("id"), 10, 32)

	tb, err := contentService.SelectBannerById(context.Background(), traceId, int32(id))
	if err != nil {
		log.Error(err)
		return Status(c, codes.ServiceError, err)
	}
	return Status(c, codes.Success, tb)

}

func PostBanner(c echo.Context) error {

	log, traceId := common.GetLoggerWithTraceId(c)
	contentService, contentClient := common.GetBisaleContentServiceClient()
	defer common.BisaleContentServicePool.Put(contentClient)

	banner := new(Banner)

	if err := c.Bind(banner); err != nil {
		return Status(c, codes.ServiceError, nil)
	}

	tb := new(content.TBanner)
	tb.AppId = banner.AppId
	tb.Title = banner.Title
	tb.Status = banner.Status
	tb.Sort = banner.Sort
	tb.Lang = "zh-CN"
	tb.LinkAddress = banner.LinkAddress
	tb.PhotoAddress = banner.PhotoAddress
	tb.CreatedAt = banner.UpAt
	tb.EndAt = banner.DownAt

	res, err := contentService.InsertBanner(context.Background(), traceId, tb)
	if err != nil {
		log.Info(err)
		return Status(c, codes.ServiceError, err)
	}

	return Status(c, codes.Success, res)

}

func PatchBannerStatus(c echo.Context) error {

	log, traceId := common.GetLoggerWithTraceId(c)
	contentService, contentClient := common.GetBisaleContentServiceClient()
	defer common.BisaleContentServicePool.Put(contentClient)

	statusChange := new(StatusChange)
	if err := c.Bind(statusChange); err != nil {
		return Status(c, codes.ServiceError, err)
	}

	res, err := contentService.UpdateBannerStatus(context.Background(), traceId, statusChange.Id, statusChange.Status)
	if err != nil {
		log.Info(err)
		return Status(c, codes.ServiceError, err)
	}

	return Status(c, codes.Success, res)

}

func PutBanner(c echo.Context) error {

	log, traceId := common.GetLoggerWithTraceId(c)
	contentService, contentClient := common.GetBisaleContentServiceClient()
	defer common.BisaleContentServicePool.Put(contentClient)

	banner := new(Banner)
	err := c.Bind(banner)
	if err != nil {
		log.Error(err)
		return Status(c, codes.ServiceError, err)
	}

	tb := new(content.TBanner)
	tb.ID = banner.Id
	tb.AppId = banner.AppId
	tb.TemplateId = banner.TemplateId
	tb.Title = banner.Title
	tb.Status = banner.Status
	tb.Sort = banner.Sort
	tb.Lang = "zh-CN"
	tb.LinkAddress = banner.LinkAddress
	tb.PhotoAddress = banner.PhotoAddress
	tb.CreatedAt = banner.UpAt
	tb.EndAt = banner.DownAt

	res, err := contentService.UpdateBanner(context.Background(), traceId, tb)
	if err != nil {
		log.Info(err)
		return Status(c, codes.ServiceError, err)
	}

	return Status(c, codes.Success, res)

}

func GetNotice(c echo.Context) error {
	id, _ := strconv.ParseInt(c.QueryParam("id"), 10, 32)

	log, traceId := common.GetLoggerWithTraceId(c)
	contentService, contentClient := common.GetBisaleContentServiceClient()
	defer common.BisaleContentServicePool.Put(contentClient)

	res, err := contentService.SelectSystemNoticeById(context.Background(), traceId, int32(id))

	if err != nil {
		log.Error(err)
		return Status(c, codes.ServiceError, err)
	}

	return Status(c, codes.Success, res)

}

func GetNoticeList(c echo.Context) error {

	page, _ := strconv.ParseInt(c.QueryParam("page"), 10, 32)
	size, _ := strconv.ParseInt(c.QueryParam("size"), 10, 32)

	if size < 10 {
		size = 10
	}

	log, traceId := common.GetLoggerWithTraceId(c)
	contentService, contentClient := common.GetBisaleContentServiceClient()
	defer common.BisaleContentServicePool.Put(contentClient)

	noticeList, err := contentService.SelectSystemNoticeList(context.Background(), traceId, int32(page), int32(size))
	if err != nil {
		log.Error(err)
		return Status(c, codes.ServiceError, err)
	}
	return Status(c, codes.Success, noticeList)
}

func GetNoticeListCount(c echo.Context) error {

	log, traceId := common.GetLoggerWithTraceId(c)
	contentService, contentClient := common.GetBisaleContentServiceClient()
	defer common.BisaleContentServicePool.Put(contentClient)

	noticeList, err := contentService.SelectSystemNoticeCount(context.Background(), traceId)
	if err != nil {
		log.Error(err)
		return Status(c, codes.ServiceError, err)
	}
	return Status(c, codes.Success, noticeList)
}

func PostNotice(c echo.Context) error {

	log, traceId := common.GetLoggerWithTraceId(c)
	contentService, contentClient := common.GetBisaleContentServiceClient()
	defer common.BisaleContentServicePool.Put(contentClient)

	notice := new(Notice)
	if err := c.Bind(notice); err != nil {
		log.Error(err)
		return Status(c, codes.ServiceError, err)
	}

	tn := new(content.TSystemNotice)
	tn.LinkUrl = notice.LinkUrl
	tn.Title = notice.Title
	tn.Sort = notice.Sort
	tn.Status = notice.Status
	tn.Lang = notice.Lang
	tn.UpTime = notice.UpTime
	tn.DownTime = notice.DownTime

	log.Info(tn)

	res, err := contentService.InsertSystemNotice(context.Background(), traceId, tn)
	if err != nil {
		log.Error(err)
		return Status(c, codes.ServiceError, err)
	}
	return Status(c, codes.Success, res)
}

func PutNotice(c echo.Context) error {

	log, traceId := common.GetLoggerWithTraceId(c)
	contentService, contentClient := common.GetBisaleContentServiceClient()
	defer common.BisaleContentServicePool.Put(contentClient)

	notice := new(Notice)
	if err := c.Bind(notice); err != nil {
		log.Error(err)
		return Status(c, codes.ServiceError, err)
	}

	tn := new(content.TSystemNotice)
	tn.ID = notice.Id
	tn.LinkUrl = notice.LinkUrl
	tn.Sort = notice.Sort
	tn.Title = notice.Title
	tn.Status = notice.Status
	tn.Lang = notice.Lang
	tn.UpTime = notice.UpTime
	tn.DownTime = notice.DownTime

	res, err := contentService.UpdateSystemNotice(context.Background(), traceId, tn)
	if err != nil {
		log.Error(err)
		return Status(c, codes.ServiceError, err)
	}
	return Status(c, codes.Success, res)
}

func PatchNoticeStatus(c echo.Context) error {

	log, traceId := common.GetLoggerWithTraceId(c)
	contentService, contentClient := common.GetBisaleContentServiceClient()
	defer common.BisaleContentServicePool.Put(contentClient)

	statusChange := new(StatusChange)
	if err := c.Bind(statusChange); err != nil {
		return Status(c, codes.ServiceError, err)
	}

	res, err := contentService.UpdateSystemNoticeStatus(context.Background(), traceId, statusChange.Id, statusChange.Status)
	if err != nil {
		log.Error(err)
		return Status(c, codes.ServiceError, err)
	}
	return Status(c, codes.Success, res)

}