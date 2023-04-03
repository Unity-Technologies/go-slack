// Package slack provides a generic interface for Slack clients
// and some basic types to enable the creation of Slack clients.
//
// Subdirectories may contain packages that implement the Slack client interface
// for specific Slack APIs, such as webhooks or bot tokens.
package slack

// Client represents a Slack client.
type Client interface {
	// Send sends the request to Slack.
	Send(url string, message, response interface{}) error
}

// SendResponse is the interface that responses implement.
type SendResponse interface {
	// Ok returns whether the response indicates success.
	Ok() bool

	// Err returns any error message included in the response.
	Err() string

	// Warn returns any warning message included in the response.
	Warn() string
}
