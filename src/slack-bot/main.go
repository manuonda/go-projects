package main

import (
	"os"

	"github.com/slack-go/slack"
)

func main() {

	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))

}
