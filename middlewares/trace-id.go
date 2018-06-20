package middlewares

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/labstack/echo"
	"github.com/satori/go.uuid"
)

func TraceId(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		s, _ := uuid.NewV4()

		hasher := md5.New()
		hasher.Write(s.Bytes())
		hash := hex.EncodeToString(hasher.Sum(nil))

		c.Request().Header.Set("X-Trace-Id", hash)

		if err := next(c); err != nil {
			c.Error(err)
		}
		return nil
	}
}
