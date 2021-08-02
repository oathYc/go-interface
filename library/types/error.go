package types

const (
	// 公共错误码 https://wiki.medlinker.com/pages/viewpage.action?pageId=17598142
	Success = 0

	// 使用 xerror/xcode/vars.go 的定义
	//ErrParam  = 1 // "请求参数错误"
	//ErrSystem = 3 // "网络繁忙，请稍后重试！"

	// 业务错误码 错误码范围: 100600 - 100699
	// https://wiki.medlinker.com/pages/viewpage.action?pageId=17602122
	ErrOrderNotCreate = 100601 // "药品订单尚未创建"
	ErrPrNotFind      = 100602 // "未找到处方"
)

func EqualByCode(err error, code int) bool {
	return true
}
