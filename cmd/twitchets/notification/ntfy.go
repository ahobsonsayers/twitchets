package notification

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/ahobsonsayers/twitchets/twickets"
	"github.com/samber/lo"
	"heckel.io/ntfy/client"
)

type NtfyClient struct {
	url      *url.URL
	user     string
	password string

	client *client.Client
}

var _ Client = NtfyClient{}

func (c NtfyClient) SendTicketNotification(ticket twickets.Ticket) error {
	notificationMessage, err := RenderMessage(ticket, nil)
	if err != nil {
		return err
	}

	_, err = c.client.Publish(
		c.url.String(),
		notificationMessage,
		client.WithTitle(ticket.Event.Name),
		client.WithActions(NtfyViewAction("Open Link", lo.ToPtr(ticket.Link()))),
		client.WithHeader("Content-Type", "text/markdown"),
		client.WithBasicAuth(c.user, c.password),
	)
	if err != nil {
		return err
	}

	return nil
}

type NtfyConfig struct {
	Url      string `json:"url"`
	Topic    string `json:"topic"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func NewNtfyClient(config NtfyConfig) (NtfyClient, error) {
	ntfyUrl, err := url.Parse(config.Url)
	if err != nil {
		return NtfyClient{}, fmt.Errorf("failed to parse ntfy url: %v", err)
	}

	ntfyUrl.Path = config.Topic

	return NtfyClient{
		url:      ntfyUrl,
		user:     config.Username,
		password: config.Password,

		client: client.New(nil),
	}, nil
}

// NtfyViewAction creates a ntfy actions string for a single view actions
// See https://docs.ntfy.sh/publish/#using-a-json-array
func NtfyViewAction(label string, link *string, params ...map[string]string) string {
	// Combine params
	combinedParams := map[string]string{}
	if link != nil {
		combinedParams["url"] = *link
	}
	for _, paramsMap := range params {
		for key, value := range paramsMap {
			combinedParams[key] = value
		}
	}

	return ntfyActionString(ntfyAction{
		action: "view",
		label:  label,
		params: combinedParams,
	})
}

type ntfyAction struct {
	action string
	label  string
	params map[string]string
}

// ntfyActionString creates a ntfy actions string from parameters
// See https://docs.ntfy.sh/publish/#using-a-json-array
func ntfyActionString(actions ...ntfyAction) string {
	if len(actions) == 0 {
		return ""
	}

	// Create a slice of maps for the action
	actionMaps := make([]map[string]string, 0, len(actions))
	for _, actions := range actions {
		actionMap := make(map[string]string, 2+len(actions.params))
		actionMap["action"] = actions.action
		actionMap["label"] = actions.label
		for name, value := range actions.params {
			actionMap[name] = value
		}

		actionMaps = append(actionMaps, actionMap)
	}

	actionsJson, _ := json.Marshal(actionMaps)
	return string(actionsJson)
}
