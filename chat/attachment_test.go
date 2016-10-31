package chat

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAttachmentNewField(t *testing.T) {
	a := &Attachment{}
	title := "key"
	val := "val"
	f := a.NewField(title, val)

	assert.Equal(t, title, f.Title)
	assert.Equal(t, val, f.Value)
	assert.True(t, f.Short)
	assert.Equal(t, 1, len(a.Fields))
}

func TestAttachmentNewFieldLong(t *testing.T) {
	a := &Attachment{}
	title := "key"
	val := "val which is longer than 20 characters"
	f := a.NewField(title, val)

	assert.Equal(t, title, f.Title)
	assert.Equal(t, val, f.Value)
	assert.False(t, f.Short)
	assert.Equal(t, 1, len(a.Fields))
}
