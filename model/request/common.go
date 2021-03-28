package request

type PageInfo struct {
	Page int `json:"page" `
	PageSize int `json:"page_size"`
}

type GetById struct {
	Id []int `json:"id"`
}