package library

import (
	"fmt"
	"github.com/pkg/errors"
	"ihtest/library/global"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type Request struct {
	ctx *gin.Context
}

// 初始化Request
func NewRequest(ctx *gin.Context) *Request {
	return &Request{
		ctx: ctx,
	}
}

// 参数解析
func (r *Request) Bind(req interface{}) error {
	if err := r.ShouldBind(req); nil != err {
		global.Log.Info(fmt.Sprintf("[ Request ]  req: %+v, err: %v", req, err))
		return errors.New("param invalid")
	}

	return nil
}

func (r *Request) ShouldBind(req interface{}) error {
	bType := binding.Default(r.ctx.Request.Method, r.ctx.ContentType())
	if binding.JSON == bType {
		err := r.ctx.ShouldBindBodyWith(req, binding.JSON)
		if nil != err {
			return err
		}
	} else {
		err := r.ctx.ShouldBindWith(req, bType)
		if nil != err {
			return err
		}
	}

	return nil
}
