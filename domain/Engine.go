package domain

type ReqType interface {
	Serialize() string
}

type ReqBase struct {
	MsgType string
	CRID    string
	Date    string
	Account string
}

type EngineReqCommon struct {
	MsgType         string
	CRID            string
	Date            string
	Account         string
	ApiKey          string
	SIG             string
	TimeStamp       int64  //心跳
	Symbol          string //推送orderbook
	OriginalMsgType string //错误响应
	ConnId          string //请求客户端id
}

type AccountInfo struct {
	ReqBase
	SIG string
}

func (a *AccountInfo) Serialize() string {
	serializeStr := a.MsgType + a.CRID + a.Date + a.Account
	return serializeStr
}