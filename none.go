package redisPagination

type NonePaginator Pagination

func NewNonePaginator(p *Pagination) *ListPaginator {
	return &ListPaginator{p.Key, p.Limit}
}

func (n *NonePaginator) Total() int64 {
	return 0
}

func (n *NonePaginator) Pages() int64 {
	return 0
}

func (n *NonePaginator) Page(skip int64) []string {
	return []string{}
}

func (n *NonePaginator) All() []string {
	return []string{}
}
