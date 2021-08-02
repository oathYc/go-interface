package app

// 响应数据体
type ResponseEntity struct {
	Code    int         `json:"errcode"`
	Message string      `json:"errmsg"`
	Data    interface{} `json:"data"`
}
