package utils

import (
	"strings"
)
func substring(source string, start int, end int) string {
	var r = []rune(source)
	length := len(r)

	if start < 0 || end > length || start > end {
		return ""
	}

	if start == 0 && end == length {
		return source
	}

	return string(r[start : end])
}



func FormatMobile (str string) string {
	if substring(str, 0, 3) != "+86" {
		return "*******" + substring(str, len(str) - 4, len(str))
	}
	return substring(str, 4, 7) + "***" + substring(str, len(str) - 4, len(str))
}

func FormatEmail(str string) string {
	dataList := strings.Split(str, "@")
	if len(dataList) < 2 {
		return "unformatted email"
	}
	prefix := dataList[0]
	suffix := dataList[1]
	if len(prefix) < 3 {
		return prefix + "*****@" + suffix
	}
	return substring(prefix, 0, 3) + "*****@" + suffix
}