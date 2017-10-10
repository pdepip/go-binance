/*
    account.go
        Account Endpoints for Binance Exchange API
*/
package binance

import (
    "fmt"
    "errors"
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


// Input for endpoint: POST /api/v3/order
type LimitOrder struct {
    Symbol      string
    Side        string
    Type        string
    TimeInForce string
    Quantity    float64
    Price       float64
    RecvWindow  int64
}

// Validating a Limit Order
func (l *LimitOrder) ValidateLimitOrder() error {
    switch {
        case len(l.Symbol) == 0:
            return errors.New("Order must contain a symbol")
        case !OrderSideEnum[l.Side]:
            return errors.New("Invalid or empty order side")
        case l.Type != "LIMIT":
            return errors.New("Invalid LIMIT order type")
        case !OrderTIFEnum[l.TimeInForce]:
            return errors.New("Invalid or empty order timeInForce")
        case l.Quantity <= 0.0:
            return errors.New("Invalud or empty order quantity")
        case l.Price <= 0.0:
           return errors.New("Invalud or empty order price")
        case l.RecvWindow == 0:
            l.RecvWindow = 5000
            return nil
        default:
            return nil
    }
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


type MarketOrder struct {
    Symbol      string
    Side        string
    Type        string
    Quantity    float64
    RecvWindow  int64
}

func (m *MarketOrder) ValidateMarketOrder() error {
    switch {
        case len(m.Symbol) == 0:
            return errors.New("Order must contain a symbol")
        case !OrderSideEnum[m.Side]:
            return errors.New("Invalid or empty or side")
        case m.Type != "MARKET":
            return errors.New("Invalid type for a market order")
        case m.Quantity <= 0.0:
            return errors.New("Invalid or empty order quantity")
        case m.RecvWindow == 0:
            m.RecvWindow = 5000
            return nil
        default:
            return nil
    }
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


// Input for endpoint: GET & DELETE /api/v3/order
type OrderQuery struct {
    Symbol     string
    OrderId    int64
    RecvWindow int64
}


func (o *OrderQuery) ValidateOrderQuery() error {
    switch {
        case len(o.Symbol) == 0:
            return errors.New("OrderQuery must contain a Symbol")
        case o.OrderId == 0:
            return errors.New("OrderQuery must contain an OrderId")
        case o.RecvWindow == 0:
            o.RecvWindow = 5000
            return nil
        default:
            return nil
    }
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


// Input for endpoint: GET /api/v3/openOrders
type OpenOrdersQuery struct {
    Symbol string
    RecvWindow int64
}

func (o *OpenOrdersQuery) ValidateOpenOrdersQuery() error {
    switch {
        case len(o.Symbol) == 0:
            return errors.New("Invalid or empty symbol")
        case o.RecvWindow == 0:
            o.RecvWindow = 5000
            return nil
        default:
            return nil
    }
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


