package notification

import (
	"github.com/ahobsonsayers/twitchets/twickets"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type TelegramClient struct {
	client *tgbotapi.BotAPI
	chatId int
}

var _ Client = TelegramClient{}

func (c TelegramClient) SendTicketNotification(ticket twickets.Ticket) error {
	messageBody, err := RenderMessage(ticket, WithHeader(), WithFooter())
	if err != nil {
		return err
	}

	message := tgbotapi.NewMessage(int64(c.chatId), messageBody)
	message.ParseMode = tgbotapi.ModeMarkdown

	_, err = c.client.Send(message)
	if err != nil {
		return err
	}

	return nil
}

type TelegramConfig struct {
	APIToken string `json:"apiToken"`
	ChatId   int    `json:"chatId"`
}

func NewTelegramClient(config TelegramConfig) (TelegramClient, error) {
	client, err := tgbotapi.NewBotAPI(config.APIToken)
	if err != nil {
		return TelegramClient{}, err
	}

	return TelegramClient{
		client: client,
		chatId: config.ChatId,
	}, nil
}
