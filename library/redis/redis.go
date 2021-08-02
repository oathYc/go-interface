package redis

import (
	"encoding/json"
	"github.com/pkg/errors"
	"ihtest/library/global"
	"math/rand"
	"time"

	"ihtest/library/utils"

	redigo "github.com/gomodule/redigo/redis"
)

const (
	NAME_DEFAULT = "default"

	DefaultPrefix = "MedIhTestCache_"

	DefaultExpiredTime = 60
)

type Task func(redisConn redigo.Conn) (interface{}, error)
type DoCache func() (interface{}, error)

func Get(name string) (redigo.Conn, error) {
	p, ok := global.RedisPoolEntry[name]
	if !ok {
		errors.New("redis 练级获取错误")
	}
	return p.Get(), nil
}

func Exec(name string, tran Task) (interface{}, error) {
	redisConn, err := Get(name)
	if nil != err {
		return nil, err
	}
	defer func() {
		if err := redisConn.Close(); nil != err {
			global.Log.Error("redis close err: " + err.Error())
		}
	}()

	return tran(redisConn)
}

func Cache(key string, expiredTime int, resp interface{}, dataTran DoCache) error {
	var (
		jsonData []byte
		err      error
	)

	rand.Seed(time.Now().UnixNano())
	expiredTime = expiredTime + rand.Intn(60)
	key = DefaultPrefix + key

	jsonData, err = redigo.Bytes(Exec(NAME_DEFAULT, func(redisConn redigo.Conn) (i interface{}, err error) {
		return redisConn.Do("GET", redigo.Args{}.Add(key)...)
	}))

	if err == nil {
		if err := json.Unmarshal(jsonData, resp); nil == err {
			return nil
		}
	}

	data, err := dataTran()
	if err != nil {
		return err
	}

	utils.CopyStruct(data, resp)
	jsonData, _ = json.Marshal(resp)
	_, err = Exec(NAME_DEFAULT, func(redisConn redigo.Conn) (i interface{}, e error) {
		return redisConn.Do("SETEX", redigo.Args{}.Add(key).Add(expiredTime).Add(jsonData)...)
	})

	return err
}
