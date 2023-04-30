package response

// 响应数据
type Response struct {
	Code    int			`json:"code"`	// 响应码
	Message string		`json:"message"`// 响应信息
	Data    interface{} `json:"data"`	// 响应数据
}

// 成功
func Success(data interface{}) *Response {
	return &Response{Code: 0, Data: data}
}

// 失败
func Fail(code int, message string) *Response {
	return &Response{Code: code, Message: message}
}

// 失败
func FastFail(message string) *Response {
	return Fail(-1, message)
}

