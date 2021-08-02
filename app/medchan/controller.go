package medchan

import (
	api_http "ihtest/app"

	"github.com/gin-gonic/gin"

	"ihtest/library"
)

type Controller struct{}

// @Summary 获得处方记录
// @Tags 合作数据
// @Accept json
// @Produce json
// @Param date query string  true "日期" required
// @Param startId query int  true "上一次返回的nextId"
// @Param limit query int  true "数据返回条数、默认15"
// @Router /api/v1/channel/list [post]
// @Success 200 {object} protocol.MedChanListResp
func (*Controller) List(ctx *gin.Context) {
	var err error = nil
	var res = new(medChanListResp)
	params := new(medChanListReq)

	response := library.NewResponse(ctx)
	if err := library.NewRequest(ctx).Bind(params); nil != err {
		response.Error(err)
		return
	}
	if params.Limit <= 0 {
		params.Limit = 15
	}
	useCache := ctx.Query("usecache")
	if 0 < len(useCache) {
		res, err = new(Service).ListWiCache(api_http.ParserMedContext(ctx), params)
	} else {
		res, err = new(Service).List(api_http.ParserMedContext(ctx), params)
	}
	if err != nil {
		response.Error(err)
		return
	}

	response.Success(res)
}
