package model

// ErrorMetadata holds error context for logging
type ErrorMetadata struct {
	Error      string `json:"error"`
	Stacktrace string `json:"stacktrace"`
}
