package response

func SucceedResponse(data map[string]interface{}, status int32, msg string) map[string]interface{} {
	return map[string]interface{}{
		"data":   data,
		"status": status,
		"msg":    msg,
	}
}

func FailedResponse(data map[string]interface{}, status int32, msg string) map[string]interface{} {
	return map[string]interface{}{
		"data":   data,
		"status": status,
		"msg":    msg,
	}
}
