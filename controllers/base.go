package controllers

import (
	"net/http"
	"encoding/json"
	"github.com/labstack/echo"
	"bisale/bisale-console-api/common"
	"bisale/bisale-console-api/locales"
)

type Result struct {
	Code    int32       `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (s Result) String() string {
	r, e := json.Marshal(s)
	if e != nil {
		panic(e)
	}
	return string(r)
}

func Status(context echo.Context, code int32, data interface{}) error {

	return Response(context, Result{
		Code: code,
		Data: data,
	})
}

func Response(c echo.Context, r Result) error {

	lang := c.Request().Header.Get("X-Accept-Lang")

	if lang == "" {
		lang = "zh-CN"
	}

	if m := locales.Get(lang, r.Code); m != "" {
		r.Message = m
	}

	if r.Data == nil || r.Data == "" {
		r.Data = struct{}{}
	}

	resultJson, _ := json.Marshal(r)

	c.Set("result-json", string(resultJson))

	if r.Code != 200 {
		if err, ok := r.Data.(error); ok {
			r.Data = err.Error()
		}
		c.Response().Header().Set("X-Trace-Position",common.GetCodePosition(3))
	}

	if r.Code < 520 {
		return c.JSON(int(r.Code), r)
	}

	return c.JSON(http.StatusOK, r)
}
