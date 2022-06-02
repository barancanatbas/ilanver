package helpers

type Response struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
}

func BasicReturn(status int, data interface{}) Response {
	return Response{
		Status: status,
		Data:   data,
	}
}

func BasicError(status int, err error) Response {
	return Response{
		Status: status,
		Data:   err.Error(),
	}
}
