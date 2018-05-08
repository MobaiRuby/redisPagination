package redisPagination

import "math"

type SortSetPaginator Pagination

func NewSortSetPaginator(p *Pagination) *SortSetPaginator {
	return &SortSetPaginator{p.Key, p.Limit}
}

func (s *SortSetPaginator) Total() int64 {
	return client.ZCard(s.Key).Val()
}

func (s *SortSetPaginator) Pages() int64 {
	return int64(math.Ceil(float64(client.ZCard(s.Key).Val()) / float64(s.Limit)))
}

func (s *SortSetPaginator) Page(skip int64) []string {
	skip = trinocular(skip < 1, 1, skip).(int64)
	return client.ZRevRange(s.Key, (skip-1)*s.Limit, skip*s.Limit-1).Val()
}

func (s *SortSetPaginator) All() []string {
	return client.ZRevRange(s.Key, 0, -1).Val()
}
