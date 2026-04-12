package apperror

type AppError struct {
	Status int               `json:"-"`
	Error  string            `json:"error"`
	Fields map[string]string `json:"fields,omitempty"`
}

func New(status int, message string) *AppError {
	return &AppError{Status: status, Error: message}
}

func WithFields(status int, message string, fields map[string]string) *AppError {
	return &AppError{Status: status, Error: message, Fields: fields}
}
