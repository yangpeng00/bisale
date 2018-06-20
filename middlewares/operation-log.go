package middlewares

import (
	"bytes"
	"context"
	"io/ioutil"
	"encoding/json"
	"github.com/labstack/echo"
	"bisale/bisale-console-api/common"
	accountInputs "bisale/bisale-console-api/thrift/inputs"
)

func OperationLog(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		var logId int32

		if memberId := c.Get("member_id"); memberId != nil {

			accountService, accountClient := common.GetAccountServiceClient()

			traceId := c.Request().Header.Get("X-Trace-Id")
			agent, _ := json.Marshal(c.Request().UserAgent())

			var input string

			if c.Request().ContentLength != 0 {
				body, _ := ioutil.ReadAll(c.Request().Body)
				c.Request().Body = ioutil.NopCloser(bytes.NewBuffer(body))
				input = string(body)
			}

			var err error

			logId, err = accountService.OperateStart(context.Background(), traceId, &accountInputs.MemberOperationInput{
				OpMethod: c.Request().Method,
				OpUrl:    c.Request().URL.String(),
				OpIp:     c.RealIP(),
				MemberId: memberId.(string),
				OpAgent:  string(agent),
				OpInput:  input,
			})

			common.AccountServicePool.Put(accountClient)

			if err != nil {
				return err
			}
		}

		if err := next(c); err != nil {
			c.Error(err)
		}

		if memberId := c.Get("member_id"); memberId != nil {

			accountService, accountClient := common.GetAccountServiceClient()
			accountService.OperateEnd(context.Background(), logId, &accountInputs.MemberOperationInput{
				OpHttpCode: int32(c.Response().Status),
				OpOutput:   c.Get("result-json").(string),
			})

			common.AccountServicePool.Put(accountClient)
		}

		return nil
	}
}
