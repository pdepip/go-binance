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
    "fmt"
    "strings"
    "net/http"
    "encoding/json"
)

type Client struct {
    key        string
    secret     string
    httpClient *http.Client
}

// Creates a new Binance HTTP Client
func NewClient(key, secret string) (c *Client) {
    client := &Client{
        key: key,
        secret: secret,
        httpClient: &http.Client{},
    }
    return client
}

func (self *Client) do(method, resource, payload string, auth bool, result interface{}) (res *http.Response, err error) {

    fullUrl := fmt.Sprintf("%s/%s/%s", BaseUrl, Version, resource)

    req, err := http.NewRequest(method, fullUrl, strings.NewReader(payload))
    if err != nil {
        return
    }

    req.Header.Add("Accept", "application/json")

    resp, err := self.httpClient.Do(req)
    if err != nil {
        return
    }
    defer resp.Body.Close()

    if resp != nil {
        decoder := json.NewDecoder(resp.Body)
        err = decoder.Decode(result)
        return
    }
    return
}


