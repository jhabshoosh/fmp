package client

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
const financialRatiosURL = "https://financialmodelingprep.com/api/v3/financial-ratios/"
const financialStatementsURL = "https://financialmodelingprep.com/api/v3/financials/income-statement/"
const cashFlowStatementsURL = "https://financialmodelingprep.com/api/v3/financials/cash-flow-statement/"

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


type FinancialRatiosResponse struct {
	Symbol string
	Ratios []FinancialRatio `json:"ratios"`
}

type FinancialRatio struct {
	Date string
	InvestmentValuation InvestmentValuationRatios `json:"investmentValuationRatios"`
	ProfitabilityIndicator ProfitabilityIndicatorRatios `json:"profitabilityIndicatorRatios"`
	OperatingPerformance OperatingPerformanceRatios `json:"operatingPerformanceRatios"`
	LiquidityMeasurement LiquidityMeasurementRatios `json:"liquidityMeasurementRatios"`
	Debt DebtRatios `json:"debtRatios"`
	CashFlowIndicator CashFlowIndicatorRatios `json:"cashFlowIndicatorRatios"`
}

type InvestmentValuationRatios struct {
	PriceBookValueRatio string
	PriceToBookRatio string
	PriceToSalesRatio string
	PriceEarningsRatio string
	ReceivablesTurnover string
	PriceToFreeCashFlowsRatio string
	PriceToOperatingCashFlowsRatio string
	PriceCashFlowRatio string
	PriceEarningsToGrowthRatio string
	PriceSalesRatio string
	DividendYield string
	EnterpriseValueMultiple string
	PriceFairValue string
}

type ProfitabilityIndicatorRatios struct {
	NiperEBT string
	EbtperEBIT string
	EbitperRevenue string
	GrossProfitMargin string
	OperatingProfitMargin string
	PretaxProfitMargin string
	NetProfitMargin string
	EffectiveTaxRate string
	ReturnOnAssets string
	ReturnOnEquity string
	ReturnOnCapitalEmployed string
	NIperEBT string
	EBTperEBIT string
	EBITperRevenue string
}

type OperatingPerformanceRatios struct {
	ReceivablesTurnover string
	PayablesTurnover string
	InventoryTurnover string
	FixedAssetTurnover string
	AssetTurnover string
}

type LiquidityMeasurementRatios struct {
	CurrentRatio string
	QuickRatio string
	CashRatio string
	DaysOfSalesOutstanding string
	DaysOfInventoryOutstanding string
	OperatingCycle string
	DaysOfPayablesOutstanding string
	CashConversionCycle string
}

type DebtRatios struct {
	DebtRatio string
	DebtEquityRatio string
	LongtermDebtToCapitalization string
	TotalDebtToCapitalization string
	InterestCoverage string
	CashFlowToDebtRatio string
	CompanyEquityMultiplier string
}

type CashFlowIndicatorRatios struct {
	OperatingCashFlowPerShare string
	FreeCashFlowPerShare string
	CashPerShare string
	PayoutRatio string
	ReceivablesTurnover string
	OperatingCashFlowSalesRatio string
	FreeCashFlowOperatingCashFlowRatio string
	CashFlowCoverageRatios string
	ShortTermCoverageRatios string
	CapitalExpenditureCoverageRatios string
	DividendpaidAndCapexCoverageRatios string
	DividendPayoutRatio string
}

type FinancialStatementsResponse struct {
	Symbol string
	Financials []FinancialStatment
}

type FinancialStatment struct {
	Date string
	Revenue string
	RevenueGrowth string `json:"Revenue Growth"`
	CostOfRevenue string `json:"Cost of Revenue"`
	GrossProfit string `json:"Gross Profit"`
	RDExpenses string `json:"R&D Expenses"`
	SGAExpense string `json:"SG&A Expense"`
	OperatingExpenses string `json:"Operating Expenses"`
	OperatingIncome string `json:"Operating Income"`
	InterestExpense string `json:"Interest Expense"`
	EarningsBeforeTax string `json:"Earnings before Tax"`
	IncomeTaxExpense string `json:"Income Tax Expense"`
	NetIncomeNonControllingInt string `json:"Net Income - Non-Controlling int"`
	NetIncomeDiscontinuedOps string `json:"Net Income - Discontinued ops"`
	NetIncome string `json:"Net Income"`
	PreferredDividends string `json:"Preferred Dividends"`
	NetIncomeCom string `json:"Net Income Com"`
	EPS string
	EPSDiluted string `json:"EPS Diluted"`
	WeightedAverageShsOut string `json:"Weighted Average Shs Out"`
	WeightedAverageShsOutDil string `json:"Weighted Average Shs Out (Dil)"`
	DividendPerShare string `json:"Dividend per Share"`
	GrossMargin string `json:"Gross Margin"`
	EBITDAMargin string `json:"EBITDA Margin"`
	EBITMargin string `json:"EBIT Margin"`
	ProfitMargin string `json:"Profit Margin"`
	FreeCashFlowMargin string `json:"Free Cash Flow margin"`
	EBITDA string
	EBIT string
	ConsolidatedIncome string `json:"Consolidated Income"`
	EarningsBeforeTaxMargin string `json:"Earnings Before Tax Margin"`
	NetProfitMargin string `json:"Net Profit Margin"`
}

type CashFlowStatementResponse struct {
	Symbol string
	Financials []CashFlowStatement
}

type CashFlowStatement struct {
	Date string 
	DepreciationAndAmortization string `json:"Depreciation & Amortization"`
	StockBasedCompensation string `json:"Stock-based compensation"`
	OperatingCashFlow string `json:"Operating Cash Flow"`
	CapitalExpenditure string `json:"Capital Expenditure"`
	AcquistionsAndDisposals string `json:"Acquisitions and disposals"`
	InvestmentPurchasesAndSales string `json:"Investment purchases and sales"`
	InvestingCashFlow string `json:"Investing Cash flow"`
	IssRepaymentOfDebt string `json:"Issuance (repayment) of debt"`
	IssBuybacksOfShares string `json:"Issuance (buybacks) of shares"`
	DividendPayments string `json:"Dividend payments"`
	FinancingCashFlow string `json:"Financing Cash Flow"`
	EffectOfForexChangesOnCash string `json:"Effect of forex changes on cash"`
	NetCashFlowOverChangeInCash string `json:"Net cash flow / Change in cash"`
	FreeCashFlow string `json:"Free Cash Flow"`
	NetCashOverMarketCap string `json:"Net Cash/Marketcap"`
}


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

func FetchFinancialRatios(symbol string) (FinancialRatiosResponse, error) {
	res, err := http.Get(financialRatiosURL + symbol)

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
			panic(err.Error())
	}

	var frr FinancialRatiosResponse
	err = json.Unmarshal(body, &frr)
	if (err != nil) {
		fmt.Println("err unmarshalling:", err)
	}
	return frr, err
}

func FetchFinancialStatements(symbol string) (FinancialStatementsResponse, error) {
	res, err := http.Get(financialStatementsURL + symbol)

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
			panic(err.Error())
	}

	var fsr FinancialStatementsResponse
	err = json.Unmarshal(body, &fsr)
	if (err != nil) {
		fmt.Println("ehttps://financialmodelingprep.com/api/v3/financials/cash-flow-statement/AAPLrr unmarshalling:", err)
	}
	return fsr, err
}


func FetchCashFlowStatements(symbol string) (CashFlowStatementResponse, error) {
	res, err := http.Get(cashFlowStatementsURL + symbol)

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
			panic(err.Error())
	}

	var cfsr CashFlowStatementResponse
	err = json.Unmarshal(body, &cfsr)
	if (err != nil) {
		fmt.Println("err unmarshalling:", err)
	}
	return cfsr, err
}