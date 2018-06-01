package controllers

import (
	"net/http"
	"encoding/json"
	"github.com/labstack/echo"
	"bisale/bisale-console-api/locales"
)

type Status struct {
	Code    int32       `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (s Status) String() string {
	r, e := json.Marshal(s)
	if e != nil {
		panic(e)
	}
	return string(r)
}

func Response(c echo.Context, s Status) error {

	lang := c.Request().Header.Get("X-Accept-Lang")

	if lang == "" {
		lang = "zh-CN"
	}

	if m := locales.Get(lang, s.Code); m != "" {
		s.Message = m
	}

	if s.Data == nil {
		s.Data = struct{}{}
	}

	if s.Code < 1000 {
		return c.JSON(int(s.Code), s)
	}

	return c.JSON(http.StatusOK, s)
}
