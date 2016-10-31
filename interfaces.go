// Package slack provides a generic interface for slack clients
// and some basic types to enable the creation of slack clients.
//
// See the webhook sub directory for an example of such a client.
package slack

// Client represents a slack client.
type Client interface {
	// Send sends the request to slack.
	Send(url string, message, response interface{}) error
}

// SendResponse is the interface that responses implement.
type SendResponse interface {
	Ok() bool
	Err() string
	Warn() string
}
