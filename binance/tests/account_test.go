package account_test

import (
    "os"
    //"fmt"
    "testing"
    "go-binance/binance"
)


func TestGetAccountInfo(t *testing.T) {

    binance := binance.New(os.Getenv("key"), os.Getenv("secret"))
    account, err := binance.GetAccountInfo()

    if err != nil {
        t.Fatal(err)
    }

    t.Logf("%+v\n", account)
}

func TestNewOrder(t *testing.T) {

    // Params 
    symbol := "BNBBTC"
    side := "BUY"
    orderType := "LIMIT"
    timeInForce := "GTC"
    quantity := 50.0
    price := 0.00025

    binance := binance.New(os.Getenv("key"), os.Getenv("secret"))
    order, err := binance.NewOrder(symbol, side, orderType, timeInForce, quantity, price)
    
    if err != nil {
        t.Fatal(err)
    }

    t.Logf("%+v\n", order)
}
