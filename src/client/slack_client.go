package client

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

const DefaultSlackTimeout = 5 * time.Second

type SlackConfig struct {
	WebhookURL string        `yaml:"webhook_url"`
	Username   string        `yaml:"username"`
	Channel    string        `yaml:"channel"`
	TimeOut    time.Duration `yaml:"timeout"`
}

type SlackRequest struct {
	Text      string
	IconEmoji string
}

type SlackMessage struct {
	Username  string `json:"username,omitempty"`
	IconEmoji string `json:"icon_emoji,omitempty"`
	Channel   string `json:"channel,omitempty"`
	Text      string `json:"text,omitempty"`
}

func (sc SlackConfig) SendSlackNotification(sr SlackRequest) error {
	slackRequest := SlackMessage{
		Text:      sr.Text,
		Username:  sc.Username,
		IconEmoji: sr.IconEmoji,
		Channel:   sc.Channel,
	}
	return sc.sendHttpRequest(slackRequest)
}

func (sc SlackConfig) sendHttpRequest(slackRequest SlackMessage) error {
	slackBody, _ := json.Marshal(slackRequest)

	if sc.TimeOut == 0 {
		sc.TimeOut = DefaultSlackTimeout
	}
	ctx, cncl := context.WithTimeout(context.Background(), sc.TimeOut)
	defer cncl()

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, sc.WebhookURL, bytes.NewBuffer(slackBody))
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/json")

	resp, err := Client.Do(req)
	if err != nil {
		return err
	}

	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(resp.Body)
	if err != nil {
		return err
	}

	if buf.String() != "ok" {
		return errors.New("non-ok response returned from Slack")
	}
	return nil
}
