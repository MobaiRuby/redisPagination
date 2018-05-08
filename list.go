package redisPagination

import "math"

type ListPaginator Pagination

func NewListPaginator(p *Pagination) *ListPaginator {
	return &ListPaginator{p.Key, p.Limit}
}

func (l *ListPaginator) Total() int64 {
	return client.LLen(l.Key).Val()
}

func (l *ListPaginator) Pages() int64 {
	return int64(math.Ceil(float64(client.LLen(l.Key).Val()) / float64(l.Limit)))
}

func (l *ListPaginator) Page(skip int64) []string {
	skip = trinocular(skip < 1, 1, skip).(int64)
	return client.LRange(l.Key, (skip-1)*l.Limit, skip*l.Limit-1).Val()
}

func (l *ListPaginator) All() []string {
	return client.LRange(l.Key, 0, -1).Val()
}
