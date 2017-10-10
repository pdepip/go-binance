/*
    market.go
        Market Data Endpoints for the Binance Exchange API

    To Do:
        1. Document Functions
        2. Optional Parameters
*/
package binance

import (
    "fmt"
    "strconv"
    "encoding/json"
)


type AggTrade struct {
    TradeId      int64   `json:"a"`
    Price        float64 `json:"p,string"`
    Quantity     float64 `json:"q,string"`
    FirstTradeId int64   `json:"f"`
    LastTradeId  int64   `json:"l"`
    Timestamp    int64   `json:"T"`
    Maker        bool    `json:"m"`
    BestMatch    bool    `json:"M"`
}

type Kline struct {
    OpenTime         int64 
    Open             float64
    High             float64
    Low              float64
    Close            float64
    Volume           float64
    CloseTime        int64
    QuoteVolume      float64
    NumTrades        int64
    TakerBaseVolume  float64
    TakerQuoteVolume float64
}

type ChangeStats struct {
    PriceChange        float64 `json:"priceChange,string"`
    PriceChangePercent float64 `json:"priceChangePercent,string"`
    WeightedAvgPrice   float64 `json:"weightedAvgPrice,string"`
    PrevClosePrice     float64 `json:"prevClosePrice,string"`
    LastPrice          float64 `json:"lastPrice,string"`
    BidPrice           float64 `json:"bidPrice,string"`
    AskPrice           float64 `json:"askPrice,string"`
    OpenPrice          float64 `json:"openPrice,string"`
    HighPrice          float64 `json:"highPrice,string"`
    LowPrice           float64 `json:"lowPrice,string"`
    Volume             float64 `json:"volume,string"`
    OpenTime           int64   `json:"openTime"`
    CloseTime          int64   `json:"closeTime"`
    FirstId            int64   `json:"firstId"`
    LastId             int64   `json:"lastId"`
    Count              int64   `json:"count"`

}

type TickerPrice struct {
    Symbol string  `json:"symbol"`
    Price  float64 `json:"price,string"`
}

type BookTicker struct {
    Symbol      string  `json:"symbol"`
    BidPrice    float64 `json:"bidPrice,string"`
    BidQuantity float64 `json:"bidQuantity,string"`
    AskPrice    float64 `json:"askPrice,string"`
    AskQuantity float64 `json:"askQuantity,string"`
}

//
// Get order book
func (b *Binance) GetOrderBook(q OrderBookQuery) (book OrderBook, err error) {

    err = q.ValidateOrderBookQuery()
    reqUrl := fmt.Sprintf("v1/depth?symbol=%s&limit=%d", q.Symbol, q.Limit)

    _, err = b.client.do("GET", reqUrl, "", false, &book)
    return
}

func (o *Order) UnmarshalJSON(b []byte) error {
    var s [2]string

    err := json.Unmarshal(b, &s)
    if err != nil {
        return err
    }

    o.Price, err = strconv.ParseFloat(s[0], 64)
    if err != nil {
        return err
    }

    o.Quantity, err = strconv.ParseFloat(s[1], 64)
    if err != nil {
        return err
    }

    return nil
}


// Get compressed, aggregate trades. Trades that fill at the time, from the same order, with the same price will have the quantity aggregated.
func (b *Binance) GetAggTrades(symbol string) (trades []AggTrade, err error) {

    reqUrl := fmt.Sprintf("v1/aggTrades?symbol=%s", symbol)

    _, err = b.client.do("GET", reqUrl, "", false, &trades)
    return
}

//
// Kline/candlestick bars for a symbol. Klines are uniquely identified by their open time.
func (b *Binance) GetKlines(symbol, interval string) (klines []Kline, err error) {

    reqUrl := fmt.Sprintf("v1/klines?symbol=%s&interval=%s", symbol, interval)

    _, err = b.client.do("GET", reqUrl, "", false, &klines)
    return
}

func (k *Kline) UnmarshalJSON(b []byte) error {
    var s [11]interface{}

    err := json.Unmarshal(b, &s)
    if err != nil {
        return err
    }

    fmt.Println(s)
    k.OpenTime = int64(s[0].(float64))

    k.Open, err = strconv.ParseFloat(s[1].(string), 64)
    if err != nil {
        return err
    }

    k.High, err = strconv.ParseFloat(s[2].(string), 64)
    if err != nil {
        return err
    }

    k.Low, err = strconv.ParseFloat(s[3].(string), 64)
    if err != nil {
        return err
    }

    k.Close, err = strconv.ParseFloat(s[4].(string), 64)
    if err != nil {
        return err
    }

    k.Volume, err = strconv.ParseFloat(s[5].(string), 64)
    if err != nil {
        return err
    }

    k.CloseTime = int64(s[6].(float64))


    k.QuoteVolume, err = strconv.ParseFloat(s[7].(string), 64)
    if err != nil {
        return err
    }

    k.NumTrades = int64(s[8].(float64))

    k.TakerBaseVolume, err = strconv.ParseFloat(s[9].(string), 64)
    if err != nil {
        return err
    }

    k.TakerQuoteVolume, err = strconv.ParseFloat(s[10].(string), 64)
    if err != nil {
        return err
    }

    return nil
}

//
// 24 hour price change statistics.
func (b *Binance) Get24Hr(symbol string) (changeStats ChangeStats, err error) {

    reqUrl := fmt.Sprintf("v1/ticker/24hr?symbol=%s", symbol)

    _, err = b.client.do("GET", reqUrl, "", false, &changeStats)
    return
}

//
// Latest price for all symbols.
func (b *Binance) GetAllPrices() (prices []TickerPrice, err error) {

    reqUrl := "v1/ticker/allPrices"

    _, err = b.client.do("GET", reqUrl, "", false, &prices)
    return
}

//
// Best price/qty on the order book for all symbols.
func (b *Binance) GetBookTickers() (booktickers []BookTicker, err error) {

    reqUrl := "v1/ticker/allBookTickers"

    _, err = b.client.do("GET", reqUrl, "", false, &booktickers)
    return
}

