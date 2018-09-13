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
	"net/http"
)

//"errors"

const (
	BaseUrl = "https://api.binance.com"
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

// Override the http.Client to use for network connections
//
// This is used for setting timeouts and other network settings and
// can also be useful for sharing connections when using multiple Binance
// instances.
func (b *Binance) SetHttpClient(client *http.Client) {
	b.client.httpClient = client
}
