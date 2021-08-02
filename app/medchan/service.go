package medchan

import (
	"encoding/json"
	"fmt"
	api_http "ihtest/app"

	"ihtest/library/redis"
	"ihtest/library/utils"

	"ihtest/grpcser"
	"ihtest/protocol"

	"ihtest/library/global"
)

const (
	CACHE_LIST_KEY = "MedChanList:v1:%s:%s"
)

type Service struct{}

// 处方详情
func (s Service) ListWiCache(ctx *api_http.MedContext, req *medChanListReq) (*medChanListResp, error) {
	var (
		err  error
		resp = new(medChanListResp)
	)

	err = redis.Cache(s.CacheKeyList("", *req), redis.DefaultExpiredTime, &resp, func() (interface{}, error) {
		r, err := s.List(ctx, req)
		if nil != err {
			return uint64(0), err
		}
		return r, nil
	})
	return resp, err
}

func (s Service) CacheKeyList(appId string, params medChanListReq) string {
	keyData, _ := json.Marshal(params)
	hash := utils.Md5(string(keyData))
	return fmt.Sprintf(CACHE_LIST_KEY, appId, hash)
}

// 处方详情
func (s Service) List(ctx *api_http.MedContext, req *medChanListReq) (*medChanListResp, error) {
	var (
		err  error
		resp = new(medChanListResp)
	)

	client := grpcser.NewMedTestSerClient(global.MedTestConn)
	modifyTime, err := utils.DatetimeStringToUnix(req.ModifyTime)
	if err != nil {
		return resp, err
	}
	list, err := client.GetMedChanList(api_http.CreateGRPCContext(ctx), &protocol.MedChanListReq{
		BaseRequest: &protocol.BaseRequest{
			Start: req.StartId,
			Limit: uint32(req.Limit + 1),
			Sorts: []string{"id"},
		},
		Filter: &protocol.MedChanListReq_Filter{
			Id:         int32(req.StartId),
			ModifyTime: modifyTime,
		},
	})
	if err != nil {
		return resp, err
	}
	hasMore := int32(0)
	start := req.StartId
	dbCount := uint32(len(list.List))
	if dbCount > req.Limit {
		list.List = list.List[:req.Limit]
		hasMore = 1
	}
	if dbCount > 1 {
		lastDate := list.List[len(list.List)-1:]
		start = lastDate[0].Id
	} else if dbCount == 1 {
		start = list.List[0].Id
	}
	for _, item := range list.List {
		row := &medChanInfo{

			Id:            item.Id,
			Title:         item.Title,
			IsSelf:        item.IsSelf,
			Remark:        item.Remark,
			ChannelStatus: item.ChannelStatus,
			StoreCount:    item.StoreCount,
			CreatedAt:     item.CreatedAt,
			UpdatedAt:     item.UpdatedAt,
			DeleteAt:      item.DeletedAt,
		}

		resp.List = append(resp.List, row)
	}
	resp.More = hasMore
	resp.Start = int32(start)
	return resp, nil
}
