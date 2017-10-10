package market_test

import (
//    "fmt"
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

func TestGetAggTrades(t *testing.T) {

    // Params
    symbol := "BNBBTC"

    binance := binance.New("", "")
    _, err := binance.GetAggTrades(symbol)
    if err != nil {
        t.Fatal(err)
    }
}

func TestGetKlines(t *testing.T) {

    // Params
    symbol := "BNBBTC"
    interval := "1m"

    binance := binance.New("", "")
    _, err := binance.GetKlines(symbol, interval)
    if err != nil {
        t.Fatal(err)
    }
}

func TestGet24Hr(t *testing.T) {

    //Params 
    symbol := "BNBBTC"

    binance := binance.New("", "")
    _, err := binance.Get24Hr(symbol)
    if err != nil {
        t.Fatal(err)
    }
}

func TestGetAllPrices(t *testing.T) {

    binance := binance.New("", "")
    _, err := binance.GetAllPrices()
    if err != nil {
        t.Fatal(err)
    }
}

func TestGetBookTickers(t *testing.T) {

    binance := binance.New("", "")
    _, err := binance.GetBookTickers()
    if err != nil {
        t.Fatal(err)
    }
}

