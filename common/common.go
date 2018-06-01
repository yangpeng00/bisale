package common

import (
	"encoding/json"
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
