package notification_test

import (
	"os"
	"testing"

	"github.com/ahobsonsayers/twitchets/cmd/twitchets/notification"
	"github.com/ahobsonsayers/twitchets/test/testutils"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/require"
)

func TestNtfySendTicketMessage(t *testing.T) {
	t.Skip("Can only be run manually locally with environment variables set. Comment to run.")

	_ = godotenv.Load(testutils.ProjectDirectoryJoin(t, ".env"))

	ntfyUrl := os.Getenv("NTFY_URL")
	require.NotEmpty(t, ntfyUrl, "NTFY_URL is not set")

	ntfyUser := os.Getenv("NTFY_USER")
	require.NotEmpty(t, ntfyUser, "NTFY_USER is not set")

	ntfyPassword := os.Getenv("NTFY_PASSWORD")
	require.NotEmpty(t, ntfyPassword, "NTFY_PASSWORD is not set")

	ntfyTopic := os.Getenv("NTFY_TOPIC")
	require.NotEmpty(t, ntfyTopic, "NTFY_PASSWORD is not set")

	client, err := notification.NewNtfyClient(notification.NtfyConfig{
		Url:      ntfyUrl,
		Username: ntfyUser,
		Password: ntfyPassword,
		Topic:    ntfyTopic,
	})
	require.NoError(t, err)

	ticket := testNotificationTicket()
	err = client.SendTicketNotification(ticket)
	require.NoError(t, err)
}
