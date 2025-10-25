package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
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

func (c NotificationConfig) Validate() error {
	if c.Ntfy != nil {
		if !beginsWithHttp(c.Ntfy.Url) {
			return errors.New("ntfy url must begin with 'http://' or 'https://'")
		}
		if c.Ntfy.Topic == "" {
			return errors.New("ntfy topic must be set")
		}
	}

	if c.Gotify != nil {
		if !beginsWithHttp(c.Gotify.Url) {
			return errors.New("gotify url must begin with 'http://' or 'https://'")
		}
		if c.Gotify.Token == "" {
			return errors.New("gotify token cannot be empty")
		}
	}

	if c.Telegram != nil {
		if c.Telegram.ChatId == 0 {
			return errors.New("telegram chat id cannot be empty")
		}
		if c.Telegram.Token == "" {
			return errors.New("telegram token cannot be empty")
		}
	}

	return nil
}

func beginsWithHttp(url string) bool {
	return strings.HasPrefix(url, "http://") || strings.HasPrefix(url, "https://")
}
