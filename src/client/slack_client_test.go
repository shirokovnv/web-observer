package client_test

import (
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/shirokovnv/web-observer/src/client"
)

var (
	mock  client.HTTPClientMock
	slack client.SlackConfig
)

func setUpSlack() {
	mock = client.HTTPClientMock{}
	client.Client = &mock
	slack = client.SlackConfig{
		WebhookURL: "https://webhook.io",
		Username:   "user",
		Channel:    "Channel",
	}
}

func TestSuccessfulResponse(t *testing.T) {
	setUpSlack()
	mock.DoFunc = func(r *http.Request) (*http.Response, error) {
		return &http.Response{
			Body:       io.NopCloser(strings.NewReader("ok")),
			StatusCode: 200,
		}, nil
	}

	sr := client.SlackRequest{Text: "Hello, Slack", IconEmoji: ":hot-coffee:"}
	err := slack.SendSlackNotification(sr)
	if err != nil {
		t.Fatal("Error occured", err)
	}
}

func TestFailedResponse(t *testing.T) {
	setUpSlack()
	mock.DoFunc = func(r *http.Request) (*http.Response, error) {
		return &http.Response{
			Body:       io.NopCloser(strings.NewReader("error")),
			StatusCode: 500,
		}, nil
	}

	sr := client.SlackRequest{Text: "Hello, Slack", IconEmoji: ":hot-coffee:"}
	err := slack.SendSlackNotification(sr)
	if err == nil {
		t.Fatal("Must be an error response")
	}
}
