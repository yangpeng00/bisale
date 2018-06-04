package common

import (
	"os"
	"fmt"
	"github.com/sirupsen/logrus"
	"bisale/bisale-console-api/config"
	"bisale/foundation/thrift/pool"
	"github.com/labstack/echo"
)

var Log *logrus.Logger
var AccountServicePool *thriftPool.ThriftPool
var MessageServicePool *thriftPool.ThriftPool
var CaptchaServicePool *thriftPool.ThriftPool
var StorageServicePool *thriftPool.ThriftPool
var BisaleUserServicePool *thriftPool.ThriftPool
var BisaleBusinessServicePool *thriftPool.ThriftPool

func init() {
	// 配置日志
	Log = logrus.New()

	logLevel, err := logrus.ParseLevel(config.Config.LogLevel)

	if err != nil {
		fmt.Printf("Set log level error: %s", err)
		os.Exit(1)
	}

	Log.Formatter = &logrus.JSONFormatter{}
	Log.Out = os.Stdout
	Log.Level = logLevel

	// 配置 Account 服务连接池
	AccountServicePool = thriftPool.NewThriftPool(
		config.Config.AccountService.Host,
		config.Config.AccountService.Port,
		config.Config.AccountService.MaxConn,
		config.Config.AccountService.ConnTimeout,
		config.Config.AccountService.IdleTimeout,
		openAccountServiceClient,
		closeAccountServiceClient,
	)

	// 配置 Captcha 服务连接池
	CaptchaServicePool = thriftPool.NewThriftPool(
		config.Config.CaptchaService.Host,
		config.Config.CaptchaService.Port,
		config.Config.CaptchaService.MaxConn,
		config.Config.CaptchaService.ConnTimeout,
		config.Config.CaptchaService.IdleTimeout,
		openCaptchaServiceClient,
		closeCaptchaServiceClient,
	)

	// 配置 Message 服务连接池
	MessageServicePool = thriftPool.NewThriftPool(
		config.Config.MessageService.Host,
		config.Config.MessageService.Port,
		config.Config.MessageService.MaxConn,
		config.Config.MessageService.ConnTimeout,
		config.Config.MessageService.IdleTimeout,
		openMessageServiceClient,
		closeMessageServiceClient,
	)

	// 配置 Storage 服务连接池
	StorageServicePool = thriftPool.NewThriftPool(
		config.Config.StorageService.Host,
		config.Config.StorageService.Port,
		config.Config.StorageService.MaxConn,
		config.Config.StorageService.ConnTimeout,
		config.Config.StorageService.IdleTimeout,
		openStorageServiceClient,
		closeStroageServiceClient,
	)

	// 配置 Bisale OSS 服务连接池
	BisaleUserServicePool = thriftPool.NewThriftPool(
		config.Config.BisaleUserService.Host,
		config.Config.BisaleUserService.Port,
		config.Config.BisaleUserService.MaxConn,
		config.Config.BisaleUserService.ConnTimeout,
		config.Config.BisaleUserService.IdleTimeout,
		openBisaleUserServiceClient,
		closeBisaleUserServiceClient,
	)

	// 配置 Bisale Business 服务连接池
	BisaleUserServicePool = thriftPool.NewThriftPool(
		config.Config.BisaleBusinessService.Host,
		config.Config.BisaleBusinessService.Port,
		config.Config.BisaleBusinessService.MaxConn,
		config.Config.BisaleBusinessService.ConnTimeout,
		config.Config.BisaleBusinessService.IdleTimeout,
		openBisaleBusinessServiceClient,
		closeBisaleBusinessServiceClient,
	)
}

func GetLoggerWithTraceId(c echo.Context) (*logrus.Entry, string) {
	traceId := c.Request().Header.Get("Trace-id")
	return Log.WithFields(logrus.Fields{
		"trace-id": traceId,
	}), traceId
}
