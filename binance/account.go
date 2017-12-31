/*

    account.go
        Account (Signed) Endpoints for Binance Exchange API

*/
package binance

import (
    "fmt"
)


func (b *Binance) GetExchangeInfo() (exchangeinfo ExchangeInfo, err error) {
    _, err = b.client.do("GET", "v1/exchangeInfo", "", false, &exchangeinfo)
    if err != nil {
        return
    }

    return
}

// Get Basic Account Information
func (b *Binance) GetAccountInfo() (account Account, err error) {

    reqUrl := fmt.Sprintf("v3/account")

    _, err = b.client.do("GET", reqUrl, "", true, &account)
    if err != nil {
        return
    }

    return
}


// Filter Basic Account Information To Retrieve Current Holdings
func (b *Binance) GetPositions() (positions []Balance, err error) {

    reqUrl := fmt.Sprintf("v3/account")
    account := Account{}

    _, err = b.client.do("GET", reqUrl, "", true, &account)
    if err != nil {
        return
    }

    positions = make([]Balance, len(account.Balances))
    i := 0

    for _, balance := range account.Balances {
        if balance.Free != 0.0 || balance.Locked != 0.0 {
            positions[i] = balance
            i++
        }
    }

    return positions[:i], nil
}


// Place a Limit Order
func (b *Binance) PlaceLimitOrder(l LimitOrder) (res PlacedOrder, err error) {

    err = l.ValidateLimitOrder()
    if err != nil {
        return
    }

    reqUrl := fmt.Sprintf("v3/order?symbol=%s&side=%s&type=%s&timeInForce=%s&quantity=%f&price=%f&recvWindow=%d", l.Symbol, l.Side, l.Type, l.TimeInForce, l.Quantity, l.Price, l.RecvWindow)

    _, err = b.client.do("POST", reqUrl, "", true, &res)
    if err != nil {
        return
    }

    return
}


// Place a Market Order
func (b *Binance) PlaceMarketOrder(m MarketOrder) (res PlacedOrder, err error) {

    err = m.ValidateMarketOrder()
    if err != nil {
        return
    }

    reqUrl := fmt.Sprintf("v3/order?symbol=%s&side=%s&type=%s&quantity=%f&recvWindow=%d", m.Symbol, m.Side, m.Type, m.Quantity, m.RecvWindow)

    _, err = b.client.do("POST", reqUrl, "", true, &res)
    if err != nil {
        return
    }

    return
}


// Cancel an Order
func (b *Binance) CancelOrder(query OrderQuery) (order CanceledOrder, err error) {

    err = query.ValidateOrderQuery()
    if err != nil {
        return
    }

    reqUrl := fmt.Sprintf("v3/order?symbol=%s&orderId=%d&recvWindow=%d", query.Symbol, query.OrderId, query.RecvWindow)

    _, err = b.client.do("DELETE", reqUrl, "", true, &order)
    if err != nil {
        return
    }

    return
}


// Check the Status of an Order
func (b *Binance) CheckOrder(query OrderQuery) (status OrderStatus, err error) {

    err = query.ValidateOrderQuery()
    if err != nil {
        return
    }

    reqUrl := fmt.Sprintf("v3/order?symbol=%s&orderId=%d&origClientOrderId=%s&recvWindow=%d", query.Symbol, query.OrderId, query.RecvWindow)

    _, err = b.client.do("GET", reqUrl, "", true, &status)
    if err != nil {
        return
    }

    return
}


// Retrieve All Open Orders
func (b *Binance) GetOpenOrders() (orders []OrderStatus, err error) {
    _, err = b.client.do("GET", "v3/openOrders", "", true, &orders)
    if err != nil {
        return
    }

    return
}

//
// Retrieves all trades
func (b *Binance) GetTrades(symbol string) (trades []Trade, err error) {
    //timeStamp := unixMillis(time.Now())
    //recvWindow := recvWindow(5 * time.Second)

    _, err = b.client.do("GET", "v3/myTrades?symbol=" + symbol, "", true, &trades)
    if err != nil {
        return
    }
    return
}

