package jsonutils

type ErrorMessage struct {
	Error       string `json:"error,omitempty"`
	ErrorCode   int    `json:"error_code,omitempty"`
	Description string `json:"description,omitempty"`
}
