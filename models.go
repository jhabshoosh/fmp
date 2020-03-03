package fmp

import (
	"fmt"
)

// KeyMetrics
type KeyMetrics struct {
	EV string `json:"Enterprise Value"`
	EvOverEbitda string `json:"Enterprise Value over EBITDA"`
	PERatio string `json:"PE Ratio"`
}

// KeyMetricsResponse 
type KeyMetricsResponse struct {
	Symbol string `json:"symbol"`
	Metrics []KeyMetrics `json:"metrics"`
}

// Stock
type Stock struct {
	Symbol string
	Name string
	Price float64
	Exchange string
}

type AllCompaniesResponse struct {
	Companies []Stock `json:"symbolsList"`
}

type CompanyQuote struct {
	Symbol string
	Price float64
	ChangesPercentage float64
	Change float64
	DayLow float64
	DayHigh float64
	YearHigh float64
	YearLow float64
	MarketCap float64
	PriceAvg50 float64
	PriceAvg200 float64
	Volume float64
	AvgVolume float64
	Exhange string

}

type CompanyQuoteResponse struct {
	Quotes []CompanyQuote
}

type CompanyProfileResponse struct {
	Symbol string
	Profile CompanyProfile `json:"profile"`
}

type CompanyProfile struct {
	Price float64
	Beta string
	VolAvg string
	MarketCap string
	LastDiv string
	Range string
	Changes float64
	ChangesPercentage string
	CompanyName string
	Exchange string
	Industry string
	Website string
	Description string
	CEO string
	Sector string
	Image string
}

func (cpr CompanyProfileResponse) String() string {
	return fmt.Sprintf("%s\t%s\t%f", cpr.Symbol, cpr.Profile.CompanyName, cpr.Profile.Price)
}