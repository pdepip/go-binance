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
    stuff ,err := binance.GetOrderBook(symbol, limit)
    fmt.Printf("%+v\n", stuff)

    for _, v := range stuff.Bids {
        fmt.Println(v)
    }

    if err != nil {
        t.Fatal(err)
    }
}

