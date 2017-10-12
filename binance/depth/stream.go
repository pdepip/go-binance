package main

import (
    "sync"
    "sort"
    "strings"
    "log"
    "fmt"
    "encoding/json"
    "github.com/gorilla/websocket"
    "go-binance/binance"
)

const (
    MaxDepth = 500 // Size of order book
    MaxQueue = 500 // Size of message queue
)

type State struct {
    EventType string `json:"e"`
    EventTime int64  `json:"E"`
    Symbol    string `json:"s"`
    UpdateId  int64  `json:"u"`
    BidDelta  []binance.Order  `json:"b"`
    AskDelta  []binance.Order  `json:"a"`
}

type OrderBook struct {
    Bids map[float64]float64
    BidMutex sync.Mutex

    Asks map[float64]float64
    AskMutex sync.Mutex
    Updates chan State
}

func (o *OrderBook) ProcessBids(bids []binance.Order) {
    for _, bid := range bids {
        o.BidMutex.Lock()
        if bid.Quantity == 0 {
            delete(o.Bids, bid.Price)
        } else {
            o.Bids[bid.Price] = bid.Quantity
        }
        o.BidMutex.Unlock()
    }
}

func (o *OrderBook) ProcessAsks(asks []binance.Order) {
    for _, ask := range asks {
        o.AskMutex.Lock()
        if ask.Quantity == 0 {
            delete(o.Asks, ask.Price)
        } else {
            o.Asks[ask.Price] = ask.Quantity
        }
        o.AskMutex.Unlock()
    }
}

func (o *OrderBook) Maintainer() {
    for {
        select {
        case job := <- o.Updates:
            if len(job.BidDelta) > 0 {
                go o.ProcessBids(job.BidDelta)
            }

            if len(job.AskDelta) > 0 {
                go o.ProcessAsks(job.AskDelta)
            }
            o.DisplayBook()
        }
    }   
}

func (o *OrderBook) DisplayBook() {

    o.BidMutex.Lock()
    o.AskMutex.Lock()
    var bidPrices []float64
    var askPrices []float64

    for k := range o.Bids {
        bidPrices = append(bidPrices, k)
    }
    
    for k := range o.Asks {
        askPrices = append(askPrices, k)
    }

    sort.Float64s(bidPrices)
    sort.Float64s(askPrices)

    fmt.Println("BIDS:", o.Bids)
    fmt.Println("ASKS:", o.Asks)
    fmt.Println("")

    o.BidMutex.Unlock()
    o.AskMutex.Unlock()
}

func main() {

    var wsDialer websocket.Dialer

    symbol := "ethbtc"

    address := fmt.Sprintf("wss://stream.binance.com:9443/ws/%s@depth", symbol)

    wsConn, _, err := wsDialer.Dial(address, nil)
    if err != nil {
        panic(err)
    }
    defer wsConn.Close()
    log.Println("Dialed:", address)

    ob := OrderBook{} 
    ob.Bids = make(map[float64]float64, MaxDepth)
    ob.Asks = make(map[float64]float64, MaxDepth)

    ob.Updates = make(chan State, 500)

    client := binance.New("", "")
    query := binance.OrderBookQuery {
        Symbol: strings.ToUpper(symbol),
    }
    orderBook, err := client.GetOrderBook(query)
    if err != nil {
        panic(err)
    }

    ob.ProcessBids(orderBook.Bids)
    ob.ProcessAsks(orderBook.Asks)

    go ob.Maintainer()

    for {
        _, message, err := wsConn.ReadMessage()
        if err != nil {
            log.Println("[ERROR] ReadMessage:", err)
        }

        msg := State{}
        err = json.Unmarshal(message, &msg)
        if err != nil {
            log.Println("[ERROR] Parsing:", err)
            continue
        }

        ob.Updates <- msg
    }

}


