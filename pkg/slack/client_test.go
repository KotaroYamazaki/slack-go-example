package slack

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"github.com/joho/godotenv"
)

var c *Client

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println(err)
	}
	webhookURL := os.Getenv("SLACK_WEBHOOK_URL")
	c = New(webhookURL)
}

func BenchmarkPostWebhook(b *testing.B) {
	ctx := context.Background()

	fmt.Printf("%#v\n", c.webhook)
	params := &MessageParams{
		Color:         ColorGreen,
		AuthorName:    "test",
		AuthorLink:    "https://github.com",
		AuthorIconURL: "https://avatars.githubusercontent.com/u/7589567?v=4",
		Text:          "hello world",
		Footer:        "github.com",
		FooterIconURL: "https://github.githubassets.com/images/modules/logos_page/Octocat.png",
		Timestamp:     time.Now(),
		ButtonText:    "Click Me",
		ButtonURL:     "https://github.com/KotaroYamazaki/slack-go-sample",
		SectionText:   "*go to repository*",
	}
	msg := c.BuildWebhookMessage(params)

	for i := 0; i < b.N; i++ {
		err := c.PostWebhook(ctx, msg)
		if err != nil {
			b.Fatal(err)
		}
	}
}
