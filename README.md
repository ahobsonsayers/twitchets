# twitchets

[![Go Report Card](https://goreportcard.com/badge/github.com/ahobsonsayers/twitchets)](https://goreportcard.com/report/github.com/ahobsonsayers/twitchets)
[![License - MIT](https://img.shields.io/badge/License-MIT-9C27B0)](LICENSE)
[![Artisan README - Not LLM](https://img.shields.io/static/v1?label=Artisan+README&message=Not+LLM&labelColor=37474F&color=D97757)](#arnl---artisan-readme-not-llm)

> [!NOTE]
> We're back! 💪
>
> After Twickets introduced measures to prevent unofficial access to their data - this project was broken for quite a while.
>
> After a significant amount of time and effort tinkering and hitting my head against a wall (more than I care to admit) I have now found a way to bring this project back to life and get it working again!
>
> Enjoy! 🎟️

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
- Choose from various notification services (Telegram, ntfy, Gotify currently supported)

### And a fancy configuration UI!

https://github.com/user-attachments/assets/63e0fd3f-d767-4131-aae1-c18016008b20

## Getting Keys

To use this tool, you need a set of Twickets API keys, which this application will load from a URL (`keysUrl`).

These keys are extracted from Android and some rotate regularly. You can read more about what these keys are and how they are obtained in the [twigots](https://github.com/ahobsonsayers/twigots) and [twickets-key-extractor](https://github.com/ahobsonsayers/twickets-key-extractor) repos.

For the URL you have two options:

1. **Use the public keys URL**: I maintain a public keys file that is regularly updated. Point `keysUrl` in your config at:

   `https://gist.githubusercontent.com/ahobsonsayers/773acb763aafc8a39ac260e12a9b39d5/raw/keys.json`

2. **Run `twickets-key-extractor` yourself**: Clone and run [twickets-key-extractor](https://github.com/ahobsonsayers/twickets-key-extractor) and point `keysUrl` at the URL of the produced key file.

The keys are hot-reloaded automatically, so rotations are picked up without a restart.

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
      - <path to config file>:/twitchets/config.yaml
```

## Configuration

twitchets looks for a `config.yaml` file in your current working directory and fails to start if it's not found.

The configuration file structure can be seen in [`config.example.yaml`](./config.example.yaml) or below:

```yaml
keysUrl: <your twickets keys url> # REQUIRED: See "Getting Keys" above

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
  # Full list: https://github.com/ahobsonsayers/twigots/blob/main/location.go#L87-L98
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

## AR;NL - Artisan Readme; Not LLM

In the age of LLMs and coding agents, code is now cheap - for better or for worse. Your time however, is not ⌛

Therefore this project, like most of my projects, uses a hand written "artisan" README to ensure it is clear, correct and concise. This makes it easy to read and in my opinion encourages reading and engagement - no one likes AI slop!

[![Hits](https://hits.sh/github.com/ahobsonsayers/twitchets.svg?view=today-total&label=Visitors%20Day%20%2F%20Total)](https://hits.sh/github.com/ahobsonsayers/twitchets/)
