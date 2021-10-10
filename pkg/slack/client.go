package slack

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/slack-go/slack"
)

type Client struct {
	webhook string
}

type MessageParams struct {
	Color         string
	AuthorName    string
	AuthorLink    string
	AuthorIconURL string
	Text          string
	Footer        string
	FooterIconURL string
	Timestamp     time.Time
	ButtonText    string
	ButtonURL     string
}

func New(webhook string) *Client {
	return &Client{
		webhook: webhook,
	}
}

func (c *Client) PostWebhook(msg *slack.WebhookMessage) error {
	return slack.PostWebhook(c.webhook, msg)
}

func (c *Client) BuildWebhookMessage(params *MessageParams) *slack.WebhookMessage {
	attachment := slack.Attachment{
		Color:      params.Color,
		AuthorName: params.AuthorName,
		AuthorLink: params.AuthorName,
		AuthorIcon: params.AuthorIconURL,
		Text:       params.AuthorName,
		Footer:     params.Footer,
		FooterIcon: params.FooterIconURL,
		Ts:         json.Number(strconv.FormatInt(params.Timestamp.Unix(), 10)),
	}
	var blk slack.Block
	if params.ButtonText != "" && params.ButtonURL != "" {
		blk = c.addClickButton(params.ButtonText, params.ButtonURL)
	}

	return &slack.WebhookMessage{
		Attachments: []slack.Attachment{attachment},
		Blocks: &slack.Blocks{
			BlockSet: []slack.Block{blk},
		},
	}
}

func (c *Client) addClickButton(text, URL string) *slack.ActionBlock {
	return &slack.ActionBlock{
		Type: slack.MBTAction,
		Elements: &slack.BlockElements{
			ElementSet: []slack.BlockElement{
				&slack.ButtonBlockElement{
					Type:  slack.METButton,
					Text:  &slack.TextBlockObject{Type: "plain_text", Text: text, Emoji: true},
					Value: "click_me_123",
					URL:   URL,
				},
			},
		},
	}
}
