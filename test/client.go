// Package test provides a slack client implementation which uses the slack api.test endpoint
// so is suitable to testing.
//
// See: https://api.slack.com/methods/api.test
package test

import (
	"net/url"

	"github.com/multiplay/go-slack/webhook"
)

// New returns a new slack.Client that can be used for testing.
func New() *webhook.Client {
	return webhook.New(Endpoint)
}

// NewError returns a new slack.Client that can be used for testing which errors with err.
func NewError(err string) *webhook.Client {
	v := url.Values{"error": {err}}
	return webhook.New(Endpoint + "?" + v.Encode())
}
