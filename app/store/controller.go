package store

import (
	api_http "ihtest/app"
	"ihtest/library"

	"github.com/gin-gonic/gin"
)

type Controller struct{}

// @Summery 获取Store列表数据
// @Tag store
// @Accept json
// @Produce json
// @param start query int true "起始ID"
// @Param limit query int true "限制输出"
// @Param usecache query string true "是否使用缓存加速"
// @Router /api/v1/store/list [post]
// @Success 200 {object} protocol.StoreListResp
func (c *Controller) List(ctx *gin.Context) {
	var (
		err  error = nil
		req        = new(storeListReq)
		resp       = new(StoreListResp)
	)

	response := library.NewResponse(ctx)
	if err := library.NewRequest(ctx).Bind(req); nil != err {
		response.Error(err)
		return
	}

	if 0 == req.Limit || req.Limit > 100 {
		req.Limit = 15
	}

	resp, err = new(Service).List(api_http.ParserMedContext(ctx), req)
	if nil != err {
		response.Error(err)
		return
	}

	response.Success(resp)
}
