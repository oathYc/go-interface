package store

import (
	"ihtest/app"
	"ihtest/grpcser"
	"ihtest/library/global"
	"ihtest/protocol"
)

type Service struct{}

func (s *Service) List(ctx *app.MedContext, req *storeListReq) (resp *StoreListResp, err error) {
	var (
		hasMore int
	)
	resp = new(StoreListResp)

	client := grpcser.NewStoreSerClient(global.MedTestConn)
	list, err := client.GetStoreList(app.CreateGRPCContext(ctx), &protocol.StoreListReq{
		BaseRequest: &protocol.BaseRequest{
			Start: req.Id,
			Limit: uint32(req.Limit),
			Sorts: []string{"id"},
		},
		Filter: &protocol.StoreListReq_Filter{
			Id:        req.Id,
			ChannelId: req.ChannelId,
			Title:     req.Title,
			Cellphone: req.Cellphone,
		},
	})

	if nil != err {
		return resp, err
	}

	start := req.Id
	dbCount := uint32(len(list.List))
	if dbCount > uint32(req.Limit) {
		list.List = list.List[:req.Limit]
		hasMore = 1
	}
	if dbCount > 1 {
		lastData := list.List[len(list.List)-1:]
		start = lastData[0].Id
	} else {
		start = list.List[0].Id
	}

	for _, item := range list.List {
		row := &StoreInfo{
			Id:          item.Id,
			Title:       item.Title,
			ChannelId:   item.ChannelId,
			Address:     item.Address,
			CreatorName: item.CreatorName,
			Cellphone:   item.Cellphone,
			StoreStatus: item.StoreStatus,
			CreatedAt:   item.CreatedAt,
			UpdatedAt:   item.UpdatedAt,
			DeleteAt:    item.DeletedAt,
		}
		resp.List = append(resp.List, row)
	}

	resp.Start = int32(start)
	resp.More = int32(hasMore)

	return resp, nil
}
