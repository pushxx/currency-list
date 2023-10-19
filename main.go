package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"github.com/gocolly/colly"
	"github.com/samber/lo"
	"log"
	"os"
	"strings"
	"sync"
	"time"
)

var (
	collector = &colly.Collector{}
	baseUrl   = "https://coinmarketcap.com/"
	maxPage   = 90

	CurrenciesInfo = sync.Map{}
)

func init() {
	collector = colly.NewCollector(
		colly.AllowURLRevisit(),
		colly.Async(true),
		colly.CacheDir("./.cache"),
		//colly.Debugger(&debug.LogDebugger{}),
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/95.0.4638.54 Safari/537.36"),
	)

}

func main() {
	collector.OnRequest(func(r *colly.Request) {
		r.Headers.Set("Referer", "https://coinmarketcap.com/") // 设置Referer头部
		r.Headers.Set("Accept-Language", "en-US,en;q=0.9")     // 设置Accept-Language头部
		fmt.Println("Visiting", r.URL.String())
	})

	collector.Limit(&colly.LimitRule{
		DomainGlob:  "*coinmarketcap.*",
		Parallelism: 2,
		Delay:       10 * time.Second,
	})

	collector.OnHTML("table > tbody > tr > td:nth-child(3) a", func(e *colly.HTMLElement) {
		href, _ := e.DOM.Attr("href")
		e.Request.Visit(href)
	})
	collector.OnHTML("body", func(element *colly.HTMLElement) {
		if strings.Contains(element.Request.URL.String(), "/currencies/") {
			jsonp := element.DOM.Find(`script#__NEXT_DATA__`).Text()
			coinMarketCapCurrencyInfo := CoinMarketCapCurrencyInfo{}
			json.Unmarshal([]byte(jsonp), &coinMarketCapCurrencyInfo)
			detail := coinMarketCapCurrencyInfo.Props.PageProps.DetailRes.Detail

			currency := &Currency{
				Logo:        CoinMarketCapIDToURL(detail.Id),
				Name:        detail.Name,
				Symbol:      detail.Symbol,
				Slug:        detail.Slug,
				Type:        detail.Category,
				Description: detail.Description,
				Links:       detail.Urls,
				Platforms: lo.Map(detail.Platforms, func(v *CoinMarketCapCurrencyInfoPlatform, _ int) *Platform {
					return &Platform{
						Logo:                fmt.Sprintf("https://s2.coinmarketcap.com/static/img/coins/128x128/%d.png", v.ContractId),
						ContractAddress:     v.ContractAddress,
						ContractPlatform:    v.ContractPlatform,
						ContractPRCURL:      v.ContractRpcUrl,
						ContractExplorerURL: v.ContractExplorerUrl,
					}
				}),
				Wallets: lo.Map(detail.Wallets, func(v *CoinMarketCapCurrencyInfoWallets, _ int) *Wallet {
					return &Wallet{Logo: CoinMarketCapIDToURL(v.Id), Name: v.Name, URL: v.Url}
				}),
				AuditInfos: lo.Map(detail.AuditInfos, func(v *CoinMarketCapCurrencyInfoAuditInfo, _ int) *AuditInfo {
					return &AuditInfo{
						Auditor:   v.Auditor,
						ReportUrl: v.ReportUrl,
					}
				}),
			}

			CurrenciesInfo.LoadOrStore(currency.Slug, currency)
		}
	})
	collector.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", string(r.Body), "\nError:", err)
	})

	for i := 1; i <= maxPage; i++ {
		collector.Visit(fmt.Sprintf("%s?page=%d", baseUrl, i))
	}

	collector.Wait()

	f, err := os.Create("./data.csv")
	if err != nil {
		log.Fatalf("create file error: %v", err)
	}
	defer f.Close()

	writerCsv := csv.NewWriter(f)

	writerCsv.Write([]string{"Name", "Symbol", "Slug", "Logo", "Type", "Description", "Links", "Platforms", "Wallets", "AuditInfos"})

	CurrenciesInfo.Range(func(key, value interface{}) bool {
		if currency, ok := value.(*Currency); ok {
			links, _ := json.Marshal(currency.Links)
			platforms, _ := json.Marshal(currency.Platforms)
			wallets, _ := json.Marshal(currency.Wallets)
			auditInfos, _ := json.Marshal(currency.AuditInfos)
			line := []string{
				currency.Name,
				currency.Symbol,
				currency.Slug,
				currency.Logo,
				currency.Type,
				fmt.Sprintf("%s", currency.Description),
				fmt.Sprintf("%s", links),
				fmt.Sprintf("%s", platforms),
				fmt.Sprintf("%s", wallets),
				fmt.Sprintf("%s", auditInfos),
			}
			writerCsv.Write(line)
		}
		return true
	})
}
