/*
    account.go
        Account Endpoints for Binance Exchange API
*/
package binance


import (
    "fmt"
    //"errors"
    //"encoding/json"
)


type Balance struct {
    Asset  string  `json:"asset"`
    Free   float64 `json:"free,string"`
    Locked float64 `json:"locked,string"`
}


// Result from endpoint: /api/v3/account
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

// Result from endpoint: /api/v3/order
type PlacedOrder struct {
    Symbol        string `json:"symbol"`
    OrderId       int64  `json:"orderId"`
    ClientOrderId string `json:"clientOrderId"`
    TransactTime  int64  `json:"transactTime"`
}


func (b *Binance) GetAccountInfo() (account Account, err error) {

    reqUrl := fmt.Sprintf("v3/account")

    _, err = b.client.do("GET", reqUrl, "", true, &account)
    if err != nil {
        return
    }

    return
}


func (b *Binance) NewOrder(symbol string, side string, orderType string, timeInForce string, quantity float64, price float64) (order PlacedOrder, err error) {

    reqUrl := fmt.Sprintf("v3/order?symbol=%s&side=%s&type=%s&timeInForce=%s&quantity=%f&price=%f", symbol, side, orderType, timeInForce, quantity, price)

    _, err = b.client.do("POST", reqUrl, "", true, &order)
    if err != nil {
        return
    }

    return
}
