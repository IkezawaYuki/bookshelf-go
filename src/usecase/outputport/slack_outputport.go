package outputport

type SlackOutputPort interface {
	SendMessage(message SlackMessage) error
}

type SlackMessage interface {
	GetTitle() string
	GetBody() string
}
