package controllers

import (
	"net/http"
	"encoding/json"
	"github.com/labstack/echo"
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

func Status(c echo.Context, code int32, data interface{}) error {
	return Response(c, Result{
		Code: code,
		Data: data,
	})
}

func Response(c echo.Context, s Result) error {

	lang := c.Request().Header.Get("X-Accept-Lang")

	if lang == "" {
		lang = "zh-CN"
	}

	if m := locales.Get(lang, s.Code); m != "" {
		s.Message = m
	}

	if s.Data == nil || s.Data == "" {
		s.Data = struct{}{}
	}

	if s.Code < 1000 {
		return c.JSON(int(s.Code), s)
	}

	return c.JSON(http.StatusOK, s)
}
