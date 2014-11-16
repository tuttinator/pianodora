package main

import "fmt"

const (
	slackEndpoint = "https://%s.slack.com/services/hooks/incoming-webhook?token=%s"
)

type SlackAttachment struct {
	Fallback string   `json:"fallback"`
	Text     string   `json:"text"`
	Color    string   `json:"color"`
	MrkdwnIn []string `json:"mrkdwn_in"`
}

type SlackMessage struct {
	Channel     string            `json:"channel"`
	Username    string            `json:"username"`
	Attachments []SlackAttachment `json:"attachments"`
	UnfurlMedia bool              `json:"unfurl_media"`
}

type Slack struct {
	Team    string `yaml:"team,omitempty"`
	Channel string `yaml:"channel,omitempty"`
	Token   string `yaml:"token,omitempty"`
	Active  bool   `yaml:"on_active,omitempty"`
	Failed  bool   `yaml:"on_failed,omitempty"`
	Restart bool   `yaml:"on_restart,omitempty"`
}

func (s *Slack) Send(message PandoraMessage) error {
	client := NewNotifierHTTPClient()
	m := s.composeMessage("good", message)
	return s.send(client, m)
}

func (s *Slack) composeMessage(color string, message PandoraMessage) *SlackMessage {
	attachments := SlackAttachment{
		message.Title, message, color, []string{"fallback", "text"},
	}

	return &SlackMessage{s.Channel, "Pianodora Bot", []SlackAttachment{attachments}}
}

func (s *Slack) send(client HTTPClient, message *SlackMessage) error {
	url := fmt.Sprintf(slackEndpoint, s.Team, s.Token)
	return client.PostJSON(url, message)
}
