package domain

type GetCertListRequest struct {
	Keyword string `query: "keyword"`
	Status  string `query: "status"`
	Page    int32  `query: "page"`
	Size    int32  `query: "size"`
}

type PostCertRequest struct {
	Id     int32  `json:"id"`
	UserId int32  `json:"userId"`
	Status string `json:"status"`
	Mark   string `json:"mark"`
}
