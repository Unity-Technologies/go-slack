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

// NewError returns a new slack error with statuscode and msg.
func NewError(statuscode int, msg string) *Error {
	return &Error{StatusCode: statuscode, Message: msg}
}

func (e *Error) Error() string {
	return fmt.Sprintf("slack: request failed statuscode: %v, message: %v", e.StatusCode, e.Message)
}
