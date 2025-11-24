package utils

type Pagination struct {
	Total    uint        `json:"total"`
	Page     uint        `json:"page"`
	PageSize uint        `json:"page_size"`
	List     interface{} `json:"list"`
}
