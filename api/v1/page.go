package v1

type PageReq struct {
	PageNumber int `json:"pageNumber"`
}

type PageRes struct {
	CurrPage   int `json:"pageNumber"`
	PageSize   int `json:"pageSize"`
	TotalCount int `json:"totalCount"`
	TotalPage  int `json:"totalPage"`
}
