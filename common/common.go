package common

import (
	"os"
	"fmt"
	"github.com/sirupsen/logrus"
	"bisale/bisale-console-api/config"
	"bisale/foundation/thrift/pool"
)

var Log *logrus.Logger
var AccountServicePool *thriftPool.ThriftPool

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

	// 配置服务连接池
	AccountServicePool = thriftPool.NewThriftPool(
		config.Config.AccountService.Host,
		config.Config.AccountService.Port,
		config.Config.AccountService.MaxConn,
		config.Config.AccountService.ConnTimeout,
		config.Config.AccountService.IdleTimeout,
		openAccountServiceClient,
		closeAccountServiceClient,
	)
}
