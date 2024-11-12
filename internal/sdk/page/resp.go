package page

// Response 分页响应
type Response struct {
	List     any   `json:"list"`
	Count    int64 `json:"count"`
	Page     int   `json:"page"`
	PageSize int   `json:"pageSize"`
}

func NewResponse(req Request, count int64, list any) Response {
	return Response{
		Page:     req.GetPage(),
		PageSize: req.GetPageSize(),
		Count:    count,
		List:     list,
	}
}
