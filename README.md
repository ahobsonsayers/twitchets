# twitchets

[![Go Report Card](https://goreportcard.com/badge/github.com/ahobsonsayers/twitchets)](https://goreportcard.com/report/github.com/ahobsonsayers/twitchets)
[![License - MIT](https://img.shields.io/badge/License-MIT-9C27B0)](LICENSE)

A tool to watch for chosen event ticket listings on [Twickets](https://www.twickets.live) that match custom filters and notify you so you can quickly snap them up!

**Why twitchets?**

I built this tool because the official Twickets app has limitations on the number of events you can watch for tickets, and lacks many features/filters I wanted and needed.

**Note**: This program does **not** buy tickets, reserve them automatically, or do anything unethical. It simply notifies you of new ticket listings!

Powered by [twigots](https://github.com/ahobsonsayers/twigots), a go package to fetch and filter event ticket listings from Twickets 🎟️

## Features

- No limit on the number of events you can watch for!
- On watch for tickets with a certain discount, number of tickets, and location
- Show more details in the notifications, such as event date/time, number of tickets, and discount
- Faster notifications than the official Twickets app
- No need to have the Twickets app or an account
- Choose from various notification services (Telegram, Ntfy, Gotify currently supported)

## Getting an API Key

To use this tool, you will need a Twickets API key. Although Twickets doesn't provide a free API, you can easily obtain a key by following these steps:

1.  Visit the [Twickets Live Feed](https://www.twickets.live/app/catalog/browse)
2.  Open your browser's Developer Tools (F12) and navigate to the Network tab
3.  Look for the GET request to `https://www.twickets.live/services/catalogue` and copy the `api_key` query parameter. You might need to refresh the page first if nothing appears in this tab.

This API key is not provided here due to liability concerns, but it appears to be a fixed, unchanging value.

## Running

The best way to run twitchets is using Docker:

```bash
docker run -d \
    --name twitchets \
    -v <path to config>:/twitchets/config.yaml
    --restart unless-stopped \
    arranhs/twitchets:latest
```

Or, use Docker Compose:

```yaml
services:
  twitchets:
    container_name: twitchets
    image: arranhs/twitchets:latest
    restart: unless-stopped
    volumes:
      - <path to config>:/twitchets
```

## Configuration

twitchets looks for a `config.yaml` file in your current working directory and fails to start if it's not found.

The configuration file structure can be seen in [`config.example.yaml`](./config.example.yaml) or below:

```yaml
apiKey: <your twickets api key> # REQUIRED: See README.md for details on how to obtain.

country: GB # Currently only GB is supported

# Configure your notification services.
# Delete the ones you do not need.
notification:
  ntfy:
    url: <your ntfy url> - you can use the public instance at https://ntfy.sh
    topic: <your ntfy topic> - if using https://ntfy.sh, make sure this is unique to you!
    username: <your ntfy username>
    password: <your ntfypassword>
  telegram:
    token: <your telegram api token>
    chatId: <your telegram chat id>
  gotify:
    url: <your gotify url>
    token: <your gotify api token>

# Global configuration that applies to all tickets.
# These can be overridden on a ticket-by-ticket basis (see below).
# The below are examples of all global configuration settings.
# Remove the settings you do not need.
global:
  # Regions which ticket listings must be in.
  # Defaults to all regions if not set.
  # Full list of regions can be seen here:
  # https://github.com/ahobsonsayers/twigots/blob/main/location.go#L79-L90
  regions:
    - GBLO # Only look for tickets in London

  # How similar the event name in the ticket listing must be to the one you specified.
  # Defaults to 0.85 if not set, allowing for naming inconsistencies.  # Between 0-1
  # Between 0-1
  eventSimilarity: 0.85

  # How many tickets must be in the listing.
  # Defaults to any number of tickets if not set.
  numTickets: 2

  # Minimum discount (including fees) of the tickets in the listing against the original price.
  # Defaults to any discount if not set.
  discount: 15

  # Notification services to send found ticket listings to.
  # Defaults to all configured services above.
  notification:
    - ntfy

# Ticket configuration.
# Settings set here take priority over global settings.
# If a setting is not set/specified here, the global setting will be used.
# To unset a global setting and use the default, set it to a negative value or an empty list, depending on the type.
# The below are examples. Tweak them as needed.
tickets:
  # Look for Lion King tickets using the global configuration.
  # For global configuration settings not set, defaults will be used.
  - event: Lion King # REQUIRED

  # Look for Taylor Swift tickets, overriding the global configuration.
  - event: Taylor Swift # REQUIRED

    # Override global configuration setting.
    # Event name must be an exact match.
    # For example, this will NOT match "Taylor Swift: The Eras Tour".
    eventSimilarity: 1

    # Override global configuration setting.
    # Reset to default - watch for tickets from any region.
    regions: []

    # Override global configuration setting.
    # Reset to default - watch for tickets with any discount.
    numTickets: -1

    # Override global configuration setting.
    # Reset to default - watch for tickets with any discount.
    discount: -1

    # Override global configuration setting:
    # Send notifications to Telegram (instead of ntfy).
    notification:
      - telegram

  # Look for Oasis tickets (mostly) using the global configuration.
  # Notifications will be sent to all configured notification services.
  - event: Oasis # REQUIRED
    notification: []
```

## How does the event name matching/similarity work?

You can see more about how this works in the [twigots readme here](https://github.com/ahobsonsayers/twigots#how-does-the-event-name-matchingsimilarity-work)

## Why the name twitchets?

Because I feel like sometimes you need to have twitch like reactions to snap up tickets on Twickets before someone else gets them - which this tool helps you do. Therefore the mangling together of **twitch** and **Twickets** seemed fun and appropriate.

[![Hits](https://hits.sh/github.com/ahobsonsayers/twitchets.svg?view=today-total&label=Visitors%20Day%20%2F%20Total)](https://hits.sh/github.com/ahobsonsayers/twitchets/)
