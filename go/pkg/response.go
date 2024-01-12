package pkg

type IResponse[T any] struct {
	Message    string `json:"message"`
	StatusCode int    `json:"statusCode"`
	Data       T      `json:"data"`
}

func Response[T any](message string, statusCode int, data T) IResponse[T] {
	response := IResponse[T]{
		Message:    message,
		StatusCode: statusCode,
		Data:       data,
	}

	return response
}