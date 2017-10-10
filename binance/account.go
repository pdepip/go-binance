/*
    account.go
        Account Endpoints for Binance Exchange API
*/
package binance

import (
    "fmt"
)


func (b *Binance) GetAccountInfo() (account Account, err error) {

    reqUrl := fmt.Sprintf("v3/account")

    _, err = b.client.do("GET", reqUrl, "", true, &account)
    if err != nil {
        return
    }

    return
}


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


func (b *Binance) CancelOrder(query OrderQuery) (order CanceledOrder, err error) {

    err = query.ValidateOrderQuery()
    if err != nil {
        return
    }

    reqUrl := fmt.Sprintf("v3/order?symbol=%s&orderId=%d&recvWindow", query.Symbol, query.OrderId, query.RecvWindow)

    _, err = b.client.do("DELETE", reqUrl, "", true, &order)
    if err != nil {
        return
    }

    return
}


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


func (b *Binance) GetOpenOrders(q OpenOrdersQuery) (orders []OrderStatus, err error) {

    err = q.ValidateOpenOrdersQuery()
    if err != nil {
        return
    }

    reqUrl := fmt.Sprintf("v3/openOrders?symbol=%s&recvWindow=%d", q.Symbol, q.RecvWindow)
 
    _, err = b.client.do("GET", reqUrl, "", true, &orders)
    if err != nil {
        return
    }

    return
}


