package codes

const (
	Success               int32 = 200
	Forbidden             int32 = 403
	InteralServerError    int32 = 500
	ValidateError         int32 = 1000
	SendCodeLock60Seconds int32 = 1010
	MemberMobileExisted   int32 = 20001
	MemberEmailExisted    int32 = 20002
	CertError             int32 = 30001
)
