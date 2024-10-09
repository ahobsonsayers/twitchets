package notification

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/ahobsonsayers/twitchets/twickets"
	"github.com/gotify/go-api-client/v2/auth"
	"github.com/gotify/go-api-client/v2/client"
	"github.com/gotify/go-api-client/v2/client/message"
	"github.com/gotify/go-api-client/v2/gotify"
	"github.com/gotify/go-api-client/v2/models"
)

type GotifyClient struct {
	url   *url.URL
	token string

	client *client.GotifyREST
}

var _ Client = GotifyClient{}

func (g GotifyClient) SendTicketNotification(ticket twickets.Ticket) error {
	notificationMessage, err := RenderMessage(ticket, WithFooter())
	if err != nil {
		return err
	}

	params := message.NewCreateMessageParams()
	params.Body = &models.MessageExternal{
		Title:   ticket.Event.Name,
		Message: notificationMessage,
		Extras: map[string]any{
			"client::display": map[string]any{
				"contentType": "text/markdown",
			},
			"client::notification": map[string]any{
				"click": map[string]any{
					"url": ticket.Link(),
				},
			},
		},
		Priority: 5,
	}

	_, err = g.client.Message.CreateMessage(
		params,
		auth.TokenAuth(g.token),
	)
	if err != nil {
		return err
	}

	return nil
}

func NewGotifyClient(gotifyUrl, gotifyToken string) (GotifyClient, error) {
	parsedGotifyUrl, err := url.Parse(gotifyUrl)
	if err != nil {
		return GotifyClient{}, fmt.Errorf("failed to parse gotify url: %v", err)
	}

	return GotifyClient{
		url:    parsedGotifyUrl,
		token:  gotifyToken, // TODO validate this token?
		client: gotify.NewClient(parsedGotifyUrl, &http.Client{}),
	}, nil
}
