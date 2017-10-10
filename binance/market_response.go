/*

    Stores response structs for API functions in market.go

*/

package binance

// Result from: GET /api/v1/depth
type OrderBook struct {
    LastUpdatedId int64 `json:"lastUpdatedId"`
    Bids []Order `json:"bids"`
    Asks []Order `json:"asks"`
}

type Order struct {
    Price    float64 `json:",string"`
    Quantity float64 `json:",string"`
}
