package test

import (
	"fmt"
	"net/url"
	"testing"

	"github.com/multiplay/go-slack"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	c := New()
	assert.Equal(t, c.URL, Endpoint)
}

func TestNewError(t *testing.T) {
	err := "my error"
	c := NewError(err)
	assert.Contains(t, c.URL, Endpoint)
	assert.Contains(t, c.URL, url.Values{"error": {err}}.Encode())
}

func ExampleNew() {
	c := New()
	msg := struct {
		Param1 string
		Param2 int
	}{Param1: "my value", Param2: 20}

	resp := struct {
		slack.Response
		Args struct {
			Param1 string
			Param2 int
		}
	}{}
	if err := c.Send("", msg, resp); err != nil {
		// No error is expected here.
		fmt.Println("error:", err)
	}
	fmt.Println("response:", resp)
}

func ExampleNewError() {
	c := NewError("my error")
	msg := struct {
		Param1 string
		Param2 int
	}{Param1: "my value", Param2: 20}

	resp := struct {
		slack.Response
		Args struct {
			Param1 string
			Param2 int
		}
	}{}
	if err := c.Send("", msg, resp); err != nil {
		// An error is expected here.
		fmt.Println("error:", err)
	}
	fmt.Println("response:", resp)
}
