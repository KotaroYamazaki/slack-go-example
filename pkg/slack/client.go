package slack

import (
	"context"
	"encoding/json"
	"strconv"
	"time"

	"github.com/slack-go/slack"
)

const (
	ColorGreen  = "good"
	ColorRed    = "danger"
	ColorYellow = "warning"

	blkNum = 1
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
	SectionText   string
}

func New(webhook string) *Client {
	return &Client{
		webhook: webhook,
	}
}

func (c *Client) PostWebhook(ctx context.Context, msg *slack.WebhookMessage) error {
	return slack.PostWebhookContext(ctx, c.webhook, msg)
}

func (c *Client) BuildWebhookMessage(params *MessageParams) *slack.WebhookMessage {
	blks := make([]slack.Block, 0, blkNum)
	if params.SectionText != "" {
		blks = append(blks, c.buildTextSectionBlk(params.SectionText, params.ButtonText, params.ButtonURL))
	}

	attachment := slack.Attachment{
		Color:      params.Color,
		AuthorName: params.AuthorName,
		AuthorLink: params.AuthorLink,
		AuthorIcon: params.AuthorIconURL,
		Text:       params.Text,
		Footer:     params.Footer,
		FooterIcon: params.FooterIconURL,
		Ts:         json.Number(strconv.FormatInt(params.Timestamp.Unix(), 10)),
	}

	return &slack.WebhookMessage{
		Attachments: []slack.Attachment{attachment},
		Blocks: &slack.Blocks{
			BlockSet: blks,
		},
	}
}

func (c *Client) buildTextSectionBlk(text, buttonText, buttonURL string) *slack.SectionBlock {
	var acs *slack.Accessory
	if buttonText != "" && buttonURL != "" {
		acs = slack.NewAccessory(
			&slack.ButtonBlockElement{
				Type: slack.METButton,
				Text: &slack.TextBlockObject{
					Type:  slack.PlainTextType,
					Text:  buttonText,
					Emoji: true,
				},
				URL: buttonURL,
			})
	}

	return &slack.SectionBlock{
		Type: slack.MBTSection,
		Text: &slack.TextBlockObject{
			Type: slack.MarkdownType,
			Text: text,
		},
		Accessory: acs,
	}
}
