package result

func Success(data any) map[string]any {
	return map[string]any{
		"code": 200,
		"msg":  "success",
		"data": data,
	}
}

func Error(message string) map[string]any {
	return map[string]any{
		"code": 0,
		"msg":  message,
		"data": nil,
	}
}
func ErrorD(message string, data any) map[string]any {
	return map[string]any{
		"code": 0,
		"msg":  message,
		"data": data,
	}
}
