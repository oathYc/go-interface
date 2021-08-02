package middleware

import (
	"context"
	api_http "ihtest/app"
	"ihtest/library/global"
	"time"

	"github.com/gin-gonic/gin"
)

func RegisterGlobalParams() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 注册自定义上下文
		registerContext(ctx)

		ctx.Next()
	}
}

func registerContext(ctx *gin.Context) {
	// 通过请求Header创建标准的context上下文变量，
	// 注入到已有的MedContext对象中，不影响该对象的原有使用。
	cx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()
	ctx.Set(
		global.HttpCustomContextKey,
		api_http.NewMedContext(cx),
	)
}
