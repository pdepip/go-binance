/*

    Stores response structs for API functions account.go

*/

package binance


// Result from: GET /api/v3/account
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


// Result from: POST /api/v3/order
type PlacedOrder struct {
    Symbol        string `json:"symbol"`
    OrderId       int64  `json:"orderId"`
    ClientOrderId string `json:"clientOrderId"`
    TransactTime  int64  `json:"transactTime"`
}


// Result from: DELETE /api/v3/order
type CanceledOrder struct {
    Symbol            string `json:"symbol"`
    OrigClientOrderId string `json:"origClientOrderId"`
    OrderId           int64  `json:"orderId"` 
    ClientOrderId     string `json:"clientOrderId"`
}


// Result from: GET /api/v3/order
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
