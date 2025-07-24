# twitchets

[![Go Report Card](https://goreportcard.com/badge/github.com/ahobsonsayers/twitchets)](https://goreportcard.com/report/github.com/ahobsonsayers/twitchets)
[![License - MIT](https://img.shields.io/badge/License-MIT-9C27B0)](LICENSE)

A tool to watch for chosen event ticket listings on [Twickets](https://www.twickets.live) that match custom filters and notify you so you can quickly snap them up!

**Why twitchets?**

I built this tool because the official Twickets app has limitations on the number of events you can watch for tickets, and lacks many features/filters I wanted and needed.

**Note**: This program does **not** buy tickets, reserve them automatically, or do anything unethical. It simply notifies you of new ticket listings!

Powered by [twigots](https://github.com/ahobsonsayers/twigots), a Go package to fetch and filter event ticket listings from Twickets 🎟️

## Features

- No limit on the number of events you can watch for!
- Watch for tickets with a certain discount, number of tickets, and location
- Show more details in the notifications, such as event date/time, number of tickets, and discount
- Faster notifications than the official Twickets app
- No need to have the Twickets app or an account
- Choose from various notification services (Telegram, Ntfy, Gotify currently supported)

## Getting an API Key

To use this tool, you will need a Twickets API key. Although Twickets doesn't provide a free API, you can easily obtain a key by following these steps:

1. Visit the [Twickets Live Feed](https://www.twickets.live/app/catalog/browse)
2. Open your browser's Developer Tools (F12) and navigate to the Network tab
3. Look for the GET request to `https://www.twickets.live/services/catalogue` and copy the `api_key` query parameter. You might need to refresh the page first if nothing appears in this tab.

This API key is not provided here due to liability concerns, but it appears to be a fixed, unchanging value.

## Installation & Running

The recommended way to run twitchets is using Docker:

```bash
docker run -d \
    --name twitchets \
    -v <path to config>:/twitchets/config.yaml \
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
apiKey: <your twickets api key> # REQUIRED: See README.md for details on how to obtain

country: GB # Currently only GB is supported

# Notification service configuration
# Remove/comment out services you don't need
notification:
  ntfy:
    url: <your ntfy url> # You can use the public instance at https://ntfy.sh
    topic: <your ntfy topic> # If using https://ntfy.sh, make sure this is unique to you!
    username: <your ntfy username> # Optional: for authenticated instances
    password: <your ntfy password> # Optional: for authenticated instances

  telegram:
    token: <your telegram api token> # Get from @BotFather on Telegram
    chatId: <your telegram chat id> # Your chat ID or group chat ID

  gotify:
    url: <your gotify url> # Your Gotify server URL
    token: <your gotify api token> # Application token from Gotify

# Global ticket configuration
# All available settings are outlined below
# These settings apply to all tickets by default
# Individual ticket configuration can override these settings
# Settings can be added and removed as needed
# Any setting not specified will use the default
global:
  # Geographic regions to search for tickets
  # Default: All regions if not specified
  # Full list: https://github.com/ahobsonsayers/twigots/blob/main/location.go#L79-L90
  regions:
    - GBLO # London only

  # Event name similarity matching (0.0 - 1.0)
  # Default: 0.9 (allows for minor naming differences)
  eventSimilarity: 0.9

  # Minimum number of tickets required in listing
  # Default: Any number of tickets
  numTickets: 2 # Exactly two tickets

  # Maximum price per ticket (including fee) in pounds (£)
  # Default: Any price
  maxTicketPrice: 50 # Maximum ticket price of £50 including fee

  # Minimum discount (including fee) on the original price as a percentage
  # Default: Any discount (including no discount)
  # discount: 10 # At least 10% off original price

  # Notification services to use
  # Default: All configured services
  notification:
    - ntfy

# Individual ticket configuration
# Available settings match the global ones
# Settings here override global settings above
# To reset a global setting to its default, use:
# - "" (empty string) for string values
# - [] (empty array) for list values
# - -1 for numeric values
tickets:
  - event: Lion King
    maxTicketPrice: 30 # Max £30 per ticket

  - event: Coldplay
    numTickets: 4 # Need exactly 4 tickets
    maxTicketPrice: -1 # Reset to default: Any max price
    discount: 25 # Must be at least 25% off

  - event: Taylor Swift
    regions: [] # Reset to default: Search all regions
    numTickets: -1 # Reset to default: Any number of tickets
    discount: -1 # Reset to default: Any discount (or no discount)

  - event: Hamilton
    regions:
      - GBSO # South only
    notification:
      - telegram # Only send to Telegram

  - event: Oasis
    notification: [] # Reset to default: Send to all configured notification services
```

## How does the event name matching/similarity work?

You can see more about how this works in the [twigots readme here](https://github.com/ahobsonsayers/twigots#how-does-the-event-name-matchingsimilarity-work).

## Why the name twitchets?

Because I feel like sometimes you need to have twitch-like reactions to snap up tickets on Twickets before someone else gets them - which this tool helps you do. Therefore the mangling together of **twitch** and **Twickets** seemed fun and appropriate.

[![Hits](https://hits.sh/github.com/ahobsonsayers/twitchets.svg?view=today-total&label=Visitors%20Day%20%2F%20Total)](https://hits.sh/github.com/ahobsonsayers/twitchets/)
