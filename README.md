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

    binance := binance.New(os.Getenv("key"), os.Getenv("secret"))
    positions, err := binance.GetPositions()

    if err != nil {
        panic(err)
    }

    for _, p := range positions {
        fmt.Println(p.Asset, p.Free, p.Locked)
    }
}
```

Get Latest Ticker Price

```go
import fmt

pair := "BTC-BNB"

price, err := binance.GetTicker(pair)

if err != nil {
	panic(err)
}

fmt.println(pair, price)
```

Place A Limit Order

```go
import fmt

orderId, err := binance.BuyLimit("BTC-BNB", 50, "0.00038588")

if err != nil {
    panic(err)
}

fmt.println(orderId)
```

Place A Market Order

```go
import fmt

orderId, err := binance.BuyMarket("BTC-BNB", 50)

if err != nil {
    panic(err)
}

fmt.println(orderId)
```

Checking Order Status

```go
import fmt

status, err := binance.CheckOrder("gd87xs6jx00v")

if err != nil {
    panic(err)
}

fmt.println(status)
```

Cancel An Order

```go
import fmt

status, err := binance.CancelOrder("BNBBTC", orderId)

if err != nil {
    panic(err)
}

fmt.println(status)
```

Get Open Orders

```go
import fmt

orders, err = binance.OpenOrders("BNBBTC")

if err != nil {
    panic(nil)
}

for _, o := range orders {
    fmt.println(order)
}

```

Get Ticker Depth

```go
import fmt


```
