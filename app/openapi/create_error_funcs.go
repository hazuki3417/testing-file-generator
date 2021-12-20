package openapi

func InternalServerError(message string) *ImplResponse {
	return &ImplResponse{
		Code: 500,
		Body: ErrorInfo{
			Message: message,
			Errors:  []Error{},
		},
	}
}
