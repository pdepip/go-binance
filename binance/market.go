/*
    market.go
        Market Data Endpoints for the Binance Exchange API

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

//
//
func (b *Binance) GetOrderBook(symbol string, limit int64) (book OrderBook, err error) {
    
    reqUrl := fmt.Sprintf("depth?symbol=%s&limit=%d", symbol, limit)

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
