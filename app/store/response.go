package store

type StoreInfo struct {
	Id          uint32 `json:"id"`
	ChannelId   int32  `json:"channel_id"`
	Title       string `json:"title"`
	Address     string `json:"address"`
	Cellphone   string `json:"cellphone"`
	CreatorName string `json:"creator_name"`
	StoreStatus int32  `json:"store_status"`
	UpdatedAt   int64  `json:"updated_at"`
	CreatedAt   int64  `json:"created_at"`
	DeleteAt    int64  `json:"delete_at"`
}

type StoreListResp struct {
	Start int32
	More  int32
	List  []*StoreInfo
}
