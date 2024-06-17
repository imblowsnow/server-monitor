package utils

import (
	"server-monitor/pkg/server/common/enum"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Result(enum enum.ResponseEnum, data interface{}) Response {
	return Response{
		Code:    enum.Code,
		Message: enum.Message,
		Data:    data,
	}
}
func ResultM(enum enum.ResponseEnum, message string, data interface{}) Response {
	return Response{
		Code:    enum.Code,
		Message: message,
		Data:    data,
	}
}
func ResultSuccessE(enum enum.ResponseEnum, data interface{}) Response {
	return Result(enum, data)
}
func ResultSuccess(data interface{}) Response {
	return ResultSuccessE(enum.RESPONSE_SUCCESS, data)
}
func ResultError(message string) Response {
	return ResultM(enum.RESPONSE_ERROR, message, nil)
}
func ResultErrorD() Response {
	return Result(enum.RESPONSE_ERROR, nil)
}
