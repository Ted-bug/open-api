package response

type AnyResponse map[string]any

func SucceedResponse(data map[string]any, status int32, msg string) AnyResponse {
	return AnyResponse{
		"data":   data,
		"status": status,
		"msg":    msg,
	}
}

func SucceedWithData(data map[string]any) AnyResponse {
	return SucceedResponse(data, 0, "success")
}

func FailedResponse(data map[string]any, status int32, msg string) AnyResponse {
	return AnyResponse{
		"data":   data,
		"status": status,
		"msg":    msg,
	}
}

func FailedWithMsg(msg string) AnyResponse {
	return FailedResponse(nil, 1, msg)
}

func FailedWithMsgAndCode(msg string, code int32) AnyResponse {
	return FailedResponse(nil, code, msg)
}
