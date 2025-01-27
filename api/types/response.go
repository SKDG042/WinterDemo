package types
// 接着为了方便返回数据，我需要按照文档构建一个结构体
type Response struct {
	Status int         `json:"status"`
	Info   string      `json:"info"`
	Data   interface{} `json:"data"`
}

// 成功的Response
func SuccessResponse(data interface{}) *Response {
	return &Response{
		Status: 10000,
		Info:   "success",
		Data:   data,
	}
}

// 失败的Response
func ErrorResponse(status int, info string) *Response {
	return &Response{
		Status: status,
		Info:   info,
		Data:   nil,
	}
}
