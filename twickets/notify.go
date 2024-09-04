package twickets

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/gotify/go-api-client/v2/auth"
	"github.com/gotify/go-api-client/v2/client"
	"github.com/gotify/go-api-client/v2/client/message"
	"github.com/gotify/go-api-client/v2/gotify"
	"github.com/gotify/go-api-client/v2/models"
)

type NotificationClient interface {
	SendTicketNotification(Ticket) error
}

func notificationMessage(ticket Ticket) string {
	lines := []string{
		fmt.Sprintf(
			"%s %s",
			ticket.Event.Time.Format("3:04pm"),
			ticket.Event.Date.Format("Monday 2 January 2006"),
		),
		fmt.Sprintf("%d ticket(s)", ticket.TicketQuantity),
		fmt.Sprintf(
			"Ticket Price: %s",
			priceString(
				(ticket.TotalSellingPrice.Number()+ticket.TotalTwicketsFee.Number())/float64(ticket.TicketQuantity),
				ticket.TotalSellingPrice.Currency,
			),
		),
		fmt.Sprintf("Original Price: %s", ticket.FaceValuePrice.PerString(ticket.TicketQuantity)),
	}

	return strings.Join(lines, "\n")
}

type GotifyClient struct {
	url    *url.URL
	token  string
	client *client.GotifyREST
}

var _ NotificationClient = GotifyClient{}

func (g GotifyClient) SendTicketNotification(ticket Ticket) error {
	params := message.NewCreateMessageParams()
	params.Body = &models.MessageExternal{
		Title:   ticket.Event.Name,
		Message: notificationMessage(ticket),
		Extras: map[string]any{
			"extras": map[string]any{
				"client::display": map[string]any{
					"contentType": "text/markdown",
				},
				"client::notification": map[string]any{
					"click": map[string]any{
						"url": ticket.Link(),
					},
				},
			},
		},
		Priority: 5,
	}

	_, err := g.client.Message.CreateMessage(
		params,
		auth.TokenAuth(g.token),
	)
	if err != nil {
		return err
	}

	return nil
}

func NewGotifyClient(gotifyUrl, gotifyToken string) (*GotifyClient, error) {
	parsedGotifyUrl, err := url.Parse(gotifyUrl)
	if err != nil {
		return nil, fmt.Errorf("failed to parse gotify url: %v", err)
	}

	return &GotifyClient{
		url:    parsedGotifyUrl,
		token:  gotifyToken, // TODO validate this token?
		client: gotify.NewClient(parsedGotifyUrl, &http.Client{}),
	}, nil
}
