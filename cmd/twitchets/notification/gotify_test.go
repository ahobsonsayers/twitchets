package notification_test

import (
	"os"
	"testing"

	"github.com/ahobsonsayers/twitchets/cmd/twitchets/notification"
	"github.com/ahobsonsayers/twitchets/test/testutils"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/require"
)

func TestGotifySendTicketMessage(t *testing.T) {
	t.Skip("Currently don't have a gotify server. Figure out how to test")
	testutils.SkipIfCI(t, "No env set in CI. Fix")

	_ = godotenv.Load(testutils.ProjectDirectoryJoin(t, ".env"))

	gotifyUrl := os.Getenv("GOTIFY_URL")
	require.NotEmpty(t, gotifyUrl, "GOTIFY_URL is not set")

	gotifyToken := os.Getenv("GOTIFY_TOKEN")
	require.NotEmpty(t, gotifyToken, "GOTIFY_TOKEN is not set")

	client, err := notification.NewGotifyClient(gotifyUrl, gotifyToken)
	require.NoError(t, err)

	ticket := testNotificationTicket()
	err = client.SendTicketNotification(ticket)
	require.NoError(t, err)
}
