// Package lrhook provides logrus hook for the Slack.
//
// It can post messages to slack based on the notification level of the
// logrus entry including the ability to rate limit messages.
//
// See: https://godoc.org/github.com/sirupsen/logrus#Hook
package lrhook

import (
	"fmt"

	"github.com/multiplay/go-slack"
	"github.com/multiplay/go-slack/chat"
	"github.com/multiplay/go-slack/webhook"

	"github.com/sirupsen/logrus"
	"golang.org/x/time/rate"
)

var (
	// DefaultLevelColors is the default level colors used if none are present in the configuration.
	DefaultLevelColors = map[string]string{
		"debug":   "#9B30FF",
		"info":    "good",
		"warning": "danger",
		"error":   "danger",
		"fatal":   "panic",
		"panic":   "panic",
	}

	// DefaultUnknownColor is the default UnknownColor if one is not present in the configuration.
	DefaultUnknownColor = "warning"
)

// Config is the configuration of a slack logrus.Hook.
type Config struct {
	// MinLevel is the minimal level at which the hook will trigger.
	MinLevel logrus.Level

	// LevelColors is a hash of logrus level names to colors used for the attachment in the messages.
	LevelColors map[string]string

	// AttachmentText is the text message used for the attachment when fields are present in the log entry.
	AttachmentText string

	// UnknownColor is the color to use if there is no match for the log level in LevelColors.
	UnknownColor string

	// Async if true then messages are sent to slack asynchronously.
	// This means that Fire will never return an error.
	Async bool

	// Limit if none zero limits the number of messages to Limit posts per second.
	Limit rate.Limit

	// Burst sets the burst limit.
	// Ignored if Limit is zero.
	Burst int

	// Message defines the details of the messages sent from the hook.
	Message chat.Message

	// Attachment defines the details of the attachment sent from the hook.
	// Field Text - Will be set to that of log entry Message.
	// Field Fields - Will be created to match the log entry Fields.
	// Field Color - Will be set according to the LevelColors or UnknownColor if a match is not found..
	Attachment chat.Attachment
}

// Hook is a logrus hook that sends messages to Slack.
type Hook struct {
	Config
	client  slack.Client
	limiter *rate.Limiter
}

// SetConfigDefaults sets defaults on the configuration if needed to ensure the cfg is valid.
func SetConfigDefaults(cfg *Config) {
	if len(cfg.LevelColors) == 0 {
		cfg.LevelColors = DefaultLevelColors
	}
	if cfg.UnknownColor == "" {
		cfg.UnknownColor = DefaultUnknownColor
	}
}

// New returns a new Hook with the given configuration that posts messages using the webhook URL.
// It ensures that the cfg is valid by calling SetConfigDefaults on the cfg.
func New(cfg Config, url string) *Hook {
	return NewClient(cfg, webhook.New(url))
}

// NewClient returns a new Hook with the given configuration using the slack.Client c.
// It ensures that the cfg is valid by calling SetConfigDefaults on the cfg.
func NewClient(cfg Config, client slack.Client) *Hook {
	SetConfigDefaults(&cfg)

	c := &Hook{Config: cfg, client: client}
	if cfg.Limit != 0 {
		c.limiter = rate.NewLimiter(cfg.Limit, cfg.Burst)
	}

	return c
}

// Levels implements logrus.Hook.
// It returns the logrus.Level's that are lower or equal to that of MinLevel.
// This means setting MinLevel to logrus.ErrorLevel will send slack messages for log entries at Error, Fatal and Panic.
func (sh *Hook) Levels() []logrus.Level {
	lvls := make([]logrus.Level, 0, len(logrus.AllLevels))
	for _, l := range logrus.AllLevels {
		if sh.MinLevel >= l {
			lvls = append(lvls, l)
		}
	}

	return lvls
}

// Fire implements logrus.Hook.
// It sends a slack message for the log entry e.
func (sh *Hook) Fire(e *logrus.Entry) error {
	if sh.limiter != nil && !sh.limiter.Allow() {
		// We've hit the configured limit, just ignore.
		return nil
	}

	m := sh.Message
	a := sh.Attachment
	m.AddAttachment(&a)
	a.Fallback = e.Message
	a.Color = sh.LevelColors[e.Level.String()]
	if a.Color == "" {
		a.Color = sh.UnknownColor
	}
	a.Text = e.Message
	for k, v := range e.Data {
		a.NewField(k, fmt.Sprint(v))
	}

	if sh.Async {
		go m.Send(sh.client)
		return nil
	}

	_, err := m.Send(sh.client)
	return err
}
