package notification_test

import (
	"os"
	"testing"

	"github.com/ahobsonsayers/twitchets/config"
	"github.com/ahobsonsayers/twitchets/notification"
	"github.com/ahobsonsayers/twitchets/test"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/require"
)

func TestGotifySendTicketMessage(t *testing.T) {
	t.Skip("Can only be run manually locally with environment variables set. Comment to run.")

	_ = godotenv.Load(test.ProjectDirectoryJoin(t, ".env"))

	gotifyUrl := os.Getenv("GOTIFY_URL")
	require.NotEmpty(t, gotifyUrl, "GOTIFY_URL is not set")

	gotifyToken := os.Getenv("GOTIFY_TOKEN")
	require.NotEmpty(t, gotifyToken, "GOTIFY_TOKEN is not set")

	client, err := notification.NewGotifyClient(config.GotifyConfig{
		Url:   gotifyUrl,
		Token: gotifyToken,
	})
	require.NoError(t, err)

	ticket := testNotificationTicket()
	err = client.SendTicketNotification(ticket)
	require.NoError(t, err)
}
