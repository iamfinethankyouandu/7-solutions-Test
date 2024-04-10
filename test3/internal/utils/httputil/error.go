package httputil

import "fmt"

type Http struct {
	Message    string `json:"Message,omitempty"`
	StatusCode int    `json:"statusCode"`
}

func (e Http) Error() string {
	return fmt.Sprintf("Message: %s", e.Message)
}

func NewHttpError(message, metadata string, statusCode int) Http {
	return Http{
		Message:    message,
		StatusCode: statusCode,
	}
}
