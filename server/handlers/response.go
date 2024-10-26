package handlers

type SuccessResponse struct {
	Data    any    `json:"data,omitempty"`
	Message string `json:"message"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}
