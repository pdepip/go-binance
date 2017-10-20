/*

    get_positions.go
        Example script showing how to call route: /api/v3/account

*/

package main

import (
    "os"
    //"fmt"
    "go-binance/binance"
)

func main() {

    query := binance.SymbolQuery {
        Symbol: "BNBBTC",
    }

    client := binance.New(os.Getenv("BINANCE_KEY"), os.Getenv("BINANCE_SECRET"))
    _, err := client.GetAllPrices(query)

    if err != nil {
        panic(err)
    }
    /*
    for _, p := range positions {
        fmt.Println(p.Asset, p.Free, p.Locked)
    }
    */
}

