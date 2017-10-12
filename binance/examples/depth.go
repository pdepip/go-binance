/*
    depth.go
        Connects to the Binance WebSocket and maintains
        local depth cache.

*/

package main

import (
    "sync"
    "strings"
    "log"
    "fmt"
    "encoding/json"
    "github.com/gorilla/websocket"
    "go-binance/binance"
)

const (
    MaxDepth = 100 // Size of order book
    MaxQueue = 100 // Size of message queue
)

// Message received from websocket
type State struct {
    EventType string           `json:"e"`
    EventTime int64            `json:"E"`
    Symbol    string           `json:"s"`
    UpdateId  int64            `json:"u"`
    BidDelta  []binance.Order  `json:"b"`
    AskDelta  []binance.Order  `json:"a"`
}


// Orderbook structure
type OrderBook struct {
    Bids map[float64]float64 // Map of all bids, key->price, value->quantity
    BidMutex sync.Mutex      // Threadsafe

    Asks map[float64]float64 // Map of all asks, key->price, value->quantity
    AskMutex sync.Mutex      // Threadsafe

    Updates chan State       // Channel of all state updates
}


// Process all incoming bids
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


// Process all incoming asks
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


// Hands off incoming messages to processing functions
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
        }
    }   
}


func main() {

    symbol := "ethbtc"
    address := fmt.Sprintf("wss://stream.binance.com:9443/ws/%s@depth", symbol)

    // Connect to websocket
    var wsDialer websocket.Dialer
    wsConn, _, err := wsDialer.Dial(address, nil)
    if err != nil {
        panic(err)
    }
    defer wsConn.Close()
    log.Println("Dialed:", address)

    // Set up Order Book
    ob := OrderBook{} 
    ob.Bids = make(map[float64]float64, MaxDepth)
    ob.Asks = make(map[float64]float64, MaxDepth)
    ob.Updates = make(chan State, 500)

    // Get initial state of orderbook from rest api
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

    // Start maintaining order book
    go ob.Maintainer()

    // Read & Process Messages from wss stream
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
