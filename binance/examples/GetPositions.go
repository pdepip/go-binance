package main

import (
    "os"
    //"fmt"
    "go-binance/binance"
)

func main() {

    query := binance.OrderBookQuery {
        Symbol: "BANBBTC",
    }

    client := binance.New(os.Getenv("key"), os.Getenv("secret"))
    _, err := client.GetOrderBook(query)

    if err != nil {
        panic(err)
    }

    //fmt.Println(res)
    /*
    for _, p := range positions {
        fmt.Println(p.Asset, p.Free, p.Locked)
    }
    */
}

