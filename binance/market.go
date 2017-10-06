/*
    market.go
        Market Data Endpoints for the Binance Exchange API

*/
package binance

import (
    "fmt"
)

type Order [2]struct {
    Price    float64 `json:",string"`
    Quantity float64 `json:",string"`
    Data     []byte
}

type OrderBook struct {
    LastUpdatedId int64 `json:"lastUpdatedId"`
    Bids [][]Order `json:"bids"`
    Asks [][]Order `json:"asks"`
}

//
//
func (b *Binance) GetOrderBook(symbol string, limit int64) (book OrderBook, err error) {
    
    reqUrl := fmt.Sprintf("depth?symbol=%s&limit=%d", symbol, limit)

    _, err = b.client.do("GET", reqUrl, "", false, &book)
    return
}

