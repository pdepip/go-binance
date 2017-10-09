package account_test

import (
    "os"
    //"fmt"
    "testing"
    "go-binance/binance"
)


/*
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
    order := binance.NewOrder {
        Symbol:      "BNBBTC",
        Side:        "BUY",
        Type:        "LIMIT",
        TimeInForce: "GTC",
        Quantity:    50.0,
        Price:       0.00025,
    }

    binance := binance.New(os.Getenv("key"), os.Getenv("secret"))
    res, err := binance.PlaceOrder(order)
    
    if err != nil {
        t.Fatal(err)
    }

    t.Logf("%+v\n", res.OrderId)

}


func TestQueryOrder(t *testing.T) {

    // Order Param
    order := binance.NewOrder {
        Symbol:      "BNBBTC",
        Side:        "BUY",
        Type:        "LIMIT",
        TimeInForce: "GTC",
        Quantity:    50.0,
        Price:       0.00025,
    }

    client := binance.New(os.Getenv("key"), os.Getenv("secret"))
    res, err := client.PlaceOrder(order)
    t.Logf("%+v\n", res)

    // Query Param
    query := binance.OrderQuery{
        Symbol:  res.Symbol,
        OrderId: res.OrderId,
    }

    status, err := client.CheckOrder(query)

    if err != nil {
        t.Fatal(err)
    }

    t.Logf("%+v\n", status)
}
*/

func TestCancelOrder(t *testing.T) {

    // Order Param
    order := binance.NewOrder {
        Symbol:      "BNBBTC",
        Side:        "BUY",
        Type:        "LIMIT",
        TimeInForce: "GTC",
        Quantity:    10.0,
        Price:       0.00025,
    }

    client := binance.New(os.Getenv("key"), os.Getenv("secret"))
    res, err := client.PlaceOrder(order)
    t.Logf("%+v\n", res)

    // Cancel Param
    query := binance.OrderQuery {
        Symbol: res.Symbol,
        OrderId: res.OrderId,
    }

    canceledOrder, err := client.CancelOrder(query)

    if err != nil {
        t.Fatal(err)
    }

    t.Logf("%+v\n", canceledOrder)
}

