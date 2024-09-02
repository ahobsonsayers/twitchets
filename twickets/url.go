package twickets

import (
	"fmt"
	"log"
	"net/url"
	"strconv"
	"strings"
)

const (
	TwicketsURL    = "https://www.twickets.live"
	TwicketsAPIKey = "83d6ec0c-54bb-4da3-b2a1-f3cb47b984f1"
)

var twicketsUrl *url.URL // https://www.twickets.live

func init() {
	// Parse twickets url
	var err error
	twicketsUrl, err = url.Parse(TwicketsURL)
	if err != nil {
		log.Fatal("failed to parse twickets url")
	}
}

// TicketURL gets a ticket url
// https://www.twickets.live/app/block/<ticketId>,<numTickets>
func TicketURL(ticketId string, numTickets int) string {
	ticketUrl := cloneURL(twicketsUrl)
	ticketUrl = ticketUrl.JoinPath("app", "block", fmt.Sprintf("%s,%d", ticketId, numTickets))
	return ticketUrl.String()
}

// FeedUrl gets the feel url. By default gets the last minute of tickets up to a maximum of 100.
// https://www.twickets.live/services/catalogue?api_key=83d6ec0c-54bb-4da3-b2a1-f3cb47b984f1&count=100&q=countryCode=GB
func FeedUrl(input GetTicketsInput) string {
	feedUrl := cloneURL(twicketsUrl)
	feedUrl = feedUrl.JoinPath("services", "catalogue")

	// Set query params
	queryParams := feedUrl.Query()
	queryParams.Set("api_key", TwicketsAPIKey)

	locationQuery := apiLocationQuery(input.Country, input.Regions...)
	if locationQuery != "" {
		queryParams.Set("q", locationQuery)
	}

	// var maxTime int64
	// if input.BeforeTime.IsZero() {
	// 	maxTime = time.Now().UnixMilli()
	// } else {
	// 	maxTime = input.BeforeTime.UnixMilli()
	// }
	// queryParams.Set("maxTime", strconv.Itoa(int(maxTime)))

	if input.MaxNumber > 0 {
		count := strconv.Itoa(input.MaxNumber)
		queryParams.Set("count", count)
	}

	// Set query
	encodedQuery := queryParams.Encode()
	encodedQuery = strings.ReplaceAll(encodedQuery, "%3D", "=")
	encodedQuery = strings.ReplaceAll(encodedQuery, "%2C", ",")
	feedUrl.RawQuery = encodedQuery

	return feedUrl.String()
}

// apiLocationQuery converts a country and selection of regions to an api query string
func apiLocationQuery(country Country, regions ...Region) string {
	if !Countries.Contains(country) {
		return ""
	}

	queryParts := make([]string, 0, len(regions)+1)

	countryQuery := fmt.Sprintf("%s=%s", countryQueryKey, country.Value)
	queryParts = append(queryParts, countryQuery)

	for _, region := range regions {
		if Regions.Contains(region) {
			regionQuery := fmt.Sprintf("%s=%s", regionQueryKey, region.Value)
			queryParts = append(queryParts, regionQuery)
		}
	}

	return strings.Join(queryParts, ",")
}

// cloneURL clones a url. Copied directly from net/http internals
// See: https://github.com/golang/go/blob/go1.19/src/net/http/clone.go#L22
func cloneURL(u *url.URL) *url.URL {
	if u == nil {
		return nil
	}
	u2 := new(url.URL)
	*u2 = *u
	if u.User != nil {
		u2.User = new(url.Userinfo)
		*u2.User = *u.User
	}
	return u2
}
