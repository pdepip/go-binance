/*

    Stores response structs for API functions in market.go

*/

package binance

// Result from: GET /api/v1/depth
type OrderBook struct {
    LastUpdatedId int64 `json:"lastUpdatedId"`
    Bids []Order `json:"bids"`
    Asks []Order `json:"asks"`
    //Msg  string  `json:"msg"`
}

type Order struct {
    Price    float64 `json:",string"`
    Quantity float64 `json:",string"`
}


// Result from: GET /api/v1/ticker/24hr
type ChangeStats struct {
    PriceChange        float64 `json:"priceChange,string"`
    PriceChangePercent float64 `json:"priceChangePercent,string"`
    WeightedAvgPrice   float64 `json:"weightedAvgPrice,string"`
    PrevClosePrice     float64 `json:"prevClosePrice,string"`
    LastPrice          float64 `json:"lastPrice,string"`
    BidPrice           float64 `json:"bidPrice,string"`
    AskPrice           float64 `json:"askPrice,string"`
    OpenPrice          float64 `json:"openPrice,string"`
    HighPrice          float64 `json:"highPrice,string"`
    LowPrice           float64 `json:"lowPrice,string"`
    Volume             float64 `json:"volume,string"`
    OpenTime           int64   `json:"openTime"`
    CloseTime          int64   `json:"closeTime"`
    FirstId            int64   `json:"firstId"`
    LastId             int64   `json:"lastId"`
    Count              int64   `json:"count"`
    Msg                string  `json:"msg"`
}


// Result from: GET /api/v1/aggTrade
type AggTrade struct {
    TradeId      int64   `json:"a"`
    Price        float64 `json:"p,string"`
    Quantity     float64 `json:"q,string"`
    FirstTradeId int64   `json:"f"`
    LastTradeId  int64   `json:"l"`
    Timestamp    int64   `json:"T"`
    Maker        bool    `json:"m"`
    BestMatch    bool    `json:"M"`
    Msg          string  `json:"msg"`
}


// Result from: GET /api/v1/allPrices
type TickerPrice struct {
    Symbol string  `json:"symbol"`
    Price  float64 `json:"price,string"`
    Msg    string  `json:"msg"`
}


// Result from: GET /api/v1/allBookTickers
type BookTicker struct {
    Symbol      string  `json:"symbol"`
    BidPrice    float64 `json:"bidPrice,string"`
    BidQuantity float64 `json:"bidQuantity,string"`
    AskPrice    float64 `json:"askPrice,string"`
    AskQuantity float64 `json:"askQuantity,string"`
    Msg         string  `json:"msg"`
}


// Result from: GET /api/v1/klines

type Kline struct {
    OpenTime         int64 
    Open             float64
    High             float64
    Low              float64
    Close            float64
    Volume           float64
    CloseTime        int64
    QuoteVolume      float64
    NumTrades        int64
    TakerBaseVolume  float64
    TakerQuoteVolume float64
    Msg              string
}
