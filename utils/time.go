package utils

import "time"

func GetCurrentTimestamp () int64 {
	return time.Now().UnixNano() / 1e6
}
