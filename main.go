package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/KotaroYamazaki/slack-go-sample/pkg/slack"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println(err)
	}
}

func main() {
	ctx := context.Background()

	webhookURL := mustGetenv("SLACK_WEBHOOK_URL")
	c := slack.New(webhookURL)
	params := &slack.MessageParams{
		Color:         slack.ColorGreen,
		AuthorName:    "KotaroYamazaki",
		AuthorLink:    "https://github.com/KotaroYamazaki",
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

	if err := c.PostWebhook(ctx, msg); err != nil {
		log.Panicln(err)
	}
}

func mustGetenv(k string) string {
	v := os.Getenv(k)
	if v == "" {
		log.Panicf("%s environment variable not set.", k)
	}
	return v
}
