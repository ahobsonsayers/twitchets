package notification_test

import (
	"os"
	"strconv"
	"testing"

	"github.com/ahobsonsayers/twitchets/config"
	"github.com/ahobsonsayers/twitchets/notification"
	"github.com/ahobsonsayers/twitchets/test"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/require"
)

func TestPushoverSendTicketMessage(t *testing.T) {
	t.Skip("Can only be run manually locally with environment variables set. Comment to run.")

	_ = godotenv.Load(test.ProjectDirectoryJoin(t, ".env"))

	pushoverToken := os.Getenv("PUSHOVER_TOKEN")
	require.NotEmpty(t, pushoverToken, "PUSHOVER_TOKEN is not set")

	pushoverUser := os.Getenv("PUSHOVER_USER")
	require.NotEmpty(t, pushoverUser, "PUSHOVER_USER is not set")

	// Optional: test with device
	pushoverDevice := os.Getenv("PUSHOVER_DEVICE")
	var devicePtr *string
	if pushoverDevice != "" {
		devicePtr = &pushoverDevice
	}

	// Optional: test with sound
	pushoverSound := os.Getenv("PUSHOVER_SOUND")
	var soundPtr *string
	if pushoverSound != "" {
		soundPtr = &pushoverSound
	}

	// Optional: test with priority
	var priorityPtr *int
	if pushoverPriorityStr := os.Getenv("PUSHOVER_PRIORITY"); pushoverPriorityStr != "" {
		priority, err := strconv.Atoi(pushoverPriorityStr)
		require.NoError(t, err, "PUSHOVER_PRIORITY is not an integer")
		priorityPtr = &priority
	}

	client, err := notification.NewPushoverClient(config.PushoverConfig{
		Token:    pushoverToken,
		User:     pushoverUser,
		Device:   devicePtr,
		Sound:    soundPtr,
		Priority: priorityPtr,
	})
	require.NoError(t, err)

	ticket := testNotificationTicket()
	err = client.SendTicketNotification(ticket)
	require.NoError(t, err)
}
