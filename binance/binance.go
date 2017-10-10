/*
    binance.go
        Wrapper for the Binance Exchange API

    Authors:
        Pat DePippo  <patrick.depippo@dcrypt.io>
        Matthew Woop <matthew.woop@dcrypt.io>

    To Do:

*/
package binance

import (
    //"errors"
)

const (
    BaseUrl = "https://www.binance.com/api"
)

type Binance struct {
    client *Client
}
/*
func handleErr(r jsonResponse) error {

    if !r.Success {
        return errors.New(r.Message)
    }

    return nil
}
*/
func New(key, secret string) *Binance {
    client := NewClient(key, secret)
    return &Binance{client}
}


