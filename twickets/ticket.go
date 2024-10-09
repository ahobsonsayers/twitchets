package twickets // nolint

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Ticket struct {
	Id string `json:"blockId"`

	CreatedAt UnixTime `json:"created"`
	ExpiresAt UnixTime `json:"expires"`

	TicketQuantity           int   `json:"ticketQuantity"`
	TicketsPrice             Price `json:"totalSellingPrice"`
	TwicketsFee              Price `json:"totalTwicketsFee"`
	OriginalTotalPrice       Price `json:"faceValuePrice"`
	SellerWillConsiderOffers bool  `json:"sellerWillConsiderOffers"`

	TicketType   string `json:"priceTier"` // Seated, Standing, Box etc.
	SeatAssigned bool   `json:"seatAssigned"`
	Section      string `json:"section"` // Can be empty
	Row          string `json:"row"`     // Can be empty

	Event Event `json:"event"`
	Tour  Tour  `json:"tour"`
}

func (t Ticket) Link() string {
	// Link: https://www.twickets.live/app/block/<ticketId>,<quanitity>
	return fmt.Sprintf("https://www.twickets.live/app/block/%s,%d", t.Id, t.TicketQuantity)
}

// TotalPrice is total price of all tickets.
// This the tickets price plus the twickets fee.
func (t Ticket) TotalPrice() Price {
	return t.TicketsPrice.Add(t.TwicketsFee)
}

// TotalTicketPrice is total price of a single ticket.
// This the tickets price plus the twickets fee divided by the number of tickets.
func (t Ticket) TotalTicketPrice() Price {
	return t.TotalPrice().Divide(t.TicketQuantity)
}

// OriginalTicketsPrice is original price of a single ticket.
// This the original tickets price divided by the number of tickets.
func (t Ticket) OriginalTicketPrice() Price {
	return t.OriginalTotalPrice.Divide(t.TicketQuantity)
}

// Discount is the price discount as a percentage
func (t Ticket) Discount() float64 {
	return (1 - t.TotalPrice().Number()/t.OriginalTotalPrice.Number()) * 100
}

func (t Ticket) DiscountString() string {
	discountString := strconv.FormatFloat(t.Discount(), 'f', 2, 64)
	return discountString + "%"
}

type Event struct {
	Id       string `json:"id"`
	Name     string `json:"eventName"`
	Category string `json:"category"`
	Date     Date   `json:"date"`
	Time     Time   `json:"showStartingTime"`

	OnSaleTime    *DateTime `json:"onSaleTime"` // 2023-11-17T10:00:00Z
	AnnouncedTime *DateTime `json:"created"`    // 2023-11-17T10:00:00Z

	Venue  Venue    `json:"venue"`
	Lineup []Lineup `json:"participants"`
}

type Lineup struct {
	Artist  Artist `json:"participant"`
	Billing int    `json:"billing"`
}

type Artist struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Slug string `json:"linkName"`
}

type Venue struct {
	Id       string   `json:"id"`
	Name     string   `json:"name"`
	Location Location `json:"location"`
	Postcode string   `json:"postcode"`
}

type Location struct {
	Id          string  `json:"id"`
	Name        string  `json:"shortName"`
	FullName    string  `json:"name"`
	CountryCode Country `json:"countryCode"`
	RegionCode  Region  `json:"regionCode"`
}

type Tour struct {
	Id           string   `json:"tourId"`
	Name         string   `json:"tourName"`
	Slug         string   `json:"slug"`
	FirstEvent   *Date    `json:"minDate"`      // 2024-06-06
	LastEvent    *Date    `json:"maxDate"`      // 2024-11-14
	CountryCodes []string `json:"countryCodes"` // TODO use enum
}

func UnmarshalTwicketsFeedJson(data []byte) ([]Ticket, error) {
	response := struct {
		ResponseData []struct { // nolint
			Tickets *Ticket `json:"catalogBlockSummary"`
		} `json:"responseData"`
	}{}
	err := json.Unmarshal(data, &response)
	if err != nil {
		return nil, err
	}

	// Get non null tickets. Tickets are null if they have been delisted
	tickets := make([]Ticket, 0, len(response.ResponseData))
	for _, responseData := range response.ResponseData {
		if responseData.Tickets != nil {
			tickets = append(tickets, *responseData.Tickets)
		}
	}

	return tickets, nil
}

type Tickets []Ticket

// GetById gets the first ticket with a matching id,
// or return nil if it does not exist.
func (t Tickets) GetById(id string) *Ticket {
	for _, ticket := range t {
		if ticket.Id == id {
			return &ticket
		}
	}
	return nil
}

// Filter filters tickets by a set of conditions
func (t Tickets) Filter(filters ...Filter) Tickets {
	if len(filters) == 0 {
		return t
	}

	filteredTickets := make(Tickets, 0, len(t))
	for _, ticket := range t {
		for _, filter := range filters {
			if filter.TicketMatches(ticket) {
				filteredTickets = append(filteredTickets, ticket)
			}
		}
	}

	return filteredTickets
}

var spaceRegex = regexp.MustCompile(`\s+`)

func normaliseEventName(eventName string) string {
	eventName = strings.TrimSpace(eventName)
	eventName = spaceRegex.ReplaceAllString(eventName, " ")
	eventName = strings.ToLower(eventName)
	eventName = strings.TrimPrefix(eventName, "the ")
	return eventName
}
