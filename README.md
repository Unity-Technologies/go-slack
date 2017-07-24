# Slack [![Go Report Card](https://goreportcard.com/badge/github.com/multiplay/go-slack)](https://goreportcard.com/report/github.com/multiplay/go-slack) [![License](https://img.shields.io/badge/license-BSD-blue.svg)](https://github.com/multiplay/go-slack/blob/master/LICENSE) [![GoDoc](https://godoc.org/github.com/multiplay/go-slack?status.svg)](https://godoc.org/github.com/multiplay/go-slack) [![Build Status](https://travis-ci.org/multiplay/go-slack.svg?branch=master)](https://travis-ci.org/multiplay/go-slack)

go-slack is a [Go](http://golang.org/) library for the [Slack API](https://api.slack.com/).

Features
--------
* [Slack Webhook](https://api.slack.com/incoming-webhooks) Support.
* [Slack chat.postMessage](https://api.slack.com/methods/chat.postMessage) Support.
* Client Interface - Use alternative implementations - currently webhook is the only client.
* [Logrus Hook](https://github.com/sirupsen/logrus) Support - Automatically send messages to [Slack](https://slack.com) when using a [Logrus](https://github.com/sirupsen/logrus) logger.

Installation
------------
```sh
go get -u github.com/multiplay/go-slack
```

Examples
--------

The simplest way to use go-slack is to create a webhook client and send chat messages using it e.g.
```go
package main

import (
	"github.com/multiplay/go-slack/chat"
	"github.com/multiplay/go-slack/webhook"
)

func main() {
	c := webhook.New("https://hooks.slack.com/services/T00000000/B00000000/XXXXXXXXXXXXXXXXXXXXXXXX")
	m := &chat.Message{Text: "test message"}
	m.Send(c)
}
```

If your using [logrus](https://github.com/sirupsen/logrus) you can use the webhook to post to slack based on your logging e.g.
```go
package main

import (
	"github.com/multiplay/go-slack/lrhook"
	"github.com/sirupsen/logrus"
)

func main() {
	cfg := lrhook.Config{
		MinLevel:       logrus.ErrorLevel,
		Message:	chat.Message{
			Channel:	"#slack-testing",
			IconEmoji:	":ghost:",
		},
	}
	h := lrhook.New(cfg, "https://hooks.slack.com/services/T00000000/B00000000/XXXXXXXXXXXXXXXXXXXXXXXX")
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetLevel(logrus.InfoLevel)
	logrus.AddHook(h)
	logrus.Error("my error)
}
```

Documentation
-------------
- [GoDoc API Reference](http://godoc.org/github.com/multiplay/go-slack).

License
-------
go-slack is available under the [BSD 2-Clause License](https://opensource.org/licenses/BSD-2-Clause).
