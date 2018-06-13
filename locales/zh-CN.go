package locales

import "bisale/bisale-console-api/codes"

var zhCN = map[int32]string{
	codes.Success:                "操作成功",
	codes.SendCodeLock60Seconds:  "请60秒后再试",
	codes.ServiceError:           "服务调用错误",
	codes.AccessTokenIsEmpty:     "AccessToken 不能为空",
	codes.AccessTokenIsInvalid:   "请重新登录",
	codes.RepeatRequestStrict:    "正在处理中",
	codes.RepeatRequestWithIp:    "重复点击",
	codes.RepeatRequestWithToken: "操作太频繁",
	codes.FormIsEmpty:            "提交数据不能为空",
	codes.ValidateError:          "数据校验错误",
	codes.SMSCodeError:           "验证码错误",
	codes.UserNotExist:           "用户不存在",
}
