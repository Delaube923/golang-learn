package v1

type CommonPaginationReq struct {
	Page int `json:"page" in:"query" d:"1" v:"min:0#分页号码错误" dc:"分页号码，默认1"`
	Size int `json:"size" in:"query" d:"10" v:"max:100#分页数量最大100条" dc:"分页数量，最大100"`
}
type CommonPaginationRes struct {
	CurrentPage int `json:"currentPage" dc:"当前页码"`
	Total       int `json:"Total" dc:"总数"`
}
