package chat

import (
	"testing"

	"github.com/multiplay/go-slack/test"

	"github.com/stretchr/testify/assert"
)

func TestMessageNewAttachment(t *testing.T) {
	m := &Message{}
	a := m.NewAttachment()

	if !assert.Equal(t, 1, len(m.Attachments)) {
		return
	}
	assert.Equal(t, a, m.Attachments[0])
}

func TestMessageSend(t *testing.T) {
	c := test.New()
	m := &Message{Text: "test message"}
	resp, err := m.Send(c)
	if !assert.NoError(t, err) {
		return
	}
	assert.True(t, resp.OK)
}

func TestMessageSendError(t *testing.T) {
	c := test.NewError("my error")
	m := &Message{Text: "test message"}
	_, err := m.Send(c)
	assert.Error(t, err)
}

func TestAddAttachment(t *testing.T) {
	m := &Message{}
	a := &Attachment{}
	m.AddAttachment(a)

	if !assert.Equal(t, 1, len(m.Attachments)) {
		return
	}
	assert.Equal(t, a, m.Attachments[0])
}
