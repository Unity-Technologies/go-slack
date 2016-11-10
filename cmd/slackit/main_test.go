package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/multiplay/go-slack/test"
	"github.com/stretchr/testify/assert"
)

func testSlackit(t *testing.T, hook, msg string) ([]byte, error) {
	args := make([]string, 0, 2)
	if hook != "" {
		args = append(args, "-hook", hook)
	}

	switch {
	case msg == "invalid-file":
		args = append(args, "-src", "invalid-file")
	case msg != "":
		tmpfile, err := ioutil.TempFile("", "slackit-test")
		if err != nil {
			t.Fatal(err)
		}

		defer os.Remove(tmpfile.Name())

		if _, err := tmpfile.Write([]byte(msg)); err != nil {
			t.Fatal(err)
		}
		if err := tmpfile.Close(); err != nil {
			t.Fatal(err)
		}

		args = append(args, "-src", tmpfile.Name())
	}

	var buf bytes.Buffer
	var err error

	exit = func(code int) {
		if code != 0 {
			err = fmt.Errorf("exit(%v)", code)
		}
	}
	os.Args = append([]string{"slackit"}, args...)
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	flag.CommandLine.SetOutput(&buf)
	log.SetOutput(&buf)
	main()

	return buf.Bytes(), err
}

func TestNoArgs(t *testing.T) {
	b, err := testSlackit(t, "", "")
	assert.Error(t, err)
	assert.Contains(t, string(b), "usage")
}

func TestMissingHook(t *testing.T) {
	b, err := testSlackit(t, "", "-")
	assert.Error(t, err)
	assert.Contains(t, string(b), "usage")
}

func TestInvalidMessage(t *testing.T) {
	b, err := testSlackit(t, test.Endpoint, "broken")
	assert.Error(t, err)
	assert.Contains(t, string(b), "failed to decode message")
}

func TestInvalidFile(t *testing.T) {
	b, err := testSlackit(t, test.Endpoint, "invalid-file")
	assert.Error(t, err)
	assert.Contains(t, string(b), "failed to decode message")
}

func TestSuccess(t *testing.T) {
	b, err := testSlackit(t, test.Endpoint, `{"text":"my message"}`)
	assert.NoError(t, err)
	assert.Empty(t, b)
}
