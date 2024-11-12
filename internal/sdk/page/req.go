package page

// Request 分页请求
type Request struct {
	Page     int `json:"page"`     // 页码 - 从1开始
	PageSize int `json:"pageSize"` // 每页条数
}

func (p *Request) GetPageSize() int {
	if p.PageSize <= 0 {
		p.PageSize = 10
	}
	return p.PageSize
}

func (p *Request) GetPage() int {
	if p.Page <= 0 {
		p.Page = 1
	}
	return p.Page
}

func (p *Request) Limit() int {
	return p.GetPageSize()
}

func (p *Request) Offset() int {
	return (p.GetPage() - 1) * p.PageSize
}
