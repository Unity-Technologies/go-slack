// Package webhook provides a slack webhook client implementation.
//
// See: https://api.slack.com/incoming-webhooks
package webhook

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/multiplay/go-slack"
)

// Client is a slack webhook client for posting messages using a webhook URL.
type Client struct {
	// URL is the webhook URL to use
	URL string
}

// New returns a new Client which sends request using the webhook URL.
func New(url string) *Client {
	return &Client{URL: url}
}

// Send sends the request to slack using the webhook protocol.
// The url parameter only exists to satisfy the slack.Client interface
// and is not used by the webhook Client.
func (c *Client) Send(url string, msg, resp interface{}) error {
	b, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	r, err := http.Post(c.URL, "application/json; charset=utf-8", bytes.NewReader(b))
	if err != nil {
		return err
	}
	defer r.Body.Close()

	if r.StatusCode != http.StatusOK {
		b, _ := ioutil.ReadAll(r.Body)
		return slack.NewError(r.StatusCode, string(b))
	}

	var body io.Reader
	if strings.HasPrefix(c.URL, "https://hooks.slack.com") {
		// Work around webhooks not returning JSON as originally documented
		// by treating all StatusOK as success.
		body = bytes.NewReader([]byte(`{"ok":true}`))
	} else {
		// This is required for compatibility with API test endpoint.
		body = r.Body
	}

	dec := json.NewDecoder(body)
	if err := dec.Decode(resp); err != nil {
		return slack.NewError(r.StatusCode, err.Error())
	}

	if sr, ok := resp.(slack.SendResponse); !ok {
		return slack.NewError(r.StatusCode, "not a response")
	} else if !sr.Ok() {
		return slack.NewError(r.StatusCode, sr.Err())
	}

	return nil
}
