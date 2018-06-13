package middlewares

import (
	"time"
	"strings"
	"crypto/md5"
	"encoding/hex"
	"github.com/labstack/echo"
	"github.com/go-redis/redis"
	"bisale/bisale-console-api/config"
	"bisale/bisale-console-api/common"
	"bisale/bisale-console-api/codes"
	"bisale/bisale-console-api/controllers"
	"github.com/sirupsen/logrus"
)

func hashString(s ...string) string {
	hasher := md5.New()
	hasher.Write([]byte(strings.Join(s, "")))
	return hex.EncodeToString(hasher.Sum(nil))
}

func FilterRequestsWithIp(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if c.Request().Method == echo.POST || c.Request().Method == echo.PUT {
			hash := c.Request().Header.Get("X-Request-Hash")
			ip := c.RealIP()
			cacheKey := hashString(config.Config.Redis.CachePrefix, "ip-filter", ip, hash)
			_, err := common.Cache.Get(cacheKey).Result()

			if err != nil && err != redis.Nil {
				return controllers.Status(c, codes.CacheError, err)
			}

			if err != redis.Nil {
				return controllers.Status(c, codes.RepeatRequestWithIp, err)
			}
			common.Log.WithFields(logrus.Fields{
				"ip":        ip,
				"cache-key": hash,
			}).Info("Filter request with ip")
			if err := common.Cache.Set(cacheKey, "", 2*time.Second).Err(); err != nil {
				common.Log.Error(err)
			}
		}
		if err := next(c); err != nil {
			c.Error(err)
		}
		return nil
	}
}

func FilterRequestsStrict(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		hash := c.Request().Header.Get("X-Request-Hash")
		cacheKey := hashString(config.Config.Redis.CachePrefix, "strict-filter", hash)
		if c.Request().Method == echo.POST || c.Request().Method == echo.PUT {
			_, err := common.Cache.Get(cacheKey).Result()

			if err != nil && err != redis.Nil {
				return controllers.Status(c, codes.CacheError, err)
			}

			if err != redis.Nil {
				return controllers.Status(c, codes.RepeatRequestStrict, err)
			}
			common.Cache.Set(cacheKey, "", 60*time.Second)
		}
		if err := next(c); err != nil {
			c.Error(err)
		}
		if c.Request().Method == echo.POST || c.Request().Method == echo.PUT {
			common.Cache.Del(cacheKey)
		}
		return nil
	}
}

func FilterIteratedRequests(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if c.Request().Method == echo.POST || c.Request().Method == echo.PUT {

			hash := c.Request().Header.Get("X-Request-Hash")
			token := c.Request().Header.Get("X-Access-Token")
			ip := c.RealIP()
			cacheKey := hashString(config.Config.Redis.CachePrefix, "token-filter", ip, token, hash)
			_, err := common.Cache.Get(cacheKey).Result()

			if err != nil && err != redis.Nil {
				return controllers.Status(c, codes.CacheError, err)
			}

			if err != redis.Nil {
				return controllers.Status(c, codes.RepeatRequestWithToken, err)
			}
			common.Cache.Set(cacheKey, "", 2*time.Second)
			if err := next(c); err != nil {
				c.Error(err)
			}
		}
		// common.Cache.Del(cacheKey)
		return nil
	}
}
