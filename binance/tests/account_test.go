package account_test

import (
    "os"
    "testing"
    "github.com/pdepip/go-binance/binance"
)



func TestGetTrades(t *testing.T) {
    client := binance.New(os.Getenv("BINANCE_KEY"), os.Getenv("BINANCE_SECRET"))

    trades, err := client.GetTrades("BNBETH")

    if err != nil {
        t.Fatal(err)
    }

    t.Logf("%+v\n", trades)
}


func TestGetWithdraws(t *testing.T) {
    client := binance.New(os.Getenv("BINANCE_KEY"), os.Getenv("BINANCE_SECRET"))

    withdraws, err := client.GetWithdrawHistory()

    if err != nil {
        t.Fatal(err)
    }

    t.Logf("%+v\n", withdraws)
}


func TestGetDeposits(t *testing.T) {
    client := binance.New(os.Getenv("BINANCE_KEY"), os.Getenv("BINANCE_SECRET"))

    deposits, err := client.GetDepositHistory()

    if err != nil {
        t.Fatal(err)
    }

    t.Logf("%+v\n", deposits)
}

func TestGetTradesFromOrder(t *testing.T) {
    client := binance.New(os.Getenv("BINANCE_KEY"), os.Getenv("BINANCE_SECRET"))

    trades, err := client.GetTradesFromOrder("LINKETH", 10107102)
    if err != nil {
        t.Fatal(err)
    }

    t.Logf("%+v\n", trades)
}

/*

func TestGetPositions(t *testing.T) {

    binance := binance.New(os.Getenv("BINANCE_KEY"), os.Getenv("BINANCE_SECRET"))
    positions, err := binance.GetPositions()

    if err != nil {
        t.Fatal(err)
    }

    t.Logf("%+v\n", positions)
}



    THE FOLLOWING TESTS WILL PLACE ACTUAL ORDERS ON THE BINANCE EXCHANGE

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

    client := binance.New(os.Getenv("BINANCE_KEY"), os.Getenv("BINANCE_SECRET"))
    res, err := client.PlaceLimitOrder(order)
    
    if err != nil {
        t.Fatal(err)
    }

    t.Logf("%+v\n", res.OrderId)

}


func TestMarketOrder(t *testing.T) {

    // Params
    order := binance.MarketOrder {
        Symbol:      "BNBBTC",
        Side:        "BUY",
        Type:        "MARKET",
        Quantity:    10.0,
    }

    client := binance.New(os.Getenv("BINANCE_KEY"), os.Getenv("BINANCE_SECRET"))
    res, err := client.PlaceMarketOrder(order)
    
    if err != nil {
        t.Fatal(err)
    }

    t.Logf("%+v\n", res.OrderId)

}


func TestQueryOrder(t *testing.T) {

    // Order Param
    order := binance.LimitOrder {
        Symbol:      "BNBBTC",
        Side:        "BUY",
        Type:        "LIMIT",
        TimeInForce: "GTC",
        Quantity:    50.0,
        Price:       0.00025,
    }

    client := binance.New(os.Getenv("BINANCE_KEY"), os.Getenv("BINANCE_SECRET"))
    res, err := client.PlaceLimitOrder(order)
    if err != nil {
        t.Fatal(err)
    }

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

func TestCancelOrder(t *testing.T) {

    // Order Param
    order := binance.LimitOrder {
        Symbol:      "BNBBTC",
        Side:        "BUY",
        Type:        "LIMIT",
        TimeInForce: "GTC",
        Quantity:    10.0,
        Price:       0.00025,
    }

    client := binance.New(os.Getenv("BINANCE_KEY"), os.Getenv("BINANCE_SECRET"))
    res, err := client.PlaceLimitOrder(order)
    if err != nil {
        t.Fatal(err)
    }

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

    client := binance.New(os.Getenv("BINANCE_KEY"), os.Getenv("BINANCE_SECRET"))
    openOrders, err := client.GetOpenOrders(query)
    if err != nil {
        t.Fatal(err)
    }

    t.Logf("%+v\n", openOrders)
}

*/
