package main

import (
	"hoge/pkg/slack"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println(err)
	}
}

func main() {
	webhookURL := mustGetenv("SLACK_WEBHOOK_URL")
	c := slack.New(webhookURL)
	params := &slack.MessageParams{
		Color:         "green",
		AuthorName:    "KotaroYamazaki",
		AuthorLink:    "https://github.com/KotaroYamazaki",
		AuthorIconURL: "https://avatars.githubusercontent.com/u/7589567?v=4",
		Text:          "hello world",
		Footer:        "",
		FooterIconURL: "",
		Timestamp:     time.Now(),
		ButtonText:    "Click Me",
		ButtonURL:     "https://github.com/KotaroYamazaki",
	}
	msg := c.BuildWebhookMessage(params)
	if err := c.PostWebhook(msg); err != nil {
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
