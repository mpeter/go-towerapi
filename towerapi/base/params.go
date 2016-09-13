package base

type ListParams struct {
	OrderBy  string `url:"order_by,omitempty"`
	PageSize int    `url:"page_size,omitempty"`
	Page     int    `url:"page,omitempty"`
	Search   string `url:"search,omitempty"`
	Name     string `url:"name,omitempty"`
}
