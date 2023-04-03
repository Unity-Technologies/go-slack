package slack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResponse(t *testing.T) {
	const (
		errMsg = "my error"
		warnMsg = "my warn"
	)

	r := &Response{
		OK:      true,
		Error:   errMsg,
		Warning: warnMsg,
	}

	assert.True(t, r.Ok())
	assert.Equal(t, errMsg, r.Err())
	assert.Equal(t, warnMsg, r.Warn())
}
