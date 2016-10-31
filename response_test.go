package slack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResponse(t *testing.T) {
	ok, err, warn := true, "my error", "my warn"
	r := &Response{OK: ok, Error: err, Warning: warn}
	assert.Equal(t, ok, r.Ok())
	assert.Equal(t, err, r.Err())
	assert.Equal(t, warn, r.Warn())
}
