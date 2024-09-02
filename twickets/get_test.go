package twickets_test

import (
	"context"
	"testing"

	"github.com/ahobsonsayers/twitchets/twickets"
	"github.com/stretchr/testify/require"
)

func TestGetLatestTickets(t *testing.T) {
	tickets, err := twickets.FetchLatestTickets(
		context.Background(),
		twickets.GetTicketsInput{
			Country:   twickets.CountryUnitedKingdom,
			MaxNumber: 10,
		},
	)
	require.NoError(t, err)
	_ = tickets
}
