/*
    client.go
        Wrapper for the Binance Exchange API

    Authors:
        Pat DePippo  <patrick.depippo@dcrypt.io>
        Matthew Woop <matthew.woop@dcrypt.io>

    To Do:

*/
package binance

import (
)

const (
    BaseUrl = "https://www.binance.com/api"
)

type Binance struct {
    client *Client
}

func New(key, secret string) *Binance {
    client := NewClient(key, secret)
    return &Binance{client}
}


