package main

import "fmt"

func CoinMarketCapIDToURL(id int) string {
	return fmt.Sprintf("https://s2.coinmarketcap.com/static/img/coins/128x128/%d.png", id)
}
