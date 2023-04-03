package slack

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewError(t *testing.T) {
	const errMsg = "my error"
	const statusCode = http.StatusNotFound
	err := NewError(statusCode, errMsg)
	assert.EqualError(t, err, fmt.Sprintf("Slack API error: status code %d, message %q", statusCode, errMsg))
	assert.Equal(t, statusCode, err.StatusCode)
}
