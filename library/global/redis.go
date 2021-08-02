package global

import "github.com/gomodule/redigo/redis"

var RedisPoolEntry map[string]*redis.Pool
