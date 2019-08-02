package controller

const (
	commonTitle string = "测试"
)

type APIJson struct {
	Code int        `json:"code"`
	Msg    interface{} `json:"msg"`
	Data   interface{} `json:"data"`
}

// APIResult 用户API 项目 忽略
func APIResult(code int, object interface{}, msg string) (apijson *APIJson) {
	// apijson 已经在返回值处 声明了，不用重复声明。
	apijson = &APIJson{Code: code, Data: object, Msg: msg}
	return apijson
}
