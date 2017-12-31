/*

    Stores response structs for API functions account.go

*/

package binance


// Result from: GET /api/v3/exchangeInfo
type ExchangeInfo struct {
    ExchangeFilters []string     `json:"ExchangeFilters"`
    RateLimits      []RateLimit  `json:"rateLimits"`
    ServerTime      int64        `json:"serverTime"`
    Symbols         []SymbolInfo `json:"symbols"`
    TimeZone        string       `json:"timezone"`
}

type SymbolInfo struct {
    Symbol             string         `json:"symbol"`
    BaseAsset          string         `json:"baseAsset"`
    QuotePrecision     int64          `json:"quotePrecision"`
    BaseAssetPrecision int64          `json:"baseAssetPrecision"`
    Status             string         `json:"status"`
    OrderTypes         []string       `json:"orderTypes"`
    Filters            []SymbolFilter `json:"filters"`
    QuoteAsset         string         `json:"quoteAsset"`
    IceBergAllowed     bool           `json:"icebergAllowed"`
}

type SymbolFilter struct {
    Type        string `json:"filterType"`
    MinPrice    float64 `json:"minPrice,string"`
    MaxPrice    float64 `json:"maxPrice,string"`
    TickSize    float64 `json:"tickSize,string"`
    StepSize    float64 `json:"stepSize,string"`
    MinQty      float64 `json:"minQty,string"`
    MaxQty      float64 `json:"maxQty,string"`
    MinNotional float64 `json:"minNotional,string"`
}

type RateLimit struct {
    Limit         int64  `json:"limit"`
    Interval      string `json:"interval"`
    RateLimitType string `json:"rateLimitType"`
}
