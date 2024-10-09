package config

import (
	"encoding/json"
	"fmt"

	"github.com/ahobsonsayers/twitchets/cmd/twitchets/notification"
	"github.com/orsinium-labs/enum"
)

type NotificationType enum.Member[string]

var (
	notificationType = enum.NewBuilder[string, NotificationType]()

	NotificationTypeNtfy     = notificationType.Add(NotificationType{"ntfy"})
	NotificationTypeTelegram = notificationType.Add(NotificationType{"telegram"})

	NotificationTypes = notificationType.Enum()
)

func (c *NotificationType) UnmarshalJSON(data []byte) error {
	var notificationTypeString string
	err := json.Unmarshal(data, &notificationTypeString)
	if err != nil {
		return err
	}

	notificationType := NotificationTypes.Parse(notificationTypeString)
	if notificationType == nil {
		return fmt.Errorf("notificationType '%s' is not valid", notificationTypeString)
	}

	*c = *notificationType
	return nil
}

func (c *NotificationType) UnmarshalText(data []byte) error {
	notificationTypeString := string(data)
	notificationType := NotificationTypes.Parse(notificationTypeString)
	if notificationType == nil {
		return fmt.Errorf("notificationType '%s' is not valid", notificationTypeString)
	}

	*c = *notificationType
	return nil
}

type NotificationConfig struct {
	Ntfy     *notification.NtfyConfig     `json:"ntfy"`
	Telegram *notification.TelegramConfig `json:"telegram"`
}

func (c NotificationConfig) Clients() (map[NotificationType]notification.Client, error) {
	clients := map[NotificationType]notification.Client{}

	if c.Ntfy != nil {
		ntfyClient, err := notification.NewNtfyClient(*c.Ntfy)
		if err != nil {
			return nil, fmt.Errorf("failed to setup ntfy client: %w", err)
		}

		clients[NotificationTypeNtfy] = ntfyClient
	}

	if c.Telegram != nil {
		telegramClient, err := notification.NewTelegramClient(*c.Telegram)
		if err != nil {
			return nil, fmt.Errorf("failed to setup telegram client: %w", err)
		}

		clients[NotificationTypeTelegram] = telegramClient
	}

	return clients, nil
}
