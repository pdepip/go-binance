/*

    Example script showing how to call GetPositions

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

