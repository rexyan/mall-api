package model

type PageReq struct {
	PageNumber int `json:"pageNumber" d:"1"`
	PageSize   int `json:"PageSize" d:"10"`
}

type PageRes struct {
	CurrentPage int `json:"pageNumber"`
	PageBarNum  int `json:"pageSize"`
	TotalSize   int `json:"totalCount"`
	TotalPage   int `json:"totalPage"`
}
