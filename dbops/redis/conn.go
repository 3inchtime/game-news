package redis

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

var redisHost = "127.0.0.1:6379"

func newRedis() redis.Conn {
	c, err := redis.Dial("tcp", redisHost)
	if err != nil {
		fmt.Printf("Redis Connected Error: %s", err)
		return nil
	} else {
		fmt.Println("Redis Connected Succssful")
		return c
	}
}

func RedisCon() redis.Conn {
	cache := newRedis()
	return cache
}
