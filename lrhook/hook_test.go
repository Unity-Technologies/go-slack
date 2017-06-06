package lrhook

import (
	"bytes"
	"io"
	"os"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/multiplay/go-slack/chat"
	"github.com/multiplay/go-slack/test"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

// sbuf is a minimal locked bytes.Buffer so its thread safe.
type sbuf struct {
	mtx sync.RWMutex
	buf bytes.Buffer
}

// Write implements io.Writer.
func (b *sbuf) Write(p []byte) (n int, err error) {
	b.mtx.Lock()
	defer b.mtx.Unlock()

	return b.buf.Write(p)
}

// String implements Stringier.
func (b *sbuf) String() string {
	b.mtx.RLock()
	defer b.mtx.RUnlock()

	return b.buf.String()
}

// Reset resets the buffer to be empty, but it retains the underlying storage for use by future writes.
func (b *sbuf) Reset() {
	b.mtx.Lock()
	defer b.mtx.Unlock()

	b.buf.Reset()
}

var (
	log    = &sbuf{}
	stderr = &sbuf{}
)

func init() {
	// Capture stderr here so logrus uses it to log fire failures.
	r, w, _ := os.Pipe()
	os.Stderr = w
	go io.Copy(stderr, r)
}

func resetBufs() {
	log.Reset()
	stderr.Reset()
}

func newHookedLogger(h logrus.Hook) *logrus.Logger {
	l := logrus.New()
	l.Out = log
	l.Hooks.Add(h)

	return l
}

// hookWait sleeps for a short period to ensure the hook has had chance to fire.
func hookWait() {
	time.Sleep(time.Millisecond * 50)
}

func TestHookError(t *testing.T) {
	cfg := Config{MinLevel: logrus.InfoLevel}
	h := New(cfg, "hdds://broken")

	logger := newHookedLogger(h)
	logger.Info("my info")

	hookWait()

	assert.Contains(t, log.String(), "my info")
	assert.Contains(t, stderr.String(), "Failed to fire hook")

	resetBufs()
}

func TestNew(t *testing.T) {
	cfg := Config{MinLevel: logrus.InfoLevel}
	h := New(cfg, test.Endpoint)
	delete(h.LevelColors, "warning")

	logger := newHookedLogger(h)
	logger.Info("my info")
	logger.WithField("myfield", "some data").Warn("my warn")

	hookWait()

	assert.Contains(t, log.String(), "my info")
	assert.Contains(t, log.String(), "my warn")
	assert.Empty(t, stderr.String())

	resetBufs()
}

func TestLimitPass(t *testing.T) {
	cfg := Config{
		MinLevel: logrus.WarnLevel,
		Limit:    10,
		Burst:    20,
	}
	h := New(cfg, "hdds://broken")

	logger := newHookedLogger(h)
	logger.Warn("my warn")
	logger.Warn("my warn")

	hookWait()

	assert.Contains(t, log.String(), "my warn")
	assert.Equal(t, 2, strings.Count(stderr.String(), "Failed to fire hook"))

	resetBufs()
}

func TestLimitLimited(t *testing.T) {
	cfg := Config{
		MinLevel: logrus.WarnLevel,
		Limit:    1,
		Burst:    1,
	}
	h := New(cfg, "hdds://broken")

	logger := newHookedLogger(h)
	logger.Warn("my warn")
	logger.Warn("my warn")

	hookWait()

	assert.Contains(t, log.String(), "my warn")
	assert.Equal(t, 1, strings.Count(stderr.String(), "Failed to fire hook"))

	resetBufs()
}

func TestAsync(t *testing.T) {
	cfg := Config{
		Async:    true,
		MinLevel: logrus.WarnLevel,
		Limit:    10,
		Burst:    20,
	}
	h := New(cfg, "hdds://broken")

	logger := newHookedLogger(h)
	logger.Warn("my warn")

	hookWait()

	assert.Contains(t, log.String(), "my warn")
	// Async doesn't return errors
	assert.Empty(t, stderr.String())

	resetBufs()
}

func ExampleNew() {
	cfg := Config{
		MinLevel: logrus.ErrorLevel,
		Message: chat.Message{
			Username:  "My App",
			Channel:   "#slack-testing",
			IconEmoji: ":ghost:",
		},
	}
	h := New(cfg, "https://hooks.slack.com/services/T00000000/B00000000/XXXXXXXXXXXXXXXXXXXXXXXX")
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetLevel(logrus.InfoLevel)
	logrus.AddHook(h)
	logrus.WithFields(logrus.Fields{"field1": "test field", "field2": 1}).Error("test error")
}
