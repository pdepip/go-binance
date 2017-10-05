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


