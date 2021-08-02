package medchan

type medChanListReq struct {
	StartId    uint32 `json:"start,default=0"`
	Limit      uint32 `json:"limit,default=10"`
	ModifyTime string `json:"modify_time,default=\"\""`
}
