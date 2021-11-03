# slack-go sample

this is an exmample of [slack-go](https://github.com/slack-go/slack)

# setup

## get package

```
$ go get -u github.com/slack-go/slack
```

## create webhook

create webhook url from https://slack.com/services/new/incoming-webhook

## set your slack webhook to .env

```
SLACK_WEBHOOK={YOUR_SLACK_WEBHOOK}
echo "SLACK_WEBHOOK_URL=$SLACK_WEBHOOK" > .env
```

# usage

1. set post params in main.go

e.g.

```
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
```

2. run

```
go run main.go
```

# result

![image](https://user-images.githubusercontent.com/7589567/136746375-7debe5b1-bd4a-45bf-8b58-65ed1eaa54f8.png)
