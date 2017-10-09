package main

import (
	"fmt"

	"github.com/kasmetski/cmcAPI"
)

//GetCoinInfo - Get full infomartion from Coin Market Cap
func GetCoinInfo(s string) (msg string, err error) {
	coin, err := cmcAPI.GetCoinInfo(s)
	if err != nil {
		msg = "Can't find anything. Type the full name of the coin"
		return
	}

	name := fmt.Sprintf("Name: %s | %s | #%d\n", coin.Name, coin.Symbol, coin.Rank)
	supply := fmt.Sprintf("Available Supply: %d\n", int(coin.AvailableSupply))
	mcap := fmt.Sprintf("MarketCap: %d\n", int(coin.MarketCapUsd))
	volume := fmt.Sprintf("Volume: %d\n", int(coin.Two4HVolumeUsd))
	price := fmt.Sprintf("PriceBTC: %f\nPriceUSD: %.2f\n", coin.PriceBtc, coin.PriceUsd)
	change := fmt.Sprintf("Change 1H/24H/7d: %.2f | %.2f | %.2f\n", coin.PercentChange1H, coin.PercentChange24H, coin.PercentChange7D)

	msg = name + supply + mcap + volume + price + change

	return
}

//GetCoinPrice - Get price information from Coin Market Cap
func GetCoinPrice(s string) (msg string, err error) {
	coin, err := cmcAPI.GetCoinInfo(s)
	if err != nil {
		msg = "Can't find anything. Type the full name of the coin"
		return
	}

	name := fmt.Sprintf("Name: %s | %s | #%d\n", coin.Name, coin.Symbol, coin.Rank)
	price := fmt.Sprintf("PriceBTC: %f\nPriceUSD: %.2f\n", coin.PriceBtc, coin.PriceUsd)
	change := fmt.Sprintf("Change 1H/24H/7d: %.2f | %.2f | %.2f\n", coin.PercentChange1H, coin.PercentChange24H, coin.PercentChange7D)

	msg = name + price + change

	return
}
