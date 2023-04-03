package slack

import (
	"fmt"
)

// Error represents an error from the Slack API.
type Error struct {
	// StatusCode is the status code returned by the request.
	StatusCode int

	// Message is the message, if any, returned in the body.
	Message string
}

// NewError returns a new slack error with status code and message.
func NewError(statusCode int, message string) *Error {
	return &Error{StatusCode: statusCode, Message: message}
}

// Error returns a string representation of the error.
func (e *Error) Error() string {
	return fmt.Sprintf("Slack API error: status code %d, message %q", e.StatusCode, e.Message)
}
