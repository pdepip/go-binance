package market_test

import (
    "testing"
    "github.com/pdepip/go-binance/binance"
)

func TestGetOrderBook(t *testing.T) {

    // Params
    query := binance.OrderBookQuery {
        Symbol: "BNBBTC",
        Limit: 100,
    }

    client := binance.New("", "")
    res, err := client.GetOrderBook(query)

    if err != nil {
        t.Fatal(err)
    }

    t.Logf("%+v\n", res)
}

func TestGetAggTrades(t *testing.T) {

    // Params
    query := binance.SymbolQuery {
        Symbol: "BNBBTC",
    }

    client := binance.New("", "")
    res, err := client.GetAggTrades(query)
    if err != nil {
        t.Fatal(err)
    }

    t.Logf("%+v\n", res)
}

func TestGetKlines(t *testing.T) {

    // Params
    query := binance.KlineQuery {
        Symbol: "BNBBTC",
        Interval: "12h",
    }

    client := binance.New("", "")
    res, err := client.GetKlines(query)
    if err != nil {
        t.Fatal(err)
    }

    t.Logf("%+v\n", res)
}

func TestGet24Hr(t *testing.T) {

    //Params
    query := binance.SymbolQuery {
        Symbol: "BNBBTC",
    }

    client := binance.New("", "")
    res, err := client.Get24Hr(query)
    if err != nil {
        t.Fatal(err)
    }

    t.Logf("%+v\n", res)
}

func TestGetAllPrices(t *testing.T) {

    client := binance.New("", "")
    res, err := client.GetAllPrices()
    if err != nil {
        t.Fatal(err)
    }

    t.Logf("%+v\n", res)
}

func TestGetBookTickers(t *testing.T) {

    client := binance.New("", "")
    res, err := client.GetBookTickers()
    if err != nil {
        t.Fatal(err)
    }

    t.Logf("%+v\n", res)
}
