package store

type storeListReq struct {
	Id        uint32 `json:"id,default=0"`
	Start     uint32 `json:"start,default=0"`
	Limit     int32  `json:"limit,default=0"`
	ChannelId int32  `json:"channel_id,default=0"`
	Title     string `json:"title,default=\"\""`
	Cellphone string `json:"cellphone,default=\"\""`
}
