package twickets_test

import (
	"context"
	"os"
	"testing"

	"github.com/ahobsonsayers/twitchets/test/testutils"
	"github.com/ahobsonsayers/twitchets/twickets"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/require"
)

func TestGetLatestTickets(t *testing.T) {
	_ = godotenv.Load(testutils.ProjectDirectoryJoin(t, ".env"))

	twicketsAPIKey := os.Getenv("TWICKETS_API_KEY")
	require.NotEmpty(t, twicketsAPIKey, "TWICKETS_API_KEY is not set")

	httpClient, err := testutils.NewProxyClient(
		testutils.RoosterKidProxyListURL,
		testutils.ProxlifyProxyListURL,
		testutils.TheSpeedXProxyListURL,
	)
	require.NoError(t, err)

	twicketsClient := twickets.NewClient(httpClient)
	tickets, err := twicketsClient.FetchTickets(
		context.Background(),
		twickets.FetchTicketsInput{
			APIKey:  twicketsAPIKey,
			Country: twickets.CountryUnitedKingdom,
			Regions: []twickets.Region{
				twickets.RegionLondon,
				twickets.RegionNorthWest,
			},
			NumTickets: 10,
		},
	)
	require.NoError(t, err)
	require.Len(t, tickets, 10)
}
