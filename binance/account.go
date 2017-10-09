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

func (b *Binance) NewOrder(symbol string, side string, orderType string, timeInForce string, quantity float64, price float64) (order PlacedOrder, err error) {

    reqUrl := fmt.Sprintf("v3/order?symbol=%s&side=%s&type=%s&timeInForce=%s&quantity=%f&price=%f", symbol, side, orderType, timeInForce, quantity, price)

    _, err = b.client.do("POST", reqUrl, "", true, &order)
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


func (b *Binance) CancelOrder(symbol string, orderId int64) (order DeletedOrder, err error) {

    reqUrl := fmt.Sprintf("v3/order?symbol=%s&orderId=%d", symbol, orderId)

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

type OrderQuery struct {
    Symbol            string
    OrderId           int64
    OrigClientOrderId string
    RecvWindow        int64
}

func (b *Binance) CheckOrder(query OrderQuery) (status OrderStatus, err error) {

    if len(query.Symbol) == 0 {
        err = errors.New("OrderQuery must contain a val symbol")
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
