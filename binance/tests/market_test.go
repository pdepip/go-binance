package market_test

import (
    "fmt"
    "testing"
    "go-binance/binance"
)

func TestGetOrderBook(t *testing.T) {

    // Params
    symbol := "BNBBTC"
    var limit int64
    limit = 100
    
    binance := binance.New("", "")
    _, err := binance.GetOrderBook(symbol, limit)

    if err != nil {
        t.Fatal(err)
    }
}

