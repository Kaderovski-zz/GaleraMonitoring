package slackApp

import (
	"fmt"
 "github.com/ashwanthkumar/slack-go-webhook"
)

// Create an app in slack API and add an Incoming Webhoosk
// more informations : https://api.slack.com/apps/
func PayloadSlack(text string) {

	// Add here your webhookurl
	webhookUrl := "https://hooks.slack.com/services/your_key"

	payload := slack.Payload{
		Text:    text,
		Username: "Galera_alerte",
		Channel:  "#infra",
	}
	err := slack.Send(webhookUrl, "", payload)
	if len(err) > 0 {
		fmt.Printf("error : %s\n", err)
	}
}
