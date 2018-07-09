package middlewares

import (
	"github.com/labstack/echo"
	"bisale/bisale-console-api/codes"
	"bisale/bisale-console-api/config"
	"bisale/bisale-console-api/common"
	"bisale/bisale-console-api/controllers"
	"context"
)

func Auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		accessToken := c.Request().Header.Get("X-Access-Token")

		if accessToken == "" {
			return controllers.Status(c, codes.AccessTokenIsEmpty, "")
		}

		traceId := c.Request().Header.Get("X-Trace-Id")

		accountService, accountClient := common.GetAccountServiceClient()
		jwtOutput, err := accountService.ValidateJWT(context.Background(), traceId, accessToken, config.Config.JWTToken)
		common.AccountServicePool.Put(accountClient)

		if err != nil {
			return err
		}

		if !jwtOutput.Valid {
			return controllers.Status(c, codes.AccessTokenIsInvalid, "")
		}

		c.Set("member_id", jwtOutput.MemberId)

		if err := next(c); err != nil {
			c.Error(err)
		}

		return nil
	}
}
