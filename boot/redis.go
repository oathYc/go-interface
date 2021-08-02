package boot

import (
	"github.com/gomodule/redigo/redis"
	"github.com/pkg/errors"
	"ihtest/library/global"
	"time"
)

func InitRedisPool() error {
	global.RedisPoolEntry = make(map[string]*redis.Pool)

	if 0 == len(global.Config.Redis.Pool) {
		return nil
	}

	for name, redisConfig := range global.Config.Redis.Pool {
		dial := func() (redis.Conn, error) {
			return redis.Dial(
				"tcp",
				redisConfig.Address,
				redis.DialPassword(redisConfig.DialPassword),
				redis.DialDatabase(redisConfig.Db),
				redis.DialConnectTimeout(time.Duration(redisConfig.Idletimeout)*time.Millisecond),
			)
		}

		conn, err := dial()
		if err != nil {
			return errors.WithMessage(err, "初始化redis连接失败"+name)
		}
		// 只是验证是否能够连接
		err = conn.Close()
		if err != nil {
			return errors.WithMessage(err, "初始化redis连接失败"+name)
		}

		global.RedisPoolEntry[name] = &redis.Pool{
			Dial:      dial,
			MaxIdle:   redisConfig.MaxIdle,
			MaxActive: redisConfig.MaxActive,
			Wait:      redisConfig.Wait,
			TestOnBorrow: func(c redis.Conn, t time.Time) error {
				_, err := c.Do("PING")
				return err
			},
		}
	}

	return nil
}
