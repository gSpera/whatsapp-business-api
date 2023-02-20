package main

import (
	"fmt"

	"github.com/gSpera/whatsapp-business-api"
)

func main() {
	c := whatsapp.NewClient(ACCESS_TOKEN)
	err := c.SendMessage(PHONE_NUMBER_ID, whatsapp.SendMessage{
		MessagingProduct: "whatsapp",
		Type:             "text",
		To:               "phone-number-to-use",

		Text: &whatsapp.SendMessageText{
			Body: "Message to send",
		},
	})

	if err != nil {
		fmt.Println("Error:", err)
	}
}
