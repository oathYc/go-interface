package library

import (
	"fmt"
	"net/http"

	api_http "ihtest/app"
	"ihtest/library/http/context"
	"ihtest/library/types"

	"github.com/gin-gonic/gin"
)

type Response struct {
	ctx *gin.Context
}

func NewResponse(c *gin.Context) *Response {
	traceId := "<empty>"
	medCtx := api_http.ParserMedContext(c)
	if nil != medCtx {
		traceId = medCtx.GetTraceId()
	}

	// 响应请求唯一标识
	c.Header("X-Trace-Id", traceId)

	return &Response{c}
}

// 响应成功请求
func (r *Response) Success(data interface{}) {
	r.SuccessMsg(data, "")
}

// 响应成功请求 自定义msg
func (r *Response) SuccessMsg(data interface{}, msg string) {
	if len(msg) == 0 {
		msg = "success"
	}

	// 响应数据
	r.ctx.JSON(http.StatusOK, api_http.ResponseEntity{
		Code:    types.Success,
		Message: msg,
		Data:    data,
	})
}

type ErrorInit struct {
	ErrorCode int
	Error     error
}

// 响应错误
func (r *Response) Error(err interface{}) {
	var errmsg string
	defer func() {
		// 供日志中间件打印
		context.SetErrMsg(r.ctx, "response err, "+errmsg)
	}()

	// 自定义错误
	if xerr, ok := err.(ErrorInit); ok {
		e := xerr.ErrorCode
		errmsg = fmt.Sprintf("[code: %d][err: (%s)]", xerr.ErrorCode, xerr.Error)

		msg := xerr.Error.Error()
		if e == 500 {
			// xerror.New() 生成的默认错误，屏蔽信息
			msg = "网络繁忙，请稍后重试！"
		}

		r.ctx.JSON(http.StatusOK, api_http.ResponseEntity{
			Code:    e,
			Message: msg,
			Data:    nil,
		})
		return
	}

	errmsg = fmt.Sprintf("err:(%s)", err)
	// 屏蔽掉其他错误
	r.ctx.JSON(http.StatusOK, api_http.ResponseEntity{
		Code:    500,
		Message: "网络繁忙，请稍后重试！",
		Data:    nil,
	})
	return
}
