package medchan

type medChanInfo struct {
	Id            uint32 `json:"id"`
	Title         string `json:"title"`
	IsSelf        uint32 `json:"is_self"`
	Remark        string `json:"remark"`
	ChannelStatus int32  `json:"channel_status"`
	StoreCount    uint32 `json:"store_count"`
	UpdatedAt     int64  `json:"updated_at"`
	CreatedAt     int64  `json:"created_at"`
	DeleteAt      int64  `json:"delete_at"`
}

type medChanListResp struct {
	Start int32          `json:"start"`
	More  int32          `json:"more"`
	List  []*medChanInfo `json:"list"`
}
