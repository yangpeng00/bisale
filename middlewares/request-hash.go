package middlewares

import (
	"github.com/labstack/echo"
	"bytes"
	"io/ioutil"
	"encoding/json"
	"crypto/md5"
	"encoding/hex"
)

func RequestHash(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var params interface{}
		var data string
		req := c.Request()
		params = c.QueryParams()
		if req.ContentLength != 0 {
			body, _ := ioutil.ReadAll(req.Body)
			req.Body = ioutil.NopCloser(bytes.NewBuffer(body))
			data = string(body)
		}
		paramsJson, _ := json.Marshal(params)
		hasher := md5.New()
		hasher.Write([]byte(string(paramsJson) + data))
		c.Request().Header.Set("X-Request-Hash", hex.EncodeToString(hasher.Sum(nil)))
		if err := next(c); err != nil {
			c.Error(err)
		}
		return nil
	}
}
