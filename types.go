package main

import "time"

type Currency struct {
	Logo        string
	Name        string
	Symbol      string
	Slug        string
	Type        string // coin, token
	Description string
	Links       *Link
	Platforms   []*Platform
	Wallets     []*Wallet
	AuditInfos  []*AuditInfo
}

type Link struct {
	Website      []string `json:"website"`
	TechnicalDoc []string `json:"technical_doc"`
	Explorer     []string `json:"explorer"`
	SourceCode   []string `json:"source_code"`
	MessageBoard []string `json:"message_board"`
	Chat         []string `json:"chat"`
	Announcement []string `json:"announcement"`
	Reddit       []string `json:"reddit"`
	Facebook     []string `json:"facebook"`
	Twitter      []string `json:"twitter"`
}

type Platform struct {
	Logo                string
	ContractAddress     string
	ContractPlatform    string
	ContractPRCURL      []string
	ContractExplorerURL string
}
type Wallet struct {
	Logo string
	Name string
	URL  string
}

type AuditInfo struct {
	Auditor   string
	ReportUrl string
}

type CoinMarketCapCurrencyInfoPlatform struct {
	ContractId                     int      `json:"contractId"`
	ContractAddress                string   `json:"contractAddress"`
	ContractPlatform               string   `json:"contractPlatform"`
	ContractPlatformId             int      `json:"contractPlatformId"`
	ContractChainId                int      `json:"contractChainId"`
	ContractRpcUrl                 []string `json:"contractRpcUrl"`
	ContractNativeCurrencyName     string   `json:"contractNativeCurrencyName"`
	ContractNativeCurrencySymbol   string   `json:"contractNativeCurrencySymbol"`
	ContractNativeCurrencyDecimals int      `json:"contractNativeCurrencyDecimals"`
	ContractBlockExplorerUrl       string   `json:"contractBlockExplorerUrl"`
	ContractExplorerUrl            string   `json:"contractExplorerUrl"`
	ContractDecimals               int      `json:"contractDecimals,omitempty"`
	PlatformCryptoId               int      `json:"platformCryptoId"`
	Sort                           int      `json:"sort"`
	Wallets                        []struct {
		Id            int    `json:"id"`
		Name          string `json:"name"`
		Tier          int    `json:"tier"`
		Url           string `json:"url"`
		Chains        string `json:"chains"`
		Types         string `json:"types"`
		Introduction  string `json:"introduction"`
		Star          int    `json:"star"`
		Security      int    `json:"security"`
		EasyToUse     int    `json:"easyToUse"`
		Decentration  bool   `json:"decentration"`
		FocusNumber   int    `json:"focusNumber"`
		Rank          int    `json:"rank"`
		Logo          string `json:"logo"`
		MultipleChain bool   `json:"multipleChain"`
	} `json:"wallets"`
}

type CoinMarketCapCurrencyInfoWallets struct {
	Id            int    `json:"id"`
	Name          string `json:"name"`
	Tier          int    `json:"tier,omitempty"`
	Url           string `json:"url"`
	Chains        string `json:"chains"`
	Types         string `json:"types"`
	Introduction  string `json:"introduction"`
	Star          int    `json:"star"`
	Security      int    `json:"security"`
	EasyToUse     int    `json:"easyToUse"`
	Decentration  bool   `json:"decentration"`
	FocusNumber   int    `json:"focusNumber"`
	Rank          int    `json:"rank,omitempty"`
	Logo          string `json:"logo"`
	MultipleChain bool   `json:"multipleChain"`
}
type CoinMarketCapCurrencyInfoAuditInfo struct {
	CoinId      string    `json:"coinId"`
	Auditor     string    `json:"auditor"`
	AuditStatus int       `json:"auditStatus"`
	AuditTime   time.Time `json:"auditTime,omitempty"`
	ReportUrl   string    `json:"reportUrl"`
}

type CoinMarketCapCurrencyInfo struct {
	Props struct {
		PageProps struct {
			DetailRes struct {
				Detail struct {
					Id          int                                   `json:"id"`
					Name        string                                `json:"name"`
					Symbol      string                                `json:"symbol"`
					Slug        string                                `json:"slug"`
					Category    string                                `json:"category"`
					Description string                                `json:"description"`
					Notice      string                                `json:"notice"`
					Urls        *Link                                 `json:"urls"`
					Platforms   []*CoinMarketCapCurrencyInfoPlatform  `json:"platforms"`
					Wallets     []*CoinMarketCapCurrencyInfoWallets   `json:"wallets"`
					IsAudited   bool                                  `json:"isAudited"`
					AuditInfos  []*CoinMarketCapCurrencyInfoAuditInfo `json:"auditInfos"`
				} `json:"detail"`
			} `json:"detailRes"`
		} `json:"pageProps"`
	} `json:"props"`
}
