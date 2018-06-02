package middlewares

import "github.com/labstack/echo"

func FilterRequests(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Request().Header.Set(echo.HeaderXRequestID, "123")
		c.Request().Header.Set("Trace-ID", "123")
		c.Response().Header().Set(echo.HeaderServer, "Bisale Console API/3.0")
		if err := next(c); err != nil {
			c.Error(err)
		}
		return nil
	}
}
