package dtos

type Response struct {
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func PrepareResponse(data interface{}, message string) Response {
	return Response{
		Message: message,
		Data:    data,
	}
}
