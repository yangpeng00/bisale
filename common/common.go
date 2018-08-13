package common

import (
	"os"
	"fmt"
	"github.com/sirupsen/logrus"
	"bisale/bisale-console-api/config"
	"bisale/foundation/thrift/pool"
	"github.com/labstack/echo"
	"github.com/go-redis/redis"
	"runtime"
)

var Log *logrus.Logger
var AccountServicePool *thriftPool.ThriftPool
var InvitationServicePool *thriftPool.ThriftPool
var MessageServicePool *thriftPool.ThriftPool
var CaptchaServicePool *thriftPool.ThriftPool
var StorageServicePool *thriftPool.ThriftPool
var BisaleContentServicePool *thriftPool.ThriftPool
var BisaleUserKycServicePool *thriftPool.ThriftPool
var BisaleOrderServicePool *thriftPool.ThriftPool
var BisaleSystemServicePool *thriftPool.ThriftPool
var BisaleWithdrawServicePool *thriftPool.ThriftPool
var BisaleBusinessServicePool *thriftPool.ThriftPool
var BisaleUserServicePool *thriftPool.ThriftPool
var WalletServicePool *thriftPool.ThriftPool

var Cache *redis.Client

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

	// 配置缓存
	Cache = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", config.Config.Redis.Host, config.Config.Redis.Port),
		Password: config.Config.Redis.Password,
	})

	if err := Cache.Ping().Err(); err != nil {
		fmt.Printf("Ping redis error: %s", err)
		os.Exit(1)
	}

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
	BisaleUserKycServicePool = thriftPool.NewThriftPool(
		config.Config.BisaleOmsService.Host,
		config.Config.BisaleOmsService.Port,
		config.Config.BisaleOmsService.MaxConn,
		config.Config.BisaleOmsService.ConnTimeout,
		config.Config.BisaleOmsService.IdleTimeout,
		openBisaleUserKycServiceClient,
		closeBisaleUserKycServiceClient,
	)

	BisaleUserServicePool = thriftPool.NewThriftPool(
		config.Config.BisaleOmsService.Host,
		config.Config.BisaleOmsService.Port,
		config.Config.BisaleOmsService.MaxConn,
		config.Config.BisaleOmsService.ConnTimeout,
		config.Config.BisaleOmsService.IdleTimeout,
		openBisaleUserServiceClient,
		closeBisaleUserServiceClient,
	)

	BisaleContentServicePool = thriftPool.NewThriftPool(
		config.Config.BisaleOmsService.Host,
		config.Config.BisaleOmsService.Port,
		config.Config.BisaleOmsService.MaxConn,
		config.Config.BisaleOmsService.ConnTimeout,
		config.Config.BisaleOmsService.IdleTimeout,
		openBisaleContentServiceClient,
		closeBisaleContentServiceClient,
	)

	BisaleOrderServicePool = thriftPool.NewThriftPool(
		config.Config.BisaleOmsService.Host,
		config.Config.BisaleOmsService.Port,
		config.Config.BisaleOmsService.MaxConn,
		config.Config.BisaleOmsService.ConnTimeout,
		config.Config.BisaleOmsService.IdleTimeout,
		openBisaleOrderServiceClient,
		closeBisaleOrderServiceClient,
	)

	BisaleSystemServicePool = thriftPool.NewThriftPool(
		config.Config.BisaleOmsService.Host,
		config.Config.BisaleOmsService.Port,
		config.Config.BisaleOmsService.MaxConn,
		config.Config.BisaleOmsService.ConnTimeout,
		config.Config.BisaleOmsService.IdleTimeout,
		openBisaleSystemServiceClient,
		closeBisaleSystemServiceClient,
	)

	// 配置 Bisale Withdraw 服务连接池
	BisaleWithdrawServicePool = thriftPool.NewThriftPool(
		config.Config.BisaleOmsService.Host,
		config.Config.BisaleOmsService.Port,
		config.Config.BisaleOmsService.MaxConn,
		config.Config.BisaleOmsService.ConnTimeout,
		config.Config.BisaleOmsService.IdleTimeout,
		openBisaleWithdrawServiceClient,
		closeBisaleWithdrawServiceClient,
	)

	// 配置 Bisale Business 服务连接池
	BisaleBusinessServicePool = thriftPool.NewThriftPool(
		config.Config.BisaleBusinessService.Host,
		config.Config.BisaleBusinessService.Port,
		config.Config.BisaleBusinessService.MaxConn,
		config.Config.BisaleBusinessService.ConnTimeout,
		config.Config.BisaleBusinessService.IdleTimeout,
		openBisaleBusinessServiceClient,
		closeBisaleBusinessServiceClient,
	)

	WalletServicePool = thriftPool.NewThriftPool(
		config.Config.WalletService.Host,
		config.Config.WalletService.Port,
		config.Config.WalletService.MaxConn,
		config.Config.WalletService.ConnTimeout,
		config.Config.WalletService.IdleTimeout,
		openWalletServiceClient,
		closeWalletServiceClient,
	)
}

func GetLoggerWithTraceId(c echo.Context) (*logrus.Entry, string) {
	traceId := c.Request().Header.Get("X-Trace-Id")
	return Log.WithFields(logrus.Fields{
		"trace-id": traceId,
	}), traceId
}

func GetCodePosition(skip int) string {
	_, file, line, ok := runtime.Caller(skip)
	if ok {
		return fmt.Sprintf("%s:%d", file, line)
	}
	return ""
}