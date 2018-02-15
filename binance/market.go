/*

   market.go
       Market Data Endpoints for the Binance Exchange API

*/
package binance

import (
	"fmt"
)

// Get order book
func (b *Binance) GetOrderBook(q OrderBookQuery) (book OrderBook, err error) {

	err = q.ValidateOrderBookQuery()
	if err != nil {
		return
	}

	reqUrl := fmt.Sprintf("api/v1/depth?symbol=%s&limit=%d", q.Symbol, q.Limit)
	_, err = b.client.do("GET", reqUrl, "", false, &book)

	return
}

// Get compressed, aggregate trades. Trades that fill at the time, from the same order, with the same price will have the quantity aggregated.
func (b *Binance) GetAggTrades(q SymbolQuery) (trades []AggTrade, err error) {

	err = q.ValidateSymbolQuery()
	if err != nil {
		return
	}

	reqUrl := fmt.Sprintf("api/v1/aggTrades?symbol=%s", q.Symbol)

	_, err = b.client.do("GET", reqUrl, "", false, &trades)
	return
}

// Kline/candlestick bars for a symbol. Klines are uniquely identified by their open time.
func (b *Binance) GetKlines(q KlineQuery) (klines []Kline, err error) {

	err = q.ValidateKlineQuery()
	if err != nil {
		return
	}

	reqUrl := fmt.Sprintf("api/v1/klines?symbol=%s&interval=%s&limit=%d", q.Symbol, q.Interval, q.Limit)

	_, err = b.client.do("GET", reqUrl, "", false, &klines)
	if err != nil {
		return
	}

	return
}

// 24 hour price change statistics.
func (b *Binance) Get24Hr(q SymbolQuery) (changeStats ChangeStats, err error) {

	err = q.ValidateSymbolQuery()
	if err != nil {
		return
	}

	reqUrl := fmt.Sprintf("api/v1/ticker/24hr?symbol=%s", q.Symbol)
	_, err = b.client.do("GET", reqUrl, "", false, &changeStats)

	return
}

// Latest price for all symbols.
func (b *Binance) GetAllPrices() (prices []TickerPrice, err error) {

	reqUrl := "api/v1/ticker/allPrices"
	_, err = b.client.do("GET", reqUrl, "", false, &prices)

	return
}

// Latest price for an individual symbol
func (b *Binance) GetLastPrice(q SymbolQuery) (price TickerPrice, err error) {

	err = q.ValidateSymbolQuery()
	if err != nil {
		return
	}

	var prices []TickerPrice
	prices, err = b.GetAllPrices()
	if err != nil {
		return
	}

	for _, p := range prices {
		if p.Symbol == q.Symbol {
			return p, nil
		}
	}

	return
}

// Best price/qty on the order book for all symbols.
func (b *Binance) GetBookTickers() (booktickers []BookTicker, err error) {

	reqUrl := "api/v1/ticker/allBookTickers"
	_, err = b.client.do("GET", reqUrl, "", false, &booktickers)

	return
}

// Exchange filters for all symbols
func (b *Binance) GetExchangeInfo() (exchangeinfo ExchangeInfo, err error) {

	_, err = b.client.do("GET", "api/v1/exchangeInfo", "", false, &exchangeinfo)
	if err != nil {
		return
	}

	return
}

// Ping Rest API. If no error is returned, API is up and running.
func (b *Binance) Ping() (pingResponse PingResponse, err error) {
	_, err = b.client.do("GET", "api/v1/ping", "", false, &pingResponse)

	return
}

// Ping Withdrawal API. Status is returned in response.
func (b *Binance) GetWithdrawalSystemStatus() (withdrawalSystemStatus WithdrawalSystemStatus, err error) {
	_, err = b.client.do("GET", "wapi/v3/systemStatus.html", "", false, &withdrawalSystemStatus)

	return
}
