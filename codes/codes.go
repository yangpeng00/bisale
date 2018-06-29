package codes

const (
	Success                int32 = 200
	Forbidden              int32 = 403
	ServiceError           int32 = 600
	CacheError             int32 = 601
	RepeatRequestStrict    int32 = 620
	RepeatRequestWithIp    int32 = 621
	RepeatRequestWithToken int32 = 622
	FormIsEmpty            int32 = 900
	AccessTokenIsEmpty     int32 = 910
	AccessTokenIsInvalid   int32 = 920
	ValidateError          int32 = 1000
	SendCodeLock60Seconds  int32 = 1010
	MemberMobileExisted    int32 = 20001
	MemberEmailExisted     int32 = 20002
	SMSCodeError           int32 = 20010
	MemberNotExist         int32 = 20011
	CertError              int32 = 30001
)
