package middlewares

import (
	"github.com/labstack/echo"
	"bytes"
	"io/ioutil"
	"crypto/md5"
	"encoding/hex"
	"bisale/bisale-console-api/common"
	"github.com/sirupsen/logrus"
)

func RequestHash(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var data string
		req := c.Request()
		url := c.Request().URL.String()
		if req.ContentLength != 0 {
			body, _ := ioutil.ReadAll(req.Body)
			req.Body = ioutil.NopCloser(bytes.NewBuffer(body))
			data = string(body)
		}

		hasher := md5.New()
		hasher.Write([]byte(url + data))
		hash := hex.EncodeToString(hasher.Sum(nil))
		c.Request().Header.Set("X-Request-Hash", hash)
		common.Log.WithFields(logrus.Fields{
			"url":  url,
			"hash": hash,
		})
		if err := next(c); err != nil {
			c.Error(err)
		}
		return nil
	}
}
