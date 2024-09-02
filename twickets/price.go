package twickets

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/orsinium-labs/enum"
)

type Currency enum.Member[string]

func (c Currency) Symbol() string {
	switch c {
	case CurrencyGBP:
		return "£"
	default:
		return ""
	}
}

func (c *Currency) UnmarshalJSON(data []byte) error {
	var currencyString string
	err := json.Unmarshal(data, &currencyString)
	if err != nil {
		return err
	}

	currency := Currencies.Parse(currencyString)
	if currency == nil {
		return fmt.Errorf("currency '%s' is not valid", currencyString)
	}
	*c = *currency
	return nil
}

var (
	currency = enum.NewBuilder[string, Currency]()

	CurrencyGBP = currency.Add(Currency{"GBP"})

	Currencies = currency.Enum()
)

type Price struct {
	Currency Currency `json:"currencyCode"`
	Amount   int      `json:"amountInCents"` // In cents, pennies etc.
}

func (p *Price) Number() float64 {
	if p == nil {
		return 0
	}
	return float64(p.Amount) / 100
}

func (p *Price) String() string {
	if p == nil {
		return ""
	}
	return priceString(p.Number(), p.Currency)
}

func (p *Price) PerString(quantity int) string {
	if p == nil {
		return ""
	}
	return priceString(p.Number()/float64(quantity), p.Currency)
}

func priceString(cost float64, currency Currency) string {
	costString := strconv.FormatFloat(cost, 'f', 2, 64)
	currencyString := currency.Symbol()
	if currencyString == "" {
		return costString + currency.Value
	}

	return currencyString + costString
}
