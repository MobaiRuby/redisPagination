package redisPagination

import "github.com/go-redis/redis"

var client *redis.Client

func getRedisClient(opt *redis.Options) *redis.Client {
	if client != nil {
		return client
	}

	client = redis.NewClient(opt)

	return client
}
