package enum

type ResponseEnum struct {
	Code    int
	Message string
}

var RESPONSE_SUCCESS = ResponseEnum{Code: 200, Message: "成功"}
var RESPONSE_ERROR = ResponseEnum{Code: 0, Message: "失败"}
var RESPONSE_NO_LOGIN = ResponseEnum{Code: 401, Message: "未登录"}
var RESPONSE_NO_PERMISSION = ResponseEnum{Code: 403, Message: "无权限"}
