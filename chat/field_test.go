package chat

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewField(t *testing.T) {
	title := "key"
	val := "val"
	f := NewField(title, val)

	assert.Equal(t, title, f.Title)
	assert.Equal(t, val, f.Value)
	assert.True(t, f.Short)
}

func TestNewFieldLong(t *testing.T) {
	title := "key"
	val := "val which is longer than 20 characters"
	f := NewField(title, val)

	assert.Equal(t, title, f.Title)
	assert.Equal(t, val, f.Value)
	assert.False(t, f.Short)
}
