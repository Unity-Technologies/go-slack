// slackit is a command line golang slack client.
//
// It reads a JSON formatted message from the file specified by -src <file>
// and posts it to slack using the webhook url specified by -hook <url>.
//
// By default slackit reads the message from stdin.
//
// Example:
//  slackit -hook https://hooks.slack.com/services/T00/B00/XXX -src msg.json
package main

import (
	"encoding/json"
	"flag"
	"log"
	"os"

	"github.com/multiplay/go-slack/chat"
	"github.com/multiplay/go-slack/webhook"
)

// exit is the function used to exit on error its a variable so it can be easily overridden for tests.
var exit = os.Exit

func main() {
	var src = flag.String("src", "-", "reads the message from the specified file or stdin if '-'")
	var hook = flag.String("hook", "", "the hook url to use")

	flag.Usage = func() {
		log.Printf("usage %s -hook <url> [-src <file>]\nflags:\n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.Parse()
	if *hook == "" {
		flag.Usage()
		exit(1)
	}

	var f *os.File

	if *src == "-" {
		f = os.Stdin
	} else {
		var err error
		if f, err = os.Open(*src); err != nil {
			log.Println("failed to open source: ", err)
			exit(1)
		}
	}

	dec := json.NewDecoder(f)
	m := &chat.Message{}
	if err := dec.Decode(m); err != nil {
		log.Println("failed to decode message: ", err)
		exit(1)
	}

	c := webhook.New(*hook)
	if _, err := m.Send(c); err != nil {
		log.Println("failed to send message:", err)
		exit(1)
	}
}
