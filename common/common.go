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
var BisaleUserServicePool *thriftPool.ThriftPool

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

	// 配置 Bisale User 服务连接池
	BisaleUserServicePool = thriftPool.NewThriftPool(
		config.Config.BisaleUserService.Host,
		config.Config.BisaleUserService.Port,
		config.Config.BisaleUserService.MaxConn,
		config.Config.BisaleUserService.ConnTimeout,
		config.Config.BisaleUserService.IdleTimeout,
		openBisaleServiceClient,
		closeBisaleServiceClient,
	)
}

func GetLoggerWithTraceId(c echo.Context) *logrus.Entry {
	return Log.WithFields(logrus.Fields{
		"trace-id": c.Request().Header.Get("Trace-id"),
	})
}
