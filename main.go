package main

import (
	"github.com/labstack/echo"
	"bisale/bisale-console-api/config"
	"github.com/labstack/echo/middleware"
	"bisale/bisale-console-api/controllers"
	"bisale/bisale-console-api/middlewares"
	"github.com/sirupsen/logrus"
	"github.com/go-playground/validator"
)

type Validator struct {
	validator *validator.Validate
}

func (v *Validator) Validate(i interface{}) error {
	return v.validator.Struct(i)
}

func main() {

	e := echo.New()
	e.Validator = &Validator{validator: validator.New()}
	e.Logger = middlewares.LogrusLogger{logrus.StandardLogger()}
	e.Use(middlewares.LogrusHook())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.Use(middlewares.TraceId)
	e.Use(middlewares.RequestHash)
	e.Use(middlewares.FilterIteratedRequests)

	e.GET("/ping", controllers.Ping)

	// auth 路由
	auth := e.Group("/api/auth")
	auth.POST("/login", controllers.PostLogin)
	auth.POST("/login/sms/code", controllers.PostLoginSMSCode)
	auth.POST("/role", controllers.PostCreateRole)

	// member 路由
	member := e.Group("/api/member", middlewares.Auth, middlewares.OperationLog)
	member.POST("", controllers.PostCreateMember)

	// bisale 业务路由
	bisale := e.Group("/api/bisale", middlewares.Auth, middlewares.OperationLog)

	bisale.GET("/cert/list", controllers.GetCertList)
	bisale.GET("/cert/list/count", controllers.GetCertListCount)
	bisale.GET("/cert/detail", controllers.GetCertDetailById)
	bisale.POST("/cert/result", controllers.PostCertResult, middlewares.FilterRequestsStrict)

	bisale.GET("/withdraw/list", controllers.GetWithdrawList)
	bisale.GET("/withdraw/list/count", controllers.GetWithdrawListCount)
	bisale.POST("/withdraw/result", controllers.PostWithdrawResult)

	bisale.GET("/user/list", controllers.GetUserList)
	bisale.GET("/user/list/count", controllers.GetUserListCount)
	bisale.GET("/user/google", controllers.GetGoogleStatusById)
	bisale.GET("/user/detail", controllers.GetUserDetailById)
	bisale.POST("/user/google", controllers.PostGoogleCodeChange)
	bisale.POST("/user/status", controllers.PostUserStatusChange)
	bisale.POST("/user/captcha", controllers.PostCaptchaCountChange)

	bisale.GET("/static/base", controllers.GetStatic)
	bisale.GET("/content/image/meta", controllers.GetImageMeta)
	bisale.GET("/content/image/url", controllers.GetImageUrl)

	bisale.GET("/content/banner", controllers.GetBanner)
	bisale.GET("/content/banner/list", controllers.GetBannerList)
	bisale.GET("/content/banner/list/count", controllers.GetBannerListCount)
	bisale.POST("/content/banner", controllers.PostBanner)
	bisale.PUT("/content/banner", controllers.PutBanner)
	bisale.PATCH("/content/banner/status", controllers.PatchBannerStatus)

	bisale.GET("/content/notice", controllers.GetNotice)
	bisale.GET("/content/notice/list", controllers.GetNoticeList)
	bisale.GET("/content/notice/list/count", controllers.GetNoticeListCount)
	bisale.POST("/content/notice", controllers.PostNotice)
	bisale.PUT("/content/notice", controllers.PutNotice)
	bisale.PATCH("/content/notice/status", controllers.PatchNoticeStatus)

	e.Logger.Fatal(e.Start(config.GetListenNetAddress()))
}
