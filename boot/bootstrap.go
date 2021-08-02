package boot

import (
	"github.com/pkg/errors"
)

func Boot() error {
	var ret error
	for ok := true; ok; ok = false {
		// 初始化配置
		if err := InitConfig(); nil != err {
			ret = errors.Wrap(err, "初始化配置出错")
			break
		}

		// 初始化日志
		if err := InitLogger(); err != nil {
			ret = errors.Wrap(err, "初始化日志出错")
			break
		}

		// 初始化业务微服务客户端
		if err := InitClient(); nil != err {
			ret = errors.Wrap(err, "初始化 grpc Client 失败")
			break
		}

		// 初始化Redis连接池
		if err := InitRedisPool(); nil != err {
			ret = errors.Wrap(err, "初始化 grpc Client 失败")
			break
		}

		// 初始化http服务
		if err := InitHttpServer(); err != nil {
			ret = errors.Wrap(err, "初始化http server错误")
			break

		}
	}

	return ret
}
