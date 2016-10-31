package webhook_test

import (
	"net/http"
	"testing"

	"github.com/multiplay/go-slack"
	"github.com/multiplay/go-slack/chat"
	"github.com/multiplay/go-slack/test"
	. "github.com/multiplay/go-slack/webhook"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	c := New(test.Endpoint)
	assert.Equal(t, c.URL, test.Endpoint)
}

func TestSend(t *testing.T) {
	c := New(test.Endpoint)
	resp := &slack.Response{}
	err := c.Send("", nil, resp)
	if !assert.NoError(t, err) {
		return
	}
	assert.True(t, resp.OK)
}

func TestSendError(t *testing.T) {
	errStr := "no_text"
	c := New(test.Endpoint + "?error=" + errStr)
	resp := &slack.Response{}
	err := c.Send("", nil, resp)
	if !assert.Error(t, err) {
		return
	}
	serr := err.(*slack.Error)
	assert.Equal(t, http.StatusOK, serr.StatusCode)
	assert.Equal(t, errStr, serr.Message)
	assert.False(t, resp.OK)
}

func TestSendHttpError(t *testing.T) {
	errStr := "no text"
	c := New(test.Endpoint + "?error=" + errStr)
	resp := &slack.Response{}
	err := c.Send("", nil, resp)
	if !assert.Error(t, err) {
		return
	}
	serr := err.(*slack.Error)
	assert.Equal(t, http.StatusBadRequest, serr.StatusCode)
	assert.False(t, resp.OK)
}

func TestSendDecodeError(t *testing.T) {
	errStr := "no_text"
	c := New(test.Endpoint + "?error=" + errStr)
	resp := struct {
		OK bool
	}{}
	err := c.Send("", nil, resp)
	if !assert.Error(t, err) {
		return
	}
	serr := err.(*slack.Error)
	assert.Equal(t, http.StatusOK, serr.StatusCode)
	assert.False(t, resp.OK)
}

func TestSendMarshalError(t *testing.T) {
	errStr := "no_text"
	c := New(test.Endpoint + "?error=" + errStr)
	resp := struct {
		OK bool
	}{}
	err := c.Send("", New, &resp)
	if !assert.Error(t, err) {
		return
	}
}

func TestSendResponsError(t *testing.T) {
	errStr := "no_text"
	c := New(test.Endpoint + "?error=" + errStr)
	resp := struct {
		OK bool
	}{}
	err := c.Send("", nil, &resp)
	if !assert.Error(t, err) {
		return
	}
	serr := err.(*slack.Error)
	assert.Equal(t, http.StatusOK, serr.StatusCode)
	assert.False(t, resp.OK)
}

func TestSendPostError(t *testing.T) {
	c := New("hhc:/broken")
	resp := &slack.Response{}
	err := c.Send("", nil, resp)
	assert.Error(t, err)
	assert.False(t, resp.OK)
}

func ExampleNew() {
	c := New("https://hooks.slack.com/services/T00000000/B00000000/XXXXXXXXXXXXXXXXXXXXXXXX")
	m := &chat.Message{Text: "test message"}
	m.Send(c)
}
