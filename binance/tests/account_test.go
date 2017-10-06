package account_test

import (
    "os"
    "fmt"
    "testing"
    "go-binance/binance"
)

func TestGetAccountInfo(t *testing.T) {

    binance := binance.New(os.Getenv("key"), os.Getenv("secret"))
    account, err := binance.GetAccountInfo()

    if err != nil {
        t.Fatal(err)
    }

    fmt.Printf("%+v\n", account)
}
