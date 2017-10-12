/*

    get_positions.go
        Example script showing how to call route: /api/v3/account

*/

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

