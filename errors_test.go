package slack

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewError(t *testing.T) {
	errStr := "my error"
	code := http.StatusNotFound
	err := NewError(code, errStr)
	assert.Error(t, err)
	assert.Equal(t, code, err.StatusCode)
	assert.Contains(t, err.Error(), "failed")
}
