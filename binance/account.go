/*
    account.go
        Account Endpoints for Binance Exchange API
*/
package binance

import (
    "fmt"
    "errors"
    //"encoding/json"
)


var OrderSideEnum = map[string]bool {
    "BUY":  true,
    "SELL": true,
}

var OrderTypeEnum = map[string]bool {
    "LIMIT":  true,
    "MARKET": true,
}

var OrderTIFEnum = map[string]bool {
    "GTC": true,
    "IOC": true,
}



// Result from endpoint: GET /api/v3/account
type Account struct {
    MakerCommission  int64     `json:"makerCommission"`
    TakerCommission  int64     `json:"takerCommission"`
    BuyerCommission  int64     `json:"buyerCommission"`
    SellerCommission int64     `json:"sellerCommission"`
    CanTrade         bool      `json:"canTrade"`
    CanWithdraw      bool      `json:"canWithdraw"`
    CanDeposit       bool      `json:"canDeposit"`
    Balances         []Balance `json:"balances"`
}

type Balance struct {
    Asset  string  `json:"asset"`
    Free   float64 `json:"free,string"`
    Locked float64 `json:"locked,string"`
}


func (b *Binance) GetAccountInfo() (account Account, err error) {

    reqUrl := fmt.Sprintf("v3/account")

    _, err = b.client.do("GET", reqUrl, "", true, &account)
    if err != nil {
        return
    }

    return
}


// Result from endpoint: POST /api/v3/order
type PlacedOrder struct {
    Symbol        string `json:"symbol"`
    OrderId       int64  `json:"orderId"`
    ClientOrderId string `json:"clientOrderId"`
    TransactTime  int64  `json:"transactTime"`
}

// Input for endpoint: POST /api/v3/order
type NewOrder struct {
    Symbol      string
    Side        string
    Type        string
    TimeInForce string
    Quantity    float64
    Price       float64
}

func (b *Binance) PlaceOrder(o NewOrder) (res PlacedOrder, err error) {

    if len(o.Symbol) == 0 {
        err = errors.New("Order must contain a symbol")
        return
    }

    if !OrderSideEnum[o.Side] {
        err = errors.New("Invalid or empty order side")
        return
    }

    if !OrderTypeEnum[o.Type] {
        err = errors.New("Invalid or empty order type")
        return
    }

    if !OrderTIFEnum[o.TimeInForce] {
        err = errors.New("Invalid or empty order timeInForce")
        return
    }

    if o.Quantity <= 0.0 {
        err = errors.New("Invalid or empty order quantity")
        return
    }

    if o.Price <= 0.0 {
        err = errors.New("Invalid or empty order price")
        return
    }

    reqUrl := fmt.Sprintf("v3/order?symbol=%s&side=%s&type=%s&timeInForce=%s&quantity=%f&price=%f", o.Symbol, o.Side, o.Type, o.TimeInForce, o.Quantity, o.Price)

    _, err = b.client.do("POST", reqUrl, "", true, &res)
    if err != nil {
        return
    }

    return
}


// Result from endpoint: DELETE /api/v3/order
type DeletedOrder struct {
    Symbol            string `json:"symbol"`
    OrigClientOrderId string `json:"origClientOrderId"`
    OrderId           int64  `json:"orderId"`
    ClientOrderId     string `json:"clientOrderId"`
}

// Input for endpoint: GET & DELETE /api/v3/order
type OrderQuery struct {
    Symbol            string
    OrderId           int64
    RecvWindow        int64
}


func (b *Binance) CancelOrder(query OrderQuery) (order DeletedOrder, err error) {

    if len(query.Symbol) == 0 {
        err = errors.New("OrderQuery must contain a symbol")
        return
    }

    if query.OrderId == 0 {
        err = errors.New("OrderQuery must contain an orderId")
        return
    }

    if query.RecvWindow == 0 {
        query.RecvWindow = 5000
    }

    reqUrl := fmt.Sprintf("v3/order?symbol=%s&orderId=%d&recvWindow", query.Symbol, query.OrderId, query.RecvWindow)

    _, err = b.client.do("DELETE", reqUrl, "", true, &order)
    if err != nil {
        return
    }

    return
}


// Result from endpoint: GET /api/v3/order
type OrderStatus struct {
    Symbol        string  `json:"symbol"`
    OrderId       int64   `json:"orderId"`
    ClientOrderId string  `json:"clientOrderId"`
    Price         float64 `json:"price,string"`
    OrigQty       float64 `json:"origQty,string"`
    ExecutedQty   float64 `json:"executedQty,string"`
    Status        string  `json:"status"`
    TimeInForce   string  `json:"timeInForce"`
    Type          string  `json:"type"`
    Side          string  `json:"side"`
    StopPrice     float64 `json:"stopPrice,string"`
    IcebergQty    float64 `json:"icebergQty,string"`
    Time          int64   `json:"time"`
}


func (b *Binance) CheckOrder(query OrderQuery) (status OrderStatus, err error) {

    if len(query.Symbol) == 0 {
        err = errors.New("OrderQuery must contain a symbol")
        return
    }

    if query.OrderId == 0 {
        err = errors.New("OrderQuery must contain orderId")
        return
    }

    if query.RecvWindow == 0 {
        query.RecvWindow = 5000
    }

    reqUrl := fmt.Sprintf("v3/order?symbol=%s&orderId=%d&origClientOrderId=%s&recvWindow=%d", query.Symbol, query.OrderId, query.RecvWindow)

    _, err = b.client.do("GET", reqUrl, "", true, &status)
    if err != nil {
        return
    }
    
    return
}
