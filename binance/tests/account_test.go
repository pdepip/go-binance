package account_test

import (
    "os"
    //"fmt"
    "testing"
    "go-binance/binance"
)


func TestGetPositions(t *testing.T) {

    binance := binance.New(os.Getenv("key"), os.Getenv("secret"))
    positions, err := binance.GetPositions()

    if err != nil {
        t.Fatal(err)
    }

    t.Logf("%+v\n", positions)
}

/*
func TestLimitOrder(t *testing.T) {

    // Params
    order := binance.LimitOrder {
        Symbol:      "BNBBTC",
        Side:        "BUY",
        Type:        "LIMIT",
        TimeInForce: "GTC",
        Quantity:    50.0,
        Price:       0.00025,
    }

    binance := binance.New(os.Getenv("key"), os.Getenv("secret"))
    res, err := binance.PlaceLimitOrder(order)
    
    if err != nil {
        t.Fatal(err)
    }

    t.Logf("%+v\n", res.OrderId)

}
/*
func TestMarketOrder(t *testing.T) {

    // Params
    order := binance.MarketOrder {
        Symbol:      "BNBBTC",
        Side:        "BUY",
        Type:        "MARKET",
        Quantity:    10.0,
    }

    binance := binance.New(os.Getenv("key"), os.Getenv("secret"))
    res, err := binance.PlaceMarketOrder(order)
    
    if err != nil {
        t.Fatal(err)
    }

    t.Logf("%+v\n", res.OrderId)

}
/*

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

/*
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

func TestGetOpenOrders(t *testing.T) {

    // Param
    query := binance.OpenOrdersQuery {
        Symbol: "BNBBTC",
    }

    client := binance.New(os.Getenv("key"), os.Getenv("secret"))
    openOrders, err := client.GetOpenOrders(query)
    if err != nil {
        t.Fatal(err)
    }

    t.Logf("%+v\n", openOrders)
}
*/
