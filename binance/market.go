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

type Order struct {
    Price    float64 `json:",string"`
    Quantity float64 `json:",string"`
}

type OrderBook struct {
    LastUpdatedId int64 `json:"lastUpdatedId"`
    Bids []Order `json:"bids"`
    Asks []Order `json:"asks"`
}

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

//
//
func (b *Binance) GetOrderBook(symbol string, limit int64) (book OrderBook, err error) {
    
    reqUrl := fmt.Sprintf("v1/depth?symbol=%s&limit=%d", symbol, limit)

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

var IntervalEnum = map[string]bool {
    "1m": true,
    "3m": true,
    "5m": true,
    "15m": true,
    "30m": true,
    "1h": true,
    "2h": true,
    "4h": true,
    "6h": true,
    "8h": true,
    "12h": true,
    "1d": true,
    "3d": true,
    "1w": true,
    "1M": true,
}       

//
//
func (b *Binance) GetAggTrades(symbol string) (trades []AggTrade, err error) {

    reqUrl := fmt.Sprintf("v1/aggTrades?symbol=%s", symbol)

    _, err = b.client.do("GET", reqUrl, "", false, &trades)
    return
}

//
//
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
