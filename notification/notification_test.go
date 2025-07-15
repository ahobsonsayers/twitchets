package notification_test

import (
	"os"
	"testing"

	"github.com/ahobsonsayers/twigots"
	"github.com/ahobsonsayers/twitchets/notification"
	"github.com/ahobsonsayers/twitchets/test"
	"github.com/stretchr/testify/require"
)

func testNotificationTicket() twigots.TicketListing {
	return twigots.TicketListing{
		Id: "test",
		Event: twigots.Event{
			Name: "Test Event",
			Venue: twigots.Venue{
				Name: "Test Venue",
				Location: twigots.Location{
					Name: "Test Location",
				},
			},
		},
		TicketType: "Standing",
		NumTickets: 2,
		TotalPriceExclFee: twigots.Price{
			Currency: twigots.CurrencyGBP,
			Amount:   200,
		},
		TwicketsFee: twigots.Price{
			Currency: twigots.CurrencyGBP,
			Amount:   100,
		},
		OriginalTotalPrice: twigots.Price{
			Currency: twigots.CurrencyGBP,
			Amount:   400,
		},
		SellerWillConsiderOffers: true,
	}
}

func TestRenderMessage(t *testing.T) {
	expectedMessagePath := test.ProjectDirectoryJoin(
		t, "test", "data", "message", "message.md",
	)
	expectedMessageBytes, err := os.ReadFile(expectedMessagePath)
	require.NoError(t, err)
	expectedMessage := string(expectedMessageBytes)

	tickets := testNotificationTicket()
	actualMessage, err := notification.RenderMessage(tickets)
	require.NoError(t, err)

	require.Equal(t, expectedMessage, actualMessage)
}

func TestRenderMessageWithHeaderAndFooter(t *testing.T) {
	expectedMessagePath := test.ProjectDirectoryJoin(
		t, "test", "data", "message", "messageWithHeaderFooter.md",
	)
	expectedMessageBytes, err := os.ReadFile(expectedMessagePath)
	require.NoError(t, err)
	expectedMessage := string(expectedMessageBytes)

	tickets := testNotificationTicket()
	actualMessage, err := notification.RenderMessage(
		tickets,
		notification.WithHeader(),
		notification.WithFooter(),
	)
	require.NoError(t, err)

	require.Equal(t, expectedMessage, actualMessage)
}
