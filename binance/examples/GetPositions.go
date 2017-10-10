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

