package fmp

import (
	"strings"
	"fmt"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const keyMetricsURL = "https://financialmodelingprep.com/api/v3/company-key-metrics/"
const allCompaniesURL = "https://financialmodelingprep.com/api/v3/company/stock/list"
const companyQuoteURL = "https://financialmodelingprep.com/api/v3/quote/"
const companyProfileURL = "https://financialmodelingprep.com/api/v3/company/profile/"

func GetSymbolsList() []Stock {
	res, err := http.Get(allCompaniesURL)
	if err != nil {
    panic(err.Error())
	}
	
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
			panic(err.Error())
	}

	acr := new(AllCompaniesResponse)
	err = json.Unmarshal(body, &acr)
	if (err != nil) {
		fmt.Println("err unmarshalling:", err)
	}
	return acr.Companies
}

func FetchKeyMetrics(symbol string) (*KeyMetricsResponse, error)  {
	res, err := http.Get(keyMetricsURL + symbol)
	if err != nil {
    panic(err.Error())
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
			panic(err.Error())
	}

	kmr := new(KeyMetricsResponse)
	err = json.Unmarshal(body, &kmr)
	if (err != nil) {
		fmt.Println("err unmarshalling:", err)
	}
	return kmr, err
}

// FetchCompanyQuote
func FetchCompanyQuote(symbols []string) ([]CompanyQuote, error) {
	symbolsParam := strings.Join(symbols, ",")
	res, err := http.Get(companyQuoteURL + symbolsParam)

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
			panic(err.Error())
	}

	var cqr []CompanyQuote
	err = json.Unmarshal(body, &cqr)
	if (err != nil) {
		fmt.Println("err unmarshalling:", err)
	}
	return cqr, err
}

func FetchCompanyProfile(symbol string) (CompanyProfileResponse, error) {
	res, err := http.Get(companyProfileURL + symbol)

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
			panic(err.Error())
	}

	var cpr CompanyProfileResponse
	err = json.Unmarshal(body, &cpr)
	if (err != nil) {
		fmt.Println("err unmarshalling:", err)
	}
	return cpr, err
}