# Go Binance API

## Summary
Go client for [Binance](https://www.binance.com)

## Installation
```go get https://github.com/pdepip/go-binance.git```

## Documentation

Full documentation can be found in the [GoDoc]() documentation

### Setup

Creating a client:

```go
import (
	"os"
	"binance"
)

// Secure method
secret := os.Getenv("BINANCE_SECRET")
key    := os.Getenv("BINANCE_KEY")

// Unsecure method
secret := "mySecret"
key    := "myKey"

client :=  binance.New(secret, key)
```

### Examples

Get Current Positions

```go
package main

import (
    "os"
    "fmt"
    "go-binance/binance"
)

func main() {

    client := binance.New(os.Getenv("key"), os.Getenv("secret"))
    positions, err := client.GetPositions()

    if err != nil {
        panic(err)
    }

    for _, p := range positions {
        fmt.Println(p.Asset, p.Free, p.Locked)
    }
}
```

Place a Limit Order

```go
package main

import (
	"os"
	"fmt"
	"go-binance/binance"
)

func main() {
    // Params
    order := binance.LimitOrder {
        Symbol:      "BNBBTC",
        Side:        "BUY",
        Type:        "LIMIT",
        TimeInForce: "GTC",
        Quantity:    50.0,
        Price:       0.00025,
    }

    client := binance.New(os.Getenv("key"), os.Getenv("secret"))
    res, err := client.PlaceLimitOrder(order)
    
    if err != nil {
    	panic(err)
    }
    
    fmt.Println(res)
}
```

Get the Order Book

```go
import (
	"fmt"
	"go-binance/binance"
)

func main() {

    // Params
    query := binance.OrderBookQuery {
        Symbol: "BNBBTC",
        Limit: 100,
    }

    client := binance.New("", "")
    res, err := client.GetOrderBook(query)

    if err != nil {
        panic(err)
    }
    
    fmt.Println(res)

}
```

### Local Depth Cache

See `examples/stream.go`