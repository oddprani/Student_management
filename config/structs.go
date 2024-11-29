package config

type ErrorResponse struct {
	Message   string `json:"message"`
	ErrorCode int    `json:"error_code"`
}
