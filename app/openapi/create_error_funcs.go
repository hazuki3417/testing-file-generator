package openapi

func InternalServerError() *ImplResponse {

	return &ImplResponse{
		Code: 500,
		Body: ErrorInfo{
			Message: "サーバーエラーが発生しました",
		},
	}
}
