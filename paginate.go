package redisPagination

import (
	"github.com/go-redis/redis"
	"github.com/pkg/errors"
)

type (
	Options struct {
		Rp *RedisOptions
		P  *Pagination
	}
	RedisOptions redis.Options

	Pagination struct {
		Key   string
		Limit int64
	}
)

type Paginator interface {
	Total() int64
	Pages() int64
	Page(skip int64) []string
	All() []string
}

const (
	dataTypeList    = "list"
	dataTypeSortSet = "zset"
	dataTypeNone    = "none"
)

func Paginate(opt *Options) Paginator {
	// get redis client
	getRedisClient(opt.Rp)

	switch client.Type(opt.P.Key).Val() {
	case dataTypeList:
		return NewListPaginator(opt.P)
	case dataTypeSortSet:
		return NewSortSetPaginator(opt.P)
	case dataTypeNone:
		return NewNonePaginator(opt.P)
	default:
		panic(errors.New("cannot handle data type when redis pagination"))
	}
}
