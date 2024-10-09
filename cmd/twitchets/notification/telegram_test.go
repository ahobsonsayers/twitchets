package notification_test

import (
	"os"
	"strconv"
	"testing"

	"github.com/ahobsonsayers/twitchets/cmd/twitchets/notification"
	"github.com/ahobsonsayers/twitchets/test/testutils"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/require"
)

func TestTelegramSendTicketMessage(t *testing.T) {
	t.Skip("Can only be run manually locally with environment variables set. Comment to run.")

	_ = godotenv.Load(testutils.ProjectDirectoryJoin(t, ".env"))

	telegramAPIKey := os.Getenv("TELEGRAM_API_KEY")
	require.NotEmpty(t, telegramAPIKey, "TELEGRAM_API_KEY is not set")

	telegramChatIdString := os.Getenv("TELEGRAM_CHAT_ID")
	require.NotEmpty(t, telegramChatIdString, "TELEGRAM_CHAT_ID is not set")

	telegramChatId, err := strconv.Atoi(telegramChatIdString)
	require.NoError(t, err, "TELEGRAM_CHAT_ID is not an integer")

	client, err := notification.NewTelegramClient(notification.TelegramConfig{
		APIToken: telegramAPIKey,
		ChatId:   telegramChatId,
	})
	require.NoError(t, err)

	ticket := testNotificationTicket()
	err = client.SendTicketNotification(ticket)
	require.NoError(t, err)
}
