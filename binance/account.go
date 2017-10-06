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
    Asset  string  `json:asset`
    Free   string  `json:free, string`
    Locked float64 `json:locked, string`
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


func (b *Binance) GetAccountInfo() (account Account, err error) {

    reqUrl := fmt.Sprintf("v3/account")

    _, err = b.client.do("GET", reqUrl, "", true, &account)

    return
}




