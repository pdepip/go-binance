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
    "time"
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

    if auth {

        if len(c.apiKey) == 0 || len(c.apiSecret) == 0 {
            err = errors.New("Private endpoints requre you to set an API Key and API Secret")
            return
        }

        req.Header.Add("X-MBX-APIKEY", c.apiKey)

        timestamp := time.Now().Unix() * 1000
        q := req.URL.Query()
        q.Set("timestamp", timestamp)

        fmt.Println(q)
        
    }   

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


