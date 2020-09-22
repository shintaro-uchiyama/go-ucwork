package presentation

type errorResponse struct {
	Message string `json:"message"`
}

func NewErrorResponse(err error) *errorResponse {
	return &errorResponse{
		Message: err.Error(),
	}
}
