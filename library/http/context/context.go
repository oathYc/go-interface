package context

import (
	api_http "ihtest/app"

	"github.com/gin-gonic/gin"
)

// gin.Context 读写

// 获取请求参数session
func GetRequestSession(ctx *gin.Context) string {
	sessionId, _ := ctx.GetQuery("sess")
	if len(sessionId) > 0 {
		return sessionId
	}

	sessionId, _ = ctx.Cookie("sess")
	if len(sessionId) > 0 {
		return sessionId
	}

	return ""
}

func GetPlatform(ctx *gin.Context) string {
	val, _ := ctx.GetQuery("x_platform")
	if len(val) > 0 {
		return val
	}

	val, _ = ctx.Cookie("x_platform")
	if len(val) > 0 {
		return val
	}

	val = ctx.GetHeader("x-platform")
	if len(val) > 0 {
		return val
	}

	return "web"
}

func GetWechatType(ctx *gin.Context) string {
	val := ctx.GetHeader("x-wechat-type")
	if len(val) > 0 {
		return val
	}

	return ""
}

// 获取合法session
func GetSession(ctx *gin.Context) string {
	return ctx.GetString("_session")
}

// 设置合法session (合法的sessionId才会存入context中)
func SetSession(ctx *gin.Context, sessionId string) {
	ctx.Set("_session", sessionId)
}

func GetErrMsg(ctx *gin.Context) string {
	return ctx.GetString("_err")
}

func SetErrMsg(ctx *gin.Context, msg string) {
	ctx.Set("_err", msg)
}

func GetTraceId(ctx *gin.Context) string {
	medCtx := api_http.ParserMedContext(ctx)
	return medCtx.GetTraceId()
}
