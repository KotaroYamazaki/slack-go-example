# Slack-go Sample

ref: https://github.com/slack-go/slack

# setup

## get package

```
$ go get -u github.com/slack-go/slack
```

## get webhook

create webhook url from https://slack.com/services/new/incoming-webhook

# usage

1. set your slack webhook to .env

```
SLACK_WEBHOOK={YOUR_SLACK_WEBHOOK}
echo "SLACK_WEBHOOK_URL=$SLACK_WEBHOOK" > .env
```

2. run

```
go run main.go
```

# result

![image](https://user-images.githubusercontent.com/7589567/136696917-ad2e1463-1f49-40b5-a8ab-16c63315e547.png)
