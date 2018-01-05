# Go Binance API

## Summary
Go client for [Binance](https://www.binance.com)

## Installation
```go get github.com/pdepip/go-binance/binance```

## Documentation
Full API Documentation can be found at https://www.binance.com/restapipub.html

## Setup

Creating a client:

```go
import (
	"os"
	"go-binance/binance"
)

// Secure method
secret := os.Getenv("BINANCE_SECRET")
key    := os.Getenv("BINANCE_KEY")

// Unsecure method
secret := "mySecret"
key    := "myKey"

client :=  binance.New(secret, key)
```

## Examples

### Get Current Positions

```go
package main

import (
    "os"
    "fmt"
    "github.com/pdepip/go-binance/binance"
)

func main() {

    client := binance.New(os.Getenv("BINANCE_KEY"), os.Getenv("BINANCE_SECRET"))
    positions, err := client.GetPositions()

    if err != nil {
        panic(err)
    }

    for _, p := range positions {
        fmt.Println(p.Asset, p.Free, p.Locked)
    }
}
```

### Place a Limit Order

```go
package main

import (
	"os"
	"fmt"
	"github.com/pdepip/go-binance/binance"
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

    client := binance.New(os.Getenv("BINANCE_KEY"), os.Getenv("BINANCE_SECRET"))
    res, err := client.PlaceLimitOrder(order)
    
    if err != nil {
    	panic(err)
    }
    
    fmt.Println(res)
}
```

### Place a Market Order

```go
package main

import (
	"os"
	"fmt"
	"github.com/pdepip/go-binance/binance"
)

func main() {
    // Params
    order := binance.MarketOrder {
        Symbol:   "BNBBTC",
        Side:     "BUY",
        Type:     "MARKET",
        Quantity: 50.0,
    }

    client := binance.New(os.Getenv("BINANCE_KEY"), os.Getenv("BINANCE_SECRET"))
    res, err := client.PlaceMarketOrder(order)
    
    if err != nil {
    	panic(err)
    }
    
    fmt.Println(res)
}
```

### Check Order Status

```go
import (
	"os"
	"fmt"
	"github.com/pdepip/go-binance/binance"
)

func main() {
    // Params
    orderQuery := binance.OrderQuery {
        Symbol:  "BNBBTC",
        OrderId: "yourOrderId",
    }

    client := binance.New(os.Getenv("BINANCE_KEY"), os.Getenv("BINANCE_SECRET"))
    res, err := client.CheckOrder(orderQuery)
    
    if err != nil {
    	panic(err)
    }
    
    fmt.Println(res)
}
```

### Cancel an Order

```go
import (
	"os"
	"fmt"
	"github.com/pdepip/go-binance/binance"
)

func main() {
    // Params
    orderQuery := binance.OrderQuery {
        Symbol:  "BNBBTC",
        OrderId: "yourOrderId",
    }

    client := binance.New(os.Getenv("BINANCE_KEY"), os.Getenv("BINANCE_SECRET"))
    res, err := client.CancelOrder(orderQuery)
    
    if err != nil {
    	panic(err)
    }
    
    fmt.Println(res)
}
```

### Get Open Orders

```go

import (
	"os"
	"fmt"
	"github.com/pdepip/go-binance/binance"
)

func main() {
    // Params
    orderQuery := binance.OpenOrdersQuery {
        Symbol: "BNBBTC",
    }

    client := binance.New(os.Getenv("BINANCE_KEY"), os.Getenv("BINANCE_SECRET"))
    res, err := client.GetOpenOrders(orderQuery)
    
    if err != nil {
    	panic(err)
    }
    
    fmt.Println(res)
}

```

### Get the Order Book

```go
import (
	"fmt"
	"github.com/pdepip/go-binance/binance"
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

### Get Latest Price of a Symbol

```go
import (
	"fmt"
	"github.com/pdepip/go-binance/binance"
)

func main() {

    // Params
    query := binance.SymbolQuery {
        Symbol: "BNBBTC",
    }

    client := binance.New("", "")
    res, err := client.GetLastPrice(query)

    if err != nil {
        panic(err)
    }
    
    fmt.Println(res)

}
```

### Local Depth Cache

See `examples/depth.go`. Script connects to Binance websocket and maintains a simple local depth cache.
