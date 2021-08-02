package app

import (
	"context"
	"fmt"
	"ihtest/library/global"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/metadata"
)

type MedContext struct {
	ctx      context.Context        // 来源的上下文接口对象，可能带有额外的参数传递
	traceId  string                 // 请求唯一标识
	clientIp string                 // 客户端 IP 信息
	userId   uint32                 // // 用户id
	storage  map[string]interface{} // 存储变量
}

func (mc *MedContext) GetTraceId() string {
	return mc.traceId
}

// 获得客户端IP
func (mc *MedContext) GetClientIp() string {
	return mc.clientIp
}

func (mc *MedContext) SetUserId(id uint32) {
	mc.userId = id
}

func (mc *MedContext) GetUserId() uint32 {
	return mc.userId
}

func (mc *MedContext) Deadline() (deadline time.Time, ok bool) {
	return
}

func (mc *MedContext) Done() <-chan struct{} {
	return nil
}

func (mc *MedContext) Err() error {
	return nil
}

func (mc *MedContext) Value(key interface{}) interface{} {
	if keyAsString, ok := key.(string); ok {
		val, _ := mc.storage[keyAsString]
		return val
	}

	return nil
}

func (mc *MedContext) Set(key string, value interface{}) {
	if len(key) == 0 {
		return
	}

	// 延迟初始化
	if mc.storage == nil {
		mc.storage = make(map[string]interface{})
	}

	mc.storage[key] = value
}

// 将已变更的数据，存储到 gin 上下文中，继续传输
func (mc *MedContext) Storage(ctx *gin.Context) bool {
	ctx.Set(global.HttpCustomContextKey, *mc)
	return true
}

// 创建自定义上下文
func NewMedContext(ctx context.Context) MedContext {
	// 这是一个兼容处理，上下文没有传递TraceId的时候自动创建一个默认的。
	// 在K8S集群环境下，正常情况上下文会有传递链路信息。
	// 当传递的TraceId不存在时默认生成一个，
	// 不影响原有的方法使用TraceId。
	traceId := strconv.Itoa(int(time.Now().UnixNano()))
	ctx = context.WithValue(ctx, fmt.Sprintf("Oath%d", traceId), traceId)
	return MedContext{
		ctx:     ctx,
		traceId: traceId,
	}
}

// 创建带 traceId 的 grpc context
func CreateGRPCContext(ctx *MedContext) context.Context {
	gRpcCtx, _ := context.WithTimeout(ctx.ctx, 15*time.Second)
	if nil != gRpcCtx {
		gRpcCtx = metadata.NewOutgoingContext(gRpcCtx, metadata.Pairs(ctx.GetTraceId()))
	}
	return gRpcCtx
}

// 从 gin 上下文中解析自定义上下文
func ParserMedContext(ctx *gin.Context) *MedContext {
	if v, ok := ctx.Get(global.HttpCustomContextKey); ok {
		rctx := v.(MedContext)
		rctx.clientIp = ctx.ClientIP()

		return &rctx
	}

	return nil
}
