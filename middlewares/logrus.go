package middlewares

import (
	"io"
	"strconv"
	"time"
	"github.com/sirupsen/logrus"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"bisale/bisale-console-api/common"
)

type LogrusLogger struct {
	*logrus.Logger
}

func (l LogrusLogger) Level() log.Lvl {
	switch l.Logger.Level {
	case logrus.DebugLevel:
		return log.DEBUG
	case logrus.WarnLevel:
		return log.WARN
	case logrus.ErrorLevel:
		return log.ERROR
	case logrus.InfoLevel:
		return log.INFO
	default:
		l.Panic("Invalid level")
	}

	return log.OFF
}

func (l LogrusLogger) SetPrefix(s string) {
	// TODO
}

func (l LogrusLogger) Prefix() string {
	// TODO.  Is this even valid?  I'm not sure it can be translated since
	// logrus uses a Formatter interface.  Which seems to me to probably be
	// a better way to do it.
	return ""
}

func (l LogrusLogger) SetLevel(lvl log.Lvl) {
	//
}


func (l LogrusLogger) Output() io.Writer {
	 return l.Out
}

func (l LogrusLogger) SetOutput(w io.Writer) {
	// common.Log.SetOutput(w)
}

func (l LogrusLogger) Printj(j log.JSON) {
	common.Log.WithFields(logrus.Fields(j)).Print()
}

func (l LogrusLogger) Debugj(j log.JSON) {
	common.Log.WithFields(logrus.Fields(j)).Debug()
}

func (l LogrusLogger) Infoj(j log.JSON) {
	common.Log.WithFields(logrus.Fields(j)).Info()
}

func (l LogrusLogger) Warnj(j log.JSON) {
	common.Log.WithFields(logrus.Fields(j)).Warn()
}

func (l LogrusLogger) Errorj(j log.JSON) {
	common.Log.WithFields(logrus.Fields(j)).Error()
}

func (l LogrusLogger) Fatalj(j log.JSON) {
	common.Log.WithFields(logrus.Fields(j)).Fatal()
}

func (l LogrusLogger) Panicj(j log.JSON) {
	common.Log.WithFields(logrus.Fields(j)).Panic()
}

func (l LogrusLogger) WithFields(j log.JSON) {
	common.Log.WithFields(logrus.Fields(j))
}

func logrusMiddlewareHandler(c echo.Context, next echo.HandlerFunc) error {
	req := c.Request()
	res := c.Response()
	start := time.Now()
	if err := next(c); err != nil {
		c.Error(err)
	}
	stop := time.Now()

	p := req.URL.Path
	if p == "" {
		p = "/"
	}

	bytesIn := req.Header.Get(echo.HeaderContentLength)
	if bytesIn == "" {
		bytesIn = "0"
	}

	common.Log.WithFields(map[string]interface{}{
		// "time_rfc3339":          time.Now().Format(time.RFC3339),
		"trace_id":      req.Header.Get("Trace-ID"),
		"remote_ip":     c.RealIP(),
		"host":          req.Host,
		"uri":           req.RequestURI,
		"method":        req.Method,
		"path":          p,
		"referer":       req.Referer(),
		"user_agent":    req.UserAgent(),
		"status":        res.Status,
		"latency":       strconv.FormatInt(stop.Sub(start).Nanoseconds()/1000, 10),
		"latency_human": stop.Sub(start).String(),
		"bytes_in":      bytesIn,
		"bytes_out":     strconv.FormatInt(res.Size, 10),
	}).Info("Handled request")

	return nil
}

func logger(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		return logrusMiddlewareHandler(c, next)
	}
}

func LogrusHook() echo.MiddlewareFunc {
	return logger
}
