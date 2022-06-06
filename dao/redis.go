package dao

import (
	"github.com/garyburd/redigo/redis"
)

var RedisPool redis.Pool

func InitRedis() {
	RedisPool = redis.Pool{
		Dial: func() (redis.Conn, error) {
			return redis.Dial(
				"tcp", "47.100.83.25:6379",
				redis.DialPassword("niyu123"),
			)
		},
		MaxIdle:         20,
		MaxActive:       50,
		IdleTimeout:     60,
		MaxConnLifetime: 60 * 5,
	}
}
