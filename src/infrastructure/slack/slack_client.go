package slack

import (
	"github.com/IkezawaYuki/bookshelf-go/src/usecase/outputport"
	"os"
)

type client struct {
	webhookUrl string
	channel    string
}

func NewSlackClient() outputport.SlackOutputPort {
	return &client{
		webhookUrl: os.Getenv("WEBHOOK_URL"),
		channel:    os.Getenv("CHANNEL"),
	}
}

func (c client) SendMessage(message outputport.SlackMessage) error {
	panic("implement me")
}
