// Package chat implements the types needed to post chat messages to slack.
//
// See: https://api.slack.com/methods/chat.postMessage
package chat

import (
	"github.com/multiplay/go-slack"
)

const (
	// PostMessageEndpoint is the slack URL endpoint for chat post Message
	PostMessageEndpoint = "https://slack.com/api/chat.postMessage"
)

// Message represents slack chat message.
type Message struct {
	// Token is the Authentication token (Requires scope: chat:write:bot or chat:write:user).
	Token string `json:"token,omitempty"`

	// Channel is the channel, private group, or IM channel to send message to.
	Channel string `json:"channel,omitempty"`

	// Text of the message to send.
	Text string `json:"text,omitempty"`

	// Markdown enables Markdown support.
	Markdown bool `json:"mrkdwn,omitempty"`

	// Parse changes how messages are treated.
	Parse string `json:"parse,omitempty"`

	// LinkNames causes link channel names and usernames to be found and linked.
	LinkNames int `json:"link_name,omitempty"`

	// Attachments is structured message attachments
	Attachments []*Attachment `json:"attachments,omitempty"`

	// UnfurLinks enables unfurling of primarily text-based content.
	UnfurlLinks bool `json:"unfurl_links,omitempty"`

	// UnfurlMedia if set to false disables unfurling of media content.
	UnfurlMedia bool `json:"unfurl_media,omitempty"`

	// Username set your bot's user name.
	// Must be used in conjunction with AsUser set to false, otherwise ignored.
	Username string `json:"username,omitempty"`

	// AsUser pass true to post the message as the authed user, instead of as a bot.
	AsUser bool `json:"as_user"`

	// IconURL is the URL to an image to use as the icon for this message.
	// Must be used in conjunction with AsUser set to false, otherwise ignored.
	IconURL string `json:"icon_url,omitempty"`

	// IconEmoji is the emoji to use as the icon for this message.
	// Overrides IconURL.
	// Must be used in conjunction with AsUser set to false, otherwise ignored.
	IconEmoji string `json:"icon_emoji,omitempty"`

	// ThreadTS is the timestamp (ts) of the parent message to reply to a thread.
	ThreadTS string `json:"thread_ts,omitempty"`

	// ReplyBroadcast used in conjunction with thread_ts and indicates whether reply
	// should be made visible to everyone in the channel or conversation.
	ReplyBroadcast bool `json:"reply_broadcast,omitempty"`
}

// NewAttachment creates a new empty attachment adds it to the message and returns it.
func (m *Message) NewAttachment() *Attachment {
	a := &Attachment{}
	m.AddAttachment(a)

	return a
}

// AddAttachment adds a to the message's attachments.
func (m *Message) AddAttachment(a *Attachment) {
	m.Attachments = append(m.Attachments, a)
}

// Send sends the msg to slack using the client c.
func (m *Message) Send(c slack.Client) (*MessageResponse, error) {
	resp := &MessageResponse{}
	if err := c.Send(PostMessageEndpoint, m, resp); err != nil {
		return nil, err
	}

	return resp, nil
}

// MessageResponse the response returned from the post message call.
type MessageResponse struct {
	slack.Response
	Timestamp string   `json:"ts,omitempty"`
	Channel   string   `json:"channel,omitempty"`
	Message   *Message `json:"message,omitempty"`
}
